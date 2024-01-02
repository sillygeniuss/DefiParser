// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package meta

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

// MetaMetaData contains all meta data concerning the Meta contract.
var MetaMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"CommitNewAdmin\",\"inputs\":[{\"name\":\"deadline\",\"type\":\"uint256\",\"indexed\":true},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewAdmin\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"_address_provider\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_registry_handler\",\"inputs\":[{\"name\":\"_registry_handler\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"update_registry_handler\",\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"},{\"name\":\"_registry_handler\",\"type\":\"address\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_registry_handlers_from_pool\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[10]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_base_registry\",\"inputs\":[{\"name\":\"registry_handler\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"find_pool_for_coins\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"find_pool_for_coins\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"find_pools_for_coins\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_admin_balances\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_admin_balances\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_balances\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_balances\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_base_pool\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_base_pool\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_coin_indices\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int128\"},{\"name\":\"\",\"type\":\"int128\"},{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_coin_indices\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int128\"},{\"name\":\"\",\"type\":\"int128\"},{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[8]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[8]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_decimals\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_decimals\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_fees\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[10]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_fees\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[10]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_gauge\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_gauge\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"gauge_idx\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_gauge\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"gauge_idx\",\"type\":\"uint256\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_gauge_type\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int128\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_gauge_type\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"gauge_idx\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int128\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_gauge_type\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"gauge_idx\",\"type\":\"uint256\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int128\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_lp_token\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_lp_token\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_n_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_n_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_n_underlying_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_n_underlying_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_pool_asset_type\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_pool_asset_type\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_pool_from_lp_token\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_pool_from_lp_token\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_pool_params\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[20]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_pool_params\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[20]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_pool_name\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_pool_name\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_underlying_balances\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_underlying_balances\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_underlying_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[8]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_underlying_coins\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[8]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_underlying_decimals\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_underlying_decimals\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[8]\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_virtual_price_from_lp_token\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_virtual_price_from_lp_token\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"is_meta\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"is_meta\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"is_registered\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"is_registered\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"},{\"name\":\"_handler_id\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pool_count\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"pool_list\",\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"address_provider\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_registry\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"registry_length\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}]}]",
}

// MetaABI is the input ABI used to generate the binding from.
// Deprecated: Use MetaMetaData.ABI instead.
var MetaABI = MetaMetaData.ABI

// Meta is an auto generated Go binding around an Ethereum contract.
type Meta struct {
	MetaCaller     // Read-only binding to the contract
	MetaTransactor // Write-only binding to the contract
	MetaFilterer   // Log filterer for contract events
}

// MetaCaller is an auto generated read-only Go binding around an Ethereum contract.
type MetaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MetaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MetaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MetaSession struct {
	Contract     *Meta             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MetaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MetaCallerSession struct {
	Contract *MetaCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MetaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MetaTransactorSession struct {
	Contract     *MetaTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MetaRaw is an auto generated low-level Go binding around an Ethereum contract.
type MetaRaw struct {
	Contract *Meta // Generic contract binding to access the raw methods on
}

// MetaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MetaCallerRaw struct {
	Contract *MetaCaller // Generic read-only contract binding to access the raw methods on
}

// MetaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MetaTransactorRaw struct {
	Contract *MetaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMeta creates a new instance of Meta, bound to a specific deployed contract.
func NewMeta(address common.Address, backend bind.ContractBackend) (*Meta, error) {
	contract, err := bindMeta(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Meta{MetaCaller: MetaCaller{contract: contract}, MetaTransactor: MetaTransactor{contract: contract}, MetaFilterer: MetaFilterer{contract: contract}}, nil
}

// NewMetaCaller creates a new read-only instance of Meta, bound to a specific deployed contract.
func NewMetaCaller(address common.Address, caller bind.ContractCaller) (*MetaCaller, error) {
	contract, err := bindMeta(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MetaCaller{contract: contract}, nil
}

// NewMetaTransactor creates a new write-only instance of Meta, bound to a specific deployed contract.
func NewMetaTransactor(address common.Address, transactor bind.ContractTransactor) (*MetaTransactor, error) {
	contract, err := bindMeta(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MetaTransactor{contract: contract}, nil
}

// NewMetaFilterer creates a new log filterer instance of Meta, bound to a specific deployed contract.
func NewMetaFilterer(address common.Address, filterer bind.ContractFilterer) (*MetaFilterer, error) {
	contract, err := bindMeta(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MetaFilterer{contract: contract}, nil
}

// bindMeta binds a generic wrapper to an already deployed contract.
func bindMeta(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MetaMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Meta *MetaRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Meta.Contract.MetaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Meta *MetaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Meta.Contract.MetaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Meta *MetaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Meta.Contract.MetaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Meta *MetaCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Meta.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Meta *MetaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Meta.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Meta *MetaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Meta.Contract.contract.Transact(opts, method, params...)
}

// AddressProvider is a free data retrieval call binding the contract method 0xce50c2e7.
//
// Solidity: function address_provider() view returns(address)
func (_Meta *MetaCaller) AddressProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "address_provider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressProvider is a free data retrieval call binding the contract method 0xce50c2e7.
//
// Solidity: function address_provider() view returns(address)
func (_Meta *MetaSession) AddressProvider() (common.Address, error) {
	return _Meta.Contract.AddressProvider(&_Meta.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0xce50c2e7.
//
// Solidity: function address_provider() view returns(address)
func (_Meta *MetaCallerSession) AddressProvider() (common.Address, error) {
	return _Meta.Contract.AddressProvider(&_Meta.CallOpts)
}

// FindPoolForCoins is a free data retrieval call binding the contract method 0xa87df06c.
//
// Solidity: function find_pool_for_coins(address _from, address _to) view returns(address)
func (_Meta *MetaCaller) FindPoolForCoins(opts *bind.CallOpts, _from common.Address, _to common.Address) (common.Address, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "find_pool_for_coins", _from, _to)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FindPoolForCoins is a free data retrieval call binding the contract method 0xa87df06c.
//
// Solidity: function find_pool_for_coins(address _from, address _to) view returns(address)
func (_Meta *MetaSession) FindPoolForCoins(_from common.Address, _to common.Address) (common.Address, error) {
	return _Meta.Contract.FindPoolForCoins(&_Meta.CallOpts, _from, _to)
}

// FindPoolForCoins is a free data retrieval call binding the contract method 0xa87df06c.
//
// Solidity: function find_pool_for_coins(address _from, address _to) view returns(address)
func (_Meta *MetaCallerSession) FindPoolForCoins(_from common.Address, _to common.Address) (common.Address, error) {
	return _Meta.Contract.FindPoolForCoins(&_Meta.CallOpts, _from, _to)
}

// FindPoolForCoins0 is a free data retrieval call binding the contract method 0x6982eb0b.
//
// Solidity: function find_pool_for_coins(address _from, address _to, uint256 i) view returns(address)
func (_Meta *MetaCaller) FindPoolForCoins0(opts *bind.CallOpts, _from common.Address, _to common.Address, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "find_pool_for_coins0", _from, _to, i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FindPoolForCoins0 is a free data retrieval call binding the contract method 0x6982eb0b.
//
// Solidity: function find_pool_for_coins(address _from, address _to, uint256 i) view returns(address)
func (_Meta *MetaSession) FindPoolForCoins0(_from common.Address, _to common.Address, i *big.Int) (common.Address, error) {
	return _Meta.Contract.FindPoolForCoins0(&_Meta.CallOpts, _from, _to, i)
}

// FindPoolForCoins0 is a free data retrieval call binding the contract method 0x6982eb0b.
//
// Solidity: function find_pool_for_coins(address _from, address _to, uint256 i) view returns(address)
func (_Meta *MetaCallerSession) FindPoolForCoins0(_from common.Address, _to common.Address, i *big.Int) (common.Address, error) {
	return _Meta.Contract.FindPoolForCoins0(&_Meta.CallOpts, _from, _to, i)
}

// FindPoolsForCoins is a free data retrieval call binding the contract method 0xa064072b.
//
// Solidity: function find_pools_for_coins(address _from, address _to) view returns(address[])
func (_Meta *MetaCaller) FindPoolsForCoins(opts *bind.CallOpts, _from common.Address, _to common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "find_pools_for_coins", _from, _to)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// FindPoolsForCoins is a free data retrieval call binding the contract method 0xa064072b.
//
// Solidity: function find_pools_for_coins(address _from, address _to) view returns(address[])
func (_Meta *MetaSession) FindPoolsForCoins(_from common.Address, _to common.Address) ([]common.Address, error) {
	return _Meta.Contract.FindPoolsForCoins(&_Meta.CallOpts, _from, _to)
}

// FindPoolsForCoins is a free data retrieval call binding the contract method 0xa064072b.
//
// Solidity: function find_pools_for_coins(address _from, address _to) view returns(address[])
func (_Meta *MetaCallerSession) FindPoolsForCoins(_from common.Address, _to common.Address) ([]common.Address, error) {
	return _Meta.Contract.FindPoolsForCoins(&_Meta.CallOpts, _from, _to)
}

// GetAdminBalances is a free data retrieval call binding the contract method 0xc11e45b8.
//
// Solidity: function get_admin_balances(address _pool) view returns(uint256[8])
func (_Meta *MetaCaller) GetAdminBalances(opts *bind.CallOpts, _pool common.Address) ([8]*big.Int, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_admin_balances", _pool)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetAdminBalances is a free data retrieval call binding the contract method 0xc11e45b8.
//
// Solidity: function get_admin_balances(address _pool) view returns(uint256[8])
func (_Meta *MetaSession) GetAdminBalances(_pool common.Address) ([8]*big.Int, error) {
	return _Meta.Contract.GetAdminBalances(&_Meta.CallOpts, _pool)
}

// GetAdminBalances is a free data retrieval call binding the contract method 0xc11e45b8.
//
// Solidity: function get_admin_balances(address _pool) view returns(uint256[8])
func (_Meta *MetaCallerSession) GetAdminBalances(_pool common.Address) ([8]*big.Int, error) {
	return _Meta.Contract.GetAdminBalances(&_Meta.CallOpts, _pool)
}

// GetAdminBalances0 is a free data retrieval call binding the contract method 0xc0bf4cbf.
//
// Solidity: function get_admin_balances(address _pool, uint256 _handler_id) view returns(uint256[8])
func (_Meta *MetaCaller) GetAdminBalances0(opts *bind.CallOpts, _pool common.Address, _handler_id *big.Int) ([8]*big.Int, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_admin_balances0", _pool, _handler_id)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetAdminBalances0 is a free data retrieval call binding the contract method 0xc0bf4cbf.
//
// Solidity: function get_admin_balances(address _pool, uint256 _handler_id) view returns(uint256[8])
func (_Meta *MetaSession) GetAdminBalances0(_pool common.Address, _handler_id *big.Int) ([8]*big.Int, error) {
	return _Meta.Contract.GetAdminBalances0(&_Meta.CallOpts, _pool, _handler_id)
}

// GetAdminBalances0 is a free data retrieval call binding the contract method 0xc0bf4cbf.
//
// Solidity: function get_admin_balances(address _pool, uint256 _handler_id) view returns(uint256[8])
func (_Meta *MetaCallerSession) GetAdminBalances0(_pool common.Address, _handler_id *big.Int) ([8]*big.Int, error) {
	return _Meta.Contract.GetAdminBalances0(&_Meta.CallOpts, _pool, _handler_id)
}

// GetBalances is a free data retrieval call binding the contract method 0x92e3cc2d.
//
// Solidity: function get_balances(address _pool) view returns(uint256[8])
func (_Meta *MetaCaller) GetBalances(opts *bind.CallOpts, _pool common.Address) ([8]*big.Int, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_balances", _pool)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetBalances is a free data retrieval call binding the contract method 0x92e3cc2d.
//
// Solidity: function get_balances(address _pool) view returns(uint256[8])
func (_Meta *MetaSession) GetBalances(_pool common.Address) ([8]*big.Int, error) {
	return _Meta.Contract.GetBalances(&_Meta.CallOpts, _pool)
}

// GetBalances is a free data retrieval call binding the contract method 0x92e3cc2d.
//
// Solidity: function get_balances(address _pool) view returns(uint256[8])
func (_Meta *MetaCallerSession) GetBalances(_pool common.Address) ([8]*big.Int, error) {
	return _Meta.Contract.GetBalances(&_Meta.CallOpts, _pool)
}

// GetBalances0 is a free data retrieval call binding the contract method 0xaa85169c.
//
// Solidity: function get_balances(address _pool, uint256 _handler_id) view returns(uint256[8])
func (_Meta *MetaCaller) GetBalances0(opts *bind.CallOpts, _pool common.Address, _handler_id *big.Int) ([8]*big.Int, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_balances0", _pool, _handler_id)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetBalances0 is a free data retrieval call binding the contract method 0xaa85169c.
//
// Solidity: function get_balances(address _pool, uint256 _handler_id) view returns(uint256[8])
func (_Meta *MetaSession) GetBalances0(_pool common.Address, _handler_id *big.Int) ([8]*big.Int, error) {
	return _Meta.Contract.GetBalances0(&_Meta.CallOpts, _pool, _handler_id)
}

// GetBalances0 is a free data retrieval call binding the contract method 0xaa85169c.
//
// Solidity: function get_balances(address _pool, uint256 _handler_id) view returns(uint256[8])
func (_Meta *MetaCallerSession) GetBalances0(_pool common.Address, _handler_id *big.Int) ([8]*big.Int, error) {
	return _Meta.Contract.GetBalances0(&_Meta.CallOpts, _pool, _handler_id)
}

// GetBasePool is a free data retrieval call binding the contract method 0x6f20d6dd.
//
// Solidity: function get_base_pool(address _pool) view returns(address)
func (_Meta *MetaCaller) GetBasePool(opts *bind.CallOpts, _pool common.Address) (common.Address, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_base_pool", _pool)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBasePool is a free data retrieval call binding the contract method 0x6f20d6dd.
//
// Solidity: function get_base_pool(address _pool) view returns(address)
func (_Meta *MetaSession) GetBasePool(_pool common.Address) (common.Address, error) {
	return _Meta.Contract.GetBasePool(&_Meta.CallOpts, _pool)
}

// GetBasePool is a free data retrieval call binding the contract method 0x6f20d6dd.
//
// Solidity: function get_base_pool(address _pool) view returns(address)
func (_Meta *MetaCallerSession) GetBasePool(_pool common.Address) (common.Address, error) {
	return _Meta.Contract.GetBasePool(&_Meta.CallOpts, _pool)
}

// GetBasePool0 is a free data retrieval call binding the contract method 0xa54e3ade.
//
// Solidity: function get_base_pool(address _pool, uint256 _handler_id) view returns(address)
func (_Meta *MetaCaller) GetBasePool0(opts *bind.CallOpts, _pool common.Address, _handler_id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_base_pool0", _pool, _handler_id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBasePool0 is a free data retrieval call binding the contract method 0xa54e3ade.
//
// Solidity: function get_base_pool(address _pool, uint256 _handler_id) view returns(address)
func (_Meta *MetaSession) GetBasePool0(_pool common.Address, _handler_id *big.Int) (common.Address, error) {
	return _Meta.Contract.GetBasePool0(&_Meta.CallOpts, _pool, _handler_id)
}

// GetBasePool0 is a free data retrieval call binding the contract method 0xa54e3ade.
//
// Solidity: function get_base_pool(address _pool, uint256 _handler_id) view returns(address)
func (_Meta *MetaCallerSession) GetBasePool0(_pool common.Address, _handler_id *big.Int) (common.Address, error) {
	return _Meta.Contract.GetBasePool0(&_Meta.CallOpts, _pool, _handler_id)
}

// GetBaseRegistry is a free data retrieval call binding the contract method 0x84e1710d.
//
// Solidity: function get_base_registry(address registry_handler) view returns(address)
func (_Meta *MetaCaller) GetBaseRegistry(opts *bind.CallOpts, registry_handler common.Address) (common.Address, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_base_registry", registry_handler)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBaseRegistry is a free data retrieval call binding the contract method 0x84e1710d.
//
// Solidity: function get_base_registry(address registry_handler) view returns(address)
func (_Meta *MetaSession) GetBaseRegistry(registry_handler common.Address) (common.Address, error) {
	return _Meta.Contract.GetBaseRegistry(&_Meta.CallOpts, registry_handler)
}

// GetBaseRegistry is a free data retrieval call binding the contract method 0x84e1710d.
//
// Solidity: function get_base_registry(address registry_handler) view returns(address)
func (_Meta *MetaCallerSession) GetBaseRegistry(registry_handler common.Address) (common.Address, error) {
	return _Meta.Contract.GetBaseRegistry(&_Meta.CallOpts, registry_handler)
}

// GetCoinIndices is a free data retrieval call binding the contract method 0xeb85226d.
//
// Solidity: function get_coin_indices(address _pool, address _from, address _to) view returns(int128, int128, bool)
func (_Meta *MetaCaller) GetCoinIndices(opts *bind.CallOpts, _pool common.Address, _from common.Address, _to common.Address) (*big.Int, *big.Int, bool, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_coin_indices", _pool, _from, _to)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)

	return out0, out1, out2, err

}

// GetCoinIndices is a free data retrieval call binding the contract method 0xeb85226d.
//
// Solidity: function get_coin_indices(address _pool, address _from, address _to) view returns(int128, int128, bool)
func (_Meta *MetaSession) GetCoinIndices(_pool common.Address, _from common.Address, _to common.Address) (*big.Int, *big.Int, bool, error) {
	return _Meta.Contract.GetCoinIndices(&_Meta.CallOpts, _pool, _from, _to)
}

// GetCoinIndices is a free data retrieval call binding the contract method 0xeb85226d.
//
// Solidity: function get_coin_indices(address _pool, address _from, address _to) view returns(int128, int128, bool)
func (_Meta *MetaCallerSession) GetCoinIndices(_pool common.Address, _from common.Address, _to common.Address) (*big.Int, *big.Int, bool, error) {
	return _Meta.Contract.GetCoinIndices(&_Meta.CallOpts, _pool, _from, _to)
}

// GetCoinIndices0 is a free data retrieval call binding the contract method 0x63fb3ddb.
//
// Solidity: function get_coin_indices(address _pool, address _from, address _to, uint256 _handler_id) view returns(int128, int128, bool)
func (_Meta *MetaCaller) GetCoinIndices0(opts *bind.CallOpts, _pool common.Address, _from common.Address, _to common.Address, _handler_id *big.Int) (*big.Int, *big.Int, bool, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_coin_indices0", _pool, _from, _to, _handler_id)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)

	return out0, out1, out2, err

}

// GetCoinIndices0 is a free data retrieval call binding the contract method 0x63fb3ddb.
//
// Solidity: function get_coin_indices(address _pool, address _from, address _to, uint256 _handler_id) view returns(int128, int128, bool)
func (_Meta *MetaSession) GetCoinIndices0(_pool common.Address, _from common.Address, _to common.Address, _handler_id *big.Int) (*big.Int, *big.Int, bool, error) {
	return _Meta.Contract.GetCoinIndices0(&_Meta.CallOpts, _pool, _from, _to, _handler_id)
}

// GetCoinIndices0 is a free data retrieval call binding the contract method 0x63fb3ddb.
//
// Solidity: function get_coin_indices(address _pool, address _from, address _to, uint256 _handler_id) view returns(int128, int128, bool)
func (_Meta *MetaCallerSession) GetCoinIndices0(_pool common.Address, _from common.Address, _to common.Address, _handler_id *big.Int) (*big.Int, *big.Int, bool, error) {
	return _Meta.Contract.GetCoinIndices0(&_Meta.CallOpts, _pool, _from, _to, _handler_id)
}

// GetCoins is a free data retrieval call binding the contract method 0x9ac90d3d.
//
// Solidity: function get_coins(address _pool) view returns(address[8])
func (_Meta *MetaCaller) GetCoins(opts *bind.CallOpts, _pool common.Address) ([8]common.Address, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_coins", _pool)

	if err != nil {
		return *new([8]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([8]common.Address)).(*[8]common.Address)

	return out0, err

}

// GetCoins is a free data retrieval call binding the contract method 0x9ac90d3d.
//
// Solidity: function get_coins(address _pool) view returns(address[8])
func (_Meta *MetaSession) GetCoins(_pool common.Address) ([8]common.Address, error) {
	return _Meta.Contract.GetCoins(&_Meta.CallOpts, _pool)
}

// GetCoins is a free data retrieval call binding the contract method 0x9ac90d3d.
//
// Solidity: function get_coins(address _pool) view returns(address[8])
func (_Meta *MetaCallerSession) GetCoins(_pool common.Address) ([8]common.Address, error) {
	return _Meta.Contract.GetCoins(&_Meta.CallOpts, _pool)
}

// GetCoins0 is a free data retrieval call binding the contract method 0x6ebe51fc.
//
// Solidity: function get_coins(address _pool, uint256 _handler_id) view returns(address[8])
func (_Meta *MetaCaller) GetCoins0(opts *bind.CallOpts, _pool common.Address, _handler_id *big.Int) ([8]common.Address, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_coins0", _pool, _handler_id)

	if err != nil {
		return *new([8]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([8]common.Address)).(*[8]common.Address)

	return out0, err

}

// GetCoins0 is a free data retrieval call binding the contract method 0x6ebe51fc.
//
// Solidity: function get_coins(address _pool, uint256 _handler_id) view returns(address[8])
func (_Meta *MetaSession) GetCoins0(_pool common.Address, _handler_id *big.Int) ([8]common.Address, error) {
	return _Meta.Contract.GetCoins0(&_Meta.CallOpts, _pool, _handler_id)
}

// GetCoins0 is a free data retrieval call binding the contract method 0x6ebe51fc.
//
// Solidity: function get_coins(address _pool, uint256 _handler_id) view returns(address[8])
func (_Meta *MetaCallerSession) GetCoins0(_pool common.Address, _handler_id *big.Int) ([8]common.Address, error) {
	return _Meta.Contract.GetCoins0(&_Meta.CallOpts, _pool, _handler_id)
}

// GetDecimals is a free data retrieval call binding the contract method 0x52b51555.
//
// Solidity: function get_decimals(address _pool) view returns(uint256[8])
func (_Meta *MetaCaller) GetDecimals(opts *bind.CallOpts, _pool common.Address) ([8]*big.Int, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_decimals", _pool)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetDecimals is a free data retrieval call binding the contract method 0x52b51555.
//
// Solidity: function get_decimals(address _pool) view returns(uint256[8])
func (_Meta *MetaSession) GetDecimals(_pool common.Address) ([8]*big.Int, error) {
	return _Meta.Contract.GetDecimals(&_Meta.CallOpts, _pool)
}

// GetDecimals is a free data retrieval call binding the contract method 0x52b51555.
//
// Solidity: function get_decimals(address _pool) view returns(uint256[8])
func (_Meta *MetaCallerSession) GetDecimals(_pool common.Address) ([8]*big.Int, error) {
	return _Meta.Contract.GetDecimals(&_Meta.CallOpts, _pool)
}

// GetDecimals0 is a free data retrieval call binding the contract method 0x403f502f.
//
// Solidity: function get_decimals(address _pool, uint256 _handler_id) view returns(uint256[8])
func (_Meta *MetaCaller) GetDecimals0(opts *bind.CallOpts, _pool common.Address, _handler_id *big.Int) ([8]*big.Int, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_decimals0", _pool, _handler_id)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetDecimals0 is a free data retrieval call binding the contract method 0x403f502f.
//
// Solidity: function get_decimals(address _pool, uint256 _handler_id) view returns(uint256[8])
func (_Meta *MetaSession) GetDecimals0(_pool common.Address, _handler_id *big.Int) ([8]*big.Int, error) {
	return _Meta.Contract.GetDecimals0(&_Meta.CallOpts, _pool, _handler_id)
}

// GetDecimals0 is a free data retrieval call binding the contract method 0x403f502f.
//
// Solidity: function get_decimals(address _pool, uint256 _handler_id) view returns(uint256[8])
func (_Meta *MetaCallerSession) GetDecimals0(_pool common.Address, _handler_id *big.Int) ([8]*big.Int, error) {
	return _Meta.Contract.GetDecimals0(&_Meta.CallOpts, _pool, _handler_id)
}

// GetFees is a free data retrieval call binding the contract method 0x7cdb72b0.
//
// Solidity: function get_fees(address _pool) view returns(uint256[10])
func (_Meta *MetaCaller) GetFees(opts *bind.CallOpts, _pool common.Address) ([10]*big.Int, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_fees", _pool)

	if err != nil {
		return *new([10]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([10]*big.Int)).(*[10]*big.Int)

	return out0, err

}

// GetFees is a free data retrieval call binding the contract method 0x7cdb72b0.
//
// Solidity: function get_fees(address _pool) view returns(uint256[10])
func (_Meta *MetaSession) GetFees(_pool common.Address) ([10]*big.Int, error) {
	return _Meta.Contract.GetFees(&_Meta.CallOpts, _pool)
}

// GetFees is a free data retrieval call binding the contract method 0x7cdb72b0.
//
// Solidity: function get_fees(address _pool) view returns(uint256[10])
func (_Meta *MetaCallerSession) GetFees(_pool common.Address) ([10]*big.Int, error) {
	return _Meta.Contract.GetFees(&_Meta.CallOpts, _pool)
}

// GetFees0 is a free data retrieval call binding the contract method 0x0ed5a427.
//
// Solidity: function get_fees(address _pool, uint256 _handler_id) view returns(uint256[10])
func (_Meta *MetaCaller) GetFees0(opts *bind.CallOpts, _pool common.Address, _handler_id *big.Int) ([10]*big.Int, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_fees0", _pool, _handler_id)

	if err != nil {
		return *new([10]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([10]*big.Int)).(*[10]*big.Int)

	return out0, err

}

// GetFees0 is a free data retrieval call binding the contract method 0x0ed5a427.
//
// Solidity: function get_fees(address _pool, uint256 _handler_id) view returns(uint256[10])
func (_Meta *MetaSession) GetFees0(_pool common.Address, _handler_id *big.Int) ([10]*big.Int, error) {
	return _Meta.Contract.GetFees0(&_Meta.CallOpts, _pool, _handler_id)
}

// GetFees0 is a free data retrieval call binding the contract method 0x0ed5a427.
//
// Solidity: function get_fees(address _pool, uint256 _handler_id) view returns(uint256[10])
func (_Meta *MetaCallerSession) GetFees0(_pool common.Address, _handler_id *big.Int) ([10]*big.Int, error) {
	return _Meta.Contract.GetFees0(&_Meta.CallOpts, _pool, _handler_id)
}

// GetGauge is a free data retrieval call binding the contract method 0xdaf297b9.
//
// Solidity: function get_gauge(address _pool) view returns(address)
func (_Meta *MetaCaller) GetGauge(opts *bind.CallOpts, _pool common.Address) (common.Address, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_gauge", _pool)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGauge is a free data retrieval call binding the contract method 0xdaf297b9.
//
// Solidity: function get_gauge(address _pool) view returns(address)
func (_Meta *MetaSession) GetGauge(_pool common.Address) (common.Address, error) {
	return _Meta.Contract.GetGauge(&_Meta.CallOpts, _pool)
}

// GetGauge is a free data retrieval call binding the contract method 0xdaf297b9.
//
// Solidity: function get_gauge(address _pool) view returns(address)
func (_Meta *MetaCallerSession) GetGauge(_pool common.Address) (common.Address, error) {
	return _Meta.Contract.GetGauge(&_Meta.CallOpts, _pool)
}

// GetGauge0 is a free data retrieval call binding the contract method 0xe4081220.
//
// Solidity: function get_gauge(address _pool, uint256 gauge_idx) view returns(address)
func (_Meta *MetaCaller) GetGauge0(opts *bind.CallOpts, _pool common.Address, gauge_idx *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_gauge0", _pool, gauge_idx)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGauge0 is a free data retrieval call binding the contract method 0xe4081220.
//
// Solidity: function get_gauge(address _pool, uint256 gauge_idx) view returns(address)
func (_Meta *MetaSession) GetGauge0(_pool common.Address, gauge_idx *big.Int) (common.Address, error) {
	return _Meta.Contract.GetGauge0(&_Meta.CallOpts, _pool, gauge_idx)
}

// GetGauge0 is a free data retrieval call binding the contract method 0xe4081220.
//
// Solidity: function get_gauge(address _pool, uint256 gauge_idx) view returns(address)
func (_Meta *MetaCallerSession) GetGauge0(_pool common.Address, gauge_idx *big.Int) (common.Address, error) {
	return _Meta.Contract.GetGauge0(&_Meta.CallOpts, _pool, gauge_idx)
}

// GetGauge1 is a free data retrieval call binding the contract method 0x773fb7e3.
//
// Solidity: function get_gauge(address _pool, uint256 gauge_idx, uint256 _handler_id) view returns(address)
func (_Meta *MetaCaller) GetGauge1(opts *bind.CallOpts, _pool common.Address, gauge_idx *big.Int, _handler_id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_gauge1", _pool, gauge_idx, _handler_id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGauge1 is a free data retrieval call binding the contract method 0x773fb7e3.
//
// Solidity: function get_gauge(address _pool, uint256 gauge_idx, uint256 _handler_id) view returns(address)
func (_Meta *MetaSession) GetGauge1(_pool common.Address, gauge_idx *big.Int, _handler_id *big.Int) (common.Address, error) {
	return _Meta.Contract.GetGauge1(&_Meta.CallOpts, _pool, gauge_idx, _handler_id)
}

// GetGauge1 is a free data retrieval call binding the contract method 0x773fb7e3.
//
// Solidity: function get_gauge(address _pool, uint256 gauge_idx, uint256 _handler_id) view returns(address)
func (_Meta *MetaCallerSession) GetGauge1(_pool common.Address, gauge_idx *big.Int, _handler_id *big.Int) (common.Address, error) {
	return _Meta.Contract.GetGauge1(&_Meta.CallOpts, _pool, gauge_idx, _handler_id)
}

// GetGaugeType is a free data retrieval call binding the contract method 0x25fa5d13.
//
// Solidity: function get_gauge_type(address _pool) view returns(int128)
func (_Meta *MetaCaller) GetGaugeType(opts *bind.CallOpts, _pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_gauge_type", _pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGaugeType is a free data retrieval call binding the contract method 0x25fa5d13.
//
// Solidity: function get_gauge_type(address _pool) view returns(int128)
func (_Meta *MetaSession) GetGaugeType(_pool common.Address) (*big.Int, error) {
	return _Meta.Contract.GetGaugeType(&_Meta.CallOpts, _pool)
}

// GetGaugeType is a free data retrieval call binding the contract method 0x25fa5d13.
//
// Solidity: function get_gauge_type(address _pool) view returns(int128)
func (_Meta *MetaCallerSession) GetGaugeType(_pool common.Address) (*big.Int, error) {
	return _Meta.Contract.GetGaugeType(&_Meta.CallOpts, _pool)
}

// GetGaugeType0 is a free data retrieval call binding the contract method 0x7c51db55.
//
// Solidity: function get_gauge_type(address _pool, uint256 gauge_idx) view returns(int128)
func (_Meta *MetaCaller) GetGaugeType0(opts *bind.CallOpts, _pool common.Address, gauge_idx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_gauge_type0", _pool, gauge_idx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGaugeType0 is a free data retrieval call binding the contract method 0x7c51db55.
//
// Solidity: function get_gauge_type(address _pool, uint256 gauge_idx) view returns(int128)
func (_Meta *MetaSession) GetGaugeType0(_pool common.Address, gauge_idx *big.Int) (*big.Int, error) {
	return _Meta.Contract.GetGaugeType0(&_Meta.CallOpts, _pool, gauge_idx)
}

// GetGaugeType0 is a free data retrieval call binding the contract method 0x7c51db55.
//
// Solidity: function get_gauge_type(address _pool, uint256 gauge_idx) view returns(int128)
func (_Meta *MetaCallerSession) GetGaugeType0(_pool common.Address, gauge_idx *big.Int) (*big.Int, error) {
	return _Meta.Contract.GetGaugeType0(&_Meta.CallOpts, _pool, gauge_idx)
}

// GetGaugeType1 is a free data retrieval call binding the contract method 0xf8b661e2.
//
// Solidity: function get_gauge_type(address _pool, uint256 gauge_idx, uint256 _handler_id) view returns(int128)
func (_Meta *MetaCaller) GetGaugeType1(opts *bind.CallOpts, _pool common.Address, gauge_idx *big.Int, _handler_id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_gauge_type1", _pool, gauge_idx, _handler_id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGaugeType1 is a free data retrieval call binding the contract method 0xf8b661e2.
//
// Solidity: function get_gauge_type(address _pool, uint256 gauge_idx, uint256 _handler_id) view returns(int128)
func (_Meta *MetaSession) GetGaugeType1(_pool common.Address, gauge_idx *big.Int, _handler_id *big.Int) (*big.Int, error) {
	return _Meta.Contract.GetGaugeType1(&_Meta.CallOpts, _pool, gauge_idx, _handler_id)
}

// GetGaugeType1 is a free data retrieval call binding the contract method 0xf8b661e2.
//
// Solidity: function get_gauge_type(address _pool, uint256 gauge_idx, uint256 _handler_id) view returns(int128)
func (_Meta *MetaCallerSession) GetGaugeType1(_pool common.Address, gauge_idx *big.Int, _handler_id *big.Int) (*big.Int, error) {
	return _Meta.Contract.GetGaugeType1(&_Meta.CallOpts, _pool, gauge_idx, _handler_id)
}

// GetLpToken is a free data retrieval call binding the contract method 0x37951049.
//
// Solidity: function get_lp_token(address _pool) view returns(address)
func (_Meta *MetaCaller) GetLpToken(opts *bind.CallOpts, _pool common.Address) (common.Address, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_lp_token", _pool)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLpToken is a free data retrieval call binding the contract method 0x37951049.
//
// Solidity: function get_lp_token(address _pool) view returns(address)
func (_Meta *MetaSession) GetLpToken(_pool common.Address) (common.Address, error) {
	return _Meta.Contract.GetLpToken(&_Meta.CallOpts, _pool)
}

// GetLpToken is a free data retrieval call binding the contract method 0x37951049.
//
// Solidity: function get_lp_token(address _pool) view returns(address)
func (_Meta *MetaCallerSession) GetLpToken(_pool common.Address) (common.Address, error) {
	return _Meta.Contract.GetLpToken(&_Meta.CallOpts, _pool)
}

// GetLpToken0 is a free data retrieval call binding the contract method 0x0881715f.
//
// Solidity: function get_lp_token(address _pool, uint256 _handler_id) view returns(address)
func (_Meta *MetaCaller) GetLpToken0(opts *bind.CallOpts, _pool common.Address, _handler_id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_lp_token0", _pool, _handler_id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLpToken0 is a free data retrieval call binding the contract method 0x0881715f.
//
// Solidity: function get_lp_token(address _pool, uint256 _handler_id) view returns(address)
func (_Meta *MetaSession) GetLpToken0(_pool common.Address, _handler_id *big.Int) (common.Address, error) {
	return _Meta.Contract.GetLpToken0(&_Meta.CallOpts, _pool, _handler_id)
}

// GetLpToken0 is a free data retrieval call binding the contract method 0x0881715f.
//
// Solidity: function get_lp_token(address _pool, uint256 _handler_id) view returns(address)
func (_Meta *MetaCallerSession) GetLpToken0(_pool common.Address, _handler_id *big.Int) (common.Address, error) {
	return _Meta.Contract.GetLpToken0(&_Meta.CallOpts, _pool, _handler_id)
}

// GetNCoins is a free data retrieval call binding the contract method 0x940494f1.
//
// Solidity: function get_n_coins(address _pool) view returns(uint256)
func (_Meta *MetaCaller) GetNCoins(opts *bind.CallOpts, _pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_n_coins", _pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNCoins is a free data retrieval call binding the contract method 0x940494f1.
//
// Solidity: function get_n_coins(address _pool) view returns(uint256)
func (_Meta *MetaSession) GetNCoins(_pool common.Address) (*big.Int, error) {
	return _Meta.Contract.GetNCoins(&_Meta.CallOpts, _pool)
}

// GetNCoins is a free data retrieval call binding the contract method 0x940494f1.
//
// Solidity: function get_n_coins(address _pool) view returns(uint256)
func (_Meta *MetaCallerSession) GetNCoins(_pool common.Address) (*big.Int, error) {
	return _Meta.Contract.GetNCoins(&_Meta.CallOpts, _pool)
}

// GetNCoins0 is a free data retrieval call binding the contract method 0x11d81279.
//
// Solidity: function get_n_coins(address _pool, uint256 _handler_id) view returns(uint256)
func (_Meta *MetaCaller) GetNCoins0(opts *bind.CallOpts, _pool common.Address, _handler_id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_n_coins0", _pool, _handler_id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNCoins0 is a free data retrieval call binding the contract method 0x11d81279.
//
// Solidity: function get_n_coins(address _pool, uint256 _handler_id) view returns(uint256)
func (_Meta *MetaSession) GetNCoins0(_pool common.Address, _handler_id *big.Int) (*big.Int, error) {
	return _Meta.Contract.GetNCoins0(&_Meta.CallOpts, _pool, _handler_id)
}

// GetNCoins0 is a free data retrieval call binding the contract method 0x11d81279.
//
// Solidity: function get_n_coins(address _pool, uint256 _handler_id) view returns(uint256)
func (_Meta *MetaCallerSession) GetNCoins0(_pool common.Address, _handler_id *big.Int) (*big.Int, error) {
	return _Meta.Contract.GetNCoins0(&_Meta.CallOpts, _pool, _handler_id)
}

// GetNUnderlyingCoins is a free data retrieval call binding the contract method 0x0a700c08.
//
// Solidity: function get_n_underlying_coins(address _pool) view returns(uint256)
func (_Meta *MetaCaller) GetNUnderlyingCoins(opts *bind.CallOpts, _pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_n_underlying_coins", _pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNUnderlyingCoins is a free data retrieval call binding the contract method 0x0a700c08.
//
// Solidity: function get_n_underlying_coins(address _pool) view returns(uint256)
func (_Meta *MetaSession) GetNUnderlyingCoins(_pool common.Address) (*big.Int, error) {
	return _Meta.Contract.GetNUnderlyingCoins(&_Meta.CallOpts, _pool)
}

// GetNUnderlyingCoins is a free data retrieval call binding the contract method 0x0a700c08.
//
// Solidity: function get_n_underlying_coins(address _pool) view returns(uint256)
func (_Meta *MetaCallerSession) GetNUnderlyingCoins(_pool common.Address) (*big.Int, error) {
	return _Meta.Contract.GetNUnderlyingCoins(&_Meta.CallOpts, _pool)
}

// GetNUnderlyingCoins0 is a free data retrieval call binding the contract method 0xdecdf68f.
//
// Solidity: function get_n_underlying_coins(address _pool, uint256 _handler_id) view returns(uint256)
func (_Meta *MetaCaller) GetNUnderlyingCoins0(opts *bind.CallOpts, _pool common.Address, _handler_id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_n_underlying_coins0", _pool, _handler_id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNUnderlyingCoins0 is a free data retrieval call binding the contract method 0xdecdf68f.
//
// Solidity: function get_n_underlying_coins(address _pool, uint256 _handler_id) view returns(uint256)
func (_Meta *MetaSession) GetNUnderlyingCoins0(_pool common.Address, _handler_id *big.Int) (*big.Int, error) {
	return _Meta.Contract.GetNUnderlyingCoins0(&_Meta.CallOpts, _pool, _handler_id)
}

// GetNUnderlyingCoins0 is a free data retrieval call binding the contract method 0xdecdf68f.
//
// Solidity: function get_n_underlying_coins(address _pool, uint256 _handler_id) view returns(uint256)
func (_Meta *MetaCallerSession) GetNUnderlyingCoins0(_pool common.Address, _handler_id *big.Int) (*big.Int, error) {
	return _Meta.Contract.GetNUnderlyingCoins0(&_Meta.CallOpts, _pool, _handler_id)
}

// GetPoolAssetType is a free data retrieval call binding the contract method 0x66d3966c.
//
// Solidity: function get_pool_asset_type(address _pool) view returns(uint256)
func (_Meta *MetaCaller) GetPoolAssetType(opts *bind.CallOpts, _pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_pool_asset_type", _pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolAssetType is a free data retrieval call binding the contract method 0x66d3966c.
//
// Solidity: function get_pool_asset_type(address _pool) view returns(uint256)
func (_Meta *MetaSession) GetPoolAssetType(_pool common.Address) (*big.Int, error) {
	return _Meta.Contract.GetPoolAssetType(&_Meta.CallOpts, _pool)
}

// GetPoolAssetType is a free data retrieval call binding the contract method 0x66d3966c.
//
// Solidity: function get_pool_asset_type(address _pool) view returns(uint256)
func (_Meta *MetaCallerSession) GetPoolAssetType(_pool common.Address) (*big.Int, error) {
	return _Meta.Contract.GetPoolAssetType(&_Meta.CallOpts, _pool)
}

// GetPoolAssetType0 is a free data retrieval call binding the contract method 0x90d1dd2d.
//
// Solidity: function get_pool_asset_type(address _pool, uint256 _handler_id) view returns(uint256)
func (_Meta *MetaCaller) GetPoolAssetType0(opts *bind.CallOpts, _pool common.Address, _handler_id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_pool_asset_type0", _pool, _handler_id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolAssetType0 is a free data retrieval call binding the contract method 0x90d1dd2d.
//
// Solidity: function get_pool_asset_type(address _pool, uint256 _handler_id) view returns(uint256)
func (_Meta *MetaSession) GetPoolAssetType0(_pool common.Address, _handler_id *big.Int) (*big.Int, error) {
	return _Meta.Contract.GetPoolAssetType0(&_Meta.CallOpts, _pool, _handler_id)
}

// GetPoolAssetType0 is a free data retrieval call binding the contract method 0x90d1dd2d.
//
// Solidity: function get_pool_asset_type(address _pool, uint256 _handler_id) view returns(uint256)
func (_Meta *MetaCallerSession) GetPoolAssetType0(_pool common.Address, _handler_id *big.Int) (*big.Int, error) {
	return _Meta.Contract.GetPoolAssetType0(&_Meta.CallOpts, _pool, _handler_id)
}

// GetPoolFromLpToken is a free data retrieval call binding the contract method 0xbdf475c3.
//
// Solidity: function get_pool_from_lp_token(address _token) view returns(address)
func (_Meta *MetaCaller) GetPoolFromLpToken(opts *bind.CallOpts, _token common.Address) (common.Address, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_pool_from_lp_token", _token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPoolFromLpToken is a free data retrieval call binding the contract method 0xbdf475c3.
//
// Solidity: function get_pool_from_lp_token(address _token) view returns(address)
func (_Meta *MetaSession) GetPoolFromLpToken(_token common.Address) (common.Address, error) {
	return _Meta.Contract.GetPoolFromLpToken(&_Meta.CallOpts, _token)
}

// GetPoolFromLpToken is a free data retrieval call binding the contract method 0xbdf475c3.
//
// Solidity: function get_pool_from_lp_token(address _token) view returns(address)
func (_Meta *MetaCallerSession) GetPoolFromLpToken(_token common.Address) (common.Address, error) {
	return _Meta.Contract.GetPoolFromLpToken(&_Meta.CallOpts, _token)
}

// GetPoolFromLpToken0 is a free data retrieval call binding the contract method 0x36678b80.
//
// Solidity: function get_pool_from_lp_token(address _token, uint256 _handler_id) view returns(address)
func (_Meta *MetaCaller) GetPoolFromLpToken0(opts *bind.CallOpts, _token common.Address, _handler_id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_pool_from_lp_token0", _token, _handler_id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPoolFromLpToken0 is a free data retrieval call binding the contract method 0x36678b80.
//
// Solidity: function get_pool_from_lp_token(address _token, uint256 _handler_id) view returns(address)
func (_Meta *MetaSession) GetPoolFromLpToken0(_token common.Address, _handler_id *big.Int) (common.Address, error) {
	return _Meta.Contract.GetPoolFromLpToken0(&_Meta.CallOpts, _token, _handler_id)
}

// GetPoolFromLpToken0 is a free data retrieval call binding the contract method 0x36678b80.
//
// Solidity: function get_pool_from_lp_token(address _token, uint256 _handler_id) view returns(address)
func (_Meta *MetaCallerSession) GetPoolFromLpToken0(_token common.Address, _handler_id *big.Int) (common.Address, error) {
	return _Meta.Contract.GetPoolFromLpToken0(&_Meta.CallOpts, _token, _handler_id)
}

// GetPoolName is a free data retrieval call binding the contract method 0x5c911741.
//
// Solidity: function get_pool_name(address _pool) view returns(string)
func (_Meta *MetaCaller) GetPoolName(opts *bind.CallOpts, _pool common.Address) (string, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_pool_name", _pool)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetPoolName is a free data retrieval call binding the contract method 0x5c911741.
//
// Solidity: function get_pool_name(address _pool) view returns(string)
func (_Meta *MetaSession) GetPoolName(_pool common.Address) (string, error) {
	return _Meta.Contract.GetPoolName(&_Meta.CallOpts, _pool)
}

// GetPoolName is a free data retrieval call binding the contract method 0x5c911741.
//
// Solidity: function get_pool_name(address _pool) view returns(string)
func (_Meta *MetaCallerSession) GetPoolName(_pool common.Address) (string, error) {
	return _Meta.Contract.GetPoolName(&_Meta.CallOpts, _pool)
}

// GetPoolName0 is a free data retrieval call binding the contract method 0x92234f45.
//
// Solidity: function get_pool_name(address _pool, uint256 _handler_id) view returns(string)
func (_Meta *MetaCaller) GetPoolName0(opts *bind.CallOpts, _pool common.Address, _handler_id *big.Int) (string, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_pool_name0", _pool, _handler_id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetPoolName0 is a free data retrieval call binding the contract method 0x92234f45.
//
// Solidity: function get_pool_name(address _pool, uint256 _handler_id) view returns(string)
func (_Meta *MetaSession) GetPoolName0(_pool common.Address, _handler_id *big.Int) (string, error) {
	return _Meta.Contract.GetPoolName0(&_Meta.CallOpts, _pool, _handler_id)
}

// GetPoolName0 is a free data retrieval call binding the contract method 0x92234f45.
//
// Solidity: function get_pool_name(address _pool, uint256 _handler_id) view returns(string)
func (_Meta *MetaCallerSession) GetPoolName0(_pool common.Address, _handler_id *big.Int) (string, error) {
	return _Meta.Contract.GetPoolName0(&_Meta.CallOpts, _pool, _handler_id)
}

// GetPoolParams is a free data retrieval call binding the contract method 0x688532aa.
//
// Solidity: function get_pool_params(address _pool) view returns(uint256[20])
func (_Meta *MetaCaller) GetPoolParams(opts *bind.CallOpts, _pool common.Address) ([20]*big.Int, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_pool_params", _pool)

	if err != nil {
		return *new([20]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([20]*big.Int)).(*[20]*big.Int)

	return out0, err

}

// GetPoolParams is a free data retrieval call binding the contract method 0x688532aa.
//
// Solidity: function get_pool_params(address _pool) view returns(uint256[20])
func (_Meta *MetaSession) GetPoolParams(_pool common.Address) ([20]*big.Int, error) {
	return _Meta.Contract.GetPoolParams(&_Meta.CallOpts, _pool)
}

// GetPoolParams is a free data retrieval call binding the contract method 0x688532aa.
//
// Solidity: function get_pool_params(address _pool) view returns(uint256[20])
func (_Meta *MetaCallerSession) GetPoolParams(_pool common.Address) ([20]*big.Int, error) {
	return _Meta.Contract.GetPoolParams(&_Meta.CallOpts, _pool)
}

// GetPoolParams0 is a free data retrieval call binding the contract method 0x7a65d2dd.
//
// Solidity: function get_pool_params(address _pool, uint256 _handler_id) view returns(uint256[20])
func (_Meta *MetaCaller) GetPoolParams0(opts *bind.CallOpts, _pool common.Address, _handler_id *big.Int) ([20]*big.Int, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_pool_params0", _pool, _handler_id)

	if err != nil {
		return *new([20]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([20]*big.Int)).(*[20]*big.Int)

	return out0, err

}

// GetPoolParams0 is a free data retrieval call binding the contract method 0x7a65d2dd.
//
// Solidity: function get_pool_params(address _pool, uint256 _handler_id) view returns(uint256[20])
func (_Meta *MetaSession) GetPoolParams0(_pool common.Address, _handler_id *big.Int) ([20]*big.Int, error) {
	return _Meta.Contract.GetPoolParams0(&_Meta.CallOpts, _pool, _handler_id)
}

// GetPoolParams0 is a free data retrieval call binding the contract method 0x7a65d2dd.
//
// Solidity: function get_pool_params(address _pool, uint256 _handler_id) view returns(uint256[20])
func (_Meta *MetaCallerSession) GetPoolParams0(_pool common.Address, _handler_id *big.Int) ([20]*big.Int, error) {
	return _Meta.Contract.GetPoolParams0(&_Meta.CallOpts, _pool, _handler_id)
}

// GetRegistry is a free data retrieval call binding the contract method 0x913d9b4d.
//
// Solidity: function get_registry(uint256 arg0) view returns(address)
func (_Meta *MetaCaller) GetRegistry(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_registry", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRegistry is a free data retrieval call binding the contract method 0x913d9b4d.
//
// Solidity: function get_registry(uint256 arg0) view returns(address)
func (_Meta *MetaSession) GetRegistry(arg0 *big.Int) (common.Address, error) {
	return _Meta.Contract.GetRegistry(&_Meta.CallOpts, arg0)
}

// GetRegistry is a free data retrieval call binding the contract method 0x913d9b4d.
//
// Solidity: function get_registry(uint256 arg0) view returns(address)
func (_Meta *MetaCallerSession) GetRegistry(arg0 *big.Int) (common.Address, error) {
	return _Meta.Contract.GetRegistry(&_Meta.CallOpts, arg0)
}

// GetRegistryHandlersFromPool is a free data retrieval call binding the contract method 0x308d1b6d.
//
// Solidity: function get_registry_handlers_from_pool(address _pool) view returns(address[10])
func (_Meta *MetaCaller) GetRegistryHandlersFromPool(opts *bind.CallOpts, _pool common.Address) ([10]common.Address, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_registry_handlers_from_pool", _pool)

	if err != nil {
		return *new([10]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([10]common.Address)).(*[10]common.Address)

	return out0, err

}

// GetRegistryHandlersFromPool is a free data retrieval call binding the contract method 0x308d1b6d.
//
// Solidity: function get_registry_handlers_from_pool(address _pool) view returns(address[10])
func (_Meta *MetaSession) GetRegistryHandlersFromPool(_pool common.Address) ([10]common.Address, error) {
	return _Meta.Contract.GetRegistryHandlersFromPool(&_Meta.CallOpts, _pool)
}

// GetRegistryHandlersFromPool is a free data retrieval call binding the contract method 0x308d1b6d.
//
// Solidity: function get_registry_handlers_from_pool(address _pool) view returns(address[10])
func (_Meta *MetaCallerSession) GetRegistryHandlersFromPool(_pool common.Address) ([10]common.Address, error) {
	return _Meta.Contract.GetRegistryHandlersFromPool(&_Meta.CallOpts, _pool)
}

// GetUnderlyingBalances is a free data retrieval call binding the contract method 0x59f4f351.
//
// Solidity: function get_underlying_balances(address _pool) view returns(uint256[8])
func (_Meta *MetaCaller) GetUnderlyingBalances(opts *bind.CallOpts, _pool common.Address) ([8]*big.Int, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_underlying_balances", _pool)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetUnderlyingBalances is a free data retrieval call binding the contract method 0x59f4f351.
//
// Solidity: function get_underlying_balances(address _pool) view returns(uint256[8])
func (_Meta *MetaSession) GetUnderlyingBalances(_pool common.Address) ([8]*big.Int, error) {
	return _Meta.Contract.GetUnderlyingBalances(&_Meta.CallOpts, _pool)
}

// GetUnderlyingBalances is a free data retrieval call binding the contract method 0x59f4f351.
//
// Solidity: function get_underlying_balances(address _pool) view returns(uint256[8])
func (_Meta *MetaCallerSession) GetUnderlyingBalances(_pool common.Address) ([8]*big.Int, error) {
	return _Meta.Contract.GetUnderlyingBalances(&_Meta.CallOpts, _pool)
}

// GetUnderlyingBalances0 is a free data retrieval call binding the contract method 0x876112e6.
//
// Solidity: function get_underlying_balances(address _pool, uint256 _handler_id) view returns(uint256[8])
func (_Meta *MetaCaller) GetUnderlyingBalances0(opts *bind.CallOpts, _pool common.Address, _handler_id *big.Int) ([8]*big.Int, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_underlying_balances0", _pool, _handler_id)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetUnderlyingBalances0 is a free data retrieval call binding the contract method 0x876112e6.
//
// Solidity: function get_underlying_balances(address _pool, uint256 _handler_id) view returns(uint256[8])
func (_Meta *MetaSession) GetUnderlyingBalances0(_pool common.Address, _handler_id *big.Int) ([8]*big.Int, error) {
	return _Meta.Contract.GetUnderlyingBalances0(&_Meta.CallOpts, _pool, _handler_id)
}

// GetUnderlyingBalances0 is a free data retrieval call binding the contract method 0x876112e6.
//
// Solidity: function get_underlying_balances(address _pool, uint256 _handler_id) view returns(uint256[8])
func (_Meta *MetaCallerSession) GetUnderlyingBalances0(_pool common.Address, _handler_id *big.Int) ([8]*big.Int, error) {
	return _Meta.Contract.GetUnderlyingBalances0(&_Meta.CallOpts, _pool, _handler_id)
}

// GetUnderlyingCoins is a free data retrieval call binding the contract method 0xa77576ef.
//
// Solidity: function get_underlying_coins(address _pool) view returns(address[8])
func (_Meta *MetaCaller) GetUnderlyingCoins(opts *bind.CallOpts, _pool common.Address) ([8]common.Address, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_underlying_coins", _pool)

	if err != nil {
		return *new([8]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([8]common.Address)).(*[8]common.Address)

	return out0, err

}

// GetUnderlyingCoins is a free data retrieval call binding the contract method 0xa77576ef.
//
// Solidity: function get_underlying_coins(address _pool) view returns(address[8])
func (_Meta *MetaSession) GetUnderlyingCoins(_pool common.Address) ([8]common.Address, error) {
	return _Meta.Contract.GetUnderlyingCoins(&_Meta.CallOpts, _pool)
}

// GetUnderlyingCoins is a free data retrieval call binding the contract method 0xa77576ef.
//
// Solidity: function get_underlying_coins(address _pool) view returns(address[8])
func (_Meta *MetaCallerSession) GetUnderlyingCoins(_pool common.Address) ([8]common.Address, error) {
	return _Meta.Contract.GetUnderlyingCoins(&_Meta.CallOpts, _pool)
}

// GetUnderlyingCoins0 is a free data retrieval call binding the contract method 0x2fc90bcf.
//
// Solidity: function get_underlying_coins(address _pool, uint256 _handler_id) view returns(address[8])
func (_Meta *MetaCaller) GetUnderlyingCoins0(opts *bind.CallOpts, _pool common.Address, _handler_id *big.Int) ([8]common.Address, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_underlying_coins0", _pool, _handler_id)

	if err != nil {
		return *new([8]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([8]common.Address)).(*[8]common.Address)

	return out0, err

}

// GetUnderlyingCoins0 is a free data retrieval call binding the contract method 0x2fc90bcf.
//
// Solidity: function get_underlying_coins(address _pool, uint256 _handler_id) view returns(address[8])
func (_Meta *MetaSession) GetUnderlyingCoins0(_pool common.Address, _handler_id *big.Int) ([8]common.Address, error) {
	return _Meta.Contract.GetUnderlyingCoins0(&_Meta.CallOpts, _pool, _handler_id)
}

// GetUnderlyingCoins0 is a free data retrieval call binding the contract method 0x2fc90bcf.
//
// Solidity: function get_underlying_coins(address _pool, uint256 _handler_id) view returns(address[8])
func (_Meta *MetaCallerSession) GetUnderlyingCoins0(_pool common.Address, _handler_id *big.Int) ([8]common.Address, error) {
	return _Meta.Contract.GetUnderlyingCoins0(&_Meta.CallOpts, _pool, _handler_id)
}

// GetUnderlyingDecimals is a free data retrieval call binding the contract method 0x4cb088f1.
//
// Solidity: function get_underlying_decimals(address _pool) view returns(uint256[8])
func (_Meta *MetaCaller) GetUnderlyingDecimals(opts *bind.CallOpts, _pool common.Address) ([8]*big.Int, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_underlying_decimals", _pool)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetUnderlyingDecimals is a free data retrieval call binding the contract method 0x4cb088f1.
//
// Solidity: function get_underlying_decimals(address _pool) view returns(uint256[8])
func (_Meta *MetaSession) GetUnderlyingDecimals(_pool common.Address) ([8]*big.Int, error) {
	return _Meta.Contract.GetUnderlyingDecimals(&_Meta.CallOpts, _pool)
}

// GetUnderlyingDecimals is a free data retrieval call binding the contract method 0x4cb088f1.
//
// Solidity: function get_underlying_decimals(address _pool) view returns(uint256[8])
func (_Meta *MetaCallerSession) GetUnderlyingDecimals(_pool common.Address) ([8]*big.Int, error) {
	return _Meta.Contract.GetUnderlyingDecimals(&_Meta.CallOpts, _pool)
}

// GetUnderlyingDecimals0 is a free data retrieval call binding the contract method 0x622d64f4.
//
// Solidity: function get_underlying_decimals(address _pool, uint256 _handler_id) view returns(uint256[8])
func (_Meta *MetaCaller) GetUnderlyingDecimals0(opts *bind.CallOpts, _pool common.Address, _handler_id *big.Int) ([8]*big.Int, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_underlying_decimals0", _pool, _handler_id)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// GetUnderlyingDecimals0 is a free data retrieval call binding the contract method 0x622d64f4.
//
// Solidity: function get_underlying_decimals(address _pool, uint256 _handler_id) view returns(uint256[8])
func (_Meta *MetaSession) GetUnderlyingDecimals0(_pool common.Address, _handler_id *big.Int) ([8]*big.Int, error) {
	return _Meta.Contract.GetUnderlyingDecimals0(&_Meta.CallOpts, _pool, _handler_id)
}

// GetUnderlyingDecimals0 is a free data retrieval call binding the contract method 0x622d64f4.
//
// Solidity: function get_underlying_decimals(address _pool, uint256 _handler_id) view returns(uint256[8])
func (_Meta *MetaCallerSession) GetUnderlyingDecimals0(_pool common.Address, _handler_id *big.Int) ([8]*big.Int, error) {
	return _Meta.Contract.GetUnderlyingDecimals0(&_Meta.CallOpts, _pool, _handler_id)
}

// GetVirtualPriceFromLpToken is a free data retrieval call binding the contract method 0xc5b7074a.
//
// Solidity: function get_virtual_price_from_lp_token(address _token) view returns(uint256)
func (_Meta *MetaCaller) GetVirtualPriceFromLpToken(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_virtual_price_from_lp_token", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPriceFromLpToken is a free data retrieval call binding the contract method 0xc5b7074a.
//
// Solidity: function get_virtual_price_from_lp_token(address _token) view returns(uint256)
func (_Meta *MetaSession) GetVirtualPriceFromLpToken(_token common.Address) (*big.Int, error) {
	return _Meta.Contract.GetVirtualPriceFromLpToken(&_Meta.CallOpts, _token)
}

// GetVirtualPriceFromLpToken is a free data retrieval call binding the contract method 0xc5b7074a.
//
// Solidity: function get_virtual_price_from_lp_token(address _token) view returns(uint256)
func (_Meta *MetaCallerSession) GetVirtualPriceFromLpToken(_token common.Address) (*big.Int, error) {
	return _Meta.Contract.GetVirtualPriceFromLpToken(&_Meta.CallOpts, _token)
}

// GetVirtualPriceFromLpToken0 is a free data retrieval call binding the contract method 0x4460f66f.
//
// Solidity: function get_virtual_price_from_lp_token(address _token, uint256 _handler_id) view returns(uint256)
func (_Meta *MetaCaller) GetVirtualPriceFromLpToken0(opts *bind.CallOpts, _token common.Address, _handler_id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "get_virtual_price_from_lp_token0", _token, _handler_id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPriceFromLpToken0 is a free data retrieval call binding the contract method 0x4460f66f.
//
// Solidity: function get_virtual_price_from_lp_token(address _token, uint256 _handler_id) view returns(uint256)
func (_Meta *MetaSession) GetVirtualPriceFromLpToken0(_token common.Address, _handler_id *big.Int) (*big.Int, error) {
	return _Meta.Contract.GetVirtualPriceFromLpToken0(&_Meta.CallOpts, _token, _handler_id)
}

// GetVirtualPriceFromLpToken0 is a free data retrieval call binding the contract method 0x4460f66f.
//
// Solidity: function get_virtual_price_from_lp_token(address _token, uint256 _handler_id) view returns(uint256)
func (_Meta *MetaCallerSession) GetVirtualPriceFromLpToken0(_token common.Address, _handler_id *big.Int) (*big.Int, error) {
	return _Meta.Contract.GetVirtualPriceFromLpToken0(&_Meta.CallOpts, _token, _handler_id)
}

// IsMeta is a free data retrieval call binding the contract method 0xe4d332a9.
//
// Solidity: function is_meta(address _pool) view returns(bool)
func (_Meta *MetaCaller) IsMeta(opts *bind.CallOpts, _pool common.Address) (bool, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "is_meta", _pool)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMeta is a free data retrieval call binding the contract method 0xe4d332a9.
//
// Solidity: function is_meta(address _pool) view returns(bool)
func (_Meta *MetaSession) IsMeta(_pool common.Address) (bool, error) {
	return _Meta.Contract.IsMeta(&_Meta.CallOpts, _pool)
}

// IsMeta is a free data retrieval call binding the contract method 0xe4d332a9.
//
// Solidity: function is_meta(address _pool) view returns(bool)
func (_Meta *MetaCallerSession) IsMeta(_pool common.Address) (bool, error) {
	return _Meta.Contract.IsMeta(&_Meta.CallOpts, _pool)
}

// IsMeta0 is a free data retrieval call binding the contract method 0xa64485a1.
//
// Solidity: function is_meta(address _pool, uint256 _handler_id) view returns(bool)
func (_Meta *MetaCaller) IsMeta0(opts *bind.CallOpts, _pool common.Address, _handler_id *big.Int) (bool, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "is_meta0", _pool, _handler_id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMeta0 is a free data retrieval call binding the contract method 0xa64485a1.
//
// Solidity: function is_meta(address _pool, uint256 _handler_id) view returns(bool)
func (_Meta *MetaSession) IsMeta0(_pool common.Address, _handler_id *big.Int) (bool, error) {
	return _Meta.Contract.IsMeta0(&_Meta.CallOpts, _pool, _handler_id)
}

// IsMeta0 is a free data retrieval call binding the contract method 0xa64485a1.
//
// Solidity: function is_meta(address _pool, uint256 _handler_id) view returns(bool)
func (_Meta *MetaCallerSession) IsMeta0(_pool common.Address, _handler_id *big.Int) (bool, error) {
	return _Meta.Contract.IsMeta0(&_Meta.CallOpts, _pool, _handler_id)
}

// IsRegistered is a free data retrieval call binding the contract method 0x619ea806.
//
// Solidity: function is_registered(address _pool) view returns(bool)
func (_Meta *MetaCaller) IsRegistered(opts *bind.CallOpts, _pool common.Address) (bool, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "is_registered", _pool)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegistered is a free data retrieval call binding the contract method 0x619ea806.
//
// Solidity: function is_registered(address _pool) view returns(bool)
func (_Meta *MetaSession) IsRegistered(_pool common.Address) (bool, error) {
	return _Meta.Contract.IsRegistered(&_Meta.CallOpts, _pool)
}

// IsRegistered is a free data retrieval call binding the contract method 0x619ea806.
//
// Solidity: function is_registered(address _pool) view returns(bool)
func (_Meta *MetaCallerSession) IsRegistered(_pool common.Address) (bool, error) {
	return _Meta.Contract.IsRegistered(&_Meta.CallOpts, _pool)
}

// IsRegistered0 is a free data retrieval call binding the contract method 0xc9c4f3ec.
//
// Solidity: function is_registered(address _pool, uint256 _handler_id) view returns(bool)
func (_Meta *MetaCaller) IsRegistered0(opts *bind.CallOpts, _pool common.Address, _handler_id *big.Int) (bool, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "is_registered0", _pool, _handler_id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegistered0 is a free data retrieval call binding the contract method 0xc9c4f3ec.
//
// Solidity: function is_registered(address _pool, uint256 _handler_id) view returns(bool)
func (_Meta *MetaSession) IsRegistered0(_pool common.Address, _handler_id *big.Int) (bool, error) {
	return _Meta.Contract.IsRegistered0(&_Meta.CallOpts, _pool, _handler_id)
}

// IsRegistered0 is a free data retrieval call binding the contract method 0xc9c4f3ec.
//
// Solidity: function is_registered(address _pool, uint256 _handler_id) view returns(bool)
func (_Meta *MetaCallerSession) IsRegistered0(_pool common.Address, _handler_id *big.Int) (bool, error) {
	return _Meta.Contract.IsRegistered0(&_Meta.CallOpts, _pool, _handler_id)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Meta *MetaCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Meta *MetaSession) Owner() (common.Address, error) {
	return _Meta.Contract.Owner(&_Meta.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Meta *MetaCallerSession) Owner() (common.Address, error) {
	return _Meta.Contract.Owner(&_Meta.CallOpts)
}

// PoolCount is a free data retrieval call binding the contract method 0x956aae3a.
//
// Solidity: function pool_count() view returns(uint256)
func (_Meta *MetaCaller) PoolCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "pool_count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolCount is a free data retrieval call binding the contract method 0x956aae3a.
//
// Solidity: function pool_count() view returns(uint256)
func (_Meta *MetaSession) PoolCount() (*big.Int, error) {
	return _Meta.Contract.PoolCount(&_Meta.CallOpts)
}

// PoolCount is a free data retrieval call binding the contract method 0x956aae3a.
//
// Solidity: function pool_count() view returns(uint256)
func (_Meta *MetaCallerSession) PoolCount() (*big.Int, error) {
	return _Meta.Contract.PoolCount(&_Meta.CallOpts)
}

// PoolList is a free data retrieval call binding the contract method 0x3a1d5d8e.
//
// Solidity: function pool_list(uint256 _index) view returns(address)
func (_Meta *MetaCaller) PoolList(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "pool_list", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolList is a free data retrieval call binding the contract method 0x3a1d5d8e.
//
// Solidity: function pool_list(uint256 _index) view returns(address)
func (_Meta *MetaSession) PoolList(_index *big.Int) (common.Address, error) {
	return _Meta.Contract.PoolList(&_Meta.CallOpts, _index)
}

// PoolList is a free data retrieval call binding the contract method 0x3a1d5d8e.
//
// Solidity: function pool_list(uint256 _index) view returns(address)
func (_Meta *MetaCallerSession) PoolList(_index *big.Int) (common.Address, error) {
	return _Meta.Contract.PoolList(&_Meta.CallOpts, _index)
}

// RegistryLength is a free data retrieval call binding the contract method 0x083297d2.
//
// Solidity: function registry_length() view returns(uint256)
func (_Meta *MetaCaller) RegistryLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Meta.contract.Call(opts, &out, "registry_length")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RegistryLength is a free data retrieval call binding the contract method 0x083297d2.
//
// Solidity: function registry_length() view returns(uint256)
func (_Meta *MetaSession) RegistryLength() (*big.Int, error) {
	return _Meta.Contract.RegistryLength(&_Meta.CallOpts)
}

// RegistryLength is a free data retrieval call binding the contract method 0x083297d2.
//
// Solidity: function registry_length() view returns(uint256)
func (_Meta *MetaCallerSession) RegistryLength() (*big.Int, error) {
	return _Meta.Contract.RegistryLength(&_Meta.CallOpts)
}

// AddRegistryHandler is a paid mutator transaction binding the contract method 0x22f595c8.
//
// Solidity: function add_registry_handler(address _registry_handler) returns()
func (_Meta *MetaTransactor) AddRegistryHandler(opts *bind.TransactOpts, _registry_handler common.Address) (*types.Transaction, error) {
	return _Meta.contract.Transact(opts, "add_registry_handler", _registry_handler)
}

// AddRegistryHandler is a paid mutator transaction binding the contract method 0x22f595c8.
//
// Solidity: function add_registry_handler(address _registry_handler) returns()
func (_Meta *MetaSession) AddRegistryHandler(_registry_handler common.Address) (*types.Transaction, error) {
	return _Meta.Contract.AddRegistryHandler(&_Meta.TransactOpts, _registry_handler)
}

// AddRegistryHandler is a paid mutator transaction binding the contract method 0x22f595c8.
//
// Solidity: function add_registry_handler(address _registry_handler) returns()
func (_Meta *MetaTransactorSession) AddRegistryHandler(_registry_handler common.Address) (*types.Transaction, error) {
	return _Meta.Contract.AddRegistryHandler(&_Meta.TransactOpts, _registry_handler)
}

// UpdateRegistryHandler is a paid mutator transaction binding the contract method 0x0dbdc9ff.
//
// Solidity: function update_registry_handler(uint256 _index, address _registry_handler) returns()
func (_Meta *MetaTransactor) UpdateRegistryHandler(opts *bind.TransactOpts, _index *big.Int, _registry_handler common.Address) (*types.Transaction, error) {
	return _Meta.contract.Transact(opts, "update_registry_handler", _index, _registry_handler)
}

// UpdateRegistryHandler is a paid mutator transaction binding the contract method 0x0dbdc9ff.
//
// Solidity: function update_registry_handler(uint256 _index, address _registry_handler) returns()
func (_Meta *MetaSession) UpdateRegistryHandler(_index *big.Int, _registry_handler common.Address) (*types.Transaction, error) {
	return _Meta.Contract.UpdateRegistryHandler(&_Meta.TransactOpts, _index, _registry_handler)
}

// UpdateRegistryHandler is a paid mutator transaction binding the contract method 0x0dbdc9ff.
//
// Solidity: function update_registry_handler(uint256 _index, address _registry_handler) returns()
func (_Meta *MetaTransactorSession) UpdateRegistryHandler(_index *big.Int, _registry_handler common.Address) (*types.Transaction, error) {
	return _Meta.Contract.UpdateRegistryHandler(&_Meta.TransactOpts, _index, _registry_handler)
}

// MetaCommitNewAdminIterator is returned from FilterCommitNewAdmin and is used to iterate over the raw logs and unpacked data for CommitNewAdmin events raised by the Meta contract.
type MetaCommitNewAdminIterator struct {
	Event *MetaCommitNewAdmin // Event containing the contract specifics and raw log

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
func (it *MetaCommitNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaCommitNewAdmin)
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
		it.Event = new(MetaCommitNewAdmin)
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
func (it *MetaCommitNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaCommitNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaCommitNewAdmin represents a CommitNewAdmin event raised by the Meta contract.
type MetaCommitNewAdmin struct {
	Deadline *big.Int
	Admin    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCommitNewAdmin is a free log retrieval operation binding the contract event 0x181aa3aa17d4cbf99265dd4443eba009433d3cde79d60164fde1d1a192beb935.
//
// Solidity: event CommitNewAdmin(uint256 indexed deadline, address indexed admin)
func (_Meta *MetaFilterer) FilterCommitNewAdmin(opts *bind.FilterOpts, deadline []*big.Int, admin []common.Address) (*MetaCommitNewAdminIterator, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Meta.contract.FilterLogs(opts, "CommitNewAdmin", deadlineRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &MetaCommitNewAdminIterator{contract: _Meta.contract, event: "CommitNewAdmin", logs: logs, sub: sub}, nil
}

// WatchCommitNewAdmin is a free log subscription operation binding the contract event 0x181aa3aa17d4cbf99265dd4443eba009433d3cde79d60164fde1d1a192beb935.
//
// Solidity: event CommitNewAdmin(uint256 indexed deadline, address indexed admin)
func (_Meta *MetaFilterer) WatchCommitNewAdmin(opts *bind.WatchOpts, sink chan<- *MetaCommitNewAdmin, deadline []*big.Int, admin []common.Address) (event.Subscription, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Meta.contract.WatchLogs(opts, "CommitNewAdmin", deadlineRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaCommitNewAdmin)
				if err := _Meta.contract.UnpackLog(event, "CommitNewAdmin", log); err != nil {
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

// ParseCommitNewAdmin is a log parse operation binding the contract event 0x181aa3aa17d4cbf99265dd4443eba009433d3cde79d60164fde1d1a192beb935.
//
// Solidity: event CommitNewAdmin(uint256 indexed deadline, address indexed admin)
func (_Meta *MetaFilterer) ParseCommitNewAdmin(log types.Log) (*MetaCommitNewAdmin, error) {
	event := new(MetaCommitNewAdmin)
	if err := _Meta.contract.UnpackLog(event, "CommitNewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the Meta contract.
type MetaNewAdminIterator struct {
	Event *MetaNewAdmin // Event containing the contract specifics and raw log

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
func (it *MetaNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaNewAdmin)
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
		it.Event = new(MetaNewAdmin)
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
func (it *MetaNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaNewAdmin represents a NewAdmin event raised by the Meta contract.
type MetaNewAdmin struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address indexed admin)
func (_Meta *MetaFilterer) FilterNewAdmin(opts *bind.FilterOpts, admin []common.Address) (*MetaNewAdminIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Meta.contract.FilterLogs(opts, "NewAdmin", adminRule)
	if err != nil {
		return nil, err
	}
	return &MetaNewAdminIterator{contract: _Meta.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address indexed admin)
func (_Meta *MetaFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *MetaNewAdmin, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Meta.contract.WatchLogs(opts, "NewAdmin", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaNewAdmin)
				if err := _Meta.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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

// ParseNewAdmin is a log parse operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address indexed admin)
func (_Meta *MetaFilterer) ParseNewAdmin(log types.Log) (*MetaNewAdmin, error) {
	event := new(MetaNewAdmin)
	if err := _Meta.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
