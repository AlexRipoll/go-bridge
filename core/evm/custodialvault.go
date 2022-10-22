package evm

import (
	"context"
	"github.com/AlexRipoll/go-bridge/contract"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Custodian interface {
	RetainToken(ctx context.Context, tokenId, destination *big.Int) (*Tx, error)
	ReleaseToken(ctx context.Context, wallet string, tokenId *big.Int) (*Tx, error)
	UpdateOwner(tokenId *big.Int, newHolder common.Address) (*Tx, error)
	Withdraw() (*Tx, error)
	EmergencyDelete(tokenId *big.Int) (*Tx, error)

}

type erc721CustodialVaultContract struct {
	contract   *contract.CustosialVault
	transactor transactor
	//Client
}

func NewErc721CustodialVaultContract(address string, transactor transactor, client Client) (Custodian, error) {
	contract, err := contract.NewCustosialVault(common.HexToAddress(address), client.conn)
	if err != nil {
		return nil, err
	}
	return erc721CustodialVaultContract{contract: contract, transactor: transactor}, nil
}

func (e erc721CustodialVaultContract) RetainToken(ctx context.Context, tokenId, destination *big.Int) (*Tx, error) {
	t, err := e.transactor.TransactOpts(ctx)
	if err != nil {
		return nil, err
	}

	gethTx, err := e.contract.RetainToken(t, tokenId, destination)
	if err != nil {
		return nil, err
	}
	tx := transactionToTx(gethTx)

	return &tx, nil
}

func (e erc721CustodialVaultContract) ReleaseToken(ctx context.Context, wallet string, tokenId *big.Int) (*Tx, error) {
	t, err := e.transactor.TransactOpts(ctx)
	if err != nil {
		return nil, err
	}

	walletAddress := common.HexToAddress(wallet)
	gethTx, err := e.contract.ReleaseToken(t, tokenId, walletAddress)
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

