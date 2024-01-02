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

// Tricrypto2MetaData contains all meta data concerning the Tricrypto2 contract.
var Tricrypto2MetaData = &bind.MetaData{
	ABI: "[{\"name\":\"TokenExchange\",\"inputs\":[{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":true},{\"name\":\"sold_id\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"tokens_sold\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"bought_id\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"tokens_bought\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"AddLiquidity\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amounts\",\"type\":\"uint256[3]\",\"indexed\":false},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidity\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amounts\",\"type\":\"uint256[3]\",\"indexed\":false},{\"name\":\"token_supply\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidityOne\",\"inputs\":[{\"name\":\"provider\",\"type\":\"address\",\"indexed\":true},{\"name\":\"token_amount\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"coin_index\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"coin_amount\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CommitNewAdmin\",\"inputs\":[{\"name\":\"deadline\",\"type\":\"uint256\",\"indexed\":true},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewAdmin\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"CommitNewParameters\",\"inputs\":[{\"name\":\"deadline\",\"type\":\"uint256\",\"indexed\":true},{\"name\":\"admin_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"mid_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"out_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"fee_gamma\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"allowed_extra_profit\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"adjustment_step\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"ma_half_time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"NewParameters\",\"inputs\":[{\"name\":\"admin_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"mid_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"out_fee\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"fee_gamma\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"allowed_extra_profit\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"adjustment_step\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"ma_half_time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RampAgamma\",\"inputs\":[{\"name\":\"initial_A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"future_A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"initial_gamma\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"future_gamma\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"initial_time\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"future_time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"StopRampA\",\"inputs\":[{\"name\":\"current_A\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"current_gamma\",\"type\":\"uint256\",\"indexed\":false},{\"name\":\"time\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"ClaimAdminFee\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true},{\"name\":\"tokens\",\"type\":\"uint256\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"admin_fee_receiver\",\"type\":\"address\"},{\"name\":\"A\",\"type\":\"uint256\"},{\"name\":\"gamma\",\"type\":\"uint256\"},{\"name\":\"mid_fee\",\"type\":\"uint256\"},{\"name\":\"out_fee\",\"type\":\"uint256\"},{\"name\":\"allowed_extra_profit\",\"type\":\"uint256\"},{\"name\":\"fee_gamma\",\"type\":\"uint256\"},{\"name\":\"adjustment_step\",\"type\":\"uint256\"},{\"name\":\"admin_fee\",\"type\":\"uint256\"},{\"name\":\"ma_half_time\",\"type\":\"uint256\"},{\"name\":\"initial_prices\",\"type\":\"uint256[2]\"}],\"outputs\":[]},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"price_oracle\",\"inputs\":[{\"name\":\"k\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3361},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"price_scale\",\"inputs\":[{\"name\":\"k\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3391},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"last_prices\",\"inputs\":[{\"name\":\"k\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3421},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":468},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"coins\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":582},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"A\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":597},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"gamma\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":11991},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":21673},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"fee_calc\",\"inputs\":[{\"name\":\"xp\",\"type\":\"uint256[3]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":11096},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_virtual_price\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":11582},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"exchange\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"dx\",\"type\":\"uint256\"},{\"name\":\"min_dy\",\"type\":\"uint256\"}],\"outputs\":[]},{\"stateMutability\":\"payable\",\"type\":\"function\",\"name\":\"exchange\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"dx\",\"type\":\"uint256\"},{\"name\":\"min_dy\",\"type\":\"uint256\"},{\"name\":\"use_eth\",\"type\":\"bool\"}],\"outputs\":[]},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"get_dy\",\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"j\",\"type\":\"uint256\"},{\"name\":\"dx\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3122},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"calc_token_fee\",\"inputs\":[{\"name\":\"amounts\",\"type\":\"uint256[3]\"},{\"name\":\"xp\",\"type\":\"uint256[3]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":26582},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"add_liquidity\",\"inputs\":[{\"name\":\"amounts\",\"type\":\"uint256[3]\"},{\"name\":\"min_mint_amount\",\"type\":\"uint256\"}],\"outputs\":[],\"gas\":738687},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"min_amounts\",\"type\":\"uint256[3]\"}],\"outputs\":[],\"gas\":233981},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"calc_token_amount\",\"inputs\":[{\"name\":\"amounts\",\"type\":\"uint256[3]\"},{\"name\":\"deposit\",\"type\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3429},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"calc_withdraw_one_coin\",\"inputs\":[{\"name\":\"token_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":13432},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"remove_liquidity_one_coin\",\"inputs\":[{\"name\":\"token_amount\",\"type\":\"uint256\"},{\"name\":\"i\",\"type\":\"uint256\"},{\"name\":\"min_amount\",\"type\":\"uint256\"}],\"outputs\":[],\"gas\":648579},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"claim_admin_fees\",\"inputs\":[],\"outputs\":[],\"gas\":389808},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"ramp_A_gamma\",\"inputs\":[{\"name\":\"future_A\",\"type\":\"uint256\"},{\"name\":\"future_gamma\",\"type\":\"uint256\"},{\"name\":\"future_time\",\"type\":\"uint256\"}],\"outputs\":[],\"gas\":163102},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"stop_ramp_A_gamma\",\"inputs\":[],\"outputs\":[],\"gas\":157247},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"commit_new_parameters\",\"inputs\":[{\"name\":\"_new_mid_fee\",\"type\":\"uint256\"},{\"name\":\"_new_out_fee\",\"type\":\"uint256\"},{\"name\":\"_new_admin_fee\",\"type\":\"uint256\"},{\"name\":\"_new_fee_gamma\",\"type\":\"uint256\"},{\"name\":\"_new_allowed_extra_profit\",\"type\":\"uint256\"},{\"name\":\"_new_adjustment_step\",\"type\":\"uint256\"},{\"name\":\"_new_ma_half_time\",\"type\":\"uint256\"}],\"outputs\":[],\"gas\":306190},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"apply_new_parameters\",\"inputs\":[],\"outputs\":[],\"gas\":683438},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"revert_new_parameters\",\"inputs\":[],\"outputs\":[],\"gas\":23222},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"commit_transfer_ownership\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"outputs\":[],\"gas\":77260},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"apply_transfer_ownership\",\"inputs\":[],\"outputs\":[],\"gas\":65937},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"revert_transfer_ownership\",\"inputs\":[],\"outputs\":[],\"gas\":23312},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"kill_me\",\"inputs\":[],\"outputs\":[],\"gas\":40535},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"unkill_me\",\"inputs\":[],\"outputs\":[],\"gas\":23372},{\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"name\":\"set_admin_fee_receiver\",\"inputs\":[{\"name\":\"_admin_fee_receiver\",\"type\":\"address\"}],\"outputs\":[],\"gas\":38505},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"last_prices_timestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3378},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"initial_A_gamma\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3408},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_A_gamma\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3438},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"initial_A_gamma_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3468},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_A_gamma_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3498},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"allowed_extra_profit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3528},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_allowed_extra_profit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3558},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"fee_gamma\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3588},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_fee_gamma\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3618},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"adjustment_step\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3648},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_adjustment_step\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3678},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"ma_half_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3708},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_ma_half_time\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3738},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"mid_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3768},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"out_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3798},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3828},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_mid_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3858},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_out_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3888},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_admin_fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3918},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"balances\",\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":4057},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"D\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":3978},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":4008},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"future_owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":4038},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"xcp_profit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":4068},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"xcp_profit_a\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":4098},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"virtual_price\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":4128},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"is_killed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"gas\":4158},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"kill_deadline\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":4188},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"transfer_ownership_deadline\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":4218},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin_actions_deadline\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"gas\":4248},{\"stateMutability\":\"view\",\"type\":\"function\",\"name\":\"admin_fee_receiver\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"gas\":4278}]",
}

// Tricrypto2ABI is the input ABI used to generate the binding from.
// Deprecated: Use Tricrypto2MetaData.ABI instead.
var Tricrypto2ABI = Tricrypto2MetaData.ABI

// Tricrypto2 is an auto generated Go binding around an Ethereum contract.
type Tricrypto2 struct {
	Tricrypto2Caller     // Read-only binding to the contract
	Tricrypto2Transactor // Write-only binding to the contract
	Tricrypto2Filterer   // Log filterer for contract events
}

// Tricrypto2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Tricrypto2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Tricrypto2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Tricrypto2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Tricrypto2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Tricrypto2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Tricrypto2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Tricrypto2Session struct {
	Contract     *Tricrypto2       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Tricrypto2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Tricrypto2CallerSession struct {
	Contract *Tricrypto2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// Tricrypto2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Tricrypto2TransactorSession struct {
	Contract     *Tricrypto2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Tricrypto2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Tricrypto2Raw struct {
	Contract *Tricrypto2 // Generic contract binding to access the raw methods on
}

// Tricrypto2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Tricrypto2CallerRaw struct {
	Contract *Tricrypto2Caller // Generic read-only contract binding to access the raw methods on
}

// Tricrypto2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Tricrypto2TransactorRaw struct {
	Contract *Tricrypto2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTricrypto2 creates a new instance of Tricrypto2, bound to a specific deployed contract.
func NewTricrypto2(address common.Address, backend bind.ContractBackend) (*Tricrypto2, error) {
	contract, err := bindTricrypto2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tricrypto2{Tricrypto2Caller: Tricrypto2Caller{contract: contract}, Tricrypto2Transactor: Tricrypto2Transactor{contract: contract}, Tricrypto2Filterer: Tricrypto2Filterer{contract: contract}}, nil
}

// NewTricrypto2Caller creates a new read-only instance of Tricrypto2, bound to a specific deployed contract.
func NewTricrypto2Caller(address common.Address, caller bind.ContractCaller) (*Tricrypto2Caller, error) {
	contract, err := bindTricrypto2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Tricrypto2Caller{contract: contract}, nil
}

// NewTricrypto2Transactor creates a new write-only instance of Tricrypto2, bound to a specific deployed contract.
func NewTricrypto2Transactor(address common.Address, transactor bind.ContractTransactor) (*Tricrypto2Transactor, error) {
	contract, err := bindTricrypto2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Tricrypto2Transactor{contract: contract}, nil
}

// NewTricrypto2Filterer creates a new log filterer instance of Tricrypto2, bound to a specific deployed contract.
func NewTricrypto2Filterer(address common.Address, filterer bind.ContractFilterer) (*Tricrypto2Filterer, error) {
	contract, err := bindTricrypto2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Tricrypto2Filterer{contract: contract}, nil
}

// bindTricrypto2 binds a generic wrapper to an already deployed contract.
func bindTricrypto2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Tricrypto2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tricrypto2 *Tricrypto2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tricrypto2.Contract.Tricrypto2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tricrypto2 *Tricrypto2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tricrypto2.Contract.Tricrypto2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tricrypto2 *Tricrypto2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tricrypto2.Contract.Tricrypto2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tricrypto2 *Tricrypto2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tricrypto2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tricrypto2 *Tricrypto2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tricrypto2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tricrypto2 *Tricrypto2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tricrypto2.Contract.contract.Transact(opts, method, params...)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Caller) A(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Session) A() (*big.Int, error) {
	return _Tricrypto2.Contract.A(&_Tricrypto2.CallOpts)
}

// A is a free data retrieval call binding the contract method 0xf446c1d0.
//
// Solidity: function A() view returns(uint256)
func (_Tricrypto2 *Tricrypto2CallerSession) A() (*big.Int, error) {
	return _Tricrypto2.Contract.A(&_Tricrypto2.CallOpts)
}

// D is a free data retrieval call binding the contract method 0x0f529ba2.
//
// Solidity: function D() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Caller) D(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "D")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// D is a free data retrieval call binding the contract method 0x0f529ba2.
//
// Solidity: function D() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Session) D() (*big.Int, error) {
	return _Tricrypto2.Contract.D(&_Tricrypto2.CallOpts)
}

// D is a free data retrieval call binding the contract method 0x0f529ba2.
//
// Solidity: function D() view returns(uint256)
func (_Tricrypto2 *Tricrypto2CallerSession) D() (*big.Int, error) {
	return _Tricrypto2.Contract.D(&_Tricrypto2.CallOpts)
}

// AdjustmentStep is a free data retrieval call binding the contract method 0x083812e5.
//
// Solidity: function adjustment_step() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Caller) AdjustmentStep(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "adjustment_step")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdjustmentStep is a free data retrieval call binding the contract method 0x083812e5.
//
// Solidity: function adjustment_step() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Session) AdjustmentStep() (*big.Int, error) {
	return _Tricrypto2.Contract.AdjustmentStep(&_Tricrypto2.CallOpts)
}

// AdjustmentStep is a free data retrieval call binding the contract method 0x083812e5.
//
// Solidity: function adjustment_step() view returns(uint256)
func (_Tricrypto2 *Tricrypto2CallerSession) AdjustmentStep() (*big.Int, error) {
	return _Tricrypto2.Contract.AdjustmentStep(&_Tricrypto2.CallOpts)
}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Caller) AdminActionsDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "admin_actions_deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Session) AdminActionsDeadline() (*big.Int, error) {
	return _Tricrypto2.Contract.AdminActionsDeadline(&_Tricrypto2.CallOpts)
}

// AdminActionsDeadline is a free data retrieval call binding the contract method 0x405e28f8.
//
// Solidity: function admin_actions_deadline() view returns(uint256)
func (_Tricrypto2 *Tricrypto2CallerSession) AdminActionsDeadline() (*big.Int, error) {
	return _Tricrypto2.Contract.AdminActionsDeadline(&_Tricrypto2.CallOpts)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Caller) AdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "admin_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Session) AdminFee() (*big.Int, error) {
	return _Tricrypto2.Contract.AdminFee(&_Tricrypto2.CallOpts)
}

// AdminFee is a free data retrieval call binding the contract method 0xfee3f7f9.
//
// Solidity: function admin_fee() view returns(uint256)
func (_Tricrypto2 *Tricrypto2CallerSession) AdminFee() (*big.Int, error) {
	return _Tricrypto2.Contract.AdminFee(&_Tricrypto2.CallOpts)
}

// AdminFeeReceiver is a free data retrieval call binding the contract method 0x6e42e4d2.
//
// Solidity: function admin_fee_receiver() view returns(address)
func (_Tricrypto2 *Tricrypto2Caller) AdminFeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "admin_fee_receiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdminFeeReceiver is a free data retrieval call binding the contract method 0x6e42e4d2.
//
// Solidity: function admin_fee_receiver() view returns(address)
func (_Tricrypto2 *Tricrypto2Session) AdminFeeReceiver() (common.Address, error) {
	return _Tricrypto2.Contract.AdminFeeReceiver(&_Tricrypto2.CallOpts)
}

// AdminFeeReceiver is a free data retrieval call binding the contract method 0x6e42e4d2.
//
// Solidity: function admin_fee_receiver() view returns(address)
func (_Tricrypto2 *Tricrypto2CallerSession) AdminFeeReceiver() (common.Address, error) {
	return _Tricrypto2.Contract.AdminFeeReceiver(&_Tricrypto2.CallOpts)
}

// AllowedExtraProfit is a free data retrieval call binding the contract method 0x49fe9e77.
//
// Solidity: function allowed_extra_profit() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Caller) AllowedExtraProfit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "allowed_extra_profit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllowedExtraProfit is a free data retrieval call binding the contract method 0x49fe9e77.
//
// Solidity: function allowed_extra_profit() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Session) AllowedExtraProfit() (*big.Int, error) {
	return _Tricrypto2.Contract.AllowedExtraProfit(&_Tricrypto2.CallOpts)
}

// AllowedExtraProfit is a free data retrieval call binding the contract method 0x49fe9e77.
//
// Solidity: function allowed_extra_profit() view returns(uint256)
func (_Tricrypto2 *Tricrypto2CallerSession) AllowedExtraProfit() (*big.Int, error) {
	return _Tricrypto2.Contract.AllowedExtraProfit(&_Tricrypto2.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_Tricrypto2 *Tricrypto2Caller) Balances(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_Tricrypto2 *Tricrypto2Session) Balances(arg0 *big.Int) (*big.Int, error) {
	return _Tricrypto2.Contract.Balances(&_Tricrypto2.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x4903b0d1.
//
// Solidity: function balances(uint256 arg0) view returns(uint256)
func (_Tricrypto2 *Tricrypto2CallerSession) Balances(arg0 *big.Int) (*big.Int, error) {
	return _Tricrypto2.Contract.Balances(&_Tricrypto2.CallOpts, arg0)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x3883e119.
//
// Solidity: function calc_token_amount(uint256[3] amounts, bool deposit) view returns(uint256)
func (_Tricrypto2 *Tricrypto2Caller) CalcTokenAmount(opts *bind.CallOpts, amounts [3]*big.Int, deposit bool) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "calc_token_amount", amounts, deposit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x3883e119.
//
// Solidity: function calc_token_amount(uint256[3] amounts, bool deposit) view returns(uint256)
func (_Tricrypto2 *Tricrypto2Session) CalcTokenAmount(amounts [3]*big.Int, deposit bool) (*big.Int, error) {
	return _Tricrypto2.Contract.CalcTokenAmount(&_Tricrypto2.CallOpts, amounts, deposit)
}

// CalcTokenAmount is a free data retrieval call binding the contract method 0x3883e119.
//
// Solidity: function calc_token_amount(uint256[3] amounts, bool deposit) view returns(uint256)
func (_Tricrypto2 *Tricrypto2CallerSession) CalcTokenAmount(amounts [3]*big.Int, deposit bool) (*big.Int, error) {
	return _Tricrypto2.Contract.CalcTokenAmount(&_Tricrypto2.CallOpts, amounts, deposit)
}

// CalcTokenFee is a free data retrieval call binding the contract method 0xcde699fa.
//
// Solidity: function calc_token_fee(uint256[3] amounts, uint256[3] xp) view returns(uint256)
func (_Tricrypto2 *Tricrypto2Caller) CalcTokenFee(opts *bind.CallOpts, amounts [3]*big.Int, xp [3]*big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "calc_token_fee", amounts, xp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcTokenFee is a free data retrieval call binding the contract method 0xcde699fa.
//
// Solidity: function calc_token_fee(uint256[3] amounts, uint256[3] xp) view returns(uint256)
func (_Tricrypto2 *Tricrypto2Session) CalcTokenFee(amounts [3]*big.Int, xp [3]*big.Int) (*big.Int, error) {
	return _Tricrypto2.Contract.CalcTokenFee(&_Tricrypto2.CallOpts, amounts, xp)
}

// CalcTokenFee is a free data retrieval call binding the contract method 0xcde699fa.
//
// Solidity: function calc_token_fee(uint256[3] amounts, uint256[3] xp) view returns(uint256)
func (_Tricrypto2 *Tricrypto2CallerSession) CalcTokenFee(amounts [3]*big.Int, xp [3]*big.Int) (*big.Int, error) {
	return _Tricrypto2.Contract.CalcTokenFee(&_Tricrypto2.CallOpts, amounts, xp)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0x4fb08c5e.
//
// Solidity: function calc_withdraw_one_coin(uint256 token_amount, uint256 i) view returns(uint256)
func (_Tricrypto2 *Tricrypto2Caller) CalcWithdrawOneCoin(opts *bind.CallOpts, token_amount *big.Int, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "calc_withdraw_one_coin", token_amount, i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0x4fb08c5e.
//
// Solidity: function calc_withdraw_one_coin(uint256 token_amount, uint256 i) view returns(uint256)
func (_Tricrypto2 *Tricrypto2Session) CalcWithdrawOneCoin(token_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _Tricrypto2.Contract.CalcWithdrawOneCoin(&_Tricrypto2.CallOpts, token_amount, i)
}

// CalcWithdrawOneCoin is a free data retrieval call binding the contract method 0x4fb08c5e.
//
// Solidity: function calc_withdraw_one_coin(uint256 token_amount, uint256 i) view returns(uint256)
func (_Tricrypto2 *Tricrypto2CallerSession) CalcWithdrawOneCoin(token_amount *big.Int, i *big.Int) (*big.Int, error) {
	return _Tricrypto2.Contract.CalcWithdrawOneCoin(&_Tricrypto2.CallOpts, token_amount, i)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 i) view returns(address)
func (_Tricrypto2 *Tricrypto2Caller) Coins(opts *bind.CallOpts, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "coins", i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 i) view returns(address)
func (_Tricrypto2 *Tricrypto2Session) Coins(i *big.Int) (common.Address, error) {
	return _Tricrypto2.Contract.Coins(&_Tricrypto2.CallOpts, i)
}

// Coins is a free data retrieval call binding the contract method 0xc6610657.
//
// Solidity: function coins(uint256 i) view returns(address)
func (_Tricrypto2 *Tricrypto2CallerSession) Coins(i *big.Int) (common.Address, error) {
	return _Tricrypto2.Contract.Coins(&_Tricrypto2.CallOpts, i)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Caller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Session) Fee() (*big.Int, error) {
	return _Tricrypto2.Contract.Fee(&_Tricrypto2.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Tricrypto2 *Tricrypto2CallerSession) Fee() (*big.Int, error) {
	return _Tricrypto2.Contract.Fee(&_Tricrypto2.CallOpts)
}

// FeeCalc is a free data retrieval call binding the contract method 0x572e5625.
//
// Solidity: function fee_calc(uint256[3] xp) view returns(uint256)
func (_Tricrypto2 *Tricrypto2Caller) FeeCalc(opts *bind.CallOpts, xp [3]*big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "fee_calc", xp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeCalc is a free data retrieval call binding the contract method 0x572e5625.
//
// Solidity: function fee_calc(uint256[3] xp) view returns(uint256)
func (_Tricrypto2 *Tricrypto2Session) FeeCalc(xp [3]*big.Int) (*big.Int, error) {
	return _Tricrypto2.Contract.FeeCalc(&_Tricrypto2.CallOpts, xp)
}

// FeeCalc is a free data retrieval call binding the contract method 0x572e5625.
//
// Solidity: function fee_calc(uint256[3] xp) view returns(uint256)
func (_Tricrypto2 *Tricrypto2CallerSession) FeeCalc(xp [3]*big.Int) (*big.Int, error) {
	return _Tricrypto2.Contract.FeeCalc(&_Tricrypto2.CallOpts, xp)
}

// FeeGamma is a free data retrieval call binding the contract method 0x72d4f0e2.
//
// Solidity: function fee_gamma() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Caller) FeeGamma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "fee_gamma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeGamma is a free data retrieval call binding the contract method 0x72d4f0e2.
//
// Solidity: function fee_gamma() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Session) FeeGamma() (*big.Int, error) {
	return _Tricrypto2.Contract.FeeGamma(&_Tricrypto2.CallOpts)
}

// FeeGamma is a free data retrieval call binding the contract method 0x72d4f0e2.
//
// Solidity: function fee_gamma() view returns(uint256)
func (_Tricrypto2 *Tricrypto2CallerSession) FeeGamma() (*big.Int, error) {
	return _Tricrypto2.Contract.FeeGamma(&_Tricrypto2.CallOpts)
}

// FutureAGamma is a free data retrieval call binding the contract method 0xf30cfad5.
//
// Solidity: function future_A_gamma() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Caller) FutureAGamma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "future_A_gamma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAGamma is a free data retrieval call binding the contract method 0xf30cfad5.
//
// Solidity: function future_A_gamma() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Session) FutureAGamma() (*big.Int, error) {
	return _Tricrypto2.Contract.FutureAGamma(&_Tricrypto2.CallOpts)
}

// FutureAGamma is a free data retrieval call binding the contract method 0xf30cfad5.
//
// Solidity: function future_A_gamma() view returns(uint256)
func (_Tricrypto2 *Tricrypto2CallerSession) FutureAGamma() (*big.Int, error) {
	return _Tricrypto2.Contract.FutureAGamma(&_Tricrypto2.CallOpts)
}

// FutureAGammaTime is a free data retrieval call binding the contract method 0xf9ed9597.
//
// Solidity: function future_A_gamma_time() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Caller) FutureAGammaTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "future_A_gamma_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAGammaTime is a free data retrieval call binding the contract method 0xf9ed9597.
//
// Solidity: function future_A_gamma_time() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Session) FutureAGammaTime() (*big.Int, error) {
	return _Tricrypto2.Contract.FutureAGammaTime(&_Tricrypto2.CallOpts)
}

// FutureAGammaTime is a free data retrieval call binding the contract method 0xf9ed9597.
//
// Solidity: function future_A_gamma_time() view returns(uint256)
func (_Tricrypto2 *Tricrypto2CallerSession) FutureAGammaTime() (*big.Int, error) {
	return _Tricrypto2.Contract.FutureAGammaTime(&_Tricrypto2.CallOpts)
}

// FutureAdjustmentStep is a free data retrieval call binding the contract method 0x4ea12c7d.
//
// Solidity: function future_adjustment_step() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Caller) FutureAdjustmentStep(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "future_adjustment_step")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAdjustmentStep is a free data retrieval call binding the contract method 0x4ea12c7d.
//
// Solidity: function future_adjustment_step() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Session) FutureAdjustmentStep() (*big.Int, error) {
	return _Tricrypto2.Contract.FutureAdjustmentStep(&_Tricrypto2.CallOpts)
}

// FutureAdjustmentStep is a free data retrieval call binding the contract method 0x4ea12c7d.
//
// Solidity: function future_adjustment_step() view returns(uint256)
func (_Tricrypto2 *Tricrypto2CallerSession) FutureAdjustmentStep() (*big.Int, error) {
	return _Tricrypto2.Contract.FutureAdjustmentStep(&_Tricrypto2.CallOpts)
}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Caller) FutureAdminFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "future_admin_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Session) FutureAdminFee() (*big.Int, error) {
	return _Tricrypto2.Contract.FutureAdminFee(&_Tricrypto2.CallOpts)
}

// FutureAdminFee is a free data retrieval call binding the contract method 0xe3824462.
//
// Solidity: function future_admin_fee() view returns(uint256)
func (_Tricrypto2 *Tricrypto2CallerSession) FutureAdminFee() (*big.Int, error) {
	return _Tricrypto2.Contract.FutureAdminFee(&_Tricrypto2.CallOpts)
}

// FutureAllowedExtraProfit is a free data retrieval call binding the contract method 0x727ced57.
//
// Solidity: function future_allowed_extra_profit() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Caller) FutureAllowedExtraProfit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "future_allowed_extra_profit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureAllowedExtraProfit is a free data retrieval call binding the contract method 0x727ced57.
//
// Solidity: function future_allowed_extra_profit() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Session) FutureAllowedExtraProfit() (*big.Int, error) {
	return _Tricrypto2.Contract.FutureAllowedExtraProfit(&_Tricrypto2.CallOpts)
}

// FutureAllowedExtraProfit is a free data retrieval call binding the contract method 0x727ced57.
//
// Solidity: function future_allowed_extra_profit() view returns(uint256)
func (_Tricrypto2 *Tricrypto2CallerSession) FutureAllowedExtraProfit() (*big.Int, error) {
	return _Tricrypto2.Contract.FutureAllowedExtraProfit(&_Tricrypto2.CallOpts)
}

// FutureFeeGamma is a free data retrieval call binding the contract method 0xd7c3dcbe.
//
// Solidity: function future_fee_gamma() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Caller) FutureFeeGamma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "future_fee_gamma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureFeeGamma is a free data retrieval call binding the contract method 0xd7c3dcbe.
//
// Solidity: function future_fee_gamma() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Session) FutureFeeGamma() (*big.Int, error) {
	return _Tricrypto2.Contract.FutureFeeGamma(&_Tricrypto2.CallOpts)
}

// FutureFeeGamma is a free data retrieval call binding the contract method 0xd7c3dcbe.
//
// Solidity: function future_fee_gamma() view returns(uint256)
func (_Tricrypto2 *Tricrypto2CallerSession) FutureFeeGamma() (*big.Int, error) {
	return _Tricrypto2.Contract.FutureFeeGamma(&_Tricrypto2.CallOpts)
}

// FutureMaHalfTime is a free data retrieval call binding the contract method 0x0c5e23d4.
//
// Solidity: function future_ma_half_time() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Caller) FutureMaHalfTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "future_ma_half_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureMaHalfTime is a free data retrieval call binding the contract method 0x0c5e23d4.
//
// Solidity: function future_ma_half_time() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Session) FutureMaHalfTime() (*big.Int, error) {
	return _Tricrypto2.Contract.FutureMaHalfTime(&_Tricrypto2.CallOpts)
}

// FutureMaHalfTime is a free data retrieval call binding the contract method 0x0c5e23d4.
//
// Solidity: function future_ma_half_time() view returns(uint256)
func (_Tricrypto2 *Tricrypto2CallerSession) FutureMaHalfTime() (*big.Int, error) {
	return _Tricrypto2.Contract.FutureMaHalfTime(&_Tricrypto2.CallOpts)
}

// FutureMidFee is a free data retrieval call binding the contract method 0x7cf9aedc.
//
// Solidity: function future_mid_fee() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Caller) FutureMidFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "future_mid_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureMidFee is a free data retrieval call binding the contract method 0x7cf9aedc.
//
// Solidity: function future_mid_fee() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Session) FutureMidFee() (*big.Int, error) {
	return _Tricrypto2.Contract.FutureMidFee(&_Tricrypto2.CallOpts)
}

// FutureMidFee is a free data retrieval call binding the contract method 0x7cf9aedc.
//
// Solidity: function future_mid_fee() view returns(uint256)
func (_Tricrypto2 *Tricrypto2CallerSession) FutureMidFee() (*big.Int, error) {
	return _Tricrypto2.Contract.FutureMidFee(&_Tricrypto2.CallOpts)
}

// FutureOutFee is a free data retrieval call binding the contract method 0x7d1b060c.
//
// Solidity: function future_out_fee() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Caller) FutureOutFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "future_out_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureOutFee is a free data retrieval call binding the contract method 0x7d1b060c.
//
// Solidity: function future_out_fee() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Session) FutureOutFee() (*big.Int, error) {
	return _Tricrypto2.Contract.FutureOutFee(&_Tricrypto2.CallOpts)
}

// FutureOutFee is a free data retrieval call binding the contract method 0x7d1b060c.
//
// Solidity: function future_out_fee() view returns(uint256)
func (_Tricrypto2 *Tricrypto2CallerSession) FutureOutFee() (*big.Int, error) {
	return _Tricrypto2.Contract.FutureOutFee(&_Tricrypto2.CallOpts)
}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() view returns(address)
func (_Tricrypto2 *Tricrypto2Caller) FutureOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "future_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() view returns(address)
func (_Tricrypto2 *Tricrypto2Session) FutureOwner() (common.Address, error) {
	return _Tricrypto2.Contract.FutureOwner(&_Tricrypto2.CallOpts)
}

// FutureOwner is a free data retrieval call binding the contract method 0x1ec0cdc1.
//
// Solidity: function future_owner() view returns(address)
func (_Tricrypto2 *Tricrypto2CallerSession) FutureOwner() (common.Address, error) {
	return _Tricrypto2.Contract.FutureOwner(&_Tricrypto2.CallOpts)
}

// Gamma is a free data retrieval call binding the contract method 0xb1373929.
//
// Solidity: function gamma() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Caller) Gamma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "gamma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Gamma is a free data retrieval call binding the contract method 0xb1373929.
//
// Solidity: function gamma() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Session) Gamma() (*big.Int, error) {
	return _Tricrypto2.Contract.Gamma(&_Tricrypto2.CallOpts)
}

// Gamma is a free data retrieval call binding the contract method 0xb1373929.
//
// Solidity: function gamma() view returns(uint256)
func (_Tricrypto2 *Tricrypto2CallerSession) Gamma() (*big.Int, error) {
	return _Tricrypto2.Contract.Gamma(&_Tricrypto2.CallOpts)
}

// GetDy is a free data retrieval call binding the contract method 0x556d6e9f.
//
// Solidity: function get_dy(uint256 i, uint256 j, uint256 dx) view returns(uint256)
func (_Tricrypto2 *Tricrypto2Caller) GetDy(opts *bind.CallOpts, i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "get_dy", i, j, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDy is a free data retrieval call binding the contract method 0x556d6e9f.
//
// Solidity: function get_dy(uint256 i, uint256 j, uint256 dx) view returns(uint256)
func (_Tricrypto2 *Tricrypto2Session) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _Tricrypto2.Contract.GetDy(&_Tricrypto2.CallOpts, i, j, dx)
}

// GetDy is a free data retrieval call binding the contract method 0x556d6e9f.
//
// Solidity: function get_dy(uint256 i, uint256 j, uint256 dx) view returns(uint256)
func (_Tricrypto2 *Tricrypto2CallerSession) GetDy(i *big.Int, j *big.Int, dx *big.Int) (*big.Int, error) {
	return _Tricrypto2.Contract.GetDy(&_Tricrypto2.CallOpts, i, j, dx)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Caller) GetVirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "get_virtual_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Session) GetVirtualPrice() (*big.Int, error) {
	return _Tricrypto2.Contract.GetVirtualPrice(&_Tricrypto2.CallOpts)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xbb7b8b80.
//
// Solidity: function get_virtual_price() view returns(uint256)
func (_Tricrypto2 *Tricrypto2CallerSession) GetVirtualPrice() (*big.Int, error) {
	return _Tricrypto2.Contract.GetVirtualPrice(&_Tricrypto2.CallOpts)
}

// InitialAGamma is a free data retrieval call binding the contract method 0x204fe3d5.
//
// Solidity: function initial_A_gamma() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Caller) InitialAGamma(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "initial_A_gamma")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialAGamma is a free data retrieval call binding the contract method 0x204fe3d5.
//
// Solidity: function initial_A_gamma() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Session) InitialAGamma() (*big.Int, error) {
	return _Tricrypto2.Contract.InitialAGamma(&_Tricrypto2.CallOpts)
}

// InitialAGamma is a free data retrieval call binding the contract method 0x204fe3d5.
//
// Solidity: function initial_A_gamma() view returns(uint256)
func (_Tricrypto2 *Tricrypto2CallerSession) InitialAGamma() (*big.Int, error) {
	return _Tricrypto2.Contract.InitialAGamma(&_Tricrypto2.CallOpts)
}

// InitialAGammaTime is a free data retrieval call binding the contract method 0xe89876ff.
//
// Solidity: function initial_A_gamma_time() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Caller) InitialAGammaTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "initial_A_gamma_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialAGammaTime is a free data retrieval call binding the contract method 0xe89876ff.
//
// Solidity: function initial_A_gamma_time() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Session) InitialAGammaTime() (*big.Int, error) {
	return _Tricrypto2.Contract.InitialAGammaTime(&_Tricrypto2.CallOpts)
}

// InitialAGammaTime is a free data retrieval call binding the contract method 0xe89876ff.
//
// Solidity: function initial_A_gamma_time() view returns(uint256)
func (_Tricrypto2 *Tricrypto2CallerSession) InitialAGammaTime() (*big.Int, error) {
	return _Tricrypto2.Contract.InitialAGammaTime(&_Tricrypto2.CallOpts)
}

// IsKilled is a free data retrieval call binding the contract method 0x9c868ac0.
//
// Solidity: function is_killed() view returns(bool)
func (_Tricrypto2 *Tricrypto2Caller) IsKilled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "is_killed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKilled is a free data retrieval call binding the contract method 0x9c868ac0.
//
// Solidity: function is_killed() view returns(bool)
func (_Tricrypto2 *Tricrypto2Session) IsKilled() (bool, error) {
	return _Tricrypto2.Contract.IsKilled(&_Tricrypto2.CallOpts)
}

// IsKilled is a free data retrieval call binding the contract method 0x9c868ac0.
//
// Solidity: function is_killed() view returns(bool)
func (_Tricrypto2 *Tricrypto2CallerSession) IsKilled() (bool, error) {
	return _Tricrypto2.Contract.IsKilled(&_Tricrypto2.CallOpts)
}

// KillDeadline is a free data retrieval call binding the contract method 0x2a426896.
//
// Solidity: function kill_deadline() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Caller) KillDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "kill_deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KillDeadline is a free data retrieval call binding the contract method 0x2a426896.
//
// Solidity: function kill_deadline() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Session) KillDeadline() (*big.Int, error) {
	return _Tricrypto2.Contract.KillDeadline(&_Tricrypto2.CallOpts)
}

// KillDeadline is a free data retrieval call binding the contract method 0x2a426896.
//
// Solidity: function kill_deadline() view returns(uint256)
func (_Tricrypto2 *Tricrypto2CallerSession) KillDeadline() (*big.Int, error) {
	return _Tricrypto2.Contract.KillDeadline(&_Tricrypto2.CallOpts)
}

// LastPrices is a free data retrieval call binding the contract method 0x59189017.
//
// Solidity: function last_prices(uint256 k) view returns(uint256)
func (_Tricrypto2 *Tricrypto2Caller) LastPrices(opts *bind.CallOpts, k *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "last_prices", k)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPrices is a free data retrieval call binding the contract method 0x59189017.
//
// Solidity: function last_prices(uint256 k) view returns(uint256)
func (_Tricrypto2 *Tricrypto2Session) LastPrices(k *big.Int) (*big.Int, error) {
	return _Tricrypto2.Contract.LastPrices(&_Tricrypto2.CallOpts, k)
}

// LastPrices is a free data retrieval call binding the contract method 0x59189017.
//
// Solidity: function last_prices(uint256 k) view returns(uint256)
func (_Tricrypto2 *Tricrypto2CallerSession) LastPrices(k *big.Int) (*big.Int, error) {
	return _Tricrypto2.Contract.LastPrices(&_Tricrypto2.CallOpts, k)
}

// LastPricesTimestamp is a free data retrieval call binding the contract method 0x6112c747.
//
// Solidity: function last_prices_timestamp() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Caller) LastPricesTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "last_prices_timestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPricesTimestamp is a free data retrieval call binding the contract method 0x6112c747.
//
// Solidity: function last_prices_timestamp() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Session) LastPricesTimestamp() (*big.Int, error) {
	return _Tricrypto2.Contract.LastPricesTimestamp(&_Tricrypto2.CallOpts)
}

// LastPricesTimestamp is a free data retrieval call binding the contract method 0x6112c747.
//
// Solidity: function last_prices_timestamp() view returns(uint256)
func (_Tricrypto2 *Tricrypto2CallerSession) LastPricesTimestamp() (*big.Int, error) {
	return _Tricrypto2.Contract.LastPricesTimestamp(&_Tricrypto2.CallOpts)
}

// MaHalfTime is a free data retrieval call binding the contract method 0x662b6274.
//
// Solidity: function ma_half_time() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Caller) MaHalfTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "ma_half_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaHalfTime is a free data retrieval call binding the contract method 0x662b6274.
//
// Solidity: function ma_half_time() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Session) MaHalfTime() (*big.Int, error) {
	return _Tricrypto2.Contract.MaHalfTime(&_Tricrypto2.CallOpts)
}

// MaHalfTime is a free data retrieval call binding the contract method 0x662b6274.
//
// Solidity: function ma_half_time() view returns(uint256)
func (_Tricrypto2 *Tricrypto2CallerSession) MaHalfTime() (*big.Int, error) {
	return _Tricrypto2.Contract.MaHalfTime(&_Tricrypto2.CallOpts)
}

// MidFee is a free data retrieval call binding the contract method 0x92526c0c.
//
// Solidity: function mid_fee() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Caller) MidFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "mid_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MidFee is a free data retrieval call binding the contract method 0x92526c0c.
//
// Solidity: function mid_fee() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Session) MidFee() (*big.Int, error) {
	return _Tricrypto2.Contract.MidFee(&_Tricrypto2.CallOpts)
}

// MidFee is a free data retrieval call binding the contract method 0x92526c0c.
//
// Solidity: function mid_fee() view returns(uint256)
func (_Tricrypto2 *Tricrypto2CallerSession) MidFee() (*big.Int, error) {
	return _Tricrypto2.Contract.MidFee(&_Tricrypto2.CallOpts)
}

// OutFee is a free data retrieval call binding the contract method 0xee8de675.
//
// Solidity: function out_fee() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Caller) OutFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "out_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OutFee is a free data retrieval call binding the contract method 0xee8de675.
//
// Solidity: function out_fee() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Session) OutFee() (*big.Int, error) {
	return _Tricrypto2.Contract.OutFee(&_Tricrypto2.CallOpts)
}

// OutFee is a free data retrieval call binding the contract method 0xee8de675.
//
// Solidity: function out_fee() view returns(uint256)
func (_Tricrypto2 *Tricrypto2CallerSession) OutFee() (*big.Int, error) {
	return _Tricrypto2.Contract.OutFee(&_Tricrypto2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tricrypto2 *Tricrypto2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tricrypto2 *Tricrypto2Session) Owner() (common.Address, error) {
	return _Tricrypto2.Contract.Owner(&_Tricrypto2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tricrypto2 *Tricrypto2CallerSession) Owner() (common.Address, error) {
	return _Tricrypto2.Contract.Owner(&_Tricrypto2.CallOpts)
}

// PriceOracle is a free data retrieval call binding the contract method 0x68727653.
//
// Solidity: function price_oracle(uint256 k) view returns(uint256)
func (_Tricrypto2 *Tricrypto2Caller) PriceOracle(opts *bind.CallOpts, k *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "price_oracle", k)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceOracle is a free data retrieval call binding the contract method 0x68727653.
//
// Solidity: function price_oracle(uint256 k) view returns(uint256)
func (_Tricrypto2 *Tricrypto2Session) PriceOracle(k *big.Int) (*big.Int, error) {
	return _Tricrypto2.Contract.PriceOracle(&_Tricrypto2.CallOpts, k)
}

// PriceOracle is a free data retrieval call binding the contract method 0x68727653.
//
// Solidity: function price_oracle(uint256 k) view returns(uint256)
func (_Tricrypto2 *Tricrypto2CallerSession) PriceOracle(k *big.Int) (*big.Int, error) {
	return _Tricrypto2.Contract.PriceOracle(&_Tricrypto2.CallOpts, k)
}

// PriceScale is a free data retrieval call binding the contract method 0xa3f7cdd5.
//
// Solidity: function price_scale(uint256 k) view returns(uint256)
func (_Tricrypto2 *Tricrypto2Caller) PriceScale(opts *bind.CallOpts, k *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "price_scale", k)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceScale is a free data retrieval call binding the contract method 0xa3f7cdd5.
//
// Solidity: function price_scale(uint256 k) view returns(uint256)
func (_Tricrypto2 *Tricrypto2Session) PriceScale(k *big.Int) (*big.Int, error) {
	return _Tricrypto2.Contract.PriceScale(&_Tricrypto2.CallOpts, k)
}

// PriceScale is a free data retrieval call binding the contract method 0xa3f7cdd5.
//
// Solidity: function price_scale(uint256 k) view returns(uint256)
func (_Tricrypto2 *Tricrypto2CallerSession) PriceScale(k *big.Int) (*big.Int, error) {
	return _Tricrypto2.Contract.PriceScale(&_Tricrypto2.CallOpts, k)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Tricrypto2 *Tricrypto2Caller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Tricrypto2 *Tricrypto2Session) Token() (common.Address, error) {
	return _Tricrypto2.Contract.Token(&_Tricrypto2.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Tricrypto2 *Tricrypto2CallerSession) Token() (common.Address, error) {
	return _Tricrypto2.Contract.Token(&_Tricrypto2.CallOpts)
}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Caller) TransferOwnershipDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "transfer_ownership_deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Session) TransferOwnershipDeadline() (*big.Int, error) {
	return _Tricrypto2.Contract.TransferOwnershipDeadline(&_Tricrypto2.CallOpts)
}

// TransferOwnershipDeadline is a free data retrieval call binding the contract method 0xe0a0b586.
//
// Solidity: function transfer_ownership_deadline() view returns(uint256)
func (_Tricrypto2 *Tricrypto2CallerSession) TransferOwnershipDeadline() (*big.Int, error) {
	return _Tricrypto2.Contract.TransferOwnershipDeadline(&_Tricrypto2.CallOpts)
}

// VirtualPrice is a free data retrieval call binding the contract method 0x0c46b72a.
//
// Solidity: function virtual_price() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Caller) VirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "virtual_price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VirtualPrice is a free data retrieval call binding the contract method 0x0c46b72a.
//
// Solidity: function virtual_price() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Session) VirtualPrice() (*big.Int, error) {
	return _Tricrypto2.Contract.VirtualPrice(&_Tricrypto2.CallOpts)
}

// VirtualPrice is a free data retrieval call binding the contract method 0x0c46b72a.
//
// Solidity: function virtual_price() view returns(uint256)
func (_Tricrypto2 *Tricrypto2CallerSession) VirtualPrice() (*big.Int, error) {
	return _Tricrypto2.Contract.VirtualPrice(&_Tricrypto2.CallOpts)
}

// XcpProfit is a free data retrieval call binding the contract method 0x7ba1a74d.
//
// Solidity: function xcp_profit() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Caller) XcpProfit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "xcp_profit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XcpProfit is a free data retrieval call binding the contract method 0x7ba1a74d.
//
// Solidity: function xcp_profit() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Session) XcpProfit() (*big.Int, error) {
	return _Tricrypto2.Contract.XcpProfit(&_Tricrypto2.CallOpts)
}

// XcpProfit is a free data retrieval call binding the contract method 0x7ba1a74d.
//
// Solidity: function xcp_profit() view returns(uint256)
func (_Tricrypto2 *Tricrypto2CallerSession) XcpProfit() (*big.Int, error) {
	return _Tricrypto2.Contract.XcpProfit(&_Tricrypto2.CallOpts)
}

// XcpProfitA is a free data retrieval call binding the contract method 0x0b7b594b.
//
// Solidity: function xcp_profit_a() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Caller) XcpProfitA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tricrypto2.contract.Call(opts, &out, "xcp_profit_a")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XcpProfitA is a free data retrieval call binding the contract method 0x0b7b594b.
//
// Solidity: function xcp_profit_a() view returns(uint256)
func (_Tricrypto2 *Tricrypto2Session) XcpProfitA() (*big.Int, error) {
	return _Tricrypto2.Contract.XcpProfitA(&_Tricrypto2.CallOpts)
}

// XcpProfitA is a free data retrieval call binding the contract method 0x0b7b594b.
//
// Solidity: function xcp_profit_a() view returns(uint256)
func (_Tricrypto2 *Tricrypto2CallerSession) XcpProfitA() (*big.Int, error) {
	return _Tricrypto2.Contract.XcpProfitA(&_Tricrypto2.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4515cef3.
//
// Solidity: function add_liquidity(uint256[3] amounts, uint256 min_mint_amount) returns()
func (_Tricrypto2 *Tricrypto2Transactor) AddLiquidity(opts *bind.TransactOpts, amounts [3]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Tricrypto2.contract.Transact(opts, "add_liquidity", amounts, min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4515cef3.
//
// Solidity: function add_liquidity(uint256[3] amounts, uint256 min_mint_amount) returns()
func (_Tricrypto2 *Tricrypto2Session) AddLiquidity(amounts [3]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Tricrypto2.Contract.AddLiquidity(&_Tricrypto2.TransactOpts, amounts, min_mint_amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4515cef3.
//
// Solidity: function add_liquidity(uint256[3] amounts, uint256 min_mint_amount) returns()
func (_Tricrypto2 *Tricrypto2TransactorSession) AddLiquidity(amounts [3]*big.Int, min_mint_amount *big.Int) (*types.Transaction, error) {
	return _Tricrypto2.Contract.AddLiquidity(&_Tricrypto2.TransactOpts, amounts, min_mint_amount)
}

// ApplyNewParameters is a paid mutator transaction binding the contract method 0x2a7dd7cd.
//
// Solidity: function apply_new_parameters() returns()
func (_Tricrypto2 *Tricrypto2Transactor) ApplyNewParameters(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tricrypto2.contract.Transact(opts, "apply_new_parameters")
}

// ApplyNewParameters is a paid mutator transaction binding the contract method 0x2a7dd7cd.
//
// Solidity: function apply_new_parameters() returns()
func (_Tricrypto2 *Tricrypto2Session) ApplyNewParameters() (*types.Transaction, error) {
	return _Tricrypto2.Contract.ApplyNewParameters(&_Tricrypto2.TransactOpts)
}

// ApplyNewParameters is a paid mutator transaction binding the contract method 0x2a7dd7cd.
//
// Solidity: function apply_new_parameters() returns()
func (_Tricrypto2 *Tricrypto2TransactorSession) ApplyNewParameters() (*types.Transaction, error) {
	return _Tricrypto2.Contract.ApplyNewParameters(&_Tricrypto2.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_Tricrypto2 *Tricrypto2Transactor) ApplyTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tricrypto2.contract.Transact(opts, "apply_transfer_ownership")
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_Tricrypto2 *Tricrypto2Session) ApplyTransferOwnership() (*types.Transaction, error) {
	return _Tricrypto2.Contract.ApplyTransferOwnership(&_Tricrypto2.TransactOpts)
}

// ApplyTransferOwnership is a paid mutator transaction binding the contract method 0x6a1c05ae.
//
// Solidity: function apply_transfer_ownership() returns()
func (_Tricrypto2 *Tricrypto2TransactorSession) ApplyTransferOwnership() (*types.Transaction, error) {
	return _Tricrypto2.Contract.ApplyTransferOwnership(&_Tricrypto2.TransactOpts)
}

// ClaimAdminFees is a paid mutator transaction binding the contract method 0xc93f49e8.
//
// Solidity: function claim_admin_fees() returns()
func (_Tricrypto2 *Tricrypto2Transactor) ClaimAdminFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tricrypto2.contract.Transact(opts, "claim_admin_fees")
}

// ClaimAdminFees is a paid mutator transaction binding the contract method 0xc93f49e8.
//
// Solidity: function claim_admin_fees() returns()
func (_Tricrypto2 *Tricrypto2Session) ClaimAdminFees() (*types.Transaction, error) {
	return _Tricrypto2.Contract.ClaimAdminFees(&_Tricrypto2.TransactOpts)
}

// ClaimAdminFees is a paid mutator transaction binding the contract method 0xc93f49e8.
//
// Solidity: function claim_admin_fees() returns()
func (_Tricrypto2 *Tricrypto2TransactorSession) ClaimAdminFees() (*types.Transaction, error) {
	return _Tricrypto2.Contract.ClaimAdminFees(&_Tricrypto2.TransactOpts)
}

// CommitNewParameters is a paid mutator transaction binding the contract method 0xa43c3351.
//
// Solidity: function commit_new_parameters(uint256 _new_mid_fee, uint256 _new_out_fee, uint256 _new_admin_fee, uint256 _new_fee_gamma, uint256 _new_allowed_extra_profit, uint256 _new_adjustment_step, uint256 _new_ma_half_time) returns()
func (_Tricrypto2 *Tricrypto2Transactor) CommitNewParameters(opts *bind.TransactOpts, _new_mid_fee *big.Int, _new_out_fee *big.Int, _new_admin_fee *big.Int, _new_fee_gamma *big.Int, _new_allowed_extra_profit *big.Int, _new_adjustment_step *big.Int, _new_ma_half_time *big.Int) (*types.Transaction, error) {
	return _Tricrypto2.contract.Transact(opts, "commit_new_parameters", _new_mid_fee, _new_out_fee, _new_admin_fee, _new_fee_gamma, _new_allowed_extra_profit, _new_adjustment_step, _new_ma_half_time)
}

// CommitNewParameters is a paid mutator transaction binding the contract method 0xa43c3351.
//
// Solidity: function commit_new_parameters(uint256 _new_mid_fee, uint256 _new_out_fee, uint256 _new_admin_fee, uint256 _new_fee_gamma, uint256 _new_allowed_extra_profit, uint256 _new_adjustment_step, uint256 _new_ma_half_time) returns()
func (_Tricrypto2 *Tricrypto2Session) CommitNewParameters(_new_mid_fee *big.Int, _new_out_fee *big.Int, _new_admin_fee *big.Int, _new_fee_gamma *big.Int, _new_allowed_extra_profit *big.Int, _new_adjustment_step *big.Int, _new_ma_half_time *big.Int) (*types.Transaction, error) {
	return _Tricrypto2.Contract.CommitNewParameters(&_Tricrypto2.TransactOpts, _new_mid_fee, _new_out_fee, _new_admin_fee, _new_fee_gamma, _new_allowed_extra_profit, _new_adjustment_step, _new_ma_half_time)
}

// CommitNewParameters is a paid mutator transaction binding the contract method 0xa43c3351.
//
// Solidity: function commit_new_parameters(uint256 _new_mid_fee, uint256 _new_out_fee, uint256 _new_admin_fee, uint256 _new_fee_gamma, uint256 _new_allowed_extra_profit, uint256 _new_adjustment_step, uint256 _new_ma_half_time) returns()
func (_Tricrypto2 *Tricrypto2TransactorSession) CommitNewParameters(_new_mid_fee *big.Int, _new_out_fee *big.Int, _new_admin_fee *big.Int, _new_fee_gamma *big.Int, _new_allowed_extra_profit *big.Int, _new_adjustment_step *big.Int, _new_ma_half_time *big.Int) (*types.Transaction, error) {
	return _Tricrypto2.Contract.CommitNewParameters(&_Tricrypto2.TransactOpts, _new_mid_fee, _new_out_fee, _new_admin_fee, _new_fee_gamma, _new_allowed_extra_profit, _new_adjustment_step, _new_ma_half_time)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_Tricrypto2 *Tricrypto2Transactor) CommitTransferOwnership(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Tricrypto2.contract.Transact(opts, "commit_transfer_ownership", _owner)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_Tricrypto2 *Tricrypto2Session) CommitTransferOwnership(_owner common.Address) (*types.Transaction, error) {
	return _Tricrypto2.Contract.CommitTransferOwnership(&_Tricrypto2.TransactOpts, _owner)
}

// CommitTransferOwnership is a paid mutator transaction binding the contract method 0x6b441a40.
//
// Solidity: function commit_transfer_ownership(address _owner) returns()
func (_Tricrypto2 *Tricrypto2TransactorSession) CommitTransferOwnership(_owner common.Address) (*types.Transaction, error) {
	return _Tricrypto2.Contract.CommitTransferOwnership(&_Tricrypto2.TransactOpts, _owner)
}

// Exchange is a paid mutator transaction binding the contract method 0x5b41b908.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy) payable returns()
func (_Tricrypto2 *Tricrypto2Transactor) Exchange(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Tricrypto2.contract.Transact(opts, "exchange", i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x5b41b908.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy) payable returns()
func (_Tricrypto2 *Tricrypto2Session) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Tricrypto2.Contract.Exchange(&_Tricrypto2.TransactOpts, i, j, dx, min_dy)
}

// Exchange is a paid mutator transaction binding the contract method 0x5b41b908.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy) payable returns()
func (_Tricrypto2 *Tricrypto2TransactorSession) Exchange(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int) (*types.Transaction, error) {
	return _Tricrypto2.Contract.Exchange(&_Tricrypto2.TransactOpts, i, j, dx, min_dy)
}

// Exchange0 is a paid mutator transaction binding the contract method 0x394747c5.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy, bool use_eth) payable returns()
func (_Tricrypto2 *Tricrypto2Transactor) Exchange0(opts *bind.TransactOpts, i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, use_eth bool) (*types.Transaction, error) {
	return _Tricrypto2.contract.Transact(opts, "exchange0", i, j, dx, min_dy, use_eth)
}

// Exchange0 is a paid mutator transaction binding the contract method 0x394747c5.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy, bool use_eth) payable returns()
func (_Tricrypto2 *Tricrypto2Session) Exchange0(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, use_eth bool) (*types.Transaction, error) {
	return _Tricrypto2.Contract.Exchange0(&_Tricrypto2.TransactOpts, i, j, dx, min_dy, use_eth)
}

// Exchange0 is a paid mutator transaction binding the contract method 0x394747c5.
//
// Solidity: function exchange(uint256 i, uint256 j, uint256 dx, uint256 min_dy, bool use_eth) payable returns()
func (_Tricrypto2 *Tricrypto2TransactorSession) Exchange0(i *big.Int, j *big.Int, dx *big.Int, min_dy *big.Int, use_eth bool) (*types.Transaction, error) {
	return _Tricrypto2.Contract.Exchange0(&_Tricrypto2.TransactOpts, i, j, dx, min_dy, use_eth)
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_Tricrypto2 *Tricrypto2Transactor) KillMe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tricrypto2.contract.Transact(opts, "kill_me")
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_Tricrypto2 *Tricrypto2Session) KillMe() (*types.Transaction, error) {
	return _Tricrypto2.Contract.KillMe(&_Tricrypto2.TransactOpts)
}

// KillMe is a paid mutator transaction binding the contract method 0xe3698853.
//
// Solidity: function kill_me() returns()
func (_Tricrypto2 *Tricrypto2TransactorSession) KillMe() (*types.Transaction, error) {
	return _Tricrypto2.Contract.KillMe(&_Tricrypto2.TransactOpts)
}

// RampAGamma is a paid mutator transaction binding the contract method 0x5e248072.
//
// Solidity: function ramp_A_gamma(uint256 future_A, uint256 future_gamma, uint256 future_time) returns()
func (_Tricrypto2 *Tricrypto2Transactor) RampAGamma(opts *bind.TransactOpts, future_A *big.Int, future_gamma *big.Int, future_time *big.Int) (*types.Transaction, error) {
	return _Tricrypto2.contract.Transact(opts, "ramp_A_gamma", future_A, future_gamma, future_time)
}

// RampAGamma is a paid mutator transaction binding the contract method 0x5e248072.
//
// Solidity: function ramp_A_gamma(uint256 future_A, uint256 future_gamma, uint256 future_time) returns()
func (_Tricrypto2 *Tricrypto2Session) RampAGamma(future_A *big.Int, future_gamma *big.Int, future_time *big.Int) (*types.Transaction, error) {
	return _Tricrypto2.Contract.RampAGamma(&_Tricrypto2.TransactOpts, future_A, future_gamma, future_time)
}

// RampAGamma is a paid mutator transaction binding the contract method 0x5e248072.
//
// Solidity: function ramp_A_gamma(uint256 future_A, uint256 future_gamma, uint256 future_time) returns()
func (_Tricrypto2 *Tricrypto2TransactorSession) RampAGamma(future_A *big.Int, future_gamma *big.Int, future_time *big.Int) (*types.Transaction, error) {
	return _Tricrypto2.Contract.RampAGamma(&_Tricrypto2.TransactOpts, future_A, future_gamma, future_time)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xecb586a5.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[3] min_amounts) returns()
func (_Tricrypto2 *Tricrypto2Transactor) RemoveLiquidity(opts *bind.TransactOpts, _amount *big.Int, min_amounts [3]*big.Int) (*types.Transaction, error) {
	return _Tricrypto2.contract.Transact(opts, "remove_liquidity", _amount, min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xecb586a5.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[3] min_amounts) returns()
func (_Tricrypto2 *Tricrypto2Session) RemoveLiquidity(_amount *big.Int, min_amounts [3]*big.Int) (*types.Transaction, error) {
	return _Tricrypto2.Contract.RemoveLiquidity(&_Tricrypto2.TransactOpts, _amount, min_amounts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xecb586a5.
//
// Solidity: function remove_liquidity(uint256 _amount, uint256[3] min_amounts) returns()
func (_Tricrypto2 *Tricrypto2TransactorSession) RemoveLiquidity(_amount *big.Int, min_amounts [3]*big.Int) (*types.Transaction, error) {
	return _Tricrypto2.Contract.RemoveLiquidity(&_Tricrypto2.TransactOpts, _amount, min_amounts)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0xf1dc3cc9.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount) returns()
func (_Tricrypto2 *Tricrypto2Transactor) RemoveLiquidityOneCoin(opts *bind.TransactOpts, token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _Tricrypto2.contract.Transact(opts, "remove_liquidity_one_coin", token_amount, i, min_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0xf1dc3cc9.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount) returns()
func (_Tricrypto2 *Tricrypto2Session) RemoveLiquidityOneCoin(token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _Tricrypto2.Contract.RemoveLiquidityOneCoin(&_Tricrypto2.TransactOpts, token_amount, i, min_amount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0xf1dc3cc9.
//
// Solidity: function remove_liquidity_one_coin(uint256 token_amount, uint256 i, uint256 min_amount) returns()
func (_Tricrypto2 *Tricrypto2TransactorSession) RemoveLiquidityOneCoin(token_amount *big.Int, i *big.Int, min_amount *big.Int) (*types.Transaction, error) {
	return _Tricrypto2.Contract.RemoveLiquidityOneCoin(&_Tricrypto2.TransactOpts, token_amount, i, min_amount)
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_Tricrypto2 *Tricrypto2Transactor) RevertNewParameters(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tricrypto2.contract.Transact(opts, "revert_new_parameters")
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_Tricrypto2 *Tricrypto2Session) RevertNewParameters() (*types.Transaction, error) {
	return _Tricrypto2.Contract.RevertNewParameters(&_Tricrypto2.TransactOpts)
}

// RevertNewParameters is a paid mutator transaction binding the contract method 0x226840fb.
//
// Solidity: function revert_new_parameters() returns()
func (_Tricrypto2 *Tricrypto2TransactorSession) RevertNewParameters() (*types.Transaction, error) {
	return _Tricrypto2.Contract.RevertNewParameters(&_Tricrypto2.TransactOpts)
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_Tricrypto2 *Tricrypto2Transactor) RevertTransferOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tricrypto2.contract.Transact(opts, "revert_transfer_ownership")
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_Tricrypto2 *Tricrypto2Session) RevertTransferOwnership() (*types.Transaction, error) {
	return _Tricrypto2.Contract.RevertTransferOwnership(&_Tricrypto2.TransactOpts)
}

// RevertTransferOwnership is a paid mutator transaction binding the contract method 0x86fbf193.
//
// Solidity: function revert_transfer_ownership() returns()
func (_Tricrypto2 *Tricrypto2TransactorSession) RevertTransferOwnership() (*types.Transaction, error) {
	return _Tricrypto2.Contract.RevertTransferOwnership(&_Tricrypto2.TransactOpts)
}

// SetAdminFeeReceiver is a paid mutator transaction binding the contract method 0x7242e524.
//
// Solidity: function set_admin_fee_receiver(address _admin_fee_receiver) returns()
func (_Tricrypto2 *Tricrypto2Transactor) SetAdminFeeReceiver(opts *bind.TransactOpts, _admin_fee_receiver common.Address) (*types.Transaction, error) {
	return _Tricrypto2.contract.Transact(opts, "set_admin_fee_receiver", _admin_fee_receiver)
}

// SetAdminFeeReceiver is a paid mutator transaction binding the contract method 0x7242e524.
//
// Solidity: function set_admin_fee_receiver(address _admin_fee_receiver) returns()
func (_Tricrypto2 *Tricrypto2Session) SetAdminFeeReceiver(_admin_fee_receiver common.Address) (*types.Transaction, error) {
	return _Tricrypto2.Contract.SetAdminFeeReceiver(&_Tricrypto2.TransactOpts, _admin_fee_receiver)
}

// SetAdminFeeReceiver is a paid mutator transaction binding the contract method 0x7242e524.
//
// Solidity: function set_admin_fee_receiver(address _admin_fee_receiver) returns()
func (_Tricrypto2 *Tricrypto2TransactorSession) SetAdminFeeReceiver(_admin_fee_receiver common.Address) (*types.Transaction, error) {
	return _Tricrypto2.Contract.SetAdminFeeReceiver(&_Tricrypto2.TransactOpts, _admin_fee_receiver)
}

// StopRampAGamma is a paid mutator transaction binding the contract method 0x244c7c2e.
//
// Solidity: function stop_ramp_A_gamma() returns()
func (_Tricrypto2 *Tricrypto2Transactor) StopRampAGamma(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tricrypto2.contract.Transact(opts, "stop_ramp_A_gamma")
}

// StopRampAGamma is a paid mutator transaction binding the contract method 0x244c7c2e.
//
// Solidity: function stop_ramp_A_gamma() returns()
func (_Tricrypto2 *Tricrypto2Session) StopRampAGamma() (*types.Transaction, error) {
	return _Tricrypto2.Contract.StopRampAGamma(&_Tricrypto2.TransactOpts)
}

// StopRampAGamma is a paid mutator transaction binding the contract method 0x244c7c2e.
//
// Solidity: function stop_ramp_A_gamma() returns()
func (_Tricrypto2 *Tricrypto2TransactorSession) StopRampAGamma() (*types.Transaction, error) {
	return _Tricrypto2.Contract.StopRampAGamma(&_Tricrypto2.TransactOpts)
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_Tricrypto2 *Tricrypto2Transactor) UnkillMe(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tricrypto2.contract.Transact(opts, "unkill_me")
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_Tricrypto2 *Tricrypto2Session) UnkillMe() (*types.Transaction, error) {
	return _Tricrypto2.Contract.UnkillMe(&_Tricrypto2.TransactOpts)
}

// UnkillMe is a paid mutator transaction binding the contract method 0x3046f972.
//
// Solidity: function unkill_me() returns()
func (_Tricrypto2 *Tricrypto2TransactorSession) UnkillMe() (*types.Transaction, error) {
	return _Tricrypto2.Contract.UnkillMe(&_Tricrypto2.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Tricrypto2 *Tricrypto2Transactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Tricrypto2.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Tricrypto2 *Tricrypto2Session) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Tricrypto2.Contract.Fallback(&_Tricrypto2.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Tricrypto2 *Tricrypto2TransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Tricrypto2.Contract.Fallback(&_Tricrypto2.TransactOpts, calldata)
}

// Tricrypto2AddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the Tricrypto2 contract.
type Tricrypto2AddLiquidityIterator struct {
	Event *Tricrypto2AddLiquidity // Event containing the contract specifics and raw log

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
func (it *Tricrypto2AddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Tricrypto2AddLiquidity)
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
		it.Event = new(Tricrypto2AddLiquidity)
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
func (it *Tricrypto2AddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Tricrypto2AddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Tricrypto2AddLiquidity represents a AddLiquidity event raised by the Tricrypto2 contract.
type Tricrypto2AddLiquidity struct {
	Provider     common.Address
	TokenAmounts [3]*big.Int
	Fee          *big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0x96b486485420b963edd3fdec0b0195730035600feb7de6f544383d7950fa97ee.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[3] token_amounts, uint256 fee, uint256 token_supply)
func (_Tricrypto2 *Tricrypto2Filterer) FilterAddLiquidity(opts *bind.FilterOpts, provider []common.Address) (*Tricrypto2AddLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Tricrypto2.contract.FilterLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &Tricrypto2AddLiquidityIterator{contract: _Tricrypto2.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0x96b486485420b963edd3fdec0b0195730035600feb7de6f544383d7950fa97ee.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[3] token_amounts, uint256 fee, uint256 token_supply)
func (_Tricrypto2 *Tricrypto2Filterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *Tricrypto2AddLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Tricrypto2.contract.WatchLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Tricrypto2AddLiquidity)
				if err := _Tricrypto2.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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

// ParseAddLiquidity is a log parse operation binding the contract event 0x96b486485420b963edd3fdec0b0195730035600feb7de6f544383d7950fa97ee.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[3] token_amounts, uint256 fee, uint256 token_supply)
func (_Tricrypto2 *Tricrypto2Filterer) ParseAddLiquidity(log types.Log) (*Tricrypto2AddLiquidity, error) {
	event := new(Tricrypto2AddLiquidity)
	if err := _Tricrypto2.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Tricrypto2ClaimAdminFeeIterator is returned from FilterClaimAdminFee and is used to iterate over the raw logs and unpacked data for ClaimAdminFee events raised by the Tricrypto2 contract.
type Tricrypto2ClaimAdminFeeIterator struct {
	Event *Tricrypto2ClaimAdminFee // Event containing the contract specifics and raw log

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
func (it *Tricrypto2ClaimAdminFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Tricrypto2ClaimAdminFee)
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
		it.Event = new(Tricrypto2ClaimAdminFee)
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
func (it *Tricrypto2ClaimAdminFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Tricrypto2ClaimAdminFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Tricrypto2ClaimAdminFee represents a ClaimAdminFee event raised by the Tricrypto2 contract.
type Tricrypto2ClaimAdminFee struct {
	Admin  common.Address
	Tokens *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimAdminFee is a free log retrieval operation binding the contract event 0x6059a38198b1dc42b3791087d1ff0fbd72b3179553c25f678cd246f52ffaaf59.
//
// Solidity: event ClaimAdminFee(address indexed admin, uint256 tokens)
func (_Tricrypto2 *Tricrypto2Filterer) FilterClaimAdminFee(opts *bind.FilterOpts, admin []common.Address) (*Tricrypto2ClaimAdminFeeIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Tricrypto2.contract.FilterLogs(opts, "ClaimAdminFee", adminRule)
	if err != nil {
		return nil, err
	}
	return &Tricrypto2ClaimAdminFeeIterator{contract: _Tricrypto2.contract, event: "ClaimAdminFee", logs: logs, sub: sub}, nil
}

// WatchClaimAdminFee is a free log subscription operation binding the contract event 0x6059a38198b1dc42b3791087d1ff0fbd72b3179553c25f678cd246f52ffaaf59.
//
// Solidity: event ClaimAdminFee(address indexed admin, uint256 tokens)
func (_Tricrypto2 *Tricrypto2Filterer) WatchClaimAdminFee(opts *bind.WatchOpts, sink chan<- *Tricrypto2ClaimAdminFee, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Tricrypto2.contract.WatchLogs(opts, "ClaimAdminFee", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Tricrypto2ClaimAdminFee)
				if err := _Tricrypto2.contract.UnpackLog(event, "ClaimAdminFee", log); err != nil {
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

// ParseClaimAdminFee is a log parse operation binding the contract event 0x6059a38198b1dc42b3791087d1ff0fbd72b3179553c25f678cd246f52ffaaf59.
//
// Solidity: event ClaimAdminFee(address indexed admin, uint256 tokens)
func (_Tricrypto2 *Tricrypto2Filterer) ParseClaimAdminFee(log types.Log) (*Tricrypto2ClaimAdminFee, error) {
	event := new(Tricrypto2ClaimAdminFee)
	if err := _Tricrypto2.contract.UnpackLog(event, "ClaimAdminFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Tricrypto2CommitNewAdminIterator is returned from FilterCommitNewAdmin and is used to iterate over the raw logs and unpacked data for CommitNewAdmin events raised by the Tricrypto2 contract.
type Tricrypto2CommitNewAdminIterator struct {
	Event *Tricrypto2CommitNewAdmin // Event containing the contract specifics and raw log

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
func (it *Tricrypto2CommitNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Tricrypto2CommitNewAdmin)
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
		it.Event = new(Tricrypto2CommitNewAdmin)
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
func (it *Tricrypto2CommitNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Tricrypto2CommitNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Tricrypto2CommitNewAdmin represents a CommitNewAdmin event raised by the Tricrypto2 contract.
type Tricrypto2CommitNewAdmin struct {
	Deadline *big.Int
	Admin    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCommitNewAdmin is a free log retrieval operation binding the contract event 0x181aa3aa17d4cbf99265dd4443eba009433d3cde79d60164fde1d1a192beb935.
//
// Solidity: event CommitNewAdmin(uint256 indexed deadline, address indexed admin)
func (_Tricrypto2 *Tricrypto2Filterer) FilterCommitNewAdmin(opts *bind.FilterOpts, deadline []*big.Int, admin []common.Address) (*Tricrypto2CommitNewAdminIterator, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Tricrypto2.contract.FilterLogs(opts, "CommitNewAdmin", deadlineRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &Tricrypto2CommitNewAdminIterator{contract: _Tricrypto2.contract, event: "CommitNewAdmin", logs: logs, sub: sub}, nil
}

// WatchCommitNewAdmin is a free log subscription operation binding the contract event 0x181aa3aa17d4cbf99265dd4443eba009433d3cde79d60164fde1d1a192beb935.
//
// Solidity: event CommitNewAdmin(uint256 indexed deadline, address indexed admin)
func (_Tricrypto2 *Tricrypto2Filterer) WatchCommitNewAdmin(opts *bind.WatchOpts, sink chan<- *Tricrypto2CommitNewAdmin, deadline []*big.Int, admin []common.Address) (event.Subscription, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Tricrypto2.contract.WatchLogs(opts, "CommitNewAdmin", deadlineRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Tricrypto2CommitNewAdmin)
				if err := _Tricrypto2.contract.UnpackLog(event, "CommitNewAdmin", log); err != nil {
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
func (_Tricrypto2 *Tricrypto2Filterer) ParseCommitNewAdmin(log types.Log) (*Tricrypto2CommitNewAdmin, error) {
	event := new(Tricrypto2CommitNewAdmin)
	if err := _Tricrypto2.contract.UnpackLog(event, "CommitNewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Tricrypto2CommitNewParametersIterator is returned from FilterCommitNewParameters and is used to iterate over the raw logs and unpacked data for CommitNewParameters events raised by the Tricrypto2 contract.
type Tricrypto2CommitNewParametersIterator struct {
	Event *Tricrypto2CommitNewParameters // Event containing the contract specifics and raw log

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
func (it *Tricrypto2CommitNewParametersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Tricrypto2CommitNewParameters)
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
		it.Event = new(Tricrypto2CommitNewParameters)
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
func (it *Tricrypto2CommitNewParametersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Tricrypto2CommitNewParametersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Tricrypto2CommitNewParameters represents a CommitNewParameters event raised by the Tricrypto2 contract.
type Tricrypto2CommitNewParameters struct {
	Deadline           *big.Int
	AdminFee           *big.Int
	MidFee             *big.Int
	OutFee             *big.Int
	FeeGamma           *big.Int
	AllowedExtraProfit *big.Int
	AdjustmentStep     *big.Int
	MaHalfTime         *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterCommitNewParameters is a free log retrieval operation binding the contract event 0x913fde9a37e1f8ab67876a4d0ce80790d764fcfc5692f4529526df9c6bdde553.
//
// Solidity: event CommitNewParameters(uint256 indexed deadline, uint256 admin_fee, uint256 mid_fee, uint256 out_fee, uint256 fee_gamma, uint256 allowed_extra_profit, uint256 adjustment_step, uint256 ma_half_time)
func (_Tricrypto2 *Tricrypto2Filterer) FilterCommitNewParameters(opts *bind.FilterOpts, deadline []*big.Int) (*Tricrypto2CommitNewParametersIterator, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}

	logs, sub, err := _Tricrypto2.contract.FilterLogs(opts, "CommitNewParameters", deadlineRule)
	if err != nil {
		return nil, err
	}
	return &Tricrypto2CommitNewParametersIterator{contract: _Tricrypto2.contract, event: "CommitNewParameters", logs: logs, sub: sub}, nil
}

// WatchCommitNewParameters is a free log subscription operation binding the contract event 0x913fde9a37e1f8ab67876a4d0ce80790d764fcfc5692f4529526df9c6bdde553.
//
// Solidity: event CommitNewParameters(uint256 indexed deadline, uint256 admin_fee, uint256 mid_fee, uint256 out_fee, uint256 fee_gamma, uint256 allowed_extra_profit, uint256 adjustment_step, uint256 ma_half_time)
func (_Tricrypto2 *Tricrypto2Filterer) WatchCommitNewParameters(opts *bind.WatchOpts, sink chan<- *Tricrypto2CommitNewParameters, deadline []*big.Int) (event.Subscription, error) {

	var deadlineRule []interface{}
	for _, deadlineItem := range deadline {
		deadlineRule = append(deadlineRule, deadlineItem)
	}

	logs, sub, err := _Tricrypto2.contract.WatchLogs(opts, "CommitNewParameters", deadlineRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Tricrypto2CommitNewParameters)
				if err := _Tricrypto2.contract.UnpackLog(event, "CommitNewParameters", log); err != nil {
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

// ParseCommitNewParameters is a log parse operation binding the contract event 0x913fde9a37e1f8ab67876a4d0ce80790d764fcfc5692f4529526df9c6bdde553.
//
// Solidity: event CommitNewParameters(uint256 indexed deadline, uint256 admin_fee, uint256 mid_fee, uint256 out_fee, uint256 fee_gamma, uint256 allowed_extra_profit, uint256 adjustment_step, uint256 ma_half_time)
func (_Tricrypto2 *Tricrypto2Filterer) ParseCommitNewParameters(log types.Log) (*Tricrypto2CommitNewParameters, error) {
	event := new(Tricrypto2CommitNewParameters)
	if err := _Tricrypto2.contract.UnpackLog(event, "CommitNewParameters", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Tricrypto2NewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the Tricrypto2 contract.
type Tricrypto2NewAdminIterator struct {
	Event *Tricrypto2NewAdmin // Event containing the contract specifics and raw log

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
func (it *Tricrypto2NewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Tricrypto2NewAdmin)
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
		it.Event = new(Tricrypto2NewAdmin)
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
func (it *Tricrypto2NewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Tricrypto2NewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Tricrypto2NewAdmin represents a NewAdmin event raised by the Tricrypto2 contract.
type Tricrypto2NewAdmin struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address indexed admin)
func (_Tricrypto2 *Tricrypto2Filterer) FilterNewAdmin(opts *bind.FilterOpts, admin []common.Address) (*Tricrypto2NewAdminIterator, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Tricrypto2.contract.FilterLogs(opts, "NewAdmin", adminRule)
	if err != nil {
		return nil, err
	}
	return &Tricrypto2NewAdminIterator{contract: _Tricrypto2.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0x71614071b88dee5e0b2ae578a9dd7b2ebbe9ae832ba419dc0242cd065a290b6c.
//
// Solidity: event NewAdmin(address indexed admin)
func (_Tricrypto2 *Tricrypto2Filterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *Tricrypto2NewAdmin, admin []common.Address) (event.Subscription, error) {

	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Tricrypto2.contract.WatchLogs(opts, "NewAdmin", adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Tricrypto2NewAdmin)
				if err := _Tricrypto2.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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
func (_Tricrypto2 *Tricrypto2Filterer) ParseNewAdmin(log types.Log) (*Tricrypto2NewAdmin, error) {
	event := new(Tricrypto2NewAdmin)
	if err := _Tricrypto2.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Tricrypto2NewParametersIterator is returned from FilterNewParameters and is used to iterate over the raw logs and unpacked data for NewParameters events raised by the Tricrypto2 contract.
type Tricrypto2NewParametersIterator struct {
	Event *Tricrypto2NewParameters // Event containing the contract specifics and raw log

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
func (it *Tricrypto2NewParametersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Tricrypto2NewParameters)
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
		it.Event = new(Tricrypto2NewParameters)
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
func (it *Tricrypto2NewParametersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Tricrypto2NewParametersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Tricrypto2NewParameters represents a NewParameters event raised by the Tricrypto2 contract.
type Tricrypto2NewParameters struct {
	AdminFee           *big.Int
	MidFee             *big.Int
	OutFee             *big.Int
	FeeGamma           *big.Int
	AllowedExtraProfit *big.Int
	AdjustmentStep     *big.Int
	MaHalfTime         *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewParameters is a free log retrieval operation binding the contract event 0x1c65bbdc939f346e5d6f0bde1f072819947438d4fc7b182cc59c2f6dc5504087.
//
// Solidity: event NewParameters(uint256 admin_fee, uint256 mid_fee, uint256 out_fee, uint256 fee_gamma, uint256 allowed_extra_profit, uint256 adjustment_step, uint256 ma_half_time)
func (_Tricrypto2 *Tricrypto2Filterer) FilterNewParameters(opts *bind.FilterOpts) (*Tricrypto2NewParametersIterator, error) {

	logs, sub, err := _Tricrypto2.contract.FilterLogs(opts, "NewParameters")
	if err != nil {
		return nil, err
	}
	return &Tricrypto2NewParametersIterator{contract: _Tricrypto2.contract, event: "NewParameters", logs: logs, sub: sub}, nil
}

// WatchNewParameters is a free log subscription operation binding the contract event 0x1c65bbdc939f346e5d6f0bde1f072819947438d4fc7b182cc59c2f6dc5504087.
//
// Solidity: event NewParameters(uint256 admin_fee, uint256 mid_fee, uint256 out_fee, uint256 fee_gamma, uint256 allowed_extra_profit, uint256 adjustment_step, uint256 ma_half_time)
func (_Tricrypto2 *Tricrypto2Filterer) WatchNewParameters(opts *bind.WatchOpts, sink chan<- *Tricrypto2NewParameters) (event.Subscription, error) {

	logs, sub, err := _Tricrypto2.contract.WatchLogs(opts, "NewParameters")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Tricrypto2NewParameters)
				if err := _Tricrypto2.contract.UnpackLog(event, "NewParameters", log); err != nil {
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

// ParseNewParameters is a log parse operation binding the contract event 0x1c65bbdc939f346e5d6f0bde1f072819947438d4fc7b182cc59c2f6dc5504087.
//
// Solidity: event NewParameters(uint256 admin_fee, uint256 mid_fee, uint256 out_fee, uint256 fee_gamma, uint256 allowed_extra_profit, uint256 adjustment_step, uint256 ma_half_time)
func (_Tricrypto2 *Tricrypto2Filterer) ParseNewParameters(log types.Log) (*Tricrypto2NewParameters, error) {
	event := new(Tricrypto2NewParameters)
	if err := _Tricrypto2.contract.UnpackLog(event, "NewParameters", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Tricrypto2RampAgammaIterator is returned from FilterRampAgamma and is used to iterate over the raw logs and unpacked data for RampAgamma events raised by the Tricrypto2 contract.
type Tricrypto2RampAgammaIterator struct {
	Event *Tricrypto2RampAgamma // Event containing the contract specifics and raw log

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
func (it *Tricrypto2RampAgammaIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Tricrypto2RampAgamma)
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
		it.Event = new(Tricrypto2RampAgamma)
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
func (it *Tricrypto2RampAgammaIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Tricrypto2RampAgammaIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Tricrypto2RampAgamma represents a RampAgamma event raised by the Tricrypto2 contract.
type Tricrypto2RampAgamma struct {
	InitialA     *big.Int
	FutureA      *big.Int
	InitialGamma *big.Int
	FutureGamma  *big.Int
	InitialTime  *big.Int
	FutureTime   *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRampAgamma is a free log retrieval operation binding the contract event 0xe35f0559b0642164e286b30df2077ec3a05426617a25db7578fd20ba39a6cd05.
//
// Solidity: event RampAgamma(uint256 initial_A, uint256 future_A, uint256 initial_gamma, uint256 future_gamma, uint256 initial_time, uint256 future_time)
func (_Tricrypto2 *Tricrypto2Filterer) FilterRampAgamma(opts *bind.FilterOpts) (*Tricrypto2RampAgammaIterator, error) {

	logs, sub, err := _Tricrypto2.contract.FilterLogs(opts, "RampAgamma")
	if err != nil {
		return nil, err
	}
	return &Tricrypto2RampAgammaIterator{contract: _Tricrypto2.contract, event: "RampAgamma", logs: logs, sub: sub}, nil
}

// WatchRampAgamma is a free log subscription operation binding the contract event 0xe35f0559b0642164e286b30df2077ec3a05426617a25db7578fd20ba39a6cd05.
//
// Solidity: event RampAgamma(uint256 initial_A, uint256 future_A, uint256 initial_gamma, uint256 future_gamma, uint256 initial_time, uint256 future_time)
func (_Tricrypto2 *Tricrypto2Filterer) WatchRampAgamma(opts *bind.WatchOpts, sink chan<- *Tricrypto2RampAgamma) (event.Subscription, error) {

	logs, sub, err := _Tricrypto2.contract.WatchLogs(opts, "RampAgamma")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Tricrypto2RampAgamma)
				if err := _Tricrypto2.contract.UnpackLog(event, "RampAgamma", log); err != nil {
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

// ParseRampAgamma is a log parse operation binding the contract event 0xe35f0559b0642164e286b30df2077ec3a05426617a25db7578fd20ba39a6cd05.
//
// Solidity: event RampAgamma(uint256 initial_A, uint256 future_A, uint256 initial_gamma, uint256 future_gamma, uint256 initial_time, uint256 future_time)
func (_Tricrypto2 *Tricrypto2Filterer) ParseRampAgamma(log types.Log) (*Tricrypto2RampAgamma, error) {
	event := new(Tricrypto2RampAgamma)
	if err := _Tricrypto2.contract.UnpackLog(event, "RampAgamma", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Tricrypto2RemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the Tricrypto2 contract.
type Tricrypto2RemoveLiquidityIterator struct {
	Event *Tricrypto2RemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *Tricrypto2RemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Tricrypto2RemoveLiquidity)
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
		it.Event = new(Tricrypto2RemoveLiquidity)
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
func (it *Tricrypto2RemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Tricrypto2RemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Tricrypto2RemoveLiquidity represents a RemoveLiquidity event raised by the Tricrypto2 contract.
type Tricrypto2RemoveLiquidity struct {
	Provider     common.Address
	TokenAmounts [3]*big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0xd6cc314a0b1e3b2579f8e64248e82434072e8271290eef8ad0886709304195f5.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[3] token_amounts, uint256 token_supply)
func (_Tricrypto2 *Tricrypto2Filterer) FilterRemoveLiquidity(opts *bind.FilterOpts, provider []common.Address) (*Tricrypto2RemoveLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Tricrypto2.contract.FilterLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &Tricrypto2RemoveLiquidityIterator{contract: _Tricrypto2.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0xd6cc314a0b1e3b2579f8e64248e82434072e8271290eef8ad0886709304195f5.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[3] token_amounts, uint256 token_supply)
func (_Tricrypto2 *Tricrypto2Filterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *Tricrypto2RemoveLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Tricrypto2.contract.WatchLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Tricrypto2RemoveLiquidity)
				if err := _Tricrypto2.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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

// ParseRemoveLiquidity is a log parse operation binding the contract event 0xd6cc314a0b1e3b2579f8e64248e82434072e8271290eef8ad0886709304195f5.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[3] token_amounts, uint256 token_supply)
func (_Tricrypto2 *Tricrypto2Filterer) ParseRemoveLiquidity(log types.Log) (*Tricrypto2RemoveLiquidity, error) {
	event := new(Tricrypto2RemoveLiquidity)
	if err := _Tricrypto2.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Tricrypto2RemoveLiquidityOneIterator is returned from FilterRemoveLiquidityOne and is used to iterate over the raw logs and unpacked data for RemoveLiquidityOne events raised by the Tricrypto2 contract.
type Tricrypto2RemoveLiquidityOneIterator struct {
	Event *Tricrypto2RemoveLiquidityOne // Event containing the contract specifics and raw log

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
func (it *Tricrypto2RemoveLiquidityOneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Tricrypto2RemoveLiquidityOne)
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
		it.Event = new(Tricrypto2RemoveLiquidityOne)
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
func (it *Tricrypto2RemoveLiquidityOneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Tricrypto2RemoveLiquidityOneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Tricrypto2RemoveLiquidityOne represents a RemoveLiquidityOne event raised by the Tricrypto2 contract.
type Tricrypto2RemoveLiquidityOne struct {
	Provider    common.Address
	TokenAmount *big.Int
	CoinIndex   *big.Int
	CoinAmount  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidityOne is a free log retrieval operation binding the contract event 0x5ad056f2e28a8cec232015406b843668c1e36cda598127ec3b8c59b8c72773a0.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 token_amount, uint256 coin_index, uint256 coin_amount)
func (_Tricrypto2 *Tricrypto2Filterer) FilterRemoveLiquidityOne(opts *bind.FilterOpts, provider []common.Address) (*Tricrypto2RemoveLiquidityOneIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Tricrypto2.contract.FilterLogs(opts, "RemoveLiquidityOne", providerRule)
	if err != nil {
		return nil, err
	}
	return &Tricrypto2RemoveLiquidityOneIterator{contract: _Tricrypto2.contract, event: "RemoveLiquidityOne", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityOne is a free log subscription operation binding the contract event 0x5ad056f2e28a8cec232015406b843668c1e36cda598127ec3b8c59b8c72773a0.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 token_amount, uint256 coin_index, uint256 coin_amount)
func (_Tricrypto2 *Tricrypto2Filterer) WatchRemoveLiquidityOne(opts *bind.WatchOpts, sink chan<- *Tricrypto2RemoveLiquidityOne, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Tricrypto2.contract.WatchLogs(opts, "RemoveLiquidityOne", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Tricrypto2RemoveLiquidityOne)
				if err := _Tricrypto2.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
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

// ParseRemoveLiquidityOne is a log parse operation binding the contract event 0x5ad056f2e28a8cec232015406b843668c1e36cda598127ec3b8c59b8c72773a0.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 token_amount, uint256 coin_index, uint256 coin_amount)
func (_Tricrypto2 *Tricrypto2Filterer) ParseRemoveLiquidityOne(log types.Log) (*Tricrypto2RemoveLiquidityOne, error) {
	event := new(Tricrypto2RemoveLiquidityOne)
	if err := _Tricrypto2.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Tricrypto2StopRampAIterator is returned from FilterStopRampA and is used to iterate over the raw logs and unpacked data for StopRampA events raised by the Tricrypto2 contract.
type Tricrypto2StopRampAIterator struct {
	Event *Tricrypto2StopRampA // Event containing the contract specifics and raw log

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
func (it *Tricrypto2StopRampAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Tricrypto2StopRampA)
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
		it.Event = new(Tricrypto2StopRampA)
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
func (it *Tricrypto2StopRampAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Tricrypto2StopRampAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Tricrypto2StopRampA represents a StopRampA event raised by the Tricrypto2 contract.
type Tricrypto2StopRampA struct {
	CurrentA     *big.Int
	CurrentGamma *big.Int
	Time         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStopRampA is a free log retrieval operation binding the contract event 0x5f0e7fba3d100c9e19446e1c92fe436f0a9a22fe99669360e4fdd6d3de2fc284.
//
// Solidity: event StopRampA(uint256 current_A, uint256 current_gamma, uint256 time)
func (_Tricrypto2 *Tricrypto2Filterer) FilterStopRampA(opts *bind.FilterOpts) (*Tricrypto2StopRampAIterator, error) {

	logs, sub, err := _Tricrypto2.contract.FilterLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return &Tricrypto2StopRampAIterator{contract: _Tricrypto2.contract, event: "StopRampA", logs: logs, sub: sub}, nil
}

// WatchStopRampA is a free log subscription operation binding the contract event 0x5f0e7fba3d100c9e19446e1c92fe436f0a9a22fe99669360e4fdd6d3de2fc284.
//
// Solidity: event StopRampA(uint256 current_A, uint256 current_gamma, uint256 time)
func (_Tricrypto2 *Tricrypto2Filterer) WatchStopRampA(opts *bind.WatchOpts, sink chan<- *Tricrypto2StopRampA) (event.Subscription, error) {

	logs, sub, err := _Tricrypto2.contract.WatchLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Tricrypto2StopRampA)
				if err := _Tricrypto2.contract.UnpackLog(event, "StopRampA", log); err != nil {
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

// ParseStopRampA is a log parse operation binding the contract event 0x5f0e7fba3d100c9e19446e1c92fe436f0a9a22fe99669360e4fdd6d3de2fc284.
//
// Solidity: event StopRampA(uint256 current_A, uint256 current_gamma, uint256 time)
func (_Tricrypto2 *Tricrypto2Filterer) ParseStopRampA(log types.Log) (*Tricrypto2StopRampA, error) {
	event := new(Tricrypto2StopRampA)
	if err := _Tricrypto2.contract.UnpackLog(event, "StopRampA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Tricrypto2TokenExchangeIterator is returned from FilterTokenExchange and is used to iterate over the raw logs and unpacked data for TokenExchange events raised by the Tricrypto2 contract.
type Tricrypto2TokenExchangeIterator struct {
	Event *Tricrypto2TokenExchange // Event containing the contract specifics and raw log

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
func (it *Tricrypto2TokenExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Tricrypto2TokenExchange)
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
		it.Event = new(Tricrypto2TokenExchange)
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
func (it *Tricrypto2TokenExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Tricrypto2TokenExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Tricrypto2TokenExchange represents a TokenExchange event raised by the Tricrypto2 contract.
type Tricrypto2TokenExchange struct {
	Buyer        common.Address
	SoldId       *big.Int
	TokensSold   *big.Int
	BoughtId     *big.Int
	TokensBought *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenExchange is a free log retrieval operation binding the contract event 0xb2e76ae99761dc136e598d4a629bb347eccb9532a5f8bbd72e18467c3c34cc98.
//
// Solidity: event TokenExchange(address indexed buyer, uint256 sold_id, uint256 tokens_sold, uint256 bought_id, uint256 tokens_bought)
func (_Tricrypto2 *Tricrypto2Filterer) FilterTokenExchange(opts *bind.FilterOpts, buyer []common.Address) (*Tricrypto2TokenExchangeIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Tricrypto2.contract.FilterLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return &Tricrypto2TokenExchangeIterator{contract: _Tricrypto2.contract, event: "TokenExchange", logs: logs, sub: sub}, nil
}

// WatchTokenExchange is a free log subscription operation binding the contract event 0xb2e76ae99761dc136e598d4a629bb347eccb9532a5f8bbd72e18467c3c34cc98.
//
// Solidity: event TokenExchange(address indexed buyer, uint256 sold_id, uint256 tokens_sold, uint256 bought_id, uint256 tokens_bought)
func (_Tricrypto2 *Tricrypto2Filterer) WatchTokenExchange(opts *bind.WatchOpts, sink chan<- *Tricrypto2TokenExchange, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Tricrypto2.contract.WatchLogs(opts, "TokenExchange", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Tricrypto2TokenExchange)
				if err := _Tricrypto2.contract.UnpackLog(event, "TokenExchange", log); err != nil {
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

// ParseTokenExchange is a log parse operation binding the contract event 0xb2e76ae99761dc136e598d4a629bb347eccb9532a5f8bbd72e18467c3c34cc98.
//
// Solidity: event TokenExchange(address indexed buyer, uint256 sold_id, uint256 tokens_sold, uint256 bought_id, uint256 tokens_bought)
func (_Tricrypto2 *Tricrypto2Filterer) ParseTokenExchange(log types.Log) (*Tricrypto2TokenExchange, error) {
	event := new(Tricrypto2TokenExchange)
	if err := _Tricrypto2.contract.UnpackLog(event, "TokenExchange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
