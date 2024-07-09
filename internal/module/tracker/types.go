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

type MyFloat float64

func (f MyFloat) MarshalJSON() ([]byte, error) {
	return []byte(`"` + fmt.Sprintf("%0.08f", f) + `"`), nil
}

type TokenRecoverEventResponse struct {
	Name            string                   `json:"name"`
	Symbol          string                   `json:"symbol"`
	Amount          *big.Int                 `json:"amount"`
	Status          store.TokenRecoverStatus `json:"status"`
	ContractAddress common.Address           `json:"contract_address"`
}

func (resp *TokenRecoverEventResponse) MarshalJSON() ([]byte, error) {
	type aliasTokenRecoverEventResponse struct {
		Name            string                   `json:"name"`
		Symbol          string                   `json:"symbol"`
		Amount          MyFloat                  `json:"amount"`
		Status          store.TokenRecoverStatus `json:"status"`
		ContractAddress string                   `json:"contract_address"`
	}
	amount, _ := resp.Amount.Float64()
	return json.Marshal(&aliasTokenRecoverEventResponse{
		Name:            resp.Name,
		Symbol:          resp.Symbol,
		Amount:          MyFloat(amount / 1e8),
		Status:          resp.Status,
		ContractAddress: resp.ContractAddress.Hex(),
	})
}
