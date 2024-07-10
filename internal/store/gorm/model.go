package gorm

import (
	"gorm.io/gorm"
)

type Proof struct {
	gorm.Model
	Address string `json:"address" gorm:"index,type:varchar(42)"` // hex encoded
	Denom   string `json:"denom" gorm:"index"`
	Proof   string `json:"proof" gorm:"type:text"` // hex encoded
	Amount  int64  `json:"amount"`
}

type ChainState struct {
	gorm.Model
	ProcessedNumber string `json:"processed_number"`
}

type TokenRecoverEvent struct {
	TokenOwner           string `json:"token_owner" gorm:"primaryKey"`
	TokenContractAddress string `json:"token_contract_address"`
	Denom                string `json:"denom" gorm:"primaryKey"`
	Amount               string `json:"amount" gorm:"primaryKey"`
	ClaimAddress         string `json:"claim_address" gorm:"index"`
	UnlockAt             int64  `json:"unlock_at"`
	Status               int8   `json:"status"`
	RecoveredBlockNumber uint64 `json:"recovered_block_number"`
	RecoveredTxHash      string `json:"recovered_tx_hash"`
	WithdrawTxHash       string `json:"withdraw_tx_hash"`
	CancelledTxHash      string `json:"cancelled_tx_hash"`
}
