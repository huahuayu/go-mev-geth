// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mevtransfer

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

// MevtransferABI is the input ABI used to generate the binding from.
const MevtransferABI = "[{\"inputs\":[],\"name\":\"chi\",\"outputs\":[{\"internalType\":\"contractIChiToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bribe\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// MevtransferBin is the compiled bytecode used for deploying new contracts.
var MevtransferBin = "0x608060405234801561001057600080fd5b50610457806100206000396000f3fe6080604052600436106100295760003560e01c806312514bba1461002e578063c92aecc41461005c575b600080fd5b61005a6004803603602081101561004457600080fd5b81019080803590602001909291905050506100b3565b005b34801561006857600080fd5b5061007161040f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b60005a90504173ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f193505050501580156100fe573d6000803e3d6000fd5b503373ffffffffffffffffffffffffffffffffffffffff166108fc479081150290604051600060405180830381858888f19350505050158015610145573d6000803e3d6000fd5b506000803690506010025a836152080103019050600061a3db61374a83018161016a57fe5b0490506000811180156102425750806d4946c0e9f43f4dee607b0ef1fa1c73ffffffffffffffffffffffffffffffffffffffff166370a08231336040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561020457600080fd5b505afa158015610218573d6000803e3d6000fd5b505050506040513d602081101561022e57600080fd5b810190808051906020019092919050505010155b80156103475750806d4946c0e9f43f4dee607b0ef1fa1c73ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e33306040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060206040518083038186803b15801561030957600080fd5b505afa15801561031d573d6000803e3d6000fd5b505050506040513d602081101561033357600080fd5b810190808051906020019092919050505010155b15610409576d4946c0e9f43f4dee607b0ef1fa1c73ffffffffffffffffffffffffffffffffffffffff1663079d229f3361a3db61374a86018161038657fe5b046040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b1580156103f057600080fd5b505af1158015610404573d6000803e3d6000fd5b505050505b50505050565b6d4946c0e9f43f4dee607b0ef1fa1c8156fea26469706673582212201def7972e4cbab6927863080b1cb0698d1d3d9a2ad624db226174a06df100d0164736f6c63430006060033"

// DeployMevtransfer deploys a new Ethereum contract, binding an instance of Mevtransfer to it.
func DeployMevtransfer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Mevtransfer, error) {
	parsed, err := abi.JSON(strings.NewReader(MevtransferABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MevtransferBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Mevtransfer{MevtransferCaller: MevtransferCaller{contract: contract}, MevtransferTransactor: MevtransferTransactor{contract: contract}, MevtransferFilterer: MevtransferFilterer{contract: contract}}, nil
}

// Mevtransfer is an auto generated Go binding around an Ethereum contract.
type Mevtransfer struct {
	MevtransferCaller     // Read-only binding to the contract
	MevtransferTransactor // Write-only binding to the contract
	MevtransferFilterer   // Log filterer for contract events
}

// MevtransferCaller is an auto generated read-only Go binding around an Ethereum contract.
type MevtransferCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MevtransferTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MevtransferTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MevtransferFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MevtransferFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MevtransferSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MevtransferSession struct {
	Contract     *Mevtransfer      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MevtransferCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MevtransferCallerSession struct {
	Contract *MevtransferCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MevtransferTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MevtransferTransactorSession struct {
	Contract     *MevtransferTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MevtransferRaw is an auto generated low-level Go binding around an Ethereum contract.
type MevtransferRaw struct {
	Contract *Mevtransfer // Generic contract binding to access the raw methods on
}

// MevtransferCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MevtransferCallerRaw struct {
	Contract *MevtransferCaller // Generic read-only contract binding to access the raw methods on
}

// MevtransferTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MevtransferTransactorRaw struct {
	Contract *MevtransferTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMevtransfer creates a new instance of Mevtransfer, bound to a specific deployed contract.
func NewMevtransfer(address common.Address, backend bind.ContractBackend) (*Mevtransfer, error) {
	contract, err := bindMevtransfer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mevtransfer{MevtransferCaller: MevtransferCaller{contract: contract}, MevtransferTransactor: MevtransferTransactor{contract: contract}, MevtransferFilterer: MevtransferFilterer{contract: contract}}, nil
}

// NewMevtransferCaller creates a new read-only instance of Mevtransfer, bound to a specific deployed contract.
func NewMevtransferCaller(address common.Address, caller bind.ContractCaller) (*MevtransferCaller, error) {
	contract, err := bindMevtransfer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MevtransferCaller{contract: contract}, nil
}

// NewMevtransferTransactor creates a new write-only instance of Mevtransfer, bound to a specific deployed contract.
func NewMevtransferTransactor(address common.Address, transactor bind.ContractTransactor) (*MevtransferTransactor, error) {
	contract, err := bindMevtransfer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MevtransferTransactor{contract: contract}, nil
}

// NewMevtransferFilterer creates a new log filterer instance of Mevtransfer, bound to a specific deployed contract.
func NewMevtransferFilterer(address common.Address, filterer bind.ContractFilterer) (*MevtransferFilterer, error) {
	contract, err := bindMevtransfer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MevtransferFilterer{contract: contract}, nil
}

// bindMevtransfer binds a generic wrapper to an already deployed contract.
func bindMevtransfer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MevtransferABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mevtransfer *MevtransferRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mevtransfer.Contract.MevtransferCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mevtransfer *MevtransferRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mevtransfer.Contract.MevtransferTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mevtransfer *MevtransferRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mevtransfer.Contract.MevtransferTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mevtransfer *MevtransferCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mevtransfer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mevtransfer *MevtransferTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mevtransfer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mevtransfer *MevtransferTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mevtransfer.Contract.contract.Transact(opts, method, params...)
}

// Chi is a free data retrieval call binding the contract method 0xc92aecc4.
//
// Solidity: function chi() view returns(address)
func (_Mevtransfer *MevtransferCaller) Chi(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mevtransfer.contract.Call(opts, &out, "chi")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Chi is a free data retrieval call binding the contract method 0xc92aecc4.
//
// Solidity: function chi() view returns(address)
func (_Mevtransfer *MevtransferSession) Chi() (common.Address, error) {
	return _Mevtransfer.Contract.Chi(&_Mevtransfer.CallOpts)
}

// Chi is a free data retrieval call binding the contract method 0xc92aecc4.
//
// Solidity: function chi() view returns(address)
func (_Mevtransfer *MevtransferCallerSession) Chi() (common.Address, error) {
	return _Mevtransfer.Contract.Chi(&_Mevtransfer.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0x12514bba.
//
// Solidity: function transfer(uint256 bribe) payable returns()
func (_Mevtransfer *MevtransferTransactor) Transfer(opts *bind.TransactOpts, bribe *big.Int) (*types.Transaction, error) {
	return _Mevtransfer.contract.Transact(opts, "transfer", bribe)
}

// Transfer is a paid mutator transaction binding the contract method 0x12514bba.
//
// Solidity: function transfer(uint256 bribe) payable returns()
func (_Mevtransfer *MevtransferSession) Transfer(bribe *big.Int) (*types.Transaction, error) {
	return _Mevtransfer.Contract.Transfer(&_Mevtransfer.TransactOpts, bribe)
}

// Transfer is a paid mutator transaction binding the contract method 0x12514bba.
//
// Solidity: function transfer(uint256 bribe) payable returns()
func (_Mevtransfer *MevtransferTransactorSession) Transfer(bribe *big.Int) (*types.Transaction, error) {
	return _Mevtransfer.Contract.Transfer(&_Mevtransfer.TransactOpts, bribe)
}
