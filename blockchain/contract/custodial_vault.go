// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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
)

// CustosialVaultMetaData contains all meta data concerning the CustosialVault contract.
var CustosialVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractERC721Enumerable\",\"name\":\"_nft\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"}],\"name\":\"NFTCustody\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"emergencyDelete\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"holdCustody\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"releaseNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"retainNFT\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"txCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newHolder\",\"type\":\"address\"}],\"name\":\"updateOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60806040526544364c5bb0006002553480156200001b57600080fd5b5060405162001531380380620015318339818101604052810190620000419190620001fd565b6001600081905550620000696200005d620000b160201b60201c565b620000b960201b60201c565b80600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506200022f565b600033905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620001b18262000184565b9050919050565b6000620001c582620001a4565b9050919050565b620001d781620001b8565b8114620001e357600080fd5b50565b600081519050620001f781620001cc565b92915050565b6000602082840312156200021657620002156200017f565b5b60006200022684828501620001e6565b91505092915050565b6112f2806200023f6000396000f3fe60806040526004361061009b5760003560e01c80637192711f116100645780637192711f146101505780638da5cb5b14610179578063b2a7e29d146101a4578063ccf678d8146101cf578063dd39bd21146101eb578063f2fde38b146102295761009b565b8062d90cd9146100a0578063150b7a02146100c95780633ccfd60b1461010657806344509d3e14610110578063715018a614610139575b600080fd5b3480156100ac57600080fd5b506100c760048036038101906100c29190610c14565b610252565b005b3480156100d557600080fd5b506100f060048036038101906100eb9190610d04565b6102f7565b6040516100fd9190610dc7565b60405180910390f35b61010e61037a565b005b34801561011c57600080fd5b5061013760048036038101906101329190610de2565b6103c2565b005b34801561014557600080fd5b5061014e6104f9565b005b34801561015c57600080fd5b5061017760048036038101906101729190610de2565b61050d565b005b34801561018557600080fd5b5061018e610638565b60405161019b9190610e31565b60405180910390f35b3480156101b057600080fd5b506101b9610662565b6040516101c69190610e5b565b60405180910390f35b6101e960048036038101906101e49190610c14565b610668565b005b3480156101f757600080fd5b50610212600480360381019061020d9190610c14565b6109c1565b604051610220929190610e76565b60405180910390f35b34801561023557600080fd5b50610250600480360381019061024b9190610e9f565b610a05565b005b600260005403610297576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161028e90610f29565b60405180910390fd5b60026000819055506102a7610a88565b600360008281526020019081526020016000206000808201600090556001820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690555050600160008190555050565b60008073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614610367576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161035e90610f95565b60405180910390fd5b63150b7a0260e01b905095945050505050565b610382610a88565b3373ffffffffffffffffffffffffffffffffffffffff166108fc479081150290604051600060405180830381858888f193505050506103c057600080fd5b565b600260005403610407576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103fe90610f29565b60405180910390fd5b6002600081905550610417610a88565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3083856040518463ffffffff1660e01b815260040161047693929190610fb5565b600060405180830381600087803b15801561049057600080fd5b505af11580156104a4573d6000803e3d6000fd5b50505050600360008381526020019081526020016000206000808201600090556001820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055505060016000819055505050565b610501610a88565b61050b6000610b06565b565b600260005403610552576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161054990610f29565b60405180910390fd5b6002600081905550610562610a88565b60405180604001604052808381526020018273ffffffffffffffffffffffffffffffffffffffff16815250600360008481526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550905050817f2eae7534e9f0d5d8a801dd79515d22a9c36f0b691692ca20f5bc04b94b055edc826040516106249190610e31565b60405180910390a260016000819055505050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60025481565b6002600054036106ad576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106a490610f29565b60405180910390fd5b600260008190555060025434146106f9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106f09061105e565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff16600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636352211e836040518263ffffffff1660e01b815260040161076b9190610e5b565b602060405180830381865afa158015610788573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107ac9190611093565b73ffffffffffffffffffffffffffffffffffffffff1614610802576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107f990611132565b60405180910390fd5b600060036000838152602001908152602001600020600001541461085b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108529061119e565b60405180910390fd5b60405180604001604052808281526020013373ffffffffffffffffffffffffffffffffffffffff16815250600360008381526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550905050600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330846040518463ffffffff1660e01b815260040161094c93929190610fb5565b600060405180830381600087803b15801561096657600080fd5b505af115801561097a573d6000803e3d6000fd5b50505050807f2eae7534e9f0d5d8a801dd79515d22a9c36f0b691692ca20f5bc04b94b055edc336040516109ae9190610e31565b60405180910390a2600160008190555050565b60036020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905082565b610a0d610a88565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610a7c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a7390611230565b60405180910390fd5b610a8581610b06565b50565b610a90610bcc565b73ffffffffffffffffffffffffffffffffffffffff16610aae610638565b73ffffffffffffffffffffffffffffffffffffffff1614610b04576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610afb9061129c565b60405180910390fd5b565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600033905090565b600080fd5b600080fd5b6000819050919050565b610bf181610bde565b8114610bfc57600080fd5b50565b600081359050610c0e81610be8565b92915050565b600060208284031215610c2a57610c29610bd4565b5b6000610c3884828501610bff565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610c6c82610c41565b9050919050565b610c7c81610c61565b8114610c8757600080fd5b50565b600081359050610c9981610c73565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f840112610cc457610cc3610c9f565b5b8235905067ffffffffffffffff811115610ce157610ce0610ca4565b5b602083019150836001820283011115610cfd57610cfc610ca9565b5b9250929050565b600080600080600060808688031215610d2057610d1f610bd4565b5b6000610d2e88828901610c8a565b9550506020610d3f88828901610c8a565b9450506040610d5088828901610bff565b935050606086013567ffffffffffffffff811115610d7157610d70610bd9565b5b610d7d88828901610cae565b92509250509295509295909350565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b610dc181610d8c565b82525050565b6000602082019050610ddc6000830184610db8565b92915050565b60008060408385031215610df957610df8610bd4565b5b6000610e0785828601610bff565b9250506020610e1885828601610c8a565b9150509250929050565b610e2b81610c61565b82525050565b6000602082019050610e466000830184610e22565b92915050565b610e5581610bde565b82525050565b6000602082019050610e706000830184610e4c565b92915050565b6000604082019050610e8b6000830185610e4c565b610e986020830184610e22565b9392505050565b600060208284031215610eb557610eb4610bd4565b5b6000610ec384828501610c8a565b91505092915050565b600082825260208201905092915050565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00600082015250565b6000610f13601f83610ecc565b9150610f1e82610edd565b602082019050919050565b60006020820190508181036000830152610f4281610f06565b9050919050565b7f43616e6e6f742052656365697665204e465473204469726563746c7900000000600082015250565b6000610f7f601c83610ecc565b9150610f8a82610f49565b602082019050919050565b60006020820190508181036000830152610fae81610f72565b9050919050565b6000606082019050610fca6000830186610e22565b610fd76020830185610e22565b610fe46040830184610e4c565b949350505050565b7f4e6f7420656e6f7567682062616c616e636520746f20636f6d706c657465207460008201527f72616e73616374696f6e2e000000000000000000000000000000000000000000602082015250565b6000611048602b83610ecc565b915061105382610fec565b604082019050919050565b600060208201905081810360008301526110778161103b565b9050919050565b60008151905061108d81610c73565b92915050565b6000602082840312156110a9576110a8610bd4565b5b60006110b78482850161107e565b91505092915050565b7f796f75206d75737420626520746865204e4654206f776e657220666f7220657860008201527f65637574696e67207468697320616374696f6e00000000000000000000000000602082015250565b600061111c603383610ecc565b9150611127826110c0565b604082019050919050565b6000602082019050818103600083015261114b8161110f565b9050919050565b7f4e465420616c72656164792073746f7265640000000000000000000000000000600082015250565b6000611188601283610ecc565b915061119382611152565b602082019050919050565b600060208201905081810360008301526111b78161117b565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b600061121a602683610ecc565b9150611225826111be565b604082019050919050565b600060208201905081810360008301526112498161120d565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000611286602083610ecc565b915061129182611250565b602082019050919050565b600060208201905081810360008301526112b581611279565b905091905056fea2646970667358221220b226d5e7c5e7cff59ee300e312d6c41a2d1ae437d316ed8a8b076f3189b94d8264736f6c63430008100033",
}

// CustosialVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use CustosialVaultMetaData.ABI instead.
var CustosialVaultABI = CustosialVaultMetaData.ABI

// CustosialVaultBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CustosialVaultMetaData.Bin instead.
var CustosialVaultBin = CustosialVaultMetaData.Bin

// DeployCustosialVault deploys a new Ethereum contract, binding an instance of CustosialVault to it.
func DeployCustosialVault(auth *bind.TransactOpts, backend bind.ContractBackend, _nft common.Address) (common.Address, *types.Transaction, *CustosialVault, error) {
	parsed, err := CustosialVaultMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CustosialVaultBin), backend, _nft)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CustosialVault{CustosialVaultCaller: CustosialVaultCaller{contract: contract}, CustosialVaultTransactor: CustosialVaultTransactor{contract: contract}, CustosialVaultFilterer: CustosialVaultFilterer{contract: contract}}, nil
}

// CustosialVault is an auto generated Go binding around an Ethereum contract.
type CustosialVault struct {
	CustosialVaultCaller     // Read-only binding to the contract
	CustosialVaultTransactor // Write-only binding to the contract
	CustosialVaultFilterer   // Log filterer for contract events
}

// CustosialVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type CustosialVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CustosialVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CustosialVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CustosialVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CustosialVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CustosialVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CustosialVaultSession struct {
	Contract     *CustosialVault   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CustosialVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CustosialVaultCallerSession struct {
	Contract *CustosialVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// CustosialVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CustosialVaultTransactorSession struct {
	Contract     *CustosialVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// CustosialVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type CustosialVaultRaw struct {
	Contract *CustosialVault // Generic contract binding to access the raw methods on
}

// CustosialVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CustosialVaultCallerRaw struct {
	Contract *CustosialVaultCaller // Generic read-only contract binding to access the raw methods on
}

// CustosialVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CustosialVaultTransactorRaw struct {
	Contract *CustosialVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCustosialVault creates a new instance of CustosialVault, bound to a specific deployed contract.
func NewCustosialVault(address common.Address, backend bind.ContractBackend) (*CustosialVault, error) {
	contract, err := bindCustosialVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CustosialVault{CustosialVaultCaller: CustosialVaultCaller{contract: contract}, CustosialVaultTransactor: CustosialVaultTransactor{contract: contract}, CustosialVaultFilterer: CustosialVaultFilterer{contract: contract}}, nil
}

// NewCustosialVaultCaller creates a new read-only instance of CustosialVault, bound to a specific deployed contract.
func NewCustosialVaultCaller(address common.Address, caller bind.ContractCaller) (*CustosialVaultCaller, error) {
	contract, err := bindCustosialVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CustosialVaultCaller{contract: contract}, nil
}

// NewCustosialVaultTransactor creates a new write-only instance of CustosialVault, bound to a specific deployed contract.
func NewCustosialVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*CustosialVaultTransactor, error) {
	contract, err := bindCustosialVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CustosialVaultTransactor{contract: contract}, nil
}

// NewCustosialVaultFilterer creates a new log filterer instance of CustosialVault, bound to a specific deployed contract.
func NewCustosialVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*CustosialVaultFilterer, error) {
	contract, err := bindCustosialVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CustosialVaultFilterer{contract: contract}, nil
}

// bindCustosialVault binds a generic wrapper to an already deployed contract.
func bindCustosialVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CustosialVaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CustosialVault *CustosialVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CustosialVault.Contract.CustosialVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CustosialVault *CustosialVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustosialVault.Contract.CustosialVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CustosialVault *CustosialVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CustosialVault.Contract.CustosialVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CustosialVault *CustosialVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CustosialVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CustosialVault *CustosialVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustosialVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CustosialVault *CustosialVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CustosialVault.Contract.contract.Transact(opts, method, params...)
}

// HoldCustody is a free data retrieval call binding the contract method 0xdd39bd21.
//
// Solidity: function holdCustody(uint256 ) view returns(uint256 tokenId, address holder)
func (_CustosialVault *CustosialVaultCaller) HoldCustody(opts *bind.CallOpts, arg0 *big.Int) (struct {
	TokenId *big.Int
	Holder  common.Address
}, error) {
	var out []interface{}
	err := _CustosialVault.contract.Call(opts, &out, "holdCustody", arg0)

	outstruct := new(struct {
		TokenId *big.Int
		Holder  common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Holder = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// HoldCustody is a free data retrieval call binding the contract method 0xdd39bd21.
//
// Solidity: function holdCustody(uint256 ) view returns(uint256 tokenId, address holder)
func (_CustosialVault *CustosialVaultSession) HoldCustody(arg0 *big.Int) (struct {
	TokenId *big.Int
	Holder  common.Address
}, error) {
	return _CustosialVault.Contract.HoldCustody(&_CustosialVault.CallOpts, arg0)
}

// HoldCustody is a free data retrieval call binding the contract method 0xdd39bd21.
//
// Solidity: function holdCustody(uint256 ) view returns(uint256 tokenId, address holder)
func (_CustosialVault *CustosialVaultCallerSession) HoldCustody(arg0 *big.Int) (struct {
	TokenId *big.Int
	Holder  common.Address
}, error) {
	return _CustosialVault.Contract.HoldCustody(&_CustosialVault.CallOpts, arg0)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address from, uint256 , bytes ) pure returns(bytes4)
func (_CustosialVault *CustosialVaultCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, from common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _CustosialVault.contract.Call(opts, &out, "onERC721Received", arg0, from, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address from, uint256 , bytes ) pure returns(bytes4)
func (_CustosialVault *CustosialVaultSession) OnERC721Received(arg0 common.Address, from common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _CustosialVault.Contract.OnERC721Received(&_CustosialVault.CallOpts, arg0, from, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address from, uint256 , bytes ) pure returns(bytes4)
func (_CustosialVault *CustosialVaultCallerSession) OnERC721Received(arg0 common.Address, from common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _CustosialVault.Contract.OnERC721Received(&_CustosialVault.CallOpts, arg0, from, arg2, arg3)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CustosialVault *CustosialVaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CustosialVault.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CustosialVault *CustosialVaultSession) Owner() (common.Address, error) {
	return _CustosialVault.Contract.Owner(&_CustosialVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CustosialVault *CustosialVaultCallerSession) Owner() (common.Address, error) {
	return _CustosialVault.Contract.Owner(&_CustosialVault.CallOpts)
}

// TxCost is a free data retrieval call binding the contract method 0xb2a7e29d.
//
// Solidity: function txCost() view returns(uint256)
func (_CustosialVault *CustosialVaultCaller) TxCost(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CustosialVault.contract.Call(opts, &out, "txCost")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TxCost is a free data retrieval call binding the contract method 0xb2a7e29d.
//
// Solidity: function txCost() view returns(uint256)
func (_CustosialVault *CustosialVaultSession) TxCost() (*big.Int, error) {
	return _CustosialVault.Contract.TxCost(&_CustosialVault.CallOpts)
}

// TxCost is a free data retrieval call binding the contract method 0xb2a7e29d.
//
// Solidity: function txCost() view returns(uint256)
func (_CustosialVault *CustosialVaultCallerSession) TxCost() (*big.Int, error) {
	return _CustosialVault.Contract.TxCost(&_CustosialVault.CallOpts)
}

// EmergencyDelete is a paid mutator transaction binding the contract method 0x00d90cd9.
//
// Solidity: function emergencyDelete(uint256 tokenId) returns()
func (_CustosialVault *CustosialVaultTransactor) EmergencyDelete(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _CustosialVault.contract.Transact(opts, "emergencyDelete", tokenId)
}

// EmergencyDelete is a paid mutator transaction binding the contract method 0x00d90cd9.
//
// Solidity: function emergencyDelete(uint256 tokenId) returns()
func (_CustosialVault *CustosialVaultSession) EmergencyDelete(tokenId *big.Int) (*types.Transaction, error) {
	return _CustosialVault.Contract.EmergencyDelete(&_CustosialVault.TransactOpts, tokenId)
}

// EmergencyDelete is a paid mutator transaction binding the contract method 0x00d90cd9.
//
// Solidity: function emergencyDelete(uint256 tokenId) returns()
func (_CustosialVault *CustosialVaultTransactorSession) EmergencyDelete(tokenId *big.Int) (*types.Transaction, error) {
	return _CustosialVault.Contract.EmergencyDelete(&_CustosialVault.TransactOpts, tokenId)
}

// ReleaseNFT is a paid mutator transaction binding the contract method 0x44509d3e.
//
// Solidity: function releaseNFT(uint256 tokenId, address wallet) returns()
func (_CustosialVault *CustosialVaultTransactor) ReleaseNFT(opts *bind.TransactOpts, tokenId *big.Int, wallet common.Address) (*types.Transaction, error) {
	return _CustosialVault.contract.Transact(opts, "releaseNFT", tokenId, wallet)
}

// ReleaseNFT is a paid mutator transaction binding the contract method 0x44509d3e.
//
// Solidity: function releaseNFT(uint256 tokenId, address wallet) returns()
func (_CustosialVault *CustosialVaultSession) ReleaseNFT(tokenId *big.Int, wallet common.Address) (*types.Transaction, error) {
	return _CustosialVault.Contract.ReleaseNFT(&_CustosialVault.TransactOpts, tokenId, wallet)
}

// ReleaseNFT is a paid mutator transaction binding the contract method 0x44509d3e.
//
// Solidity: function releaseNFT(uint256 tokenId, address wallet) returns()
func (_CustosialVault *CustosialVaultTransactorSession) ReleaseNFT(tokenId *big.Int, wallet common.Address) (*types.Transaction, error) {
	return _CustosialVault.Contract.ReleaseNFT(&_CustosialVault.TransactOpts, tokenId, wallet)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CustosialVault *CustosialVaultTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustosialVault.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CustosialVault *CustosialVaultSession) RenounceOwnership() (*types.Transaction, error) {
	return _CustosialVault.Contract.RenounceOwnership(&_CustosialVault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CustosialVault *CustosialVaultTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CustosialVault.Contract.RenounceOwnership(&_CustosialVault.TransactOpts)
}

// RetainNFT is a paid mutator transaction binding the contract method 0xccf678d8.
//
// Solidity: function retainNFT(uint256 tokenId) payable returns()
func (_CustosialVault *CustosialVaultTransactor) RetainNFT(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _CustosialVault.contract.Transact(opts, "retainNFT", tokenId)
}

// RetainNFT is a paid mutator transaction binding the contract method 0xccf678d8.
//
// Solidity: function retainNFT(uint256 tokenId) payable returns()
func (_CustosialVault *CustosialVaultSession) RetainNFT(tokenId *big.Int) (*types.Transaction, error) {
	return _CustosialVault.Contract.RetainNFT(&_CustosialVault.TransactOpts, tokenId)
}

// RetainNFT is a paid mutator transaction binding the contract method 0xccf678d8.
//
// Solidity: function retainNFT(uint256 tokenId) payable returns()
func (_CustosialVault *CustosialVaultTransactorSession) RetainNFT(tokenId *big.Int) (*types.Transaction, error) {
	return _CustosialVault.Contract.RetainNFT(&_CustosialVault.TransactOpts, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CustosialVault *CustosialVaultTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CustosialVault.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CustosialVault *CustosialVaultSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CustosialVault.Contract.TransferOwnership(&_CustosialVault.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CustosialVault *CustosialVaultTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CustosialVault.Contract.TransferOwnership(&_CustosialVault.TransactOpts, newOwner)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x7192711f.
//
// Solidity: function updateOwner(uint256 tokenId, address newHolder) returns()
func (_CustosialVault *CustosialVaultTransactor) UpdateOwner(opts *bind.TransactOpts, tokenId *big.Int, newHolder common.Address) (*types.Transaction, error) {
	return _CustosialVault.contract.Transact(opts, "updateOwner", tokenId, newHolder)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x7192711f.
//
// Solidity: function updateOwner(uint256 tokenId, address newHolder) returns()
func (_CustosialVault *CustosialVaultSession) UpdateOwner(tokenId *big.Int, newHolder common.Address) (*types.Transaction, error) {
	return _CustosialVault.Contract.UpdateOwner(&_CustosialVault.TransactOpts, tokenId, newHolder)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x7192711f.
//
// Solidity: function updateOwner(uint256 tokenId, address newHolder) returns()
func (_CustosialVault *CustosialVaultTransactorSession) UpdateOwner(tokenId *big.Int, newHolder common.Address) (*types.Transaction, error) {
	return _CustosialVault.Contract.UpdateOwner(&_CustosialVault.TransactOpts, tokenId, newHolder)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_CustosialVault *CustosialVaultTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustosialVault.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_CustosialVault *CustosialVaultSession) Withdraw() (*types.Transaction, error) {
	return _CustosialVault.Contract.Withdraw(&_CustosialVault.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_CustosialVault *CustosialVaultTransactorSession) Withdraw() (*types.Transaction, error) {
	return _CustosialVault.Contract.Withdraw(&_CustosialVault.TransactOpts)
}

// CustosialVaultNFTCustodyIterator is returned from FilterNFTCustody and is used to iterate over the raw logs and unpacked data for NFTCustody events raised by the CustosialVault contract.
type CustosialVaultNFTCustodyIterator struct {
	Event *CustosialVaultNFTCustody // Event containing the contract specifics and raw log

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
func (it *CustosialVaultNFTCustodyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustosialVaultNFTCustody)
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
		it.Event = new(CustosialVaultNFTCustody)
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
func (it *CustosialVaultNFTCustodyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CustosialVaultNFTCustodyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CustosialVaultNFTCustody represents a NFTCustody event raised by the CustosialVault contract.
type CustosialVaultNFTCustody struct {
	TokenId *big.Int
	Holder  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNFTCustody is a free log retrieval operation binding the contract event 0x2eae7534e9f0d5d8a801dd79515d22a9c36f0b691692ca20f5bc04b94b055edc.
//
// Solidity: event NFTCustody(uint256 indexed tokenId, address holder)
func (_CustosialVault *CustosialVaultFilterer) FilterNFTCustody(opts *bind.FilterOpts, tokenId []*big.Int) (*CustosialVaultNFTCustodyIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CustosialVault.contract.FilterLogs(opts, "NFTCustody", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CustosialVaultNFTCustodyIterator{contract: _CustosialVault.contract, event: "NFTCustody", logs: logs, sub: sub}, nil
}

// WatchNFTCustody is a free log subscription operation binding the contract event 0x2eae7534e9f0d5d8a801dd79515d22a9c36f0b691692ca20f5bc04b94b055edc.
//
// Solidity: event NFTCustody(uint256 indexed tokenId, address holder)
func (_CustosialVault *CustosialVaultFilterer) WatchNFTCustody(opts *bind.WatchOpts, sink chan<- *CustosialVaultNFTCustody, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CustosialVault.contract.WatchLogs(opts, "NFTCustody", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CustosialVaultNFTCustody)
				if err := _CustosialVault.contract.UnpackLog(event, "NFTCustody", log); err != nil {
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

// ParseNFTCustody is a log parse operation binding the contract event 0x2eae7534e9f0d5d8a801dd79515d22a9c36f0b691692ca20f5bc04b94b055edc.
//
// Solidity: event NFTCustody(uint256 indexed tokenId, address holder)
func (_CustosialVault *CustosialVaultFilterer) ParseNFTCustody(log types.Log) (*CustosialVaultNFTCustody, error) {
	event := new(CustosialVaultNFTCustody)
	if err := _CustosialVault.contract.UnpackLog(event, "NFTCustody", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CustosialVaultOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CustosialVault contract.
type CustosialVaultOwnershipTransferredIterator struct {
	Event *CustosialVaultOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CustosialVaultOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustosialVaultOwnershipTransferred)
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
		it.Event = new(CustosialVaultOwnershipTransferred)
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
func (it *CustosialVaultOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CustosialVaultOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CustosialVaultOwnershipTransferred represents a OwnershipTransferred event raised by the CustosialVault contract.
type CustosialVaultOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CustosialVault *CustosialVaultFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CustosialVaultOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CustosialVault.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CustosialVaultOwnershipTransferredIterator{contract: _CustosialVault.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CustosialVault *CustosialVaultFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CustosialVaultOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CustosialVault.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CustosialVaultOwnershipTransferred)
				if err := _CustosialVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_CustosialVault *CustosialVaultFilterer) ParseOwnershipTransferred(log types.Log) (*CustosialVaultOwnershipTransferred, error) {
	event := new(CustosialVaultOwnershipTransferred)
	if err := _CustosialVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
