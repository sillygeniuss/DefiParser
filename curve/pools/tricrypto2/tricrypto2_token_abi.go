// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tricrypto2

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

// Tricrypto2TokenMetaData contains all meta data concerning the Tricrypto2Token contract.
var Tricrypto2TokenMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"Transfer\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\",\"indexed\":true},{\"name\":\"_to\",\"type\":\"address\",\"indexed\":true},{\"name\":\"_value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Approval\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"indexed\":true},{\"name\":\"_spender\",\"type\":\"address\",\"indexed\":true},{\"name\":\"_value\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"SetName\",\"inputs\":[{\"name\":\"old_name\",\"type\":\"string\",\"indexed\":false},{\"name\":\"old_symbol\",\"type\":\"string\",\"indexed\":false},{\"name\":\"name\",\"type\":\"string\",\"indexed\":false},{\"name\":\"symbol\",\"type\":\"string\",\"indexed\":false},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":false},{\"name\":\"time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":288},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":77340},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":115282},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":37821},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"increaseAllowance\",\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_added_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":40365},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"decreaseAllowance\",\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtracted_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":40389},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":79579},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"mint_relative\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"frac\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":79983},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"burnFrom\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":79627},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_minter\",\"inputs\":[{\"name\":\"_minter\",\"type\":\"address\"}],\"outputs\":[],\"gas\":37815},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_name\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"}],\"outputs\":[],\"gas\":226002},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"gas\":13020},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"gas\":10773},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":2993},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"},{\"name\":\"arg1\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3238},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":2838},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"minter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":2868}]",
}

// Tricrypto2TokenABI is the input ABI used to generate the binding from.
// Deprecated: Use Tricrypto2TokenMetaData.ABI instead.
var Tricrypto2TokenABI = Tricrypto2TokenMetaData.ABI

// Tricrypto2Token is an auto generated Go binding around an Ethereum contract.
type Tricrypto2Token struct {
	Tricrypto2TokenCaller     // Read-only binding to the contract
	Tricrypto2TokenTransactor // Write-only binding to the contract
	Tricrypto2TokenFilterer   // Log filterer for contract events
}

// Tricrypto2TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type Tricrypto2TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Tricrypto2TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Tricrypto2TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Tricrypto2TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Tricrypto2TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Tricrypto2TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Tricrypto2TokenSession struct {
	Contract     *Tricrypto2Token  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Tricrypto2TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Tricrypto2TokenCallerSession struct {
	Contract *Tricrypto2TokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// Tricrypto2TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Tricrypto2TokenTransactorSession struct {
	Contract     *Tricrypto2TokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// Tricrypto2TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type Tricrypto2TokenRaw struct {
	Contract *Tricrypto2Token // Generic contract binding to access the raw methods on
}

// Tricrypto2TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Tricrypto2TokenCallerRaw struct {
	Contract *Tricrypto2TokenCaller // Generic read-only contract binding to access the raw methods on
}

// Tricrypto2TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Tricrypto2TokenTransactorRaw struct {
	Contract *Tricrypto2TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTricrypto2Token creates a new instance of Tricrypto2Token, bound to a specific deployed contract.
func NewTricrypto2Token(address common.Address, backend bind.ContractBackend) (*Tricrypto2Token, error) {
	contract, err := bindTricrypto2Token(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tricrypto2Token{Tricrypto2TokenCaller: Tricrypto2TokenCaller{contract: contract}, Tricrypto2TokenTransactor: Tricrypto2TokenTransactor{contract: contract}, Tricrypto2TokenFilterer: Tricrypto2TokenFilterer{contract: contract}}, nil
}

// NewTricrypto2TokenCaller creates a new read-only instance of Tricrypto2Token, bound to a specific deployed contract.
func NewTricrypto2TokenCaller(address common.Address, caller bind.ContractCaller) (*Tricrypto2TokenCaller, error) {
	contract, err := bindTricrypto2Token(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Tricrypto2TokenCaller{contract: contract}, nil
}

// NewTricrypto2TokenTransactor creates a new write-only instance of Tricrypto2Token, bound to a specific deployed contract.
func NewTricrypto2TokenTransactor(address common.Address, transactor bind.ContractTransactor) (*Tricrypto2TokenTransactor, error) {
	contract, err := bindTricrypto2Token(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Tricrypto2TokenTransactor{contract: contract}, nil
}

// NewTricrypto2TokenFilterer creates a new log filterer instance of Tricrypto2Token, bound to a specific deployed contract.
func NewTricrypto2TokenFilterer(address common.Address, filterer bind.ContractFilterer) (*Tricrypto2TokenFilterer, error) {
	contract, err := bindTricrypto2Token(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Tricrypto2TokenFilterer{contract: contract}, nil
}

// bindTricrypto2Token binds a generic wrapper to an already deployed contract.
func bindTricrypto2Token(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Tricrypto2TokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tricrypto2Token *Tricrypto2TokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tricrypto2Token.Contract.Tricrypto2TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tricrypto2Token *Tricrypto2TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tricrypto2Token.Contract.Tricrypto2TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tricrypto2Token *Tricrypto2TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tricrypto2Token.Contract.Tricrypto2TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tricrypto2Token *Tricrypto2TokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tricrypto2Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tricrypto2Token *Tricrypto2TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tricrypto2Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tricrypto2Token *Tricrypto2TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tricrypto2Token.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_Tricrypto2Token *Tricrypto2TokenCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2Token.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_Tricrypto2Token *Tricrypto2TokenSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Tricrypto2Token.Contract.Allowance(&_Tricrypto2Token.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_Tricrypto2Token *Tricrypto2TokenCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Tricrypto2Token.Contract.Allowance(&_Tricrypto2Token.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_Tricrypto2Token *Tricrypto2TokenCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2Token.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_Tricrypto2Token *Tricrypto2TokenSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Tricrypto2Token.Contract.BalanceOf(&_Tricrypto2Token.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_Tricrypto2Token *Tricrypto2TokenCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Tricrypto2Token.Contract.BalanceOf(&_Tricrypto2Token.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Tricrypto2Token *Tricrypto2TokenCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2Token.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Tricrypto2Token *Tricrypto2TokenSession) Decimals() (*big.Int, error) {
	return _Tricrypto2Token.Contract.Decimals(&_Tricrypto2Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Tricrypto2Token *Tricrypto2TokenCallerSession) Decimals() (*big.Int, error) {
	return _Tricrypto2Token.Contract.Decimals(&_Tricrypto2Token.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Tricrypto2Token *Tricrypto2TokenCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tricrypto2Token.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Tricrypto2Token *Tricrypto2TokenSession) Minter() (common.Address, error) {
	return _Tricrypto2Token.Contract.Minter(&_Tricrypto2Token.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_Tricrypto2Token *Tricrypto2TokenCallerSession) Minter() (common.Address, error) {
	return _Tricrypto2Token.Contract.Minter(&_Tricrypto2Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Tricrypto2Token *Tricrypto2TokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Tricrypto2Token.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Tricrypto2Token *Tricrypto2TokenSession) Name() (string, error) {
	return _Tricrypto2Token.Contract.Name(&_Tricrypto2Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Tricrypto2Token *Tricrypto2TokenCallerSession) Name() (string, error) {
	return _Tricrypto2Token.Contract.Name(&_Tricrypto2Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Tricrypto2Token *Tricrypto2TokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Tricrypto2Token.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Tricrypto2Token *Tricrypto2TokenSession) Symbol() (string, error) {
	return _Tricrypto2Token.Contract.Symbol(&_Tricrypto2Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Tricrypto2Token *Tricrypto2TokenCallerSession) Symbol() (string, error) {
	return _Tricrypto2Token.Contract.Symbol(&_Tricrypto2Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Tricrypto2Token *Tricrypto2TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2Token.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Tricrypto2Token *Tricrypto2TokenSession) TotalSupply() (*big.Int, error) {
	return _Tricrypto2Token.Contract.TotalSupply(&_Tricrypto2Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Tricrypto2Token *Tricrypto2TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Tricrypto2Token.Contract.TotalSupply(&_Tricrypto2Token.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Tricrypto2Token *Tricrypto2TokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Tricrypto2Token.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Tricrypto2Token *Tricrypto2TokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Tricrypto2Token.Contract.Approve(&_Tricrypto2Token.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_Tricrypto2Token *Tricrypto2TokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Tricrypto2Token.Contract.Approve(&_Tricrypto2Token.TransactOpts, _spender, _value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address _to, uint256 _value) returns(bool)
func (_Tricrypto2Token *Tricrypto2TokenTransactor) BurnFrom(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Tricrypto2Token.contract.Transact(opts, "burnFrom", _to, _value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address _to, uint256 _value) returns(bool)
func (_Tricrypto2Token *Tricrypto2TokenSession) BurnFrom(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Tricrypto2Token.Contract.BurnFrom(&_Tricrypto2Token.TransactOpts, _to, _value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address _to, uint256 _value) returns(bool)
func (_Tricrypto2Token *Tricrypto2TokenTransactorSession) BurnFrom(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Tricrypto2Token.Contract.BurnFrom(&_Tricrypto2Token.TransactOpts, _to, _value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtracted_value) returns(bool)
func (_Tricrypto2Token *Tricrypto2TokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _subtracted_value *big.Int) (*types.Transaction, error) {
	return _Tricrypto2Token.contract.Transact(opts, "decreaseAllowance", _spender, _subtracted_value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtracted_value) returns(bool)
func (_Tricrypto2Token *Tricrypto2TokenSession) DecreaseAllowance(_spender common.Address, _subtracted_value *big.Int) (*types.Transaction, error) {
	return _Tricrypto2Token.Contract.DecreaseAllowance(&_Tricrypto2Token.TransactOpts, _spender, _subtracted_value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtracted_value) returns(bool)
func (_Tricrypto2Token *Tricrypto2TokenTransactorSession) DecreaseAllowance(_spender common.Address, _subtracted_value *big.Int) (*types.Transaction, error) {
	return _Tricrypto2Token.Contract.DecreaseAllowance(&_Tricrypto2Token.TransactOpts, _spender, _subtracted_value)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _added_value) returns(bool)
func (_Tricrypto2Token *Tricrypto2TokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _added_value *big.Int) (*types.Transaction, error) {
	return _Tricrypto2Token.contract.Transact(opts, "increaseAllowance", _spender, _added_value)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _added_value) returns(bool)
func (_Tricrypto2Token *Tricrypto2TokenSession) IncreaseAllowance(_spender common.Address, _added_value *big.Int) (*types.Transaction, error) {
	return _Tricrypto2Token.Contract.IncreaseAllowance(&_Tricrypto2Token.TransactOpts, _spender, _added_value)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _added_value) returns(bool)
func (_Tricrypto2Token *Tricrypto2TokenTransactorSession) IncreaseAllowance(_spender common.Address, _added_value *big.Int) (*types.Transaction, error) {
	return _Tricrypto2Token.Contract.IncreaseAllowance(&_Tricrypto2Token.TransactOpts, _spender, _added_value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _value) returns(bool)
func (_Tricrypto2Token *Tricrypto2TokenTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Tricrypto2Token.contract.Transact(opts, "mint", _to, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _value) returns(bool)
func (_Tricrypto2Token *Tricrypto2TokenSession) Mint(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Tricrypto2Token.Contract.Mint(&_Tricrypto2Token.TransactOpts, _to, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _value) returns(bool)
func (_Tricrypto2Token *Tricrypto2TokenTransactorSession) Mint(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Tricrypto2Token.Contract.Mint(&_Tricrypto2Token.TransactOpts, _to, _value)
}

// MintRelative is a paid mutator transaction binding the contract method 0x6962f845.
//
// Solidity: function mint_relative(address _to, uint256 frac) returns(uint256)
func (_Tricrypto2Token *Tricrypto2TokenTransactor) MintRelative(opts *bind.TransactOpts, _to common.Address, frac *big.Int) (*types.Transaction, error) {
	return _Tricrypto2Token.contract.Transact(opts, "mint_relative", _to, frac)
}

// MintRelative is a paid mutator transaction binding the contract method 0x6962f845.
//
// Solidity: function mint_relative(address _to, uint256 frac) returns(uint256)
func (_Tricrypto2Token *Tricrypto2TokenSession) MintRelative(_to common.Address, frac *big.Int) (*types.Transaction, error) {
	return _Tricrypto2Token.Contract.MintRelative(&_Tricrypto2Token.TransactOpts, _to, frac)
}

// MintRelative is a paid mutator transaction binding the contract method 0x6962f845.
//
// Solidity: function mint_relative(address _to, uint256 frac) returns(uint256)
func (_Tricrypto2Token *Tricrypto2TokenTransactorSession) MintRelative(_to common.Address, frac *big.Int) (*types.Transaction, error) {
	return _Tricrypto2Token.Contract.MintRelative(&_Tricrypto2Token.TransactOpts, _to, frac)
}

// SetMinter is a paid mutator transaction binding the contract method 0x1652e9fc.
//
// Solidity: function set_minter(address _minter) returns()
func (_Tricrypto2Token *Tricrypto2TokenTransactor) SetMinter(opts *bind.TransactOpts, _minter common.Address) (*types.Transaction, error) {
	return _Tricrypto2Token.contract.Transact(opts, "set_minter", _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0x1652e9fc.
//
// Solidity: function set_minter(address _minter) returns()
func (_Tricrypto2Token *Tricrypto2TokenSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _Tricrypto2Token.Contract.SetMinter(&_Tricrypto2Token.TransactOpts, _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0x1652e9fc.
//
// Solidity: function set_minter(address _minter) returns()
func (_Tricrypto2Token *Tricrypto2TokenTransactorSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _Tricrypto2Token.Contract.SetMinter(&_Tricrypto2Token.TransactOpts, _minter)
}

// SetName is a paid mutator transaction binding the contract method 0xe1430e06.
//
// Solidity: function set_name(string _name, string _symbol) returns()
func (_Tricrypto2Token *Tricrypto2TokenTransactor) SetName(opts *bind.TransactOpts, _name string, _symbol string) (*types.Transaction, error) {
	return _Tricrypto2Token.contract.Transact(opts, "set_name", _name, _symbol)
}

// SetName is a paid mutator transaction binding the contract method 0xe1430e06.
//
// Solidity: function set_name(string _name, string _symbol) returns()
func (_Tricrypto2Token *Tricrypto2TokenSession) SetName(_name string, _symbol string) (*types.Transaction, error) {
	return _Tricrypto2Token.Contract.SetName(&_Tricrypto2Token.TransactOpts, _name, _symbol)
}

// SetName is a paid mutator transaction binding the contract method 0xe1430e06.
//
// Solidity: function set_name(string _name, string _symbol) returns()
func (_Tricrypto2Token *Tricrypto2TokenTransactorSession) SetName(_name string, _symbol string) (*types.Transaction, error) {
	return _Tricrypto2Token.Contract.SetName(&_Tricrypto2Token.TransactOpts, _name, _symbol)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Tricrypto2Token *Tricrypto2TokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Tricrypto2Token.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Tricrypto2Token *Tricrypto2TokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Tricrypto2Token.Contract.Transfer(&_Tricrypto2Token.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_Tricrypto2Token *Tricrypto2TokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Tricrypto2Token.Contract.Transfer(&_Tricrypto2Token.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Tricrypto2Token *Tricrypto2TokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Tricrypto2Token.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Tricrypto2Token *Tricrypto2TokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Tricrypto2Token.Contract.TransferFrom(&_Tricrypto2Token.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_Tricrypto2Token *Tricrypto2TokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Tricrypto2Token.Contract.TransferFrom(&_Tricrypto2Token.TransactOpts, _from, _to, _value)
}

// Tricrypto2TokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Tricrypto2Token contract.
type Tricrypto2TokenApprovalIterator struct {
	Event *Tricrypto2TokenApproval // Event containing the contract specifics and raw log

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
func (it *Tricrypto2TokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Tricrypto2TokenApproval)
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
		it.Event = new(Tricrypto2TokenApproval)
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
func (it *Tricrypto2TokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Tricrypto2TokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Tricrypto2TokenApproval represents a Approval event raised by the Tricrypto2Token contract.
type Tricrypto2TokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Tricrypto2Token *Tricrypto2TokenFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*Tricrypto2TokenApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Tricrypto2Token.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &Tricrypto2TokenApprovalIterator{contract: _Tricrypto2Token.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Tricrypto2Token *Tricrypto2TokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Tricrypto2TokenApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Tricrypto2Token.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Tricrypto2TokenApproval)
				if err := _Tricrypto2Token.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Tricrypto2Token *Tricrypto2TokenFilterer) ParseApproval(log types.Log) (*Tricrypto2TokenApproval, error) {
	event := new(Tricrypto2TokenApproval)
	if err := _Tricrypto2Token.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Tricrypto2TokenSetNameIterator is returned from FilterSetName and is used to iterate over the raw logs and unpacked data for SetName events raised by the Tricrypto2Token contract.
type Tricrypto2TokenSetNameIterator struct {
	Event *Tricrypto2TokenSetName // Event containing the contract specifics and raw log

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
func (it *Tricrypto2TokenSetNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Tricrypto2TokenSetName)
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
		it.Event = new(Tricrypto2TokenSetName)
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
func (it *Tricrypto2TokenSetNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Tricrypto2TokenSetNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Tricrypto2TokenSetName represents a SetName event raised by the Tricrypto2Token contract.
type Tricrypto2TokenSetName struct {
	OldName   string
	OldSymbol string
	Name      string
	Symbol    string
	Owner     common.Address
	Time      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetName is a free log retrieval operation binding the contract event 0x68ed9e6681c98d0e2744ce6c08d46c045e098a479b120b5b7253fa95e4c48954.
//
// Solidity: event SetName(string old_name, string old_symbol, string name, string symbol, address owner, uint256 time)
func (_Tricrypto2Token *Tricrypto2TokenFilterer) FilterSetName(opts *bind.FilterOpts) (*Tricrypto2TokenSetNameIterator, error) {

	logs, sub, err := _Tricrypto2Token.contract.FilterLogs(opts, "SetName")
	if err != nil {
		return nil, err
	}
	return &Tricrypto2TokenSetNameIterator{contract: _Tricrypto2Token.contract, event: "SetName", logs: logs, sub: sub}, nil
}

// WatchSetName is a free log subscription operation binding the contract event 0x68ed9e6681c98d0e2744ce6c08d46c045e098a479b120b5b7253fa95e4c48954.
//
// Solidity: event SetName(string old_name, string old_symbol, string name, string symbol, address owner, uint256 time)
func (_Tricrypto2Token *Tricrypto2TokenFilterer) WatchSetName(opts *bind.WatchOpts, sink chan<- *Tricrypto2TokenSetName) (event.Subscription, error) {

	logs, sub, err := _Tricrypto2Token.contract.WatchLogs(opts, "SetName")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Tricrypto2TokenSetName)
				if err := _Tricrypto2Token.contract.UnpackLog(event, "SetName", log); err != nil {
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

// ParseSetName is a log parse operation binding the contract event 0x68ed9e6681c98d0e2744ce6c08d46c045e098a479b120b5b7253fa95e4c48954.
//
// Solidity: event SetName(string old_name, string old_symbol, string name, string symbol, address owner, uint256 time)
func (_Tricrypto2Token *Tricrypto2TokenFilterer) ParseSetName(log types.Log) (*Tricrypto2TokenSetName, error) {
	event := new(Tricrypto2TokenSetName)
	if err := _Tricrypto2Token.contract.UnpackLog(event, "SetName", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Tricrypto2TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Tricrypto2Token contract.
type Tricrypto2TokenTransferIterator struct {
	Event *Tricrypto2TokenTransfer // Event containing the contract specifics and raw log

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
func (it *Tricrypto2TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Tricrypto2TokenTransfer)
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
		it.Event = new(Tricrypto2TokenTransfer)
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
func (it *Tricrypto2TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Tricrypto2TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Tricrypto2TokenTransfer represents a Transfer event raised by the Tricrypto2Token contract.
type Tricrypto2TokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Tricrypto2Token *Tricrypto2TokenFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*Tricrypto2TokenTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Tricrypto2Token.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &Tricrypto2TokenTransferIterator{contract: _Tricrypto2Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Tricrypto2Token *Tricrypto2TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Tricrypto2TokenTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Tricrypto2Token.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Tricrypto2TokenTransfer)
				if err := _Tricrypto2Token.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Tricrypto2Token *Tricrypto2TokenFilterer) ParseTransfer(log types.Log) (*Tricrypto2TokenTransfer, error) {
	event := new(Tricrypto2TokenTransfer)
	if err := _Tricrypto2Token.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
