package store

import (
	"errors"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	// ErrInvalidToken is returned when the token is not found in the merkle tree
	ErrInvalidToken   = errors.New("invalid token")
	ErrRecordNotFound = errors.New("record not found")
)

// Proof is a merkle proof of an account
type Proof struct {
	Address sdk.AccAddress `json:"address"`
	Denom   string         `json:"denom"`
	Amount  int64          `json:"amount"`
	Proof   [][]byte       `json:"proof"`
}

// Serialize implements merkle tree data Serialize method.
func (p *Proof) Serialize() ([]byte, error) {
	if p.Amount == 0 {
		return nil, ErrInvalidToken
	}

	var symbol [32]byte
	copy(symbol[:], p.Denom)
	return crypto.Keccak256Hash(
		p.Address.Bytes(),
		symbol[:],
		big.NewInt(p.Amount).FillBytes(make([]byte, 32)),
	).Bytes(), nil
}

type ChainState struct {
	ProcessedNumber *big.Int `json:"processed_number"`
}

type TokenRecoverStatus uint8

const (
	NotBounded  TokenRecoverStatus = 1
	Pending     TokenRecoverStatus = 2
	Requested   TokenRecoverStatus = 3
	Locked      TokenRecoverStatus = 4
	Withdrawing TokenRecoverStatus = 5
	Unlocked    TokenRecoverStatus = 6
	Cancelled   TokenRecoverStatus = 7
)

type TokenRecoverEvent struct {
	TokenOwner           sdk.AccAddress     `json:"token_owner" gorm:"index"`
	TokenContractAddress common.Address     `json:"token_contract_address"`
	Denom                string             `json:"denom"`
	Amount               *big.Int           `json:"amount"`
	ClaimAddress         common.Address     `json:"claim_address" gorm:"index"`
	UnlockAt             int64              `json:"unlock_at"`
	Status               TokenRecoverStatus `json:"status"`
	WithdrawTxHash       common.Hash        `json:"withdraw_tx_hash"`
}
