package store

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type GeneralStore interface {
	ProofStore() ProofStore
	BscBlockStore() BscBlockStore
	TokenRecoverEventStore() TokenRecoverEventStore
	Close() error
}

type Pagination struct {
	Offset int
	Limit  int
}

type ProofStore interface {
	GetAccountAssetProof(address sdk.AccAddress, symbol string) (proofs *Proof, err error)
	GetAccountAssetProofs(address sdk.AccAddress, pagination Pagination) (proofs []*Proof, count int64, err error)
	CountAccountAssetProofs() (count int64, err error)
	Close() error
}

type BscBlockStore interface {
	GetProcessedBlockNumber() (*big.Int, error)
	SaveProcessedBlockNumber(number *big.Int) error
}

type TokenRecoverEventStore interface {
	GetTokenRecoverEvent(condition TokenRecoverEvent) (*TokenRecoverEvent, error)
	GetTokenRecoverEvents(condition TokenRecoverEvent, pagination Pagination) ([]*TokenRecoverEvent, int64, error)
	BatchSaveTokenRecoverEvent(events []*TokenRecoverEvent) error
}
