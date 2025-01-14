// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokena

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

// TokenaMetaData contains all meta data concerning the Tokena contract.
var TokenaMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405180604001604052806006815260200165546f6b656e4160d01b81525060405180604001604052806006815260200165546f6b656e4160d01b8152506200006a620000646200009e60201b60201c565b620000a2565b81516200007f906004906020850190620000f2565b50805162000095906005906020840190620000f2565b505050620001d5565b3390565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b828054620001009062000198565b90600052602060002090601f0160209004810192826200012457600085556200016f565b82601f106200013f57805160ff19168380011785556200016f565b828001600101855582156200016f579182015b828111156200016f57825182559160200191906001019062000152565b506200017d92915062000181565b5090565b5b808211156200017d576000815560010162000182565b600181811c90821680620001ad57607f821691505b60208210811415620001cf57634e487b7160e01b600052602260045260246000fd5b50919050565b610b1680620001e56000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c8063715018a611610097578063a457c2d711610066578063a457c2d7146101eb578063a9059cbb146101fe578063dd62ed3e14610211578063f2fde38b1461022457600080fd5b8063715018a6146101ab5780638da5cb5b146101b557806395d89b41146101d0578063a0712d68146101d857600080fd5b806323b872dd116100d357806323b872dd1461014d578063313ce56714610160578063395093511461016f57806370a082311461018257600080fd5b806306fdde03146100fa578063095ea7b31461011857806318160ddd1461013b575b600080fd5b610102610237565b60405161010f919061093a565b60405180910390f35b61012b6101263660046109ab565b6102c9565b604051901515815260200161010f565b6003545b60405190815260200161010f565b61012b61015b3660046109d5565b6102e1565b6040516012815260200161010f565b61012b61017d3660046109ab565b610305565b61013f610190366004610a11565b6001600160a01b031660009081526001602052604090205490565b6101b3610327565b005b6000546040516001600160a01b03909116815260200161010f565b61010261033b565b6101b36101e6366004610a33565b61034a565b61012b6101f93660046109ab565b610357565b61012b61020c3660046109ab565b6103d7565b61013f61021f366004610a4c565b6103e5565b6101b3610232366004610a11565b610410565b60606004805461024690610a7f565b80601f016020809104026020016040519081016040528092919081815260200182805461027290610a7f565b80156102bf5780601f10610294576101008083540402835291602001916102bf565b820191906000526020600020905b8154815290600101906020018083116102a257829003601f168201915b5050505050905090565b6000336102d7818585610486565b5060019392505050565b6000336102ef8582856105aa565b6102fa858585610624565b506001949350505050565b6000336102d781858561031883836103e5565b6103229190610aba565b610486565b61032f6107cf565b6103396000610829565b565b60606005805461024690610a7f565b6103543382610879565b50565b6000338161036582866103e5565b9050838110156103ca5760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b60648201526084015b60405180910390fd5b6102fa8286868403610486565b6000336102d7818585610624565b6001600160a01b03918216600090815260026020908152604080832093909416825291909152205490565b6104186107cf565b6001600160a01b03811661047d5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016103c1565b61035481610829565b6001600160a01b0383166104e85760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b60648201526084016103c1565b6001600160a01b0382166105495760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b60648201526084016103c1565b6001600160a01b0383811660008181526002602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b60006105b684846103e5565b9050600019811461061e57818110156106115760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e636500000060448201526064016103c1565b61061e8484848403610486565b50505050565b6001600160a01b0383166106885760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b60648201526084016103c1565b6001600160a01b0382166106ea5760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b60648201526084016103c1565b6001600160a01b038316600090815260016020526040902054818110156107625760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b60648201526084016103c1565b6001600160a01b0380851660008181526001602052604080822086860390559286168082529083902080548601905591517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906107c29086815260200190565b60405180910390a361061e565b6000546001600160a01b031633146103395760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103c1565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b0382166108cf5760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064016103c1565b80600360008282546108e19190610aba565b90915550506001600160a01b0382166000818152600160209081526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b600060208083528351808285015260005b818110156109675785810183015185820160400152820161094b565b81811115610979576000604083870101525b50601f01601f1916929092016040019392505050565b80356001600160a01b03811681146109a657600080fd5b919050565b600080604083850312156109be57600080fd5b6109c78361098f565b946020939093013593505050565b6000806000606084860312156109ea57600080fd5b6109f38461098f565b9250610a016020850161098f565b9150604084013590509250925092565b600060208284031215610a2357600080fd5b610a2c8261098f565b9392505050565b600060208284031215610a4557600080fd5b5035919050565b60008060408385031215610a5f57600080fd5b610a688361098f565b9150610a766020840161098f565b90509250929050565b600181811c90821680610a9357607f821691505b60208210811415610ab457634e487b7160e01b600052602260045260246000fd5b50919050565b60008219821115610adb57634e487b7160e01b600052601160045260246000fd5b50019056fea2646970667358221220038938a7d2c41b000cc9f7a1cd07011efd58501900dcbc009f2d4dae9e92358664736f6c634300080a0033",
}

// TokenaABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenaMetaData.ABI instead.
var TokenaABI = TokenaMetaData.ABI

// TokenaBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TokenaMetaData.Bin instead.
var TokenaBin = TokenaMetaData.Bin

// DeployTokena deploys a new Ethereum contract, binding an instance of Tokena to it.
func DeployTokena(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Tokena, error) {
	parsed, err := TokenaMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenaBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Tokena{TokenaCaller: TokenaCaller{contract: contract}, TokenaTransactor: TokenaTransactor{contract: contract}, TokenaFilterer: TokenaFilterer{contract: contract}}, nil
}

// Tokena is an auto generated Go binding around an Ethereum contract.
type Tokena struct {
	TokenaCaller     // Read-only binding to the contract
	TokenaTransactor // Write-only binding to the contract
	TokenaFilterer   // Log filterer for contract events
}

// TokenaCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenaSession struct {
	Contract     *Tokena           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenaCallerSession struct {
	Contract *TokenaCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenaTransactorSession struct {
	Contract     *TokenaTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenaRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenaRaw struct {
	Contract *Tokena // Generic contract binding to access the raw methods on
}

// TokenaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenaCallerRaw struct {
	Contract *TokenaCaller // Generic read-only contract binding to access the raw methods on
}

// TokenaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenaTransactorRaw struct {
	Contract *TokenaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokena creates a new instance of Tokena, bound to a specific deployed contract.
func NewTokena(address common.Address, backend bind.ContractBackend) (*Tokena, error) {
	contract, err := bindTokena(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tokena{TokenaCaller: TokenaCaller{contract: contract}, TokenaTransactor: TokenaTransactor{contract: contract}, TokenaFilterer: TokenaFilterer{contract: contract}}, nil
}

// NewTokenaCaller creates a new read-only instance of Tokena, bound to a specific deployed contract.
func NewTokenaCaller(address common.Address, caller bind.ContractCaller) (*TokenaCaller, error) {
	contract, err := bindTokena(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenaCaller{contract: contract}, nil
}

// NewTokenaTransactor creates a new write-only instance of Tokena, bound to a specific deployed contract.
func NewTokenaTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenaTransactor, error) {
	contract, err := bindTokena(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenaTransactor{contract: contract}, nil
}

// NewTokenaFilterer creates a new log filterer instance of Tokena, bound to a specific deployed contract.
func NewTokenaFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenaFilterer, error) {
	contract, err := bindTokena(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenaFilterer{contract: contract}, nil
}

// bindTokena binds a generic wrapper to an already deployed contract.
func bindTokena(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenaMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokena *TokenaRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokena.Contract.TokenaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokena *TokenaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokena.Contract.TokenaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokena *TokenaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokena.Contract.TokenaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokena *TokenaCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokena.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokena *TokenaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokena.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokena *TokenaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokena.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Tokena *TokenaCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Tokena.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Tokena *TokenaSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Tokena.Contract.Allowance(&_Tokena.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Tokena *TokenaCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Tokena.Contract.Allowance(&_Tokena.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Tokena *TokenaCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Tokena.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Tokena *TokenaSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Tokena.Contract.BalanceOf(&_Tokena.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Tokena *TokenaCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Tokena.Contract.BalanceOf(&_Tokena.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Tokena *TokenaCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokena.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Tokena *TokenaSession) Decimals() (uint8, error) {
	return _Tokena.Contract.Decimals(&_Tokena.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Tokena *TokenaCallerSession) Decimals() (uint8, error) {
	return _Tokena.Contract.Decimals(&_Tokena.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Tokena *TokenaCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Tokena.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Tokena *TokenaSession) Name() (string, error) {
	return _Tokena.Contract.Name(&_Tokena.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Tokena *TokenaCallerSession) Name() (string, error) {
	return _Tokena.Contract.Name(&_Tokena.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tokena *TokenaCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokena.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tokena *TokenaSession) Owner() (common.Address, error) {
	return _Tokena.Contract.Owner(&_Tokena.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tokena *TokenaCallerSession) Owner() (common.Address, error) {
	return _Tokena.Contract.Owner(&_Tokena.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Tokena *TokenaCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Tokena.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Tokena *TokenaSession) Symbol() (string, error) {
	return _Tokena.Contract.Symbol(&_Tokena.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Tokena *TokenaCallerSession) Symbol() (string, error) {
	return _Tokena.Contract.Symbol(&_Tokena.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Tokena *TokenaCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokena.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Tokena *TokenaSession) TotalSupply() (*big.Int, error) {
	return _Tokena.Contract.TotalSupply(&_Tokena.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Tokena *TokenaCallerSession) TotalSupply() (*big.Int, error) {
	return _Tokena.Contract.TotalSupply(&_Tokena.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Tokena *TokenaTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokena.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Tokena *TokenaSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokena.Contract.Approve(&_Tokena.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Tokena *TokenaTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokena.Contract.Approve(&_Tokena.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Tokena *TokenaTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Tokena.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Tokena *TokenaSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Tokena.Contract.DecreaseAllowance(&_Tokena.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Tokena *TokenaTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Tokena.Contract.DecreaseAllowance(&_Tokena.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Tokena *TokenaTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Tokena.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Tokena *TokenaSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Tokena.Contract.IncreaseAllowance(&_Tokena.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Tokena *TokenaTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Tokena.Contract.IncreaseAllowance(&_Tokena.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_Tokena *TokenaTransactor) Mint(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Tokena.contract.Transact(opts, "mint", amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_Tokena *TokenaSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _Tokena.Contract.Mint(&_Tokena.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_Tokena *TokenaTransactorSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _Tokena.Contract.Mint(&_Tokena.TransactOpts, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tokena *TokenaTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokena.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tokena *TokenaSession) RenounceOwnership() (*types.Transaction, error) {
	return _Tokena.Contract.RenounceOwnership(&_Tokena.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tokena *TokenaTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Tokena.Contract.RenounceOwnership(&_Tokena.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Tokena *TokenaTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokena.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Tokena *TokenaSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokena.Contract.Transfer(&_Tokena.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Tokena *TokenaTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokena.Contract.Transfer(&_Tokena.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Tokena *TokenaTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokena.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Tokena *TokenaSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokena.Contract.TransferFrom(&_Tokena.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Tokena *TokenaTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokena.Contract.TransferFrom(&_Tokena.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tokena *TokenaTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Tokena.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tokena *TokenaSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Tokena.Contract.TransferOwnership(&_Tokena.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tokena *TokenaTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Tokena.Contract.TransferOwnership(&_Tokena.TransactOpts, newOwner)
}

// TokenaApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Tokena contract.
type TokenaApprovalIterator struct {
	Event *TokenaApproval // Event containing the contract specifics and raw log

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
func (it *TokenaApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenaApproval)
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
		it.Event = new(TokenaApproval)
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
func (it *TokenaApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenaApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenaApproval represents a Approval event raised by the Tokena contract.
type TokenaApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Tokena *TokenaFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TokenaApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Tokena.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TokenaApprovalIterator{contract: _Tokena.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Tokena *TokenaFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TokenaApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Tokena.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenaApproval)
				if err := _Tokena.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Tokena *TokenaFilterer) ParseApproval(log types.Log) (*TokenaApproval, error) {
	event := new(TokenaApproval)
	if err := _Tokena.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenaOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Tokena contract.
type TokenaOwnershipTransferredIterator struct {
	Event *TokenaOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenaOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenaOwnershipTransferred)
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
		it.Event = new(TokenaOwnershipTransferred)
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
func (it *TokenaOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenaOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenaOwnershipTransferred represents a OwnershipTransferred event raised by the Tokena contract.
type TokenaOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tokena *TokenaFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenaOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Tokena.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenaOwnershipTransferredIterator{contract: _Tokena.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tokena *TokenaFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenaOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Tokena.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenaOwnershipTransferred)
				if err := _Tokena.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tokena *TokenaFilterer) ParseOwnershipTransferred(log types.Log) (*TokenaOwnershipTransferred, error) {
	event := new(TokenaOwnershipTransferred)
	if err := _Tokena.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenaTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Tokena contract.
type TokenaTransferIterator struct {
	Event *TokenaTransfer // Event containing the contract specifics and raw log

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
func (it *TokenaTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenaTransfer)
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
		it.Event = new(TokenaTransfer)
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
func (it *TokenaTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenaTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenaTransfer represents a Transfer event raised by the Tokena contract.
type TokenaTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Tokena *TokenaFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenaTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Tokena.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenaTransferIterator{contract: _Tokena.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Tokena *TokenaFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenaTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Tokena.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenaTransfer)
				if err := _Tokena.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Tokena *TokenaFilterer) ParseTransfer(log types.Log) (*TokenaTransfer, error) {
	event := new(TokenaTransfer)
	if err := _Tokena.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
