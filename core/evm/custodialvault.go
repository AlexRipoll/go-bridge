package evm

import (
	"context"
	"fmt"
	"github.com/AlexRipoll/go-bridge/contract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type Custodian interface {
	RetainToken(ctx context.Context, tokenId, destination *big.Int) (*Tx, error)
	ReleaseToken(ctx context.Context, wallet string, tokenId *big.Int) (*Tx, error)
	UpdateOwner(ctx context.Context, tokenId *big.Int, newHolder string) (*Tx, error)
	Withdraw() (*Tx, error)
	EmergencyDelete(tokenId *big.Int) (*Tx, error)
	Address() string
	TokenHolder(ctx context.Context, tokenId *big.Int) (string, error)
}

type erc721CustodialVaultContract struct {
	address    string
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
	fmt.Println(">>>> ReleaseToken nonce: ", txOpts.Nonce)

	walletAddress := common.HexToAddress(wallet)
	gethTx, err := e.contract.ReleaseToken(txOpts, tokenId, walletAddress)
	if err != nil {
		return nil, err
	}
	tx := transactionToTx(gethTx)

	return &tx, nil
}

func (e erc721CustodialVaultContract) UpdateOwner(ctx context.Context, tokenId *big.Int, newHolder string) (*Tx, error) {
	txOpts, err := e.transactor.TransactOpts(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println(">>>> Update owner nonce: ", txOpts.Nonce)

	holderAddress := common.HexToAddress(newHolder)
	gethTx, err := e.contract.UpdateOwner(txOpts, tokenId, holderAddress)
	if err != nil {
		return nil, err
	}
	tx := transactionToTx(gethTx)

	return &tx, nil
}

func (e erc721CustodialVaultContract) TokenHolder(ctx context.Context, tokenId *big.Int) (string, error) {
	callOpts, err := e.transactor.CallOpts(ctx)
	if err != nil {
		return "", err
	}

	holder, err := e.contract.HoldCustody(callOpts, tokenId)
	if err != nil {
		return "", err
	}

	return holder.String(), nil
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
