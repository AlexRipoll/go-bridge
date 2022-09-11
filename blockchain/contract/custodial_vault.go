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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractERC721Enumerable\",\"name\":\"_nft\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"}],\"name\":\"NFTCustody\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"costCustom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"costNative\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"emergencyDelete\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"holdCustody\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"releaseNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"retainNFT\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newHolder\",\"type\":\"address\"}],\"name\":\"updateOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x6080604052670de0b6b3a76400006002556544364c5bb0006003553480156200002757600080fd5b50604051620015793803806200157983398181016040528101906200004d919062000209565b60016000819055506200007562000069620000bd60201b60201c565b620000c560201b60201c565b80600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506200023b565b600033905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620001bd8262000190565b9050919050565b6000620001d182620001b0565b9050919050565b620001e381620001c4565b8114620001ef57600080fd5b50565b6000815190506200020381620001d8565b92915050565b6000602082840312156200022257620002216200018b565b5b60006200023284828501620001f2565b91505092915050565b61132e806200024b6000396000f3fe6080604052600436106100a65760003560e01c80637192711f116100645780637192711f146101865780638da5cb5b146101af578063a8643a23146101da578063ccf678d814610205578063dd39bd2114610221578063f2fde38b1461025f576100a6565b8062d90cd9146100ab578063150b7a02146100d45780633ccfd60b1461011157806344509d3e1461011b57806353ef035d14610144578063715018a61461016f575b600080fd5b3480156100b757600080fd5b506100d260048036038101906100cd9190610c50565b610288565b005b3480156100e057600080fd5b506100fb60048036038101906100f69190610d40565b61032d565b6040516101089190610e03565b60405180910390f35b6101196103b0565b005b34801561012757600080fd5b50610142600480360381019061013d9190610e1e565b6103f8565b005b34801561015057600080fd5b5061015961052f565b6040516101669190610e6d565b60405180910390f35b34801561017b57600080fd5b50610184610535565b005b34801561019257600080fd5b506101ad60048036038101906101a89190610e1e565b610549565b005b3480156101bb57600080fd5b506101c4610674565b6040516101d19190610e97565b60405180910390f35b3480156101e657600080fd5b506101ef61069e565b6040516101fc9190610e6d565b60405180910390f35b61021f600480360381019061021a9190610c50565b6106a4565b005b34801561022d57600080fd5b5061024860048036038101906102439190610c50565b6109fd565b604051610256929190610eb2565b60405180910390f35b34801561026b57600080fd5b5061028660048036038101906102819190610edb565b610a41565b005b6002600054036102cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102c490610f65565b60405180910390fd5b60026000819055506102dd610ac4565b600460008281526020019081526020016000206000808201600090556001820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690555050600160008190555050565b60008073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff161461039d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161039490610fd1565b60405180910390fd5b63150b7a0260e01b905095945050505050565b6103b8610ac4565b3373ffffffffffffffffffffffffffffffffffffffff166108fc479081150290604051600060405180830381858888f193505050506103f657600080fd5b565b60026000540361043d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161043490610f65565b60405180910390fd5b600260008190555061044d610ac4565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3083856040518463ffffffff1660e01b81526004016104ac93929190610ff1565b600060405180830381600087803b1580156104c657600080fd5b505af11580156104da573d6000803e3d6000fd5b50505050600460008381526020019081526020016000206000808201600090556001820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055505060016000819055505050565b60025481565b61053d610ac4565b6105476000610b42565b565b60026000540361058e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161058590610f65565b60405180910390fd5b600260008190555061059e610ac4565b60405180604001604052808381526020018273ffffffffffffffffffffffffffffffffffffffff16815250600460008481526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550905050817f2eae7534e9f0d5d8a801dd79515d22a9c36f0b691692ca20f5bc04b94b055edc826040516106609190610e97565b60405180910390a260016000819055505050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60035481565b6002600054036106e9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e090610f65565b60405180910390fd5b60026000819055506003543414610735576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161072c9061109a565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff16600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636352211e836040518263ffffffff1660e01b81526004016107a79190610e6d565b602060405180830381865afa1580156107c4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107e891906110cf565b73ffffffffffffffffffffffffffffffffffffffff161461083e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108359061116e565b60405180910390fd5b6000600460008381526020019081526020016000206000015414610897576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161088e906111da565b60405180910390fd5b60405180604001604052808281526020013373ffffffffffffffffffffffffffffffffffffffff16815250600460008381526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550905050600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330846040518463ffffffff1660e01b815260040161098893929190610ff1565b600060405180830381600087803b1580156109a257600080fd5b505af11580156109b6573d6000803e3d6000fd5b50505050807f2eae7534e9f0d5d8a801dd79515d22a9c36f0b691692ca20f5bc04b94b055edc336040516109ea9190610e97565b60405180910390a2600160008190555050565b60046020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905082565b610a49610ac4565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610ab8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aaf9061126c565b60405180910390fd5b610ac181610b42565b50565b610acc610c08565b73ffffffffffffffffffffffffffffffffffffffff16610aea610674565b73ffffffffffffffffffffffffffffffffffffffff1614610b40576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b37906112d8565b60405180910390fd5b565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600033905090565b600080fd5b600080fd5b6000819050919050565b610c2d81610c1a565b8114610c3857600080fd5b50565b600081359050610c4a81610c24565b92915050565b600060208284031215610c6657610c65610c10565b5b6000610c7484828501610c3b565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610ca882610c7d565b9050919050565b610cb881610c9d565b8114610cc357600080fd5b50565b600081359050610cd581610caf565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f840112610d0057610cff610cdb565b5b8235905067ffffffffffffffff811115610d1d57610d1c610ce0565b5b602083019150836001820283011115610d3957610d38610ce5565b5b9250929050565b600080600080600060808688031215610d5c57610d5b610c10565b5b6000610d6a88828901610cc6565b9550506020610d7b88828901610cc6565b9450506040610d8c88828901610c3b565b935050606086013567ffffffffffffffff811115610dad57610dac610c15565b5b610db988828901610cea565b92509250509295509295909350565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b610dfd81610dc8565b82525050565b6000602082019050610e186000830184610df4565b92915050565b60008060408385031215610e3557610e34610c10565b5b6000610e4385828601610c3b565b9250506020610e5485828601610cc6565b9150509250929050565b610e6781610c1a565b82525050565b6000602082019050610e826000830184610e5e565b92915050565b610e9181610c9d565b82525050565b6000602082019050610eac6000830184610e88565b92915050565b6000604082019050610ec76000830185610e5e565b610ed46020830184610e88565b9392505050565b600060208284031215610ef157610ef0610c10565b5b6000610eff84828501610cc6565b91505092915050565b600082825260208201905092915050565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00600082015250565b6000610f4f601f83610f08565b9150610f5a82610f19565b602082019050919050565b60006020820190508181036000830152610f7e81610f42565b9050919050565b7f43616e6e6f742052656365697665204e465473204469726563746c7900000000600082015250565b6000610fbb601c83610f08565b9150610fc682610f85565b602082019050919050565b60006020820190508181036000830152610fea81610fae565b9050919050565b60006060820190506110066000830186610e88565b6110136020830185610e88565b6110206040830184610e5e565b949350505050565b7f4e6f7420656e6f7567682062616c616e636520746f20636f6d706c657465207460008201527f72616e73616374696f6e2e000000000000000000000000000000000000000000602082015250565b6000611084602b83610f08565b915061108f82611028565b604082019050919050565b600060208201905081810360008301526110b381611077565b9050919050565b6000815190506110c981610caf565b92915050565b6000602082840312156110e5576110e4610c10565b5b60006110f3848285016110ba565b91505092915050565b7f796f75206d75737420626520746865204e4654206f776e657220666f7220657860008201527f65637574696e67207468697320616374696f6e00000000000000000000000000602082015250565b6000611158603383610f08565b9150611163826110fc565b604082019050919050565b600060208201905081810360008301526111878161114b565b9050919050565b7f4e465420616c72656164792073746f7265640000000000000000000000000000600082015250565b60006111c4601283610f08565b91506111cf8261118e565b602082019050919050565b600060208201905081810360008301526111f3816111b7565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000611256602683610f08565b9150611261826111fa565b604082019050919050565b6000602082019050818103600083015261128581611249565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b60006112c2602083610f08565b91506112cd8261128c565b602082019050919050565b600060208201905081810360008301526112f1816112b5565b905091905056fea26469706673582212203b0922bad691881633bc3cb68b5d2a668303c776edbcb3eaa961b985f010444e64736f6c63430008100033",
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

// CostCustom is a free data retrieval call binding the contract method 0x53ef035d.
//
// Solidity: function costCustom() view returns(uint256)
func (_CustosialVault *CustosialVaultCaller) CostCustom(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CustosialVault.contract.Call(opts, &out, "costCustom")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CostCustom is a free data retrieval call binding the contract method 0x53ef035d.
//
// Solidity: function costCustom() view returns(uint256)
func (_CustosialVault *CustosialVaultSession) CostCustom() (*big.Int, error) {
	return _CustosialVault.Contract.CostCustom(&_CustosialVault.CallOpts)
}

// CostCustom is a free data retrieval call binding the contract method 0x53ef035d.
//
// Solidity: function costCustom() view returns(uint256)
func (_CustosialVault *CustosialVaultCallerSession) CostCustom() (*big.Int, error) {
	return _CustosialVault.Contract.CostCustom(&_CustosialVault.CallOpts)
}

// CostNative is a free data retrieval call binding the contract method 0xa8643a23.
//
// Solidity: function costNative() view returns(uint256)
func (_CustosialVault *CustosialVaultCaller) CostNative(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CustosialVault.contract.Call(opts, &out, "costNative")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CostNative is a free data retrieval call binding the contract method 0xa8643a23.
//
// Solidity: function costNative() view returns(uint256)
func (_CustosialVault *CustosialVaultSession) CostNative() (*big.Int, error) {
	return _CustosialVault.Contract.CostNative(&_CustosialVault.CallOpts)
}

// CostNative is a free data retrieval call binding the contract method 0xa8643a23.
//
// Solidity: function costNative() view returns(uint256)
func (_CustosialVault *CustosialVaultCallerSession) CostNative() (*big.Int, error) {
	return _CustosialVault.Contract.CostNative(&_CustosialVault.CallOpts)
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
