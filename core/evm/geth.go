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
	finality   uint64
	contracts  Contracts
	Transactor ContractTransactor
}

type Contracts struct {
	custodianVault Custodian
	erc721Token    Erc721Token
}

func NewClient(
	conn *ethclient.Client,
	chainId, finality uint64,
	custodialVaultContract Custodian,
	erc721TokenContract Erc721Token,
	) Client {
	return Client{
		conn:       conn,
		chainId:    chainId,
		finality:   finality,
		contracts:  Contracts{
			custodianVault: custodialVaultContract,
			erc721Token:    erc721TokenContract,
		},
	}
}

func (c Client) ChainId() uint64 {
	return c.chainId
}

func (c Client) Finality() uint64 {
	return c.finality
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

type transactor struct {
	conn       *ethclient.Client
	privateKey string
}

func NewTransactor(conn *ethclient.Client, privateKey string) ContractTransactor {
	return &transactor{conn: conn}
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

	txOpts, err := bind.NewKeyedTransactorWithChainID(privateKeyECDSA, chainId)
	if err != nil {
		return nil, err
	}
	txOpts.Nonce = big.NewInt(int64(nonce))
	txOpts.Value = big.NewInt(0)
	//txOpts.GasLimit = t.GasLimit
	txOpts.GasPrice = gasPrice

	return txOpts, nil
}

func (t transactor) CallOpts(ctx context.Context) (*bind.CallOpts, error) {
	caller, err := t.Address()
	if err != nil {
		return nil, err
	}

	callOpts := bind.CallOpts{
		From:    caller,
		Context: ctx,
	}

	return &callOpts, nil
}
