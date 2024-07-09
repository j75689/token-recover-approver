package util

import (
	"bytes"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func EncodeBytesArrayToHex(data [][]byte) []string {
	hex := make([]string, 0, len(data))
	for _, v := range data {
		hex = append(hex, hexutil.Encode(v))
	}
	return hex
}

func EncodeBytesToHex(data []byte) string {
	return hexutil.Encode(data)
}

func MustDecodeHexToBytes(hex string) []byte {
	data, _ := hexutil.Decode(hex)
	return data
}

func MustDecodeHexArrayToBytes(hexArray []string) [][]byte {
	data := make([][]byte, 0, len(hexArray))
	for _, v := range hexArray {
		data = append(data, MustDecodeHexToBytes(v))
	}
	return data
}

func DecodeBytesToSymbol(data []byte) string {
	return string(bytes.Trim(data, "\x00"))
}

func MustDecodeStringToBigInt(data string) *big.Int {
	result, _ := new(big.Int).SetString(string(data), 10)
	return result
}
