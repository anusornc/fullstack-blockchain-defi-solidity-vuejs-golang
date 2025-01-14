// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package thepool

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

// ThepoolMetaData contains all meta data concerning the Thepool contract.
var ThepoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AddLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RemoveLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountAIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBOut\",\"type\":\"uint256\"}],\"name\":\"swapAForB\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"}],\"name\":\"getAmountBByA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountBIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAOut\",\"type\":\"uint256\"}],\"name\":\"swapBForA\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"name\":\"getAmountAByB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001fc438038062001fc4833981016040819052620000349162000209565b6040518060400160405280600781526020016626282a37b5b2b760c91b8152506040518060400160405280600781526020016626282a37b5b2b760c91b8152506200008e62000088620000f260201b60201c565b620000f6565b8151620000a390600490602085019062000146565b508051620000b990600590602084019062000146565b5050600680546001600160a01b039485166001600160a01b0319918216179091556007805493909416921691909117909155506200027e565b3390565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b828054620001549062000241565b90600052602060002090601f016020900481019282620001785760008555620001c3565b82601f106200019357805160ff1916838001178555620001c3565b82800160010185558215620001c3579182015b82811115620001c3578251825591602001919060010190620001a6565b50620001d1929150620001d5565b5090565b5b80821115620001d15760008155600101620001d6565b80516001600160a01b03811681146200020457600080fd5b919050565b600080604083850312156200021d57600080fd5b6200022883620001ec565b91506200023860208401620001ec565b90509250929050565b600181811c908216806200025657607f821691505b602082108114156200027857634e487b7160e01b600052602260045260246000fd5b50919050565b611d36806200028e6000396000f3fe608060405234801561001057600080fd5b506004361061012c5760003560e01c8063715018a6116100ad578063a457c2d711610071578063a457c2d71461026e578063a9059cbb14610281578063dab4e35414610294578063dd62ed3e146102a7578063f2fde38b146102ba57600080fd5b8063715018a61461021d5780638da5cb5b1461022557806395d89b41146102405780639c8f9f23146102485780639cd441da1461025b57600080fd5b806339509351116100f457806339509351146101a6578063553a1db7146101b95780635938c86b146101ce57806362d830df146101e157806370a08231146101f457600080fd5b806306fdde0314610131578063095ea7b31461014f57806318160ddd1461017257806323b872dd14610184578063313ce56714610197575b600080fd5b6101396102cd565b6040516101469190611a97565b60405180910390f35b61016261015d366004611b08565b61035f565b6040519015158152602001610146565b6003545b604051908152602001610146565b610162610192366004611b32565b610377565b60405160128152602001610146565b6101626101b4366004611b08565b61039b565b6101cc6101c7366004611b6e565b6103bd565b005b6101cc6101dc366004611b6e565b610762565b6101766101ef366004611b90565b610981565b610176610202366004611ba9565b6001600160a01b031660009081526001602052604090205490565b6101cc610b23565b6000546040516001600160a01b039091168152602001610146565b610139610b37565b6101cc610256366004611b90565b610b46565b6101cc610269366004611b6e565b610df8565b61016261027c366004611b08565b6111e2565b61016261028f366004611b08565b61125d565b6101766102a2366004611b90565b61126b565b6101766102b5366004611bc4565b6113fa565b6101cc6102c8366004611ba9565b611425565b6060600480546102dc90611bf7565b80601f016020809104026020016040519081016040528092919081815260200182805461030890611bf7565b80156103555780601f1061032a57610100808354040283529160200191610355565b820191906000526020600020905b81548152906001019060200180831161033857829003601f168201915b5050505050905090565b60003361036d81858561149e565b5060019392505050565b6000336103858582856115c3565b61039085858561163d565b506001949350505050565b60003361036d8185856103ae83836113fa565b6103b89190611c48565b61149e565b600082116104095760405162461bcd60e51b8152602060048201526014602482015273696e73756666696369656e7420616d6f756e744160601b60448201526064015b60405180910390fd5b600081116104505760405162461bcd60e51b815260206004820152601460248201527334b739bab33334b1b4b2b73a1030b6b7bab73a2160611b6044820152606401610400565b6006546040516370a0823160e01b81523060048201819052916000916001600160a01b03909116906370a0823190602401602060405180830381865afa15801561049e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104c29190611c60565b6007546040516370a0823160e01b81526001600160a01b038581166004830152929350600092909116906370a0823190602401602060405180830381865afa158015610512573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105369190611c60565b6006546040516323b872dd60e01b81529192506001600160a01b0316906323b872dd9061056b90339030908a90600401611c79565b6020604051808303816000875af115801561058a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105ae9190611c9d565b5060075460405163a9059cbb60e01b8152336004820152602481018690526001600160a01b039091169063a9059cbb906044015b6020604051808303816000875af1158015610601573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106259190611c9d565b506006546040516370a0823160e01b81526001600160a01b03858116600483015260009216906370a0823190602401602060405180830381865afa158015610671573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106959190611c60565b6007546040516370a0823160e01b81526001600160a01b038781166004830152929350600092909116906370a0823190602401602060405180830381865afa1580156106e5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107099190611c60565b90506107158385611cbf565b61071f8284611cbf565b10156107595760405162461bcd60e51b8152602060048201526009602482015268696e76616c6964204b60b81b6044820152606401610400565b50505050505050565b600082116107a95760405162461bcd60e51b815260206004820152601460248201527334b739bab33334b1b4b2b73a1030b6b7bab73a2160611b6044820152606401610400565b600081116107f05760405162461bcd60e51b8152602060048201526014602482015273696e73756666696369656e7420616d6f756e744160601b6044820152606401610400565b6006546040516370a0823160e01b81523060048201819052916000916001600160a01b03909116906370a0823190602401602060405180830381865afa15801561083e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108629190611c60565b6007546040516370a0823160e01b81526001600160a01b038581166004830152929350600092909116906370a0823190602401602060405180830381865afa1580156108b2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108d69190611c60565b60065460405163a9059cbb60e01b8152336004820152602481018790529192506001600160a01b03169063a9059cbb906044016020604051808303816000875af1158015610928573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061094c9190611c9d565b506007546040516323b872dd60e01b81526001600160a01b03909116906323b872dd906105e290339030908a90600401611c79565b60008082116109c85760405162461bcd60e51b81526020600482015260136024820152721a5b9cdd59999a58da595b9d08185b5bdd5b9d606a1b6044820152606401610400565b6006546040516370a0823160e01b81523060048201819052916000916001600160a01b03909116906370a0823190602401602060405180830381865afa158015610a16573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a3a9190611c60565b6007546040516370a0823160e01b81526001600160a01b038581166004830152929350600092909116906370a0823190602401602060405180830381865afa158015610a8a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aae9190611c60565b9050600082118015610ac05750600081115b610b055760405162461bcd60e51b8152602060048201526016602482015275696e73756666696369656e74206c697175696469747960501b6044820152606401610400565b80610b108387611cbf565b610b1a9190611cde565b95945050505050565b610b2b6117e8565b610b356000611842565b565b6060600580546102dc90611bf7565b33600090815260016020526040902054811115610b9c5760405162461bcd60e51b8152602060048201526014602482015273696e73756666696369656e742062616c616e636560601b6044820152606401610400565b6006546040516370a0823160e01b81523060048201819052916000916001600160a01b03909116906370a0823190602401602060405180830381865afa158015610bea573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c0e9190611c60565b6007546040516370a0823160e01b81526001600160a01b038581166004830152929350600092909116906370a0823190602401602060405180830381865afa158015610c5e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c829190611c60565b90506000610c8f60035490565b9050600081610c9e8588611cbf565b610ca89190611cde565b9050600082610cb78589611cbf565b610cc19190611cde565b60405188815290915033907fdfdd120ded9b7afc0c23dd5310e93aaa3e1c3b9f75af9b805fab3030842439f29060200160405180910390a2610d033388611892565b60065460405163a9059cbb60e01b8152336004820152602481018490526001600160a01b039091169063a9059cbb906044016020604051808303816000875af1158015610d54573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d789190611c9d565b5060075460405163a9059cbb60e01b8152336004820152602481018390526001600160a01b039091169063a9059cbb906044016020604051808303816000875af1158015610dca573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dee9190611c9d565b5050505050505050565b6006546040516370a0823160e01b81523060048201819052916000916001600160a01b03909116906370a0823190602401602060405180830381865afa158015610e46573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e6a9190611c60565b6007546040516370a0823160e01b81526001600160a01b038581166004830152929350600092909116906370a0823190602401602060405180830381865afa158015610eba573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ede9190611c60565b6006546040516323b872dd60e01b81529192506001600160a01b0316906323b872dd90610f1390339087908a90600401611c79565b6020604051808303816000875af1158015610f32573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f569190611c9d565b506007546040516323b872dd60e01b81526001600160a01b03909116906323b872dd90610f8b90339087908990600401611c79565b6020604051808303816000875af1158015610faa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fce9190611c9d565b506006546040516370a0823160e01b81526001600160a01b03858116600483015260009216906370a0823190602401602060405180830381865afa15801561101a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061103e9190611c60565b6007546040516370a0823160e01b81526001600160a01b038781166004830152929350600092909116906370a0823190602401602060405180830381865afa15801561108e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110b29190611c60565b90506110be8385611cbf565b6110c88284611cbf565b10156111025760405162461bcd60e51b8152602060048201526009602482015268696e76616c6964204b60b81b6044820152606401610400565b600061110d60035490565b905060008161111d575087611152565b61114f8661112b848c611cbf565b6111359190611cde565b86611140858c611cbf565b61114a9190611cde565b6119be565b90505b600081116111985760405162461bcd60e51b815260206004820152601360248201527206c70746f6b656e20616d6f756e74203d3d203606c1b6044820152606401610400565b60405181815233907fbb37879e252460856212dc4e8c6edf174e473df9423e3a7feccd21f8c28d587a9060200160405180910390a26111d733826119d6565b505050505050505050565b600033816111f082866113fa565b9050838110156112505760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b6064820152608401610400565b610390828686840361149e565b60003361036d81858561163d565b60008082116112b25760405162461bcd60e51b81526020600482015260136024820152721a5b9cdd59999a58da595b9d08185b5bdd5b9d606a1b6044820152606401610400565b6006546040516370a0823160e01b81523060048201819052916000916001600160a01b03909116906370a0823190602401602060405180830381865afa158015611300573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113249190611c60565b6007546040516370a0823160e01b81526001600160a01b038581166004830152929350600092909116906370a0823190602401602060405180830381865afa158015611374573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113989190611c60565b90506000821180156113aa5750600081115b6113ef5760405162461bcd60e51b8152602060048201526016602482015275696e73756666696369656e74206c697175696469747960501b6044820152606401610400565b81610b108287611cbf565b6001600160a01b03918216600090815260026020908152604080832093909416825291909152205490565b61142d6117e8565b6001600160a01b0381166114925760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610400565b61149b81611842565b50565b6001600160a01b0383166115005760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b6064820152608401610400565b6001600160a01b0382166115615760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b6064820152608401610400565b6001600160a01b0383811660008181526002602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b60006115cf84846113fa565b90506000198114611637578181101561162a5760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610400565b611637848484840361149e565b50505050565b6001600160a01b0383166116a15760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b6064820152608401610400565b6001600160a01b0382166117035760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b6064820152608401610400565b6001600160a01b0383166000908152600160205260409020548181101561177b5760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b6064820152608401610400565b6001600160a01b0380851660008181526001602052604080822086860390559286168082529083902080548601905591517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906117db9086815260200190565b60405180910390a3611637565b6000546001600160a01b03163314610b355760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610400565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b0382166118f25760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b6064820152608401610400565b6001600160a01b038216600090815260016020526040902054818110156119665760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b6064820152608401610400565b6001600160a01b03831660008181526001602090815260408083208686039055600380548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91016115b6565b60008183106119cd57816119cf565b825b9392505050565b6001600160a01b038216611a2c5760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610400565b8060036000828254611a3e9190611c48565b90915550506001600160a01b0382166000818152600160209081526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b600060208083528351808285015260005b81811015611ac457858101830151858201604001528201611aa8565b81811115611ad6576000604083870101525b50601f01601f1916929092016040019392505050565b80356001600160a01b0381168114611b0357600080fd5b919050565b60008060408385031215611b1b57600080fd5b611b2483611aec565b946020939093013593505050565b600080600060608486031215611b4757600080fd5b611b5084611aec565b9250611b5e60208501611aec565b9150604084013590509250925092565b60008060408385031215611b8157600080fd5b50508035926020909101359150565b600060208284031215611ba257600080fd5b5035919050565b600060208284031215611bbb57600080fd5b6119cf82611aec565b60008060408385031215611bd757600080fd5b611be083611aec565b9150611bee60208401611aec565b90509250929050565b600181811c90821680611c0b57607f821691505b60208210811415611c2c57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b60008219821115611c5b57611c5b611c32565b500190565b600060208284031215611c7257600080fd5b5051919050565b6001600160a01b039384168152919092166020820152604081019190915260600190565b600060208284031215611caf57600080fd5b815180151581146119cf57600080fd5b6000816000190483118215151615611cd957611cd9611c32565b500290565b600082611cfb57634e487b7160e01b600052601260045260246000fd5b50049056fea2646970667358221220ee6fea90ed5f7ff0c3cbdab926afcdf29bb96b35e8d204d3aad16a1c3e84e16364736f6c634300080a0033",
}

// ThepoolABI is the input ABI used to generate the binding from.
// Deprecated: Use ThepoolMetaData.ABI instead.
var ThepoolABI = ThepoolMetaData.ABI

// ThepoolBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ThepoolMetaData.Bin instead.
var ThepoolBin = ThepoolMetaData.Bin

// DeployThepool deploys a new Ethereum contract, binding an instance of Thepool to it.
func DeployThepool(auth *bind.TransactOpts, backend bind.ContractBackend, tokenA common.Address, tokenB common.Address) (common.Address, *types.Transaction, *Thepool, error) {
	parsed, err := ThepoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ThepoolBin), backend, tokenA, tokenB)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Thepool{ThepoolCaller: ThepoolCaller{contract: contract}, ThepoolTransactor: ThepoolTransactor{contract: contract}, ThepoolFilterer: ThepoolFilterer{contract: contract}}, nil
}

// Thepool is an auto generated Go binding around an Ethereum contract.
type Thepool struct {
	ThepoolCaller     // Read-only binding to the contract
	ThepoolTransactor // Write-only binding to the contract
	ThepoolFilterer   // Log filterer for contract events
}

// ThepoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type ThepoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ThepoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ThepoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ThepoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ThepoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ThepoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ThepoolSession struct {
	Contract     *Thepool          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ThepoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ThepoolCallerSession struct {
	Contract *ThepoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ThepoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ThepoolTransactorSession struct {
	Contract     *ThepoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ThepoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type ThepoolRaw struct {
	Contract *Thepool // Generic contract binding to access the raw methods on
}

// ThepoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ThepoolCallerRaw struct {
	Contract *ThepoolCaller // Generic read-only contract binding to access the raw methods on
}

// ThepoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ThepoolTransactorRaw struct {
	Contract *ThepoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewThepool creates a new instance of Thepool, bound to a specific deployed contract.
func NewThepool(address common.Address, backend bind.ContractBackend) (*Thepool, error) {
	contract, err := bindThepool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Thepool{ThepoolCaller: ThepoolCaller{contract: contract}, ThepoolTransactor: ThepoolTransactor{contract: contract}, ThepoolFilterer: ThepoolFilterer{contract: contract}}, nil
}

// NewThepoolCaller creates a new read-only instance of Thepool, bound to a specific deployed contract.
func NewThepoolCaller(address common.Address, caller bind.ContractCaller) (*ThepoolCaller, error) {
	contract, err := bindThepool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ThepoolCaller{contract: contract}, nil
}

// NewThepoolTransactor creates a new write-only instance of Thepool, bound to a specific deployed contract.
func NewThepoolTransactor(address common.Address, transactor bind.ContractTransactor) (*ThepoolTransactor, error) {
	contract, err := bindThepool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ThepoolTransactor{contract: contract}, nil
}

// NewThepoolFilterer creates a new log filterer instance of Thepool, bound to a specific deployed contract.
func NewThepoolFilterer(address common.Address, filterer bind.ContractFilterer) (*ThepoolFilterer, error) {
	contract, err := bindThepool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ThepoolFilterer{contract: contract}, nil
}

// bindThepool binds a generic wrapper to an already deployed contract.
func bindThepool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ThepoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Thepool *ThepoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Thepool.Contract.ThepoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Thepool *ThepoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Thepool.Contract.ThepoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Thepool *ThepoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Thepool.Contract.ThepoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Thepool *ThepoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Thepool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Thepool *ThepoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Thepool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Thepool *ThepoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Thepool.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Thepool *ThepoolCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Thepool.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Thepool *ThepoolSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Thepool.Contract.Allowance(&_Thepool.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Thepool *ThepoolCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Thepool.Contract.Allowance(&_Thepool.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Thepool *ThepoolCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Thepool.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Thepool *ThepoolSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Thepool.Contract.BalanceOf(&_Thepool.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Thepool *ThepoolCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Thepool.Contract.BalanceOf(&_Thepool.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Thepool *ThepoolCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Thepool.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Thepool *ThepoolSession) Decimals() (uint8, error) {
	return _Thepool.Contract.Decimals(&_Thepool.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Thepool *ThepoolCallerSession) Decimals() (uint8, error) {
	return _Thepool.Contract.Decimals(&_Thepool.CallOpts)
}

// GetAmountAByB is a free data retrieval call binding the contract method 0x62d830df.
//
// Solidity: function getAmountAByB(uint256 amountB) view returns(uint256)
func (_Thepool *ThepoolCaller) GetAmountAByB(opts *bind.CallOpts, amountB *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Thepool.contract.Call(opts, &out, "getAmountAByB", amountB)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountAByB is a free data retrieval call binding the contract method 0x62d830df.
//
// Solidity: function getAmountAByB(uint256 amountB) view returns(uint256)
func (_Thepool *ThepoolSession) GetAmountAByB(amountB *big.Int) (*big.Int, error) {
	return _Thepool.Contract.GetAmountAByB(&_Thepool.CallOpts, amountB)
}

// GetAmountAByB is a free data retrieval call binding the contract method 0x62d830df.
//
// Solidity: function getAmountAByB(uint256 amountB) view returns(uint256)
func (_Thepool *ThepoolCallerSession) GetAmountAByB(amountB *big.Int) (*big.Int, error) {
	return _Thepool.Contract.GetAmountAByB(&_Thepool.CallOpts, amountB)
}

// GetAmountBByA is a free data retrieval call binding the contract method 0xdab4e354.
//
// Solidity: function getAmountBByA(uint256 amountA) view returns(uint256)
func (_Thepool *ThepoolCaller) GetAmountBByA(opts *bind.CallOpts, amountA *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Thepool.contract.Call(opts, &out, "getAmountBByA", amountA)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountBByA is a free data retrieval call binding the contract method 0xdab4e354.
//
// Solidity: function getAmountBByA(uint256 amountA) view returns(uint256)
func (_Thepool *ThepoolSession) GetAmountBByA(amountA *big.Int) (*big.Int, error) {
	return _Thepool.Contract.GetAmountBByA(&_Thepool.CallOpts, amountA)
}

// GetAmountBByA is a free data retrieval call binding the contract method 0xdab4e354.
//
// Solidity: function getAmountBByA(uint256 amountA) view returns(uint256)
func (_Thepool *ThepoolCallerSession) GetAmountBByA(amountA *big.Int) (*big.Int, error) {
	return _Thepool.Contract.GetAmountBByA(&_Thepool.CallOpts, amountA)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Thepool *ThepoolCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Thepool.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Thepool *ThepoolSession) Name() (string, error) {
	return _Thepool.Contract.Name(&_Thepool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Thepool *ThepoolCallerSession) Name() (string, error) {
	return _Thepool.Contract.Name(&_Thepool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Thepool *ThepoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Thepool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Thepool *ThepoolSession) Owner() (common.Address, error) {
	return _Thepool.Contract.Owner(&_Thepool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Thepool *ThepoolCallerSession) Owner() (common.Address, error) {
	return _Thepool.Contract.Owner(&_Thepool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Thepool *ThepoolCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Thepool.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Thepool *ThepoolSession) Symbol() (string, error) {
	return _Thepool.Contract.Symbol(&_Thepool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Thepool *ThepoolCallerSession) Symbol() (string, error) {
	return _Thepool.Contract.Symbol(&_Thepool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Thepool *ThepoolCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Thepool.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Thepool *ThepoolSession) TotalSupply() (*big.Int, error) {
	return _Thepool.Contract.TotalSupply(&_Thepool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Thepool *ThepoolCallerSession) TotalSupply() (*big.Int, error) {
	return _Thepool.Contract.TotalSupply(&_Thepool.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9cd441da.
//
// Solidity: function addLiquidity(uint256 amountA, uint256 amountB) returns()
func (_Thepool *ThepoolTransactor) AddLiquidity(opts *bind.TransactOpts, amountA *big.Int, amountB *big.Int) (*types.Transaction, error) {
	return _Thepool.contract.Transact(opts, "addLiquidity", amountA, amountB)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9cd441da.
//
// Solidity: function addLiquidity(uint256 amountA, uint256 amountB) returns()
func (_Thepool *ThepoolSession) AddLiquidity(amountA *big.Int, amountB *big.Int) (*types.Transaction, error) {
	return _Thepool.Contract.AddLiquidity(&_Thepool.TransactOpts, amountA, amountB)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9cd441da.
//
// Solidity: function addLiquidity(uint256 amountA, uint256 amountB) returns()
func (_Thepool *ThepoolTransactorSession) AddLiquidity(amountA *big.Int, amountB *big.Int) (*types.Transaction, error) {
	return _Thepool.Contract.AddLiquidity(&_Thepool.TransactOpts, amountA, amountB)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Thepool *ThepoolTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Thepool.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Thepool *ThepoolSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Thepool.Contract.Approve(&_Thepool.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Thepool *ThepoolTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Thepool.Contract.Approve(&_Thepool.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Thepool *ThepoolTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Thepool.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Thepool *ThepoolSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Thepool.Contract.DecreaseAllowance(&_Thepool.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Thepool *ThepoolTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Thepool.Contract.DecreaseAllowance(&_Thepool.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Thepool *ThepoolTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Thepool.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Thepool *ThepoolSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Thepool.Contract.IncreaseAllowance(&_Thepool.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Thepool *ThepoolTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Thepool.Contract.IncreaseAllowance(&_Thepool.TransactOpts, spender, addedValue)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x9c8f9f23.
//
// Solidity: function removeLiquidity(uint256 amount) returns()
func (_Thepool *ThepoolTransactor) RemoveLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Thepool.contract.Transact(opts, "removeLiquidity", amount)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x9c8f9f23.
//
// Solidity: function removeLiquidity(uint256 amount) returns()
func (_Thepool *ThepoolSession) RemoveLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _Thepool.Contract.RemoveLiquidity(&_Thepool.TransactOpts, amount)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x9c8f9f23.
//
// Solidity: function removeLiquidity(uint256 amount) returns()
func (_Thepool *ThepoolTransactorSession) RemoveLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _Thepool.Contract.RemoveLiquidity(&_Thepool.TransactOpts, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Thepool *ThepoolTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Thepool.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Thepool *ThepoolSession) RenounceOwnership() (*types.Transaction, error) {
	return _Thepool.Contract.RenounceOwnership(&_Thepool.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Thepool *ThepoolTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Thepool.Contract.RenounceOwnership(&_Thepool.TransactOpts)
}

// SwapAForB is a paid mutator transaction binding the contract method 0x553a1db7.
//
// Solidity: function swapAForB(uint256 amountAIn, uint256 amountBOut) returns()
func (_Thepool *ThepoolTransactor) SwapAForB(opts *bind.TransactOpts, amountAIn *big.Int, amountBOut *big.Int) (*types.Transaction, error) {
	return _Thepool.contract.Transact(opts, "swapAForB", amountAIn, amountBOut)
}

// SwapAForB is a paid mutator transaction binding the contract method 0x553a1db7.
//
// Solidity: function swapAForB(uint256 amountAIn, uint256 amountBOut) returns()
func (_Thepool *ThepoolSession) SwapAForB(amountAIn *big.Int, amountBOut *big.Int) (*types.Transaction, error) {
	return _Thepool.Contract.SwapAForB(&_Thepool.TransactOpts, amountAIn, amountBOut)
}

// SwapAForB is a paid mutator transaction binding the contract method 0x553a1db7.
//
// Solidity: function swapAForB(uint256 amountAIn, uint256 amountBOut) returns()
func (_Thepool *ThepoolTransactorSession) SwapAForB(amountAIn *big.Int, amountBOut *big.Int) (*types.Transaction, error) {
	return _Thepool.Contract.SwapAForB(&_Thepool.TransactOpts, amountAIn, amountBOut)
}

// SwapBForA is a paid mutator transaction binding the contract method 0x5938c86b.
//
// Solidity: function swapBForA(uint256 amountBIn, uint256 amountAOut) returns()
func (_Thepool *ThepoolTransactor) SwapBForA(opts *bind.TransactOpts, amountBIn *big.Int, amountAOut *big.Int) (*types.Transaction, error) {
	return _Thepool.contract.Transact(opts, "swapBForA", amountBIn, amountAOut)
}

// SwapBForA is a paid mutator transaction binding the contract method 0x5938c86b.
//
// Solidity: function swapBForA(uint256 amountBIn, uint256 amountAOut) returns()
func (_Thepool *ThepoolSession) SwapBForA(amountBIn *big.Int, amountAOut *big.Int) (*types.Transaction, error) {
	return _Thepool.Contract.SwapBForA(&_Thepool.TransactOpts, amountBIn, amountAOut)
}

// SwapBForA is a paid mutator transaction binding the contract method 0x5938c86b.
//
// Solidity: function swapBForA(uint256 amountBIn, uint256 amountAOut) returns()
func (_Thepool *ThepoolTransactorSession) SwapBForA(amountBIn *big.Int, amountAOut *big.Int) (*types.Transaction, error) {
	return _Thepool.Contract.SwapBForA(&_Thepool.TransactOpts, amountBIn, amountAOut)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Thepool *ThepoolTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Thepool.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Thepool *ThepoolSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Thepool.Contract.Transfer(&_Thepool.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Thepool *ThepoolTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Thepool.Contract.Transfer(&_Thepool.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Thepool *ThepoolTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Thepool.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Thepool *ThepoolSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Thepool.Contract.TransferFrom(&_Thepool.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Thepool *ThepoolTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Thepool.Contract.TransferFrom(&_Thepool.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Thepool *ThepoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Thepool.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Thepool *ThepoolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Thepool.Contract.TransferOwnership(&_Thepool.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Thepool *ThepoolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Thepool.Contract.TransferOwnership(&_Thepool.TransactOpts, newOwner)
}

// ThepoolAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the Thepool contract.
type ThepoolAddLiquidityIterator struct {
	Event *ThepoolAddLiquidity // Event containing the contract specifics and raw log

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
func (it *ThepoolAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThepoolAddLiquidity)
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
		it.Event = new(ThepoolAddLiquidity)
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
func (it *ThepoolAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThepoolAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThepoolAddLiquidity represents a AddLiquidity event raised by the Thepool contract.
type ThepoolAddLiquidity struct {
	Provider common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0xbb37879e252460856212dc4e8c6edf174e473df9423e3a7feccd21f8c28d587a.
//
// Solidity: event AddLiquidity(address indexed provider, uint256 amount)
func (_Thepool *ThepoolFilterer) FilterAddLiquidity(opts *bind.FilterOpts, provider []common.Address) (*ThepoolAddLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Thepool.contract.FilterLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &ThepoolAddLiquidityIterator{contract: _Thepool.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0xbb37879e252460856212dc4e8c6edf174e473df9423e3a7feccd21f8c28d587a.
//
// Solidity: event AddLiquidity(address indexed provider, uint256 amount)
func (_Thepool *ThepoolFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *ThepoolAddLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Thepool.contract.WatchLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThepoolAddLiquidity)
				if err := _Thepool.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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

// ParseAddLiquidity is a log parse operation binding the contract event 0xbb37879e252460856212dc4e8c6edf174e473df9423e3a7feccd21f8c28d587a.
//
// Solidity: event AddLiquidity(address indexed provider, uint256 amount)
func (_Thepool *ThepoolFilterer) ParseAddLiquidity(log types.Log) (*ThepoolAddLiquidity, error) {
	event := new(ThepoolAddLiquidity)
	if err := _Thepool.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThepoolApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Thepool contract.
type ThepoolApprovalIterator struct {
	Event *ThepoolApproval // Event containing the contract specifics and raw log

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
func (it *ThepoolApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThepoolApproval)
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
		it.Event = new(ThepoolApproval)
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
func (it *ThepoolApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThepoolApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThepoolApproval represents a Approval event raised by the Thepool contract.
type ThepoolApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Thepool *ThepoolFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ThepoolApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Thepool.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ThepoolApprovalIterator{contract: _Thepool.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Thepool *ThepoolFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ThepoolApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Thepool.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThepoolApproval)
				if err := _Thepool.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Thepool *ThepoolFilterer) ParseApproval(log types.Log) (*ThepoolApproval, error) {
	event := new(ThepoolApproval)
	if err := _Thepool.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThepoolOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Thepool contract.
type ThepoolOwnershipTransferredIterator struct {
	Event *ThepoolOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ThepoolOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThepoolOwnershipTransferred)
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
		it.Event = new(ThepoolOwnershipTransferred)
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
func (it *ThepoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThepoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThepoolOwnershipTransferred represents a OwnershipTransferred event raised by the Thepool contract.
type ThepoolOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Thepool *ThepoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ThepoolOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Thepool.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ThepoolOwnershipTransferredIterator{contract: _Thepool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Thepool *ThepoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ThepoolOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Thepool.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThepoolOwnershipTransferred)
				if err := _Thepool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Thepool *ThepoolFilterer) ParseOwnershipTransferred(log types.Log) (*ThepoolOwnershipTransferred, error) {
	event := new(ThepoolOwnershipTransferred)
	if err := _Thepool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThepoolRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the Thepool contract.
type ThepoolRemoveLiquidityIterator struct {
	Event *ThepoolRemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *ThepoolRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThepoolRemoveLiquidity)
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
		it.Event = new(ThepoolRemoveLiquidity)
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
func (it *ThepoolRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThepoolRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThepoolRemoveLiquidity represents a RemoveLiquidity event raised by the Thepool contract.
type ThepoolRemoveLiquidity struct {
	Provider common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0xdfdd120ded9b7afc0c23dd5310e93aaa3e1c3b9f75af9b805fab3030842439f2.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256 amount)
func (_Thepool *ThepoolFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts, provider []common.Address) (*ThepoolRemoveLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Thepool.contract.FilterLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &ThepoolRemoveLiquidityIterator{contract: _Thepool.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0xdfdd120ded9b7afc0c23dd5310e93aaa3e1c3b9f75af9b805fab3030842439f2.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256 amount)
func (_Thepool *ThepoolFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *ThepoolRemoveLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Thepool.contract.WatchLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThepoolRemoveLiquidity)
				if err := _Thepool.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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

// ParseRemoveLiquidity is a log parse operation binding the contract event 0xdfdd120ded9b7afc0c23dd5310e93aaa3e1c3b9f75af9b805fab3030842439f2.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256 amount)
func (_Thepool *ThepoolFilterer) ParseRemoveLiquidity(log types.Log) (*ThepoolRemoveLiquidity, error) {
	event := new(ThepoolRemoveLiquidity)
	if err := _Thepool.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThepoolTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Thepool contract.
type ThepoolTransferIterator struct {
	Event *ThepoolTransfer // Event containing the contract specifics and raw log

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
func (it *ThepoolTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ThepoolTransfer)
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
		it.Event = new(ThepoolTransfer)
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
func (it *ThepoolTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ThepoolTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ThepoolTransfer represents a Transfer event raised by the Thepool contract.
type ThepoolTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Thepool *ThepoolFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ThepoolTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Thepool.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ThepoolTransferIterator{contract: _Thepool.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Thepool *ThepoolFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ThepoolTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Thepool.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ThepoolTransfer)
				if err := _Thepool.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Thepool *ThepoolFilterer) ParseTransfer(log types.Log) (*ThepoolTransfer, error) {
	event := new(ThepoolTransfer)
	if err := _Thepool.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
