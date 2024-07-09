package tokenrecoverportal

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type TokenRecoverRequestedEvent struct {
	OwnerAddress common.Address
	TokenSymbol  [32]byte
	Account      common.Address
	Amount       *big.Int
}
