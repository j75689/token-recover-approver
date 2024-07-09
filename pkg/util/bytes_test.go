package util

import (
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/cosmos/cosmos-sdk/types"
)

func TestDecodeBytesToSymbol(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test DecodeBytesToSymbol",
			args: args{
				data: MustDecodeHexToBytes("0x424e420000000000000000000000000000000000000000000000000000000000"),
			},
			want: "BNB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeBytesToSymbol(tt.args.data); got != tt.want {
				t.Errorf("DecodeBytesToSymbol() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeAccAccount(t *testing.T) {
	sdkConfig := types.GetConfig()
	sdkConfig.SetBech32PrefixForAccount("tbnb", "bnbp")
	hexData := "0xa10a940119f06b56fef35d77c13414ddfd0093b9"
	data := MustDecodeHexToBytes(hexData)
	expected := "tbnb15y9fgqge7p44dlhnt4muzdq5mh7spyaejjvajm"

	accAccount := types.AccAddress(data)
	fmt.Println(accAccount.String())
	if accAccount.String() != expected {
		t.Fail()
	}
}

type MyFloat float64

func (f MyFloat) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%0.08f", f)), nil
}

func Test_T(t *testing.T) {
	amount := big.NewInt(10)

	b, _ := amount.Float64()
	j, _ := json.Marshal(MyFloat(b))
	fmt.Println(string(j))
}
