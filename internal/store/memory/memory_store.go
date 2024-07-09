package memory

import (
	"math/big"

	"github.com/cosmos/cosmos-sdk/types"

	"github.com/bnb-chain/token-recover-app/internal/store"
	"github.com/bnb-chain/token-recover-app/pkg/util"
)

var _ store.ProofStore = (*MemoryStore)(nil)
var _ store.BscBlockStore = (*MemoryStore)(nil)
var _ store.TokenRecoverEventStore = (*MemoryStore)(nil)
var _ store.GeneralStore = (*MemoryStore)(nil)

func NewMemoryStore(proofsPath string) (*MemoryStore, error) {
	errChan := make(chan error, 1)
	defer close(errChan)
	// load proofs
	stream := util.NewJSONStream(func() any {
		return &Proof{}
	})

	proofs := make(map[string]*Proof)
	proofsByOwner := make(map[string][]*Proof)
	go func() {
		for data := range stream.Watch() {
			if data.Error != nil {
				errChan <- data.Error
				return
			}
			proof := data.Data.(*Proof)
			index := proof.Address.String() + ":" + proof.Coin.Denom
			proofs[index] = proof
			proofsByOwner[proof.Address.String()] = append(proofsByOwner[proof.Address.String()], proof)
		}
		errChan <- nil
	}()
	stream.Start(proofsPath)
	err := <-errChan
	if err != nil {
		return nil, err
	}

	return &MemoryStore{
		proofs:        proofs,
		proofsByOwner: proofsByOwner,
	}, nil
}

// MemoryStore implements store.Store.
type MemoryStore struct {
	proofs        map[string]*Proof   // address:index:symbol -> proofs
	proofsByOwner map[string][]*Proof // address -> proofs
}

// BscBlockStore implements store.GeneralStore.
func (ss *MemoryStore) BscBlockStore() store.BscBlockStore {
	panic("unsupported")
}

// ProofStore implements store.GeneralStore.
func (ss *MemoryStore) ProofStore() store.ProofStore {
	return ss
}

// TokenRecoverEventStore implements store.GeneralStore.
func (ss *MemoryStore) TokenRecoverEventStore() store.TokenRecoverEventStore {
	panic("unsupported")
}

// BatchSaveTokenRecoverEvent implements store.TokenRecoverEventStore.
func (ss *MemoryStore) BatchSaveTokenRecoverEvent(events []*store.TokenRecoverEvent) error {
	panic("unsupported")
}

// GetTokenRecoverEvents implements store.TokenRecoverEventStore.
func (ss *MemoryStore) GetTokenRecoverEvents(condition store.TokenRecoverEvent, pagination store.Pagination) ([]*store.TokenRecoverEvent, int64, error) {
	panic("unsupported")
}

// GetTokenRecoverEvent implements store.TokenRecoverEventStore.
func (ss *MemoryStore) GetTokenRecoverEvent(condition store.TokenRecoverEvent) (*store.TokenRecoverEvent, error) {
	panic("unsupported")
}

// GetProcessedBlockNumber implements store.BscBlockStore.
func (ss *MemoryStore) GetProcessedBlockNumber() (*big.Int, error) {
	panic("unsupported")
}

// SaveProcessedBlockNumber implements store.BscBlockStore.
func (ss *MemoryStore) SaveProcessedBlockNumber(number *big.Int) error {
	panic("unsupported")
}

// GetAccountAssetProofs implements store.ProofStore.
func (ss *MemoryStore) GetAccountAssetProofs(address types.AccAddress, pagination store.Pagination) ([]*store.Proof, int64, error) {
	data, exist := ss.proofsByOwner[address.String()]
	if !exist {
		return nil, 0, nil
	}
	proofs := make([]*store.Proof, 0, len(data))
	for i, proof := range data {
		if i < pagination.Offset {
			continue
		}
		proofs = append(proofs, &store.Proof{
			Address: proof.Address,
			Denom:   proof.Coin.Denom,
			Amount:  proof.Coin.Amount,
			Proof:   proof.Proof,
		})
		if i >= pagination.Offset+pagination.Limit {
			break
		}
	}
	return proofs, int64(len(data)), nil
}

// GetAccountProofs implements store.Store.
func (ss *MemoryStore) GetAccountAssetProof(address types.AccAddress, symbol string) (*store.Proof, error) {
	index := address.String() + ":" + symbol
	proofs, exist := ss.proofs[index]
	if !exist {
		return nil, ErrProofNotFound
	}
	return &store.Proof{
		Address: proofs.Address,
		Denom:   proofs.Coin.Denom,
		Amount:  proofs.Coin.Amount,
		Proof:   proofs.Proof,
	}, nil
}

// CountAccountAssetProofs implements store.Store.
func (ss *MemoryStore) CountAccountAssetProofs() (count int64, err error) {
	return int64(len(ss.proofs)), nil
}

// Close implements store.Store.
func (ss *MemoryStore) Close() error {
	return nil
}
