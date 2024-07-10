package gorm

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	// database driver for gorm
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"

	"github.com/bnb-chain/token-recover-app/internal/config"
	"github.com/bnb-chain/token-recover-app/internal/store"
	"github.com/bnb-chain/token-recover-app/pkg/util"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
)

type DataSourceTypeName string

const (
	Mysql      DataSourceTypeName = "mysql"
	Postgresql DataSourceTypeName = "postgres"
	Sqlite     DataSourceTypeName = "sqlite"
)

var _supportedDataSource = map[DataSourceTypeName]func(port uint, host, dbname, user, password, connectTimeout, readTimeout, writeTimeout string, sslmode bool) gorm.Dialector{
	Mysql: func(port uint, host, dbname, user, password, connectTimeout, readTimeout, writeTimeout string, sslmode bool) gorm.Dialector {
		return mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=UTC&time_zone=UTC&timeout=%s&readTimeout=%s&writeTimeout=%s", user, password, host, port, dbname, connectTimeout, readTimeout, writeTimeout))
	},
	Postgresql: func(port uint, host, dbname, user, password, connectTimeout, readTimeout, writeTimeout string, sslmode bool) gorm.Dialector {
		ssl := "disable"
		if sslmode {
			ssl = "allow"
		}
		return postgres.Open(fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s timezone=UTC", host, port, user, dbname, password, ssl))
	},
	Sqlite: func(port uint, host, dbname, user, password, connectTimeout, readTimeout, writeTimeout string, sslmode bool) gorm.Dialector {
		return sqlite.Open(fmt.Sprintf("%s.db", dbname))
	},
}

var _ store.ProofStore = (*SQLStore)(nil)
var _ store.BscBlockStore = (*SQLStore)(nil)
var _ store.TokenRecoverEventStore = (*SQLStore)(nil)
var _ store.GeneralStore = (*SQLStore)(nil)

func NewSQLStore(config *config.Config, options ...Option) (*SQLStore, error) {
	supported, ok := _supportedDataSource[DataSourceTypeName(config.Store.SqlStore.SQLDriver)]
	if !ok {
		return nil, fmt.Errorf("unsupported database driver: %s", config.Store.SqlStore.SQLDriver)
	}

	sqlDriver := supported(
		config.Store.SqlStore.Port, config.Store.SqlStore.Host, config.Store.SqlStore.DBName,
		config.Store.SqlStore.User, config.Store.SqlStore.Password,
		config.Store.SqlStore.ConnectTimeout, config.Store.SqlStore.ReadTimeout, config.Store.SqlStore.WriteTimeout,
		config.Store.SqlStore.SSLMode)

	engine, err := gorm.Open(sqlDriver, &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.Store.SqlStore.DialTimeout)
	defer cancel()

	sql, err := engine.DB()
	if err != nil {
		return nil, err
	}
	err = sql.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	for _, opt := range options {
		err = opt.Apply(engine)
		if err != nil {
			return nil, err
		}
	}

	for _, migration := range Migrations {
		err := migration.Migrate(engine)
		if err != nil {
			return nil, err
		}
	}

	return &SQLStore{
		db: engine,
	}, nil
}

// SQLStore implements store.Store.
type SQLStore struct {
	db *gorm.DB
}

// BscBlockStore implements store.GeneralStore.
func (s *SQLStore) BscBlockStore() store.BscBlockStore {
	return s
}

// ProofStore implements store.GeneralStore.
func (s *SQLStore) ProofStore() store.ProofStore {
	return s
}

// TokenRecoverEventStore implements store.GeneralStore.
func (s *SQLStore) TokenRecoverEventStore() store.TokenRecoverEventStore {
	return s
}

// GetAccountProof implements store.Store.
func (s *SQLStore) GetAccountAssetProof(address types.AccAddress, symbol string) (*store.Proof, error) {
	var proof Proof
	result := s.db.Where("address = ? AND denom = ?", util.EncodeBytesToHex(address), symbol).First(&proof)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, store.ErrRecordNotFound
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return &store.Proof{
		Address: types.AccAddress(util.MustDecodeHexToBytes(proof.Address)),
		Denom:   proof.Denom,
		Amount:  proof.Amount,
		Proof:   util.MustDecodeHexArrayToBytes(strings.Split(proof.Proof, ",")),
	}, nil
}

// GetAccountAssetProofs implements store.ProofStore.
func (s *SQLStore) GetAccountAssetProofs(address types.AccAddress, pagination store.Pagination) ([]*store.Proof, int64, error) {
	var count int64
	result := s.db.Model(&Proof{}).Where("address = ?", util.EncodeBytesToHex(address)).Count(&count)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	var data []*Proof
	result = s.db.Where("address = ?", util.EncodeBytesToHex(address)).Offset(pagination.Offset).Limit(pagination.Limit).Find(&data)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	proofs := make([]*store.Proof, 0, len(data))
	for _, proof := range data {
		proofs = append(proofs, &store.Proof{
			Address: types.AccAddress(util.MustDecodeHexToBytes(proof.Address)),
			Denom:   proof.Denom,
			Amount:  proof.Amount,
			Proof:   util.MustDecodeHexArrayToBytes(strings.Split(proof.Proof, ",")),
		})
	}
	return proofs, count, nil
}

// InsertAccountProof implements a function to insert account proof.
func (s *SQLStore) InsertAccountAssetProof(proof *store.Proof) error {
	dbProof := &Proof{
		Address: util.EncodeBytesToHex(proof.Address[:]),
		Denom:   proof.Denom,
		Amount:  proof.Amount,
		Proof:   strings.Join(util.EncodeBytesArrayToHex(proof.Proof), ","),
	}

	result := s.db.Create(dbProof)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// CountAccountAssetProofs implements store.Store.
func (s *SQLStore) CountAccountAssetProofs() (int64, error) {
	var count int64
	result := s.db.Model(&Proof{}).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

// BatchSaveTokenRecoverEvent implements store.TokenRecoverEventStore.
func (s *SQLStore) BatchSaveTokenRecoverEvent(events []*store.TokenRecoverEvent) error {
	tokenRecoverEvents := make([]TokenRecoverEvent, 0, len(events))
	for _, event := range events {
		tokenRecoverEvents = append(tokenRecoverEvents, TokenRecoverEvent{
			TokenOwner:           util.EncodeBytesToHex(event.TokenOwner),
			TokenContractAddress: event.TokenContractAddress.Hex(),
			Denom:                event.Denom,
			Amount:               event.Amount.String(),
			ClaimAddress:         event.ClaimAddress.Hex(),
			UnlockAt:             event.UnlockAt,
			Status:               int8(event.Status),
			RecoveredBlockNumber: event.RecoveredBlockNumber,
			RecoveredTxHash:      event.RecoveredTxHash.Hex(),
			WithdrawTxHash:       event.WithdrawTxHash.Hex(),
			CancelledTxHash:      event.CancelledTxHash.Hex(),
		})
	}

	return s.db.Model(&TokenRecoverEvent{}).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "token_owner"}, {Name: "denom"}, {Name: "amount"}}, // key colume
		DoUpdates: clause.AssignmentColumns([]string{
			"token_owner",
			"token_contract_address",
			"denom",
			"amount",
			"claim_address",
			"unlock_at",
			"status",
			"recovered_block_number",
			"recovered_tx_hash",
			"withdraw_tx_hash",
			"cancelled_tx_hash",
		}), // column needed to be updated
	}).Create(&tokenRecoverEvents).Error
}

func (s *SQLStore) convertCondition(condition store.TokenRecoverEvent) *TokenRecoverEvent {
	result := &TokenRecoverEvent{}
	if !condition.TokenOwner.Empty() {
		result.TokenOwner = util.EncodeBytesToHex(condition.TokenOwner)
	}
	if condition.TokenContractAddress != store.EmptyAccount {
		result.TokenContractAddress = condition.TokenContractAddress.Hex()
	}
	if condition.Denom != "" {
		result.Denom = condition.Denom
	}
	if condition.Amount != nil {
		result.Amount = condition.Amount.String()
	}
	if condition.ClaimAddress != store.EmptyAccount {
		result.ClaimAddress = condition.ClaimAddress.Hex()
	}
	if condition.UnlockAt != 0 {
		result.UnlockAt = condition.UnlockAt
	}
	if condition.Status != 0 {
		result.Status = int8(condition.Status)
	}
	if condition.RecoveredBlockNumber != 0 {
		result.RecoveredBlockNumber = condition.RecoveredBlockNumber
	}
	if condition.RecoveredTxHash != store.EmptyTxHash {
		result.RecoveredTxHash = condition.RecoveredTxHash.Hex()
	}
	if condition.WithdrawTxHash != store.EmptyTxHash {
		result.WithdrawTxHash = condition.WithdrawTxHash.Hex()
	}
	if condition.CancelledTxHash != store.EmptyTxHash {
		result.CancelledTxHash = condition.CancelledTxHash.Hex()
	}
	return result
}

// GetTokenRecoverEvents implements store.TokenRecoverEventStore.
func (s *SQLStore) GetTokenRecoverEvents(condition store.TokenRecoverEvent, pagination store.Pagination, extraCondition *store.ExtraCondition) ([]*store.TokenRecoverEvent, int64, error) {
	tx := s.db.Model(&TokenRecoverEvent{}).Where(s.convertCondition(condition))
	if extraCondition != nil && extraCondition.AllowUnlocked {
		tx.Where("unlock_at < ? AND status < ?", time.Now().Unix(), store.Withdrawing)
	}

	var count int64
	result := tx.Count(&count)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	var events []*TokenRecoverEvent
	result = tx.
		Offset(pagination.Offset).Limit(pagination.Limit).Find(&events)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	tokenRecoverEvents := make([]*store.TokenRecoverEvent, 0, len(events))
	for _, event := range events {
		tokenRecoverEvents = append(tokenRecoverEvents, &store.TokenRecoverEvent{
			TokenOwner:           types.AccAddress(util.MustDecodeHexToBytes(event.TokenOwner)),
			TokenContractAddress: common.HexToAddress(event.TokenContractAddress),
			Denom:                event.Denom,
			Amount:               util.MustDecodeStringToBigInt(event.Amount),
			ClaimAddress:         common.HexToAddress(event.ClaimAddress),
			UnlockAt:             event.UnlockAt,
			Status:               store.TokenRecoverStatus(event.Status),
			RecoveredBlockNumber: event.RecoveredBlockNumber,
			RecoveredTxHash:      common.HexToHash(event.RecoveredTxHash),
			WithdrawTxHash:       common.HexToHash(event.WithdrawTxHash),
			CancelledTxHash:      common.HexToHash(event.CancelledTxHash),
		})
	}
	return tokenRecoverEvents, count, nil
}

// GetTokenRecoverEvent implements store.TokenRecoverEventStore.
func (s *SQLStore) GetTokenRecoverEvent(condition store.TokenRecoverEvent) (*store.TokenRecoverEvent, error) {
	var event TokenRecoverEvent
	result := s.db.Model(&TokenRecoverEvent{}).Where(s.convertCondition(condition)).First(&event)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, store.ErrRecordNotFound
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &store.TokenRecoverEvent{
		TokenOwner:           types.AccAddress(util.MustDecodeHexToBytes(event.TokenOwner)),
		TokenContractAddress: common.HexToAddress(event.TokenContractAddress),
		Denom:                event.Denom,
		Amount:               util.MustDecodeStringToBigInt(event.Amount),
		ClaimAddress:         common.HexToAddress(event.ClaimAddress),
		UnlockAt:             event.UnlockAt,
		Status:               store.TokenRecoverStatus(event.Status),
		RecoveredBlockNumber: event.RecoveredBlockNumber,
		RecoveredTxHash:      common.HexToHash(event.RecoveredTxHash),
		WithdrawTxHash:       common.HexToHash(event.WithdrawTxHash),
		CancelledTxHash:      common.HexToHash(event.CancelledTxHash),
	}, nil
}

// GetProcessedBlockNumber implements store.BscBlockStore.
func (s *SQLStore) GetProcessedBlockNumber() (*big.Int, error) {
	var chainState ChainState
	err := s.db.Model(&ChainState{}).First(&chainState).Where("id = ?", 1).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return common.Big0, nil
	}
	if err != nil {
		return nil, err
	}
	number, _ := new(big.Int).SetString(chainState.ProcessedNumber, 10)
	return number, nil
}

// SaveProcessedBlockNumber implements store.BscBlockStore.
func (s *SQLStore) SaveProcessedBlockNumber(number *big.Int) error {
	return s.db.Model(&ChainState{}).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},                          // key colume
		DoUpdates: clause.AssignmentColumns([]string{"processed_number"}), // column needed to be updated
	}).Create(&ChainState{
		Model:           gorm.Model{ID: 1},
		ProcessedNumber: number.String(),
	}).Error
}

// Close implements store.Store.
func (s *SQLStore) Close() error {
	db, err := s.db.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
