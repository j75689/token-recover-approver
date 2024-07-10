package tracker

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/bnb-chain/token-recover-app/internal/store"
)

type TokenInfo struct {
	Name             string `json:"name"`
	Symbol           string `json:"symbol"`
	ContractDecimals int    `json:"contract_decimals"`
	ContractAddress  string `json:"contract_address"`
}

type TokenList map[string]TokenInfo

type Decimal big.Int

var decimalFromBNBChain = big.NewInt(1e8)

func (f Decimal) MarshalJSON() ([]byte, error) {
	return []byte(`"` +
		new(big.Int).Div(f.BigInt(), decimalFromBNBChain).String() +
		"." +
		fmt.Sprintf("%08d", new(big.Int).Mod(f.BigInt(), decimalFromBNBChain).Uint64()) +
		`"`), nil
}

func (f Decimal) BigInt() *big.Int {
	return (*big.Int)(&f)
}

type TokenRecoverEventResponse struct {
	Name            string                   `json:"name"`
	Symbol          string                   `json:"symbol"`
	Amount          *big.Int                 `json:"amount"`
	Status          store.TokenRecoverStatus `json:"status"`
	UnlockAt        int64                    `json:"unlock_at"`
	ContractAddress common.Address           `json:"contract_address,omitempty"`
}

func (resp TokenRecoverEventResponse) MarshalJSON() ([]byte, error) {
	type aliasTokenRecoverEventResponse struct {
		Name            string                   `json:"name"`
		Symbol          string                   `json:"symbol"`
		Amount          Decimal                  `json:"amount"`
		Status          store.TokenRecoverStatus `json:"status"`
		UnlockAt        int64                    `json:"unlock_at,omitempty"`
		ContractAddress string                   `json:"contract_address,omitempty"`
	}

	contractAddr := ""
	if resp.ContractAddress != store.EmptyAccount {
		contractAddr = resp.ContractAddress.Hex()
	}
	return json.Marshal(&aliasTokenRecoverEventResponse{
		Name:            resp.Name,
		Symbol:          resp.Symbol,
		Amount:          Decimal(*resp.Amount),
		Status:          resp.Status,
		UnlockAt:        resp.UnlockAt,
		ContractAddress: contractAddr,
	})
}
