// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERCs

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ERCsABI is the input ABI used to generate the binding from.
const ERCsABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"name\":\"magicValue\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ERCs is an auto generated Go binding around an Ethereum contract.
type ERCs struct {
	ERCsCaller     // Read-only binding to the contract
	ERCsTransactor // Write-only binding to the contract
	ERCsFilterer   // Log filterer for contract events
}

// ERCsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERCsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERCsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERCsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERCsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERCsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERCsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERCsSession struct {
	Contract     *ERCs             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERCsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERCsCallerSession struct {
	Contract *ERCsCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERCsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERCsTransactorSession struct {
	Contract     *ERCsTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERCsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERCsRaw struct {
	Contract *ERCs // Generic contract binding to access the raw methods on
}

// ERCsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERCsCallerRaw struct {
	Contract *ERCsCaller // Generic read-only contract binding to access the raw methods on
}

// ERCsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERCsTransactorRaw struct {
	Contract *ERCsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERCs creates a new instance of ERCs, bound to a specific deployed contract.
func NewERCs(address common.Address, backend bind.ContractBackend) (*ERCs, error) {
	contract, err := bindERCs(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERCs{ERCsCaller: ERCsCaller{contract: contract}, ERCsTransactor: ERCsTransactor{contract: contract}, ERCsFilterer: ERCsFilterer{contract: contract}}, nil
}

// NewERCsCaller creates a new read-only instance of ERCs, bound to a specific deployed contract.
func NewERCsCaller(address common.Address, caller bind.ContractCaller) (*ERCsCaller, error) {
	contract, err := bindERCs(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERCsCaller{contract: contract}, nil
}

// NewERCsTransactor creates a new write-only instance of ERCs, bound to a specific deployed contract.
func NewERCsTransactor(address common.Address, transactor bind.ContractTransactor) (*ERCsTransactor, error) {
	contract, err := bindERCs(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERCsTransactor{contract: contract}, nil
}

// NewERCsFilterer creates a new log filterer instance of ERCs, bound to a specific deployed contract.
func NewERCsFilterer(address common.Address, filterer bind.ContractFilterer) (*ERCsFilterer, error) {
	contract, err := bindERCs(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERCsFilterer{contract: contract}, nil
}

// bindERCs binds a generic wrapper to an already deployed contract.
func bindERCs(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERCsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERCs *ERCsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERCs.Contract.ERCsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERCs *ERCsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERCs.Contract.ERCsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERCs *ERCsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERCs.Contract.ERCsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERCs *ERCsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERCs.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERCs *ERCsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERCs.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERCs *ERCsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERCs.Contract.contract.Transact(opts, method, params...)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 hash, bytes _signature) view returns(bytes4 magicValue)
func (_ERCs *ERCsCaller) IsValidSignature(opts *bind.CallOpts, hash [32]byte, _signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _ERCs.contract.Call(opts, &out, "isValidSignature", hash, _signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 hash, bytes _signature) view returns(bytes4 magicValue)
func (_ERCs *ERCsSession) IsValidSignature(hash [32]byte, _signature []byte) ([4]byte, error) {
	return _ERCs.Contract.IsValidSignature(&_ERCs.CallOpts, hash, _signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 hash, bytes _signature) view returns(bytes4 magicValue)
func (_ERCs *ERCsCallerSession) IsValidSignature(hash [32]byte, _signature []byte) ([4]byte, error) {
	return _ERCs.Contract.IsValidSignature(&_ERCs.CallOpts, hash, _signature)
}
