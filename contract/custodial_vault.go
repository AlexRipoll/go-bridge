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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractERC721Enumerable\",\"name\":\"_erc721token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destinationChainId\",\"type\":\"uint256\"}],\"name\":\"TokenCustody\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"emergencyDelete\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"holdCustody\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"releaseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destination\",\"type\":\"uint256\"}],\"name\":\"retainToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"txCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newHolder\",\"type\":\"address\"}],\"name\":\"updateOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60806040526103e86002553480156200001757600080fd5b506040516200156c3803806200156c83398181016040528101906200003d9190620001f9565b60016000819055506200006562000059620000ad60201b60201c565b620000b560201b60201c565b80600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506200022b565b600033905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620001ad8262000180565b9050919050565b6000620001c182620001a0565b9050919050565b620001d381620001b4565b8114620001df57600080fd5b50565b600081519050620001f381620001c8565b92915050565b6000602082840312156200021257620002116200017b565b5b60006200022284828501620001e2565b91505092915050565b611331806200023b6000396000f3fe60806040526004361061009b5760003560e01c80637192711f116100645780637192711f146101505780638da5cb5b14610179578063982ebbb9146101a4578063b2a7e29d146101c0578063dd39bd21146101eb578063f2fde38b146102285761009b565b8062d90cd9146100a0578063050ed154146100c9578063150b7a02146100f25780633ccfd60b1461012f578063715018a614610139575b600080fd5b3480156100ac57600080fd5b506100c760048036038101906100c29190610c13565b610251565b005b3480156100d557600080fd5b506100f060048036038101906100eb9190610c9e565b6102e7565b005b3480156100fe57600080fd5b5061011960048036038101906101149190610d43565b61040f565b6040516101269190610e06565b60405180910390f35b610137610492565b005b34801561014557600080fd5b5061014e6104da565b005b34801561015c57600080fd5b5061017760048036038101906101729190610c9e565b6104ee565b005b34801561018557600080fd5b5061018e6105a1565b60405161019b9190610e30565b60405180910390f35b6101be60048036038101906101b99190610e4b565b6105cb565b005b3480156101cc57600080fd5b506101d56109cb565b6040516101e29190610e9a565b60405180910390f35b3480156101f757600080fd5b50610212600480360381019061020d9190610c13565b6109d1565b60405161021f9190610e30565b60405180910390f35b34801561023457600080fd5b5061024f600480360381019061024a9190610eb5565b610a04565b005b600260005403610296576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161028d90610f3f565b60405180910390fd5b60026000819055506102a6610a87565b6003600082815260200190815260200160002060006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600160008190555050565b60026000540361032c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161032390610f3f565b60405180910390fd5b600260008190555061033c610a87565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3083856040518463ffffffff1660e01b815260040161039b93929190610f5f565b600060405180830381600087803b1580156103b557600080fd5b505af11580156103c9573d6000803e3d6000fd5b505050506003600083815260200190815260200160002060006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905560016000819055505050565b60008073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff161461047f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161047690610fe2565b60405180910390fd5b63150b7a0260e01b905095945050505050565b61049a610a87565b3373ffffffffffffffffffffffffffffffffffffffff166108fc479081150290604051600060405180830381858888f193505050506104d857600080fd5b565b6104e2610a87565b6104ec6000610b05565b565b600260005403610533576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161052a90610f3f565b60405180910390fd5b6002600081905550610543610a87565b806003600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060016000819055505050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600260005403610610576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161060790610f3f565b60405180910390fd5b6002600081905550600254341461065c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161065390611074565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff16600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636352211e846040518263ffffffff1660e01b81526004016106ce9190610e9a565b602060405180830381865afa1580156106eb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061070f91906110a9565b73ffffffffffffffffffffffffffffffffffffffff1614610765576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161075c90611148565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166003600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610807576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107fe906111b4565b60405180910390fd5b336003600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636352211e856040518263ffffffff1660e01b81526004016108f29190610e9a565b602060405180830381865afa15801561090f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061093391906110a9565b30856040518463ffffffff1660e01b815260040161095393929190610f5f565b600060405180830381600087803b15801561096d57600080fd5b505af1158015610981573d6000803e3d6000fd5b50505050817f9bea55fe89520950ade3cb8f5b298d336f4ee0eca04a544b1b0760c40b4423b533836040516109b79291906111d4565b60405180910390a260016000819055505050565b60025481565b60036020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b610a0c610a87565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610a7b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a729061126f565b60405180910390fd5b610a8481610b05565b50565b610a8f610bcb565b73ffffffffffffffffffffffffffffffffffffffff16610aad6105a1565b73ffffffffffffffffffffffffffffffffffffffff1614610b03576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610afa906112db565b60405180910390fd5b565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600033905090565b600080fd5b600080fd5b6000819050919050565b610bf081610bdd565b8114610bfb57600080fd5b50565b600081359050610c0d81610be7565b92915050565b600060208284031215610c2957610c28610bd3565b5b6000610c3784828501610bfe565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610c6b82610c40565b9050919050565b610c7b81610c60565b8114610c8657600080fd5b50565b600081359050610c9881610c72565b92915050565b60008060408385031215610cb557610cb4610bd3565b5b6000610cc385828601610bfe565b9250506020610cd485828601610c89565b9150509250929050565b600080fd5b600080fd5b600080fd5b60008083601f840112610d0357610d02610cde565b5b8235905067ffffffffffffffff811115610d2057610d1f610ce3565b5b602083019150836001820283011115610d3c57610d3b610ce8565b5b9250929050565b600080600080600060808688031215610d5f57610d5e610bd3565b5b6000610d6d88828901610c89565b9550506020610d7e88828901610c89565b9450506040610d8f88828901610bfe565b935050606086013567ffffffffffffffff811115610db057610daf610bd8565b5b610dbc88828901610ced565b92509250509295509295909350565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b610e0081610dcb565b82525050565b6000602082019050610e1b6000830184610df7565b92915050565b610e2a81610c60565b82525050565b6000602082019050610e456000830184610e21565b92915050565b60008060408385031215610e6257610e61610bd3565b5b6000610e7085828601610bfe565b9250506020610e8185828601610bfe565b9150509250929050565b610e9481610bdd565b82525050565b6000602082019050610eaf6000830184610e8b565b92915050565b600060208284031215610ecb57610eca610bd3565b5b6000610ed984828501610c89565b91505092915050565b600082825260208201905092915050565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00600082015250565b6000610f29601f83610ee2565b9150610f3482610ef3565b602082019050919050565b60006020820190508181036000830152610f5881610f1c565b9050919050565b6000606082019050610f746000830186610e21565b610f816020830185610e21565b610f8e6040830184610e8b565b949350505050565b7f43616e6e6f74205265636569766520546f6b656e73204469726563746c790000600082015250565b6000610fcc601e83610ee2565b9150610fd782610f96565b602082019050919050565b60006020820190508181036000830152610ffb81610fbf565b9050919050565b7f4e6f7420656e6f7567682062616c616e636520746f20636f6d706c657465207460008201527f72616e73616374696f6e2e000000000000000000000000000000000000000000602082015250565b600061105e602b83610ee2565b915061106982611002565b604082019050919050565b6000602082019050818103600083015261108d81611051565b9050919050565b6000815190506110a381610c72565b92915050565b6000602082840312156110bf576110be610bd3565b5b60006110cd84828501611094565b91505092915050565b7f796f75206d7573742062652074686520546f6b656e206f776e657220666f722060008201527f657865637574696e67207468697320616374696f6e0000000000000000000000602082015250565b6000611132603583610ee2565b915061113d826110d6565b604082019050919050565b6000602082019050818103600083015261116181611125565b9050919050565b7f546f6b656e20616c72656164792073746f726564000000000000000000000000600082015250565b600061119e601483610ee2565b91506111a982611168565b602082019050919050565b600060208201905081810360008301526111cd81611191565b9050919050565b60006040820190506111e96000830185610e21565b6111f66020830184610e8b565b9392505050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000611259602683610ee2565b9150611264826111fd565b604082019050919050565b600060208201905081810360008301526112888161124c565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b60006112c5602083610ee2565b91506112d08261128f565b602082019050919050565b600060208201905081810360008301526112f4816112b8565b905091905056fea2646970667358221220a0c7b668129381d7f5c26f150c9521e9778b74307a0d0151683196365ece1fce64736f6c63430008110033",
}

// CustosialVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use CustosialVaultMetaData.ABI instead.
var CustosialVaultABI = CustosialVaultMetaData.ABI

// CustosialVaultBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CustosialVaultMetaData.Bin instead.
var CustosialVaultBin = CustosialVaultMetaData.Bin

// DeployCustosialVault deploys a new Ethereum contract, binding an instance of CustosialVault to it.
func DeployCustosialVault(auth *bind.TransactOpts, backend bind.ContractBackend, _erc721token common.Address) (common.Address, *types.Transaction, *CustosialVault, error) {
	parsed, err := CustosialVaultMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CustosialVaultBin), backend, _erc721token)
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
// Solidity: function holdCustody(uint256 ) view returns(address)
func (_CustosialVault *CustosialVaultCaller) HoldCustody(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CustosialVault.contract.Call(opts, &out, "holdCustody", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HoldCustody is a free data retrieval call binding the contract method 0xdd39bd21.
//
// Solidity: function holdCustody(uint256 ) view returns(address)
func (_CustosialVault *CustosialVaultSession) HoldCustody(arg0 *big.Int) (common.Address, error) {
	return _CustosialVault.Contract.HoldCustody(&_CustosialVault.CallOpts, arg0)
}

// HoldCustody is a free data retrieval call binding the contract method 0xdd39bd21.
//
// Solidity: function holdCustody(uint256 ) view returns(address)
func (_CustosialVault *CustosialVaultCallerSession) HoldCustody(arg0 *big.Int) (common.Address, error) {
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

// ReleaseToken is a paid mutator transaction binding the contract method 0x050ed154.
//
// Solidity: function releaseToken(uint256 tokenId, address wallet) returns()
func (_CustosialVault *CustosialVaultTransactor) ReleaseToken(opts *bind.TransactOpts, tokenId *big.Int, wallet common.Address) (*types.Transaction, error) {
	return _CustosialVault.contract.Transact(opts, "releaseToken", tokenId, wallet)
}

// ReleaseToken is a paid mutator transaction binding the contract method 0x050ed154.
//
// Solidity: function releaseToken(uint256 tokenId, address wallet) returns()
func (_CustosialVault *CustosialVaultSession) ReleaseToken(tokenId *big.Int, wallet common.Address) (*types.Transaction, error) {
	return _CustosialVault.Contract.ReleaseToken(&_CustosialVault.TransactOpts, tokenId, wallet)
}

// ReleaseToken is a paid mutator transaction binding the contract method 0x050ed154.
//
// Solidity: function releaseToken(uint256 tokenId, address wallet) returns()
func (_CustosialVault *CustosialVaultTransactorSession) ReleaseToken(tokenId *big.Int, wallet common.Address) (*types.Transaction, error) {
	return _CustosialVault.Contract.ReleaseToken(&_CustosialVault.TransactOpts, tokenId, wallet)
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

// RetainToken is a paid mutator transaction binding the contract method 0x982ebbb9.
//
// Solidity: function retainToken(uint256 tokenId, uint256 destination) payable returns()
func (_CustosialVault *CustosialVaultTransactor) RetainToken(opts *bind.TransactOpts, tokenId *big.Int, destination *big.Int) (*types.Transaction, error) {
	return _CustosialVault.contract.Transact(opts, "retainToken", tokenId, destination)
}

// RetainToken is a paid mutator transaction binding the contract method 0x982ebbb9.
//
// Solidity: function retainToken(uint256 tokenId, uint256 destination) payable returns()
func (_CustosialVault *CustosialVaultSession) RetainToken(tokenId *big.Int, destination *big.Int) (*types.Transaction, error) {
	return _CustosialVault.Contract.RetainToken(&_CustosialVault.TransactOpts, tokenId, destination)
}

// RetainToken is a paid mutator transaction binding the contract method 0x982ebbb9.
//
// Solidity: function retainToken(uint256 tokenId, uint256 destination) payable returns()
func (_CustosialVault *CustosialVaultTransactorSession) RetainToken(tokenId *big.Int, destination *big.Int) (*types.Transaction, error) {
	return _CustosialVault.Contract.RetainToken(&_CustosialVault.TransactOpts, tokenId, destination)
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

// CustosialVaultTokenCustodyIterator is returned from FilterTokenCustody and is used to iterate over the raw logs and unpacked data for TokenCustody events raised by the CustosialVault contract.
type CustosialVaultTokenCustodyIterator struct {
	Event *CustosialVaultTokenCustody // Event containing the contract specifics and raw log

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
func (it *CustosialVaultTokenCustodyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustosialVaultTokenCustody)
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
		it.Event = new(CustosialVaultTokenCustody)
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
func (it *CustosialVaultTokenCustodyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CustosialVaultTokenCustodyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CustosialVaultTokenCustody represents a TokenCustody event raised by the CustosialVault contract.
type CustosialVaultTokenCustody struct {
	TokenId            *big.Int
	Holder             common.Address
	DestinationChainId *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterTokenCustody is a free log retrieval operation binding the contract event 0x9bea55fe89520950ade3cb8f5b298d336f4ee0eca04a544b1b0760c40b4423b5.
//
// Solidity: event TokenCustody(uint256 indexed tokenId, address holder, uint256 destinationChainId)
func (_CustosialVault *CustosialVaultFilterer) FilterTokenCustody(opts *bind.FilterOpts, tokenId []*big.Int) (*CustosialVaultTokenCustodyIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CustosialVault.contract.FilterLogs(opts, "TokenCustody", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CustosialVaultTokenCustodyIterator{contract: _CustosialVault.contract, event: "TokenCustody", logs: logs, sub: sub}, nil
}

// WatchTokenCustody is a free log subscription operation binding the contract event 0x9bea55fe89520950ade3cb8f5b298d336f4ee0eca04a544b1b0760c40b4423b5.
//
// Solidity: event TokenCustody(uint256 indexed tokenId, address holder, uint256 destinationChainId)
func (_CustosialVault *CustosialVaultFilterer) WatchTokenCustody(opts *bind.WatchOpts, sink chan<- *CustosialVaultTokenCustody, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CustosialVault.contract.WatchLogs(opts, "TokenCustody", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CustosialVaultTokenCustody)
				if err := _CustosialVault.contract.UnpackLog(event, "TokenCustody", log); err != nil {
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

// ParseTokenCustody is a log parse operation binding the contract event 0x9bea55fe89520950ade3cb8f5b298d336f4ee0eca04a544b1b0760c40b4423b5.
//
// Solidity: event TokenCustody(uint256 indexed tokenId, address holder, uint256 destinationChainId)
func (_CustosialVault *CustosialVaultFilterer) ParseTokenCustody(log types.Log) (*CustosialVaultTokenCustody, error) {
	event := new(CustosialVaultTokenCustody)
	if err := _CustosialVault.contract.UnpackLog(event, "TokenCustody", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
