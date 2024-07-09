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
	gorm.Model
	TokenOwner           string `json:"token_owner" gorm:"index"`
	TokenContractAddress string `json:"token_contract_address"`
	Denom                string `json:"denom"`
	Amount               string `json:"amount"`
	ClaimAddress         string `json:"claim_address" gorm:"index"`
	UnlockAt             int64  `json:"unlock_at"`
	Status               int8   `json:"status"`
	WithdrawTxHash       string `json:"withdraw_tx_hash"`
}
