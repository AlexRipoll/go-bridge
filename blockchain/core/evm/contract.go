package evm

import (
	"context"
	"github.com/AlexRipoll/go-bridge/blockchain/contract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type Blockchain interface {
	RetainNFT(ctx context.Context, tokenId *big.Int) (*Tx, error)
	ReleaseNFT(tokenId *big.Int, wallet common.Address) (*Tx, error)
	UpdateOwner(tokenId *big.Int, newHolder common.Address) (*Tx, error)
	Withdraw() (*Tx, error)
	EmergencyDelete(tokenId *big.Int) (*Tx, error)

	HoldCustody(idx *big.Int) (*Custody, error)
}

type Tx struct {
	Hash string
	Size float64
	ChainID   *big.Int
	Data      []byte
	Gas       uint64
	GasPrice  *big.Int
	GasTipCap *big.Int
	GasFeeCap *big.Int
	Value     *big.Int
	Nonce     uint64
	To        string
}

type Custody struct {
	TokenId *big.Int
	Holder  string
}

func (geth geth) RetainNFT(ctx context.Context, tokenId *big.Int) (*Tx, error) {
	t, err := geth.prepareTransactor(ctx)
	if err != nil {
		return nil, err
	}

	gethTx, err := geth.custodialVault.RetainNFT(t, tokenId)
	if err != nil {
		return nil, err
	}
	tx := transactionToTx(gethTx)

	return &tx, nil
}

func (geth geth) ReleaseNFT(tokenId *big.Int, wallet common.Address) (*Tx, error) {
	panic("implement me")
}

func (geth geth) UpdateOwner(tokenId *big.Int, newHolder common.Address) (*Tx, error) {
	panic("implement me")
}

func (geth geth) Withdraw() (*Tx, error) {
	panic("implement me")
}

func (geth geth) EmergencyDelete(tokenId *big.Int) (*Tx, error) {
	panic("implement me")
}

func (geth geth) HoldCustody(idx *big.Int) (*Custody, error) {
	panic("implement me")
}

type DeployRx struct {
	Address common.Address
	Tx      *types.Transaction
}

func (geth geth) Deploy(ctx context.Context) (*DeployRx, error) {
	transactor, err := geth.prepareTransactor(ctx)
	if err != nil {
		return nil, err
	}

	// TODO check nft construct argument
	address, tx, _, err := contract.DeployCustosialVault(transactor, geth.client, common.Address{})
	if err != nil {
		return nil, err
	}

	return &DeployRx{
		Address: address,
		Tx:      tx,
	}, nil
}

func transactionToTx(transaction *types.Transaction) Tx {
	return Tx{
		Hash: transaction.Hash().Hex(),
		Size: float64(transaction.Size()),
		//From:    transaction.,
		ChainID:   transaction.ChainId(),
		Data:      transaction.Data(),
		Gas:       transaction.Gas(),
		GasPrice:  transaction.GasPrice(),
		GasTipCap: transaction.GasTipCap(),
		GasFeeCap: transaction.GasFeeCap(),
		Value:     transaction.Value(),
		Nonce:     transaction.Nonce(),
		To:        transaction.To().Hex(),
	}
}
