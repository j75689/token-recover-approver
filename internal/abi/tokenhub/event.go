package tokenhub

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type TokenRecoverLockedEvent struct {
	TokenSymbol [32]byte
	TokenAddr   common.Address
	Recipient   common.Address
	Amount      *big.Int
	UnlockAt    *big.Int
}

type WithdrawUnlockedTokenEvent struct {
	TokenAddr common.Address
	Recipient common.Address
	Amount    *big.Int
}

type CancelTokenRecoverLockEvent struct {
	TokenSymbol [32]byte
	TokenAddr   common.Address
	Attacker    common.Address
	Amount      *big.Int
}
