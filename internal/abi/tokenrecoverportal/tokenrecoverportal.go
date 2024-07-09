// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokenrecoverportal

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// TokenrecoverportalMetaData contains all meta data concerning the Tokenrecoverportal contract.
var TokenrecoverportalMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"BC_FUSION_CHANNELID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SOURCE_CHAIN_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKING_CHANNELID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addToBlackList\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"approvalAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blackList\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelTokenRecover\",\"inputs\":[{\"name\":\"tokenSymbol\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"attacker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isPaused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isRecovered\",\"inputs\":[{\"name\":\"node\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"merkleRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"merkleRootAlreadyInit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"recover\",\"inputs\":[{\"name\":\"tokenSymbol\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ownerPubKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"ownerSignature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"approvalSignature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"merkleProof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeFromBlackList\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resume\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateParam\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BlackListed\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ParamChange\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProtectorChanged\",\"inputs\":[{\"name\":\"oldProtector\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newProtector\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Resumed\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenRecoverRequested\",\"inputs\":[{\"name\":\"ownerAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"tokenSymbol\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnBlackListed\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyRecovered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ApprovalAddressNotInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InBlackList\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidApprovalSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOwnerPubKeyLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOwnerSignatureLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidValue\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"MerkleRootAlreadyInitiated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MerkleRootNotInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCoinbase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyProtector\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlySystemContract\",\"inputs\":[{\"name\":\"systemContract\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OnlyZeroGasPrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenRecoverPortalPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnknownParam\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]",
}

// TokenrecoverportalABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenrecoverportalMetaData.ABI instead.
var TokenrecoverportalABI = TokenrecoverportalMetaData.ABI

// Tokenrecoverportal is an auto generated Go binding around an Ethereum contract.
type Tokenrecoverportal struct {
	TokenrecoverportalCaller     // Read-only binding to the contract
	TokenrecoverportalTransactor // Write-only binding to the contract
	TokenrecoverportalFilterer   // Log filterer for contract events
}

// TokenrecoverportalCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenrecoverportalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenrecoverportalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenrecoverportalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenrecoverportalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenrecoverportalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenrecoverportalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenrecoverportalSession struct {
	Contract     *Tokenrecoverportal // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TokenrecoverportalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenrecoverportalCallerSession struct {
	Contract *TokenrecoverportalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// TokenrecoverportalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenrecoverportalTransactorSession struct {
	Contract     *TokenrecoverportalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// TokenrecoverportalRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenrecoverportalRaw struct {
	Contract *Tokenrecoverportal // Generic contract binding to access the raw methods on
}

// TokenrecoverportalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenrecoverportalCallerRaw struct {
	Contract *TokenrecoverportalCaller // Generic read-only contract binding to access the raw methods on
}

// TokenrecoverportalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenrecoverportalTransactorRaw struct {
	Contract *TokenrecoverportalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenrecoverportal creates a new instance of Tokenrecoverportal, bound to a specific deployed contract.
func NewTokenrecoverportal(address common.Address, backend bind.ContractBackend) (*Tokenrecoverportal, error) {
	contract, err := bindTokenrecoverportal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tokenrecoverportal{TokenrecoverportalCaller: TokenrecoverportalCaller{contract: contract}, TokenrecoverportalTransactor: TokenrecoverportalTransactor{contract: contract}, TokenrecoverportalFilterer: TokenrecoverportalFilterer{contract: contract}}, nil
}

// NewTokenrecoverportalCaller creates a new read-only instance of Tokenrecoverportal, bound to a specific deployed contract.
func NewTokenrecoverportalCaller(address common.Address, caller bind.ContractCaller) (*TokenrecoverportalCaller, error) {
	contract, err := bindTokenrecoverportal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenrecoverportalCaller{contract: contract}, nil
}

// NewTokenrecoverportalTransactor creates a new write-only instance of Tokenrecoverportal, bound to a specific deployed contract.
func NewTokenrecoverportalTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenrecoverportalTransactor, error) {
	contract, err := bindTokenrecoverportal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenrecoverportalTransactor{contract: contract}, nil
}

// NewTokenrecoverportalFilterer creates a new log filterer instance of Tokenrecoverportal, bound to a specific deployed contract.
func NewTokenrecoverportalFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenrecoverportalFilterer, error) {
	contract, err := bindTokenrecoverportal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenrecoverportalFilterer{contract: contract}, nil
}

// bindTokenrecoverportal binds a generic wrapper to an already deployed contract.
func bindTokenrecoverportal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenrecoverportalMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokenrecoverportal *TokenrecoverportalRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokenrecoverportal.Contract.TokenrecoverportalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokenrecoverportal *TokenrecoverportalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenrecoverportal.Contract.TokenrecoverportalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokenrecoverportal *TokenrecoverportalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokenrecoverportal.Contract.TokenrecoverportalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokenrecoverportal *TokenrecoverportalCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokenrecoverportal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokenrecoverportal *TokenrecoverportalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenrecoverportal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokenrecoverportal *TokenrecoverportalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokenrecoverportal.Contract.contract.Transact(opts, method, params...)
}

// BCFUSIONCHANNELID is a free data retrieval call binding the contract method 0xf1fad104.
//
// Solidity: function BC_FUSION_CHANNELID() view returns(uint8)
func (_Tokenrecoverportal *TokenrecoverportalCaller) BCFUSIONCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokenrecoverportal.contract.Call(opts, &out, "BC_FUSION_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BCFUSIONCHANNELID is a free data retrieval call binding the contract method 0xf1fad104.
//
// Solidity: function BC_FUSION_CHANNELID() view returns(uint8)
func (_Tokenrecoverportal *TokenrecoverportalSession) BCFUSIONCHANNELID() (uint8, error) {
	return _Tokenrecoverportal.Contract.BCFUSIONCHANNELID(&_Tokenrecoverportal.CallOpts)
}

// BCFUSIONCHANNELID is a free data retrieval call binding the contract method 0xf1fad104.
//
// Solidity: function BC_FUSION_CHANNELID() view returns(uint8)
func (_Tokenrecoverportal *TokenrecoverportalCallerSession) BCFUSIONCHANNELID() (uint8, error) {
	return _Tokenrecoverportal.Contract.BCFUSIONCHANNELID(&_Tokenrecoverportal.CallOpts)
}

// SOURCECHAINID is a free data retrieval call binding the contract method 0x74be2150.
//
// Solidity: function SOURCE_CHAIN_ID() view returns(string)
func (_Tokenrecoverportal *TokenrecoverportalCaller) SOURCECHAINID(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Tokenrecoverportal.contract.Call(opts, &out, "SOURCE_CHAIN_ID")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SOURCECHAINID is a free data retrieval call binding the contract method 0x74be2150.
//
// Solidity: function SOURCE_CHAIN_ID() view returns(string)
func (_Tokenrecoverportal *TokenrecoverportalSession) SOURCECHAINID() (string, error) {
	return _Tokenrecoverportal.Contract.SOURCECHAINID(&_Tokenrecoverportal.CallOpts)
}

// SOURCECHAINID is a free data retrieval call binding the contract method 0x74be2150.
//
// Solidity: function SOURCE_CHAIN_ID() view returns(string)
func (_Tokenrecoverportal *TokenrecoverportalCallerSession) SOURCECHAINID() (string, error) {
	return _Tokenrecoverportal.Contract.SOURCECHAINID(&_Tokenrecoverportal.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() view returns(uint8)
func (_Tokenrecoverportal *TokenrecoverportalCaller) STAKINGCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokenrecoverportal.contract.Call(opts, &out, "STAKING_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() view returns(uint8)
func (_Tokenrecoverportal *TokenrecoverportalSession) STAKINGCHANNELID() (uint8, error) {
	return _Tokenrecoverportal.Contract.STAKINGCHANNELID(&_Tokenrecoverportal.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() view returns(uint8)
func (_Tokenrecoverportal *TokenrecoverportalCallerSession) STAKINGCHANNELID() (uint8, error) {
	return _Tokenrecoverportal.Contract.STAKINGCHANNELID(&_Tokenrecoverportal.CallOpts)
}

// ApprovalAddress is a free data retrieval call binding the contract method 0xe842426a.
//
// Solidity: function approvalAddress() view returns(address)
func (_Tokenrecoverportal *TokenrecoverportalCaller) ApprovalAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenrecoverportal.contract.Call(opts, &out, "approvalAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ApprovalAddress is a free data retrieval call binding the contract method 0xe842426a.
//
// Solidity: function approvalAddress() view returns(address)
func (_Tokenrecoverportal *TokenrecoverportalSession) ApprovalAddress() (common.Address, error) {
	return _Tokenrecoverportal.Contract.ApprovalAddress(&_Tokenrecoverportal.CallOpts)
}

// ApprovalAddress is a free data retrieval call binding the contract method 0xe842426a.
//
// Solidity: function approvalAddress() view returns(address)
func (_Tokenrecoverportal *TokenrecoverportalCallerSession) ApprovalAddress() (common.Address, error) {
	return _Tokenrecoverportal.Contract.ApprovalAddress(&_Tokenrecoverportal.CallOpts)
}

// BlackList is a free data retrieval call binding the contract method 0x4838d165.
//
// Solidity: function blackList(address ) view returns(bool)
func (_Tokenrecoverportal *TokenrecoverportalCaller) BlackList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Tokenrecoverportal.contract.Call(opts, &out, "blackList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BlackList is a free data retrieval call binding the contract method 0x4838d165.
//
// Solidity: function blackList(address ) view returns(bool)
func (_Tokenrecoverportal *TokenrecoverportalSession) BlackList(arg0 common.Address) (bool, error) {
	return _Tokenrecoverportal.Contract.BlackList(&_Tokenrecoverportal.CallOpts, arg0)
}

// BlackList is a free data retrieval call binding the contract method 0x4838d165.
//
// Solidity: function blackList(address ) view returns(bool)
func (_Tokenrecoverportal *TokenrecoverportalCallerSession) BlackList(arg0 common.Address) (bool, error) {
	return _Tokenrecoverportal.Contract.BlackList(&_Tokenrecoverportal.CallOpts, arg0)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Tokenrecoverportal *TokenrecoverportalCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Tokenrecoverportal.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Tokenrecoverportal *TokenrecoverportalSession) IsPaused() (bool, error) {
	return _Tokenrecoverportal.Contract.IsPaused(&_Tokenrecoverportal.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Tokenrecoverportal *TokenrecoverportalCallerSession) IsPaused() (bool, error) {
	return _Tokenrecoverportal.Contract.IsPaused(&_Tokenrecoverportal.CallOpts)
}

// IsRecovered is a free data retrieval call binding the contract method 0xe33f8d32.
//
// Solidity: function isRecovered(bytes32 node) view returns(bool)
func (_Tokenrecoverportal *TokenrecoverportalCaller) IsRecovered(opts *bind.CallOpts, node [32]byte) (bool, error) {
	var out []interface{}
	err := _Tokenrecoverportal.contract.Call(opts, &out, "isRecovered", node)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRecovered is a free data retrieval call binding the contract method 0xe33f8d32.
//
// Solidity: function isRecovered(bytes32 node) view returns(bool)
func (_Tokenrecoverportal *TokenrecoverportalSession) IsRecovered(node [32]byte) (bool, error) {
	return _Tokenrecoverportal.Contract.IsRecovered(&_Tokenrecoverportal.CallOpts, node)
}

// IsRecovered is a free data retrieval call binding the contract method 0xe33f8d32.
//
// Solidity: function isRecovered(bytes32 node) view returns(bool)
func (_Tokenrecoverportal *TokenrecoverportalCallerSession) IsRecovered(node [32]byte) (bool, error) {
	return _Tokenrecoverportal.Contract.IsRecovered(&_Tokenrecoverportal.CallOpts, node)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_Tokenrecoverportal *TokenrecoverportalCaller) MerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Tokenrecoverportal.contract.Call(opts, &out, "merkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_Tokenrecoverportal *TokenrecoverportalSession) MerkleRoot() ([32]byte, error) {
	return _Tokenrecoverportal.Contract.MerkleRoot(&_Tokenrecoverportal.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_Tokenrecoverportal *TokenrecoverportalCallerSession) MerkleRoot() ([32]byte, error) {
	return _Tokenrecoverportal.Contract.MerkleRoot(&_Tokenrecoverportal.CallOpts)
}

// MerkleRootAlreadyInit is a free data retrieval call binding the contract method 0x9fcb5012.
//
// Solidity: function merkleRootAlreadyInit() view returns(bool)
func (_Tokenrecoverportal *TokenrecoverportalCaller) MerkleRootAlreadyInit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Tokenrecoverportal.contract.Call(opts, &out, "merkleRootAlreadyInit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MerkleRootAlreadyInit is a free data retrieval call binding the contract method 0x9fcb5012.
//
// Solidity: function merkleRootAlreadyInit() view returns(bool)
func (_Tokenrecoverportal *TokenrecoverportalSession) MerkleRootAlreadyInit() (bool, error) {
	return _Tokenrecoverportal.Contract.MerkleRootAlreadyInit(&_Tokenrecoverportal.CallOpts)
}

// MerkleRootAlreadyInit is a free data retrieval call binding the contract method 0x9fcb5012.
//
// Solidity: function merkleRootAlreadyInit() view returns(bool)
func (_Tokenrecoverportal *TokenrecoverportalCallerSession) MerkleRootAlreadyInit() (bool, error) {
	return _Tokenrecoverportal.Contract.MerkleRootAlreadyInit(&_Tokenrecoverportal.CallOpts)
}

// AddToBlackList is a paid mutator transaction binding the contract method 0x417c73a7.
//
// Solidity: function addToBlackList(address account) returns()
func (_Tokenrecoverportal *TokenrecoverportalTransactor) AddToBlackList(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Tokenrecoverportal.contract.Transact(opts, "addToBlackList", account)
}

// AddToBlackList is a paid mutator transaction binding the contract method 0x417c73a7.
//
// Solidity: function addToBlackList(address account) returns()
func (_Tokenrecoverportal *TokenrecoverportalSession) AddToBlackList(account common.Address) (*types.Transaction, error) {
	return _Tokenrecoverportal.Contract.AddToBlackList(&_Tokenrecoverportal.TransactOpts, account)
}

// AddToBlackList is a paid mutator transaction binding the contract method 0x417c73a7.
//
// Solidity: function addToBlackList(address account) returns()
func (_Tokenrecoverportal *TokenrecoverportalTransactorSession) AddToBlackList(account common.Address) (*types.Transaction, error) {
	return _Tokenrecoverportal.Contract.AddToBlackList(&_Tokenrecoverportal.TransactOpts, account)
}

// CancelTokenRecover is a paid mutator transaction binding the contract method 0x572c9980.
//
// Solidity: function cancelTokenRecover(bytes32 tokenSymbol, address attacker) returns()
func (_Tokenrecoverportal *TokenrecoverportalTransactor) CancelTokenRecover(opts *bind.TransactOpts, tokenSymbol [32]byte, attacker common.Address) (*types.Transaction, error) {
	return _Tokenrecoverportal.contract.Transact(opts, "cancelTokenRecover", tokenSymbol, attacker)
}

// CancelTokenRecover is a paid mutator transaction binding the contract method 0x572c9980.
//
// Solidity: function cancelTokenRecover(bytes32 tokenSymbol, address attacker) returns()
func (_Tokenrecoverportal *TokenrecoverportalSession) CancelTokenRecover(tokenSymbol [32]byte, attacker common.Address) (*types.Transaction, error) {
	return _Tokenrecoverportal.Contract.CancelTokenRecover(&_Tokenrecoverportal.TransactOpts, tokenSymbol, attacker)
}

// CancelTokenRecover is a paid mutator transaction binding the contract method 0x572c9980.
//
// Solidity: function cancelTokenRecover(bytes32 tokenSymbol, address attacker) returns()
func (_Tokenrecoverportal *TokenrecoverportalTransactorSession) CancelTokenRecover(tokenSymbol [32]byte, attacker common.Address) (*types.Transaction, error) {
	return _Tokenrecoverportal.Contract.CancelTokenRecover(&_Tokenrecoverportal.TransactOpts, tokenSymbol, attacker)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Tokenrecoverportal *TokenrecoverportalTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenrecoverportal.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Tokenrecoverportal *TokenrecoverportalSession) Initialize() (*types.Transaction, error) {
	return _Tokenrecoverportal.Contract.Initialize(&_Tokenrecoverportal.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Tokenrecoverportal *TokenrecoverportalTransactorSession) Initialize() (*types.Transaction, error) {
	return _Tokenrecoverportal.Contract.Initialize(&_Tokenrecoverportal.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Tokenrecoverportal *TokenrecoverportalTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenrecoverportal.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Tokenrecoverportal *TokenrecoverportalSession) Pause() (*types.Transaction, error) {
	return _Tokenrecoverportal.Contract.Pause(&_Tokenrecoverportal.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Tokenrecoverportal *TokenrecoverportalTransactorSession) Pause() (*types.Transaction, error) {
	return _Tokenrecoverportal.Contract.Pause(&_Tokenrecoverportal.TransactOpts)
}

// Recover is a paid mutator transaction binding the contract method 0xbfb5a6a1.
//
// Solidity: function recover(bytes32 tokenSymbol, uint256 amount, bytes ownerPubKey, bytes ownerSignature, bytes approvalSignature, bytes32[] merkleProof) returns()
func (_Tokenrecoverportal *TokenrecoverportalTransactor) Recover(opts *bind.TransactOpts, tokenSymbol [32]byte, amount *big.Int, ownerPubKey []byte, ownerSignature []byte, approvalSignature []byte, merkleProof [][32]byte) (*types.Transaction, error) {
	return _Tokenrecoverportal.contract.Transact(opts, "recover", tokenSymbol, amount, ownerPubKey, ownerSignature, approvalSignature, merkleProof)
}

// Recover is a paid mutator transaction binding the contract method 0xbfb5a6a1.
//
// Solidity: function recover(bytes32 tokenSymbol, uint256 amount, bytes ownerPubKey, bytes ownerSignature, bytes approvalSignature, bytes32[] merkleProof) returns()
func (_Tokenrecoverportal *TokenrecoverportalSession) Recover(tokenSymbol [32]byte, amount *big.Int, ownerPubKey []byte, ownerSignature []byte, approvalSignature []byte, merkleProof [][32]byte) (*types.Transaction, error) {
	return _Tokenrecoverportal.Contract.Recover(&_Tokenrecoverportal.TransactOpts, tokenSymbol, amount, ownerPubKey, ownerSignature, approvalSignature, merkleProof)
}

// Recover is a paid mutator transaction binding the contract method 0xbfb5a6a1.
//
// Solidity: function recover(bytes32 tokenSymbol, uint256 amount, bytes ownerPubKey, bytes ownerSignature, bytes approvalSignature, bytes32[] merkleProof) returns()
func (_Tokenrecoverportal *TokenrecoverportalTransactorSession) Recover(tokenSymbol [32]byte, amount *big.Int, ownerPubKey []byte, ownerSignature []byte, approvalSignature []byte, merkleProof [][32]byte) (*types.Transaction, error) {
	return _Tokenrecoverportal.Contract.Recover(&_Tokenrecoverportal.TransactOpts, tokenSymbol, amount, ownerPubKey, ownerSignature, approvalSignature, merkleProof)
}

// RemoveFromBlackList is a paid mutator transaction binding the contract method 0x4a49ac4c.
//
// Solidity: function removeFromBlackList(address account) returns()
func (_Tokenrecoverportal *TokenrecoverportalTransactor) RemoveFromBlackList(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Tokenrecoverportal.contract.Transact(opts, "removeFromBlackList", account)
}

// RemoveFromBlackList is a paid mutator transaction binding the contract method 0x4a49ac4c.
//
// Solidity: function removeFromBlackList(address account) returns()
func (_Tokenrecoverportal *TokenrecoverportalSession) RemoveFromBlackList(account common.Address) (*types.Transaction, error) {
	return _Tokenrecoverportal.Contract.RemoveFromBlackList(&_Tokenrecoverportal.TransactOpts, account)
}

// RemoveFromBlackList is a paid mutator transaction binding the contract method 0x4a49ac4c.
//
// Solidity: function removeFromBlackList(address account) returns()
func (_Tokenrecoverportal *TokenrecoverportalTransactorSession) RemoveFromBlackList(account common.Address) (*types.Transaction, error) {
	return _Tokenrecoverportal.Contract.RemoveFromBlackList(&_Tokenrecoverportal.TransactOpts, account)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Tokenrecoverportal *TokenrecoverportalTransactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenrecoverportal.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Tokenrecoverportal *TokenrecoverportalSession) Resume() (*types.Transaction, error) {
	return _Tokenrecoverportal.Contract.Resume(&_Tokenrecoverportal.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Tokenrecoverportal *TokenrecoverportalTransactorSession) Resume() (*types.Transaction, error) {
	return _Tokenrecoverportal.Contract.Resume(&_Tokenrecoverportal.TransactOpts)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Tokenrecoverportal *TokenrecoverportalTransactor) UpdateParam(opts *bind.TransactOpts, key string, value []byte) (*types.Transaction, error) {
	return _Tokenrecoverportal.contract.Transact(opts, "updateParam", key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Tokenrecoverportal *TokenrecoverportalSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _Tokenrecoverportal.Contract.UpdateParam(&_Tokenrecoverportal.TransactOpts, key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Tokenrecoverportal *TokenrecoverportalTransactorSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _Tokenrecoverportal.Contract.UpdateParam(&_Tokenrecoverportal.TransactOpts, key, value)
}

// TokenrecoverportalBlackListedIterator is returned from FilterBlackListed and is used to iterate over the raw logs and unpacked data for BlackListed events raised by the Tokenrecoverportal contract.
type TokenrecoverportalBlackListedIterator struct {
	Event *TokenrecoverportalBlackListed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenrecoverportalBlackListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenrecoverportalBlackListed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenrecoverportalBlackListed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenrecoverportalBlackListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenrecoverportalBlackListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenrecoverportalBlackListed represents a BlackListed event raised by the Tokenrecoverportal contract.
type TokenrecoverportalBlackListed struct {
	Target common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBlackListed is a free log retrieval operation binding the contract event 0x7fd26be6fc92aff63f1f4409b2b2ddeb272a888031d7f55ec830485ec6194186.
//
// Solidity: event BlackListed(address indexed target)
func (_Tokenrecoverportal *TokenrecoverportalFilterer) FilterBlackListed(opts *bind.FilterOpts, target []common.Address) (*TokenrecoverportalBlackListedIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Tokenrecoverportal.contract.FilterLogs(opts, "BlackListed", targetRule)
	if err != nil {
		return nil, err
	}
	return &TokenrecoverportalBlackListedIterator{contract: _Tokenrecoverportal.contract, event: "BlackListed", logs: logs, sub: sub}, nil
}

// WatchBlackListed is a free log subscription operation binding the contract event 0x7fd26be6fc92aff63f1f4409b2b2ddeb272a888031d7f55ec830485ec6194186.
//
// Solidity: event BlackListed(address indexed target)
func (_Tokenrecoverportal *TokenrecoverportalFilterer) WatchBlackListed(opts *bind.WatchOpts, sink chan<- *TokenrecoverportalBlackListed, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Tokenrecoverportal.contract.WatchLogs(opts, "BlackListed", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenrecoverportalBlackListed)
				if err := _Tokenrecoverportal.contract.UnpackLog(event, "BlackListed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBlackListed is a log parse operation binding the contract event 0x7fd26be6fc92aff63f1f4409b2b2ddeb272a888031d7f55ec830485ec6194186.
//
// Solidity: event BlackListed(address indexed target)
func (_Tokenrecoverportal *TokenrecoverportalFilterer) ParseBlackListed(log types.Log) (*TokenrecoverportalBlackListed, error) {
	event := new(TokenrecoverportalBlackListed)
	if err := _Tokenrecoverportal.contract.UnpackLog(event, "BlackListed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenrecoverportalInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Tokenrecoverportal contract.
type TokenrecoverportalInitializedIterator struct {
	Event *TokenrecoverportalInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenrecoverportalInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenrecoverportalInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenrecoverportalInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenrecoverportalInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenrecoverportalInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenrecoverportalInitialized represents a Initialized event raised by the Tokenrecoverportal contract.
type TokenrecoverportalInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Tokenrecoverportal *TokenrecoverportalFilterer) FilterInitialized(opts *bind.FilterOpts) (*TokenrecoverportalInitializedIterator, error) {

	logs, sub, err := _Tokenrecoverportal.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TokenrecoverportalInitializedIterator{contract: _Tokenrecoverportal.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Tokenrecoverportal *TokenrecoverportalFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TokenrecoverportalInitialized) (event.Subscription, error) {

	logs, sub, err := _Tokenrecoverportal.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenrecoverportalInitialized)
				if err := _Tokenrecoverportal.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Tokenrecoverportal *TokenrecoverportalFilterer) ParseInitialized(log types.Log) (*TokenrecoverportalInitialized, error) {
	event := new(TokenrecoverportalInitialized)
	if err := _Tokenrecoverportal.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenrecoverportalParamChangeIterator is returned from FilterParamChange and is used to iterate over the raw logs and unpacked data for ParamChange events raised by the Tokenrecoverportal contract.
type TokenrecoverportalParamChangeIterator struct {
	Event *TokenrecoverportalParamChange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenrecoverportalParamChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenrecoverportalParamChange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenrecoverportalParamChange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenrecoverportalParamChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenrecoverportalParamChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenrecoverportalParamChange represents a ParamChange event raised by the Tokenrecoverportal contract.
type TokenrecoverportalParamChange struct {
	Key   string
	Value []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterParamChange is a free log retrieval operation binding the contract event 0xf1ce9b2cbf50eeb05769a29e2543fd350cab46894a7dd9978a12d534bb20e633.
//
// Solidity: event ParamChange(string key, bytes value)
func (_Tokenrecoverportal *TokenrecoverportalFilterer) FilterParamChange(opts *bind.FilterOpts) (*TokenrecoverportalParamChangeIterator, error) {

	logs, sub, err := _Tokenrecoverportal.contract.FilterLogs(opts, "ParamChange")
	if err != nil {
		return nil, err
	}
	return &TokenrecoverportalParamChangeIterator{contract: _Tokenrecoverportal.contract, event: "ParamChange", logs: logs, sub: sub}, nil
}

// WatchParamChange is a free log subscription operation binding the contract event 0xf1ce9b2cbf50eeb05769a29e2543fd350cab46894a7dd9978a12d534bb20e633.
//
// Solidity: event ParamChange(string key, bytes value)
func (_Tokenrecoverportal *TokenrecoverportalFilterer) WatchParamChange(opts *bind.WatchOpts, sink chan<- *TokenrecoverportalParamChange) (event.Subscription, error) {

	logs, sub, err := _Tokenrecoverportal.contract.WatchLogs(opts, "ParamChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenrecoverportalParamChange)
				if err := _Tokenrecoverportal.contract.UnpackLog(event, "ParamChange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseParamChange is a log parse operation binding the contract event 0xf1ce9b2cbf50eeb05769a29e2543fd350cab46894a7dd9978a12d534bb20e633.
//
// Solidity: event ParamChange(string key, bytes value)
func (_Tokenrecoverportal *TokenrecoverportalFilterer) ParseParamChange(log types.Log) (*TokenrecoverportalParamChange, error) {
	event := new(TokenrecoverportalParamChange)
	if err := _Tokenrecoverportal.contract.UnpackLog(event, "ParamChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenrecoverportalPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Tokenrecoverportal contract.
type TokenrecoverportalPausedIterator struct {
	Event *TokenrecoverportalPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenrecoverportalPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenrecoverportalPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenrecoverportalPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenrecoverportalPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenrecoverportalPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenrecoverportalPaused represents a Paused event raised by the Tokenrecoverportal contract.
type TokenrecoverportalPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_Tokenrecoverportal *TokenrecoverportalFilterer) FilterPaused(opts *bind.FilterOpts) (*TokenrecoverportalPausedIterator, error) {

	logs, sub, err := _Tokenrecoverportal.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &TokenrecoverportalPausedIterator{contract: _Tokenrecoverportal.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_Tokenrecoverportal *TokenrecoverportalFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *TokenrecoverportalPaused) (event.Subscription, error) {

	logs, sub, err := _Tokenrecoverportal.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenrecoverportalPaused)
				if err := _Tokenrecoverportal.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_Tokenrecoverportal *TokenrecoverportalFilterer) ParsePaused(log types.Log) (*TokenrecoverportalPaused, error) {
	event := new(TokenrecoverportalPaused)
	if err := _Tokenrecoverportal.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenrecoverportalProtectorChangedIterator is returned from FilterProtectorChanged and is used to iterate over the raw logs and unpacked data for ProtectorChanged events raised by the Tokenrecoverportal contract.
type TokenrecoverportalProtectorChangedIterator struct {
	Event *TokenrecoverportalProtectorChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenrecoverportalProtectorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenrecoverportalProtectorChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenrecoverportalProtectorChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenrecoverportalProtectorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenrecoverportalProtectorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenrecoverportalProtectorChanged represents a ProtectorChanged event raised by the Tokenrecoverportal contract.
type TokenrecoverportalProtectorChanged struct {
	OldProtector common.Address
	NewProtector common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProtectorChanged is a free log retrieval operation binding the contract event 0x44fc1b38a4abaa91ebd1b628a5b259a698f86238c8217d68f516e87769c60c0b.
//
// Solidity: event ProtectorChanged(address indexed oldProtector, address indexed newProtector)
func (_Tokenrecoverportal *TokenrecoverportalFilterer) FilterProtectorChanged(opts *bind.FilterOpts, oldProtector []common.Address, newProtector []common.Address) (*TokenrecoverportalProtectorChangedIterator, error) {

	var oldProtectorRule []interface{}
	for _, oldProtectorItem := range oldProtector {
		oldProtectorRule = append(oldProtectorRule, oldProtectorItem)
	}
	var newProtectorRule []interface{}
	for _, newProtectorItem := range newProtector {
		newProtectorRule = append(newProtectorRule, newProtectorItem)
	}

	logs, sub, err := _Tokenrecoverportal.contract.FilterLogs(opts, "ProtectorChanged", oldProtectorRule, newProtectorRule)
	if err != nil {
		return nil, err
	}
	return &TokenrecoverportalProtectorChangedIterator{contract: _Tokenrecoverportal.contract, event: "ProtectorChanged", logs: logs, sub: sub}, nil
}

// WatchProtectorChanged is a free log subscription operation binding the contract event 0x44fc1b38a4abaa91ebd1b628a5b259a698f86238c8217d68f516e87769c60c0b.
//
// Solidity: event ProtectorChanged(address indexed oldProtector, address indexed newProtector)
func (_Tokenrecoverportal *TokenrecoverportalFilterer) WatchProtectorChanged(opts *bind.WatchOpts, sink chan<- *TokenrecoverportalProtectorChanged, oldProtector []common.Address, newProtector []common.Address) (event.Subscription, error) {

	var oldProtectorRule []interface{}
	for _, oldProtectorItem := range oldProtector {
		oldProtectorRule = append(oldProtectorRule, oldProtectorItem)
	}
	var newProtectorRule []interface{}
	for _, newProtectorItem := range newProtector {
		newProtectorRule = append(newProtectorRule, newProtectorItem)
	}

	logs, sub, err := _Tokenrecoverportal.contract.WatchLogs(opts, "ProtectorChanged", oldProtectorRule, newProtectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenrecoverportalProtectorChanged)
				if err := _Tokenrecoverportal.contract.UnpackLog(event, "ProtectorChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProtectorChanged is a log parse operation binding the contract event 0x44fc1b38a4abaa91ebd1b628a5b259a698f86238c8217d68f516e87769c60c0b.
//
// Solidity: event ProtectorChanged(address indexed oldProtector, address indexed newProtector)
func (_Tokenrecoverportal *TokenrecoverportalFilterer) ParseProtectorChanged(log types.Log) (*TokenrecoverportalProtectorChanged, error) {
	event := new(TokenrecoverportalProtectorChanged)
	if err := _Tokenrecoverportal.contract.UnpackLog(event, "ProtectorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenrecoverportalResumedIterator is returned from FilterResumed and is used to iterate over the raw logs and unpacked data for Resumed events raised by the Tokenrecoverportal contract.
type TokenrecoverportalResumedIterator struct {
	Event *TokenrecoverportalResumed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenrecoverportalResumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenrecoverportalResumed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenrecoverportalResumed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenrecoverportalResumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenrecoverportalResumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenrecoverportalResumed represents a Resumed event raised by the Tokenrecoverportal contract.
type TokenrecoverportalResumed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterResumed is a free log retrieval operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_Tokenrecoverportal *TokenrecoverportalFilterer) FilterResumed(opts *bind.FilterOpts) (*TokenrecoverportalResumedIterator, error) {

	logs, sub, err := _Tokenrecoverportal.contract.FilterLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return &TokenrecoverportalResumedIterator{contract: _Tokenrecoverportal.contract, event: "Resumed", logs: logs, sub: sub}, nil
}

// WatchResumed is a free log subscription operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_Tokenrecoverportal *TokenrecoverportalFilterer) WatchResumed(opts *bind.WatchOpts, sink chan<- *TokenrecoverportalResumed) (event.Subscription, error) {

	logs, sub, err := _Tokenrecoverportal.contract.WatchLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenrecoverportalResumed)
				if err := _Tokenrecoverportal.contract.UnpackLog(event, "Resumed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseResumed is a log parse operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_Tokenrecoverportal *TokenrecoverportalFilterer) ParseResumed(log types.Log) (*TokenrecoverportalResumed, error) {
	event := new(TokenrecoverportalResumed)
	if err := _Tokenrecoverportal.contract.UnpackLog(event, "Resumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenrecoverportalTokenRecoverRequestedIterator is returned from FilterTokenRecoverRequested and is used to iterate over the raw logs and unpacked data for TokenRecoverRequested events raised by the Tokenrecoverportal contract.
type TokenrecoverportalTokenRecoverRequestedIterator struct {
	Event *TokenrecoverportalTokenRecoverRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenrecoverportalTokenRecoverRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenrecoverportalTokenRecoverRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenrecoverportalTokenRecoverRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenrecoverportalTokenRecoverRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenrecoverportalTokenRecoverRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenrecoverportalTokenRecoverRequested represents a TokenRecoverRequested event raised by the Tokenrecoverportal contract.
type TokenrecoverportalTokenRecoverRequested struct {
	OwnerAddress []byte
	TokenSymbol  [32]byte
	Account      common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenRecoverRequested is a free log retrieval operation binding the contract event 0x39cc0b7297a0ef9102d75ebc4919ffec0347d50008c2b865eda4125d5812cb64.
//
// Solidity: event TokenRecoverRequested(bytes ownerAddress, bytes32 tokenSymbol, address account, uint256 amount)
func (_Tokenrecoverportal *TokenrecoverportalFilterer) FilterTokenRecoverRequested(opts *bind.FilterOpts) (*TokenrecoverportalTokenRecoverRequestedIterator, error) {

	logs, sub, err := _Tokenrecoverportal.contract.FilterLogs(opts, "TokenRecoverRequested")
	if err != nil {
		return nil, err
	}
	return &TokenrecoverportalTokenRecoverRequestedIterator{contract: _Tokenrecoverportal.contract, event: "TokenRecoverRequested", logs: logs, sub: sub}, nil
}

// WatchTokenRecoverRequested is a free log subscription operation binding the contract event 0x39cc0b7297a0ef9102d75ebc4919ffec0347d50008c2b865eda4125d5812cb64.
//
// Solidity: event TokenRecoverRequested(bytes ownerAddress, bytes32 tokenSymbol, address account, uint256 amount)
func (_Tokenrecoverportal *TokenrecoverportalFilterer) WatchTokenRecoverRequested(opts *bind.WatchOpts, sink chan<- *TokenrecoverportalTokenRecoverRequested) (event.Subscription, error) {

	logs, sub, err := _Tokenrecoverportal.contract.WatchLogs(opts, "TokenRecoverRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenrecoverportalTokenRecoverRequested)
				if err := _Tokenrecoverportal.contract.UnpackLog(event, "TokenRecoverRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenRecoverRequested is a log parse operation binding the contract event 0x39cc0b7297a0ef9102d75ebc4919ffec0347d50008c2b865eda4125d5812cb64.
//
// Solidity: event TokenRecoverRequested(bytes ownerAddress, bytes32 tokenSymbol, address account, uint256 amount)
func (_Tokenrecoverportal *TokenrecoverportalFilterer) ParseTokenRecoverRequested(log types.Log) (*TokenrecoverportalTokenRecoverRequested, error) {
	event := new(TokenrecoverportalTokenRecoverRequested)
	if err := _Tokenrecoverportal.contract.UnpackLog(event, "TokenRecoverRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenrecoverportalUnBlackListedIterator is returned from FilterUnBlackListed and is used to iterate over the raw logs and unpacked data for UnBlackListed events raised by the Tokenrecoverportal contract.
type TokenrecoverportalUnBlackListedIterator struct {
	Event *TokenrecoverportalUnBlackListed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenrecoverportalUnBlackListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenrecoverportalUnBlackListed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenrecoverportalUnBlackListed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenrecoverportalUnBlackListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenrecoverportalUnBlackListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenrecoverportalUnBlackListed represents a UnBlackListed event raised by the Tokenrecoverportal contract.
type TokenrecoverportalUnBlackListed struct {
	Target common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnBlackListed is a free log retrieval operation binding the contract event 0xe0db3499b7fdc3da4cddff5f45d694549c19835e7f719fb5606d3ad1a5de4011.
//
// Solidity: event UnBlackListed(address indexed target)
func (_Tokenrecoverportal *TokenrecoverportalFilterer) FilterUnBlackListed(opts *bind.FilterOpts, target []common.Address) (*TokenrecoverportalUnBlackListedIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Tokenrecoverportal.contract.FilterLogs(opts, "UnBlackListed", targetRule)
	if err != nil {
		return nil, err
	}
	return &TokenrecoverportalUnBlackListedIterator{contract: _Tokenrecoverportal.contract, event: "UnBlackListed", logs: logs, sub: sub}, nil
}

// WatchUnBlackListed is a free log subscription operation binding the contract event 0xe0db3499b7fdc3da4cddff5f45d694549c19835e7f719fb5606d3ad1a5de4011.
//
// Solidity: event UnBlackListed(address indexed target)
func (_Tokenrecoverportal *TokenrecoverportalFilterer) WatchUnBlackListed(opts *bind.WatchOpts, sink chan<- *TokenrecoverportalUnBlackListed, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Tokenrecoverportal.contract.WatchLogs(opts, "UnBlackListed", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenrecoverportalUnBlackListed)
				if err := _Tokenrecoverportal.contract.UnpackLog(event, "UnBlackListed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnBlackListed is a log parse operation binding the contract event 0xe0db3499b7fdc3da4cddff5f45d694549c19835e7f719fb5606d3ad1a5de4011.
//
// Solidity: event UnBlackListed(address indexed target)
func (_Tokenrecoverportal *TokenrecoverportalFilterer) ParseUnBlackListed(log types.Log) (*TokenrecoverportalUnBlackListed, error) {
	event := new(TokenrecoverportalUnBlackListed)
	if err := _Tokenrecoverportal.contract.UnpackLog(event, "UnBlackListed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
