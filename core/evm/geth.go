package evm

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type Client struct {
	conn       *ethclient.Client
	chainId    uint64
	contracts  Contracts
	Transactor ContractTransactor
}

type Contracts struct {
	custodianVault Custodian
	erc721Token    Erc721Token
}

func NewTransactor(conn *ethclient.Client, privateKey string) (ContractTransactor, error) {
	return &transactor{conn: conn}, nil
}

type transactor struct {
	conn *ethclient.Client
	privateKey string
}

type ContractTransactor interface {
	Address() (common.Address, error)
	Nonce(ctx context.Context) (uint64, error)
	EstimatedGasPrice(ctx context.Context) (*big.Int, error)
	ChainId(ctx context.Context) (*big.Int, error)
	CurrentBlock(ctx context.Context) (uint64, error)
	TransactOpts(ctx context.Context) (*bind.TransactOpts, error)
	CallOpts(ctx context.Context) (*bind.CallOpts, error)
}

func (t transactor) privateKeyECDSA() (*ecdsa.PrivateKey, error) {
	return crypto.HexToECDSA(t.privateKey)
}

func (t transactor) Address() (common.Address, error) {
	privateKey, err := t.privateKeyECDSA()
	if err != nil {
		return common.Address{}, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	return address, nil
}

func (t transactor) Nonce(ctx context.Context) (uint64, error) {
	address, err := t.Address()
	if err != nil {
		return 0, err
	}
	return t.conn.PendingNonceAt(ctx, address)
}

func (t transactor) EstimatedGasPrice(ctx context.Context) (*big.Int, error) {
	return t.conn.SuggestGasPrice(ctx)
}

func (t transactor) ChainId(ctx context.Context) (*big.Int, error) {
	return t.conn.ChainID(ctx)
}

func (t transactor) CurrentBlock(ctx context.Context) (uint64, error) {
	return t.conn.BlockNumber(ctx)
}

func (t transactor) TransactOpts(ctx context.Context) (*bind.TransactOpts, error) {
	privateKeyECDSA, err := t.privateKeyECDSA()
	if err != nil {
		return nil, err
	}

	chainId, err := t.ChainId(ctx)
	if err != nil {
		return nil, err
	}

	nonce, err := t.Nonce(ctx)
	if err != nil {
		return nil, err
	}

	gasPrice, err := t.EstimatedGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	txOpt, err := bind.NewKeyedTransactorWithChainID(privateKeyECDSA, chainId)
	if err != nil {
		return nil, err
	}
	txOpt.Nonce = big.NewInt(int64(nonce))
	txOpt.Value = big.NewInt(0)
	//txOpt.GasLimit = t.GasLimit
	txOpt.GasPrice = gasPrice

	return txOpt, nil
}

func (t transactor) CallOpts(ctx context.Context) (*bind.CallOpts, error) {
	caller, err := t.Address()
	if err != nil {
		return nil, err
	}

	callOpts := bind.CallOpts{
		From:        caller,
		Context:     ctx,
	}

	return &callOpts, nil
}
