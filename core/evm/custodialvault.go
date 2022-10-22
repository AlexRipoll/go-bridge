package evm

import (
	"context"
	"github.com/AlexRipoll/go-bridge/contract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type Custodian interface {
	RetainToken(ctx context.Context, tokenId, destination *big.Int) (*Tx, error)
	ReleaseToken(ctx context.Context, wallet string, tokenId *big.Int) (*Tx, error)
	UpdateOwner(tokenId *big.Int, newHolder common.Address) (*Tx, error)
	Withdraw() (*Tx, error)
	EmergencyDelete(tokenId *big.Int) (*Tx, error)
	Address() string
}

type erc721CustodialVaultContract struct {
	address	string
	contract   *contract.CustosialVault
	transactor ContractTransactor
}

func NewErc721CustodialVaultContract(address string, transactor ContractTransactor, conn *ethclient.Client) (Custodian, error) {
	c, err := contract.NewCustosialVault(common.HexToAddress(address), conn)
	if err != nil {
		return nil, err
	}
	return erc721CustodialVaultContract{address: address, contract: c, transactor: transactor}, nil
}

func (e erc721CustodialVaultContract) RetainToken(ctx context.Context, tokenId, destination *big.Int) (*Tx, error) {
	txOpts, err := e.transactor.TransactOpts(ctx)
	if err != nil {
		return nil, err
	}

	gethTx, err := e.contract.RetainToken(txOpts, tokenId, destination)
	if err != nil {
		return nil, err
	}
	tx := transactionToTx(gethTx)

	return &tx, nil
}

func (e erc721CustodialVaultContract) ReleaseToken(ctx context.Context, wallet string, tokenId *big.Int) (*Tx, error) {
	txOpts, err := e.transactor.TransactOpts(ctx)
	if err != nil {
		return nil, err
	}

	walletAddress := common.HexToAddress(wallet)
	gethTx, err := e.contract.ReleaseToken(txOpts, tokenId, walletAddress)
	if err != nil {
		return nil, err
	}
	tx := transactionToTx(gethTx)

	return &tx, nil
}

func (e erc721CustodialVaultContract) UpdateOwner(tokenId *big.Int, newHolder common.Address) (*Tx, error) {
	panic("implement me")
}

func (e erc721CustodialVaultContract) Withdraw() (*Tx, error) {
	panic("implement me")
}

func (e erc721CustodialVaultContract) EmergencyDelete(tokenId *big.Int) (*Tx, error) {
	panic("implement me")
}

// Address returns the contact address in hexadecimal format
func (e erc721CustodialVaultContract) Address() string {
	return common.HexToAddress(e.address).String()
}