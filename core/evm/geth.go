package evm

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/AlexRipoll/go-bridge/config"
	contract2 "github.com/AlexRipoll/go-bridge/contract"
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

func NewClient(conn *ethclient.Client, custodialVaultAddress, erc721TokenAddress string) (*Client, error) {

	return &Client{
		conn: conn,
	}, nil
}

type geth struct {
	conn *ethclient.Client
}

type ContractTransactor interface {
	Address() (common.Address, error)
	Nonce(ctx context.Context) (uint64, error)
	EstimatedGasPrice(ctx context.Context) (*big.Int, error)
	ChainId(ctx context.Context) (*big.Int, error)
}

type Reader interface {
	CurrentBlock(ctx context.Context) (uint64, error)
}

type custodian struct {
	Client
	Contract *contract2.CustosialVault
}

func NewCustodian(conn *ethclient.Client, contractAddr string, config config.Config, network string) (Custodian, error) {
	geth, err := NewClient(conn, config, network)
	if err != nil {
		return nil, err
	}

	var vault *contract2.CustosialVault
	if contractAddr != "" {
		vault, err = contract2.NewCustosialVault(common.HexToAddress(contractAddr), conn)
		if err != nil {
			return nil, err
		}
	}

	return custodian{
		Client:   *geth,
		Contract: vault,
	}, nil
}

func NewBridger(conn *ethclient.Client, contractAddr string, config config.Config, network string) (Erc721Token, error) {
	geth, err := NewClient(conn, config, network)
	if err != nil {
		return nil, err
	}

	var nft *contract2.NFT
	if contractAddr != "" {
		nft, err = contract2.NewNFT(common.HexToAddress(contractAddr), conn)
		if err != nil {
			return nil, err
		}
	}
	return bridger{
		Client:   *geth,
		Contract: nft,
	}, nil
}

func (g geth) privateKey() (*ecdsa.PrivateKey, error) {
	return crypto.HexToECDSA(g.PrivateKey)
}

func (g geth) Address() (common.Address, error) {
	privateKey, err := g.privateKey()
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

func (g geth) Nonce(ctx context.Context) (uint64, error) {
	address, err := g.Address()
	if err != nil {
		return 0, err
	}
	return g.conn.PendingNonceAt(ctx, address)
}

func (g geth) EstimatedGasPrice(ctx context.Context) (*big.Int, error) {
	return g.conn.SuggestGasPrice(ctx)
}

func (g geth) ChainId(ctx context.Context) (*big.Int, error) {
	return g.conn.ChainID(ctx)
}

func (g geth) CurrentBlock(ctx context.Context) (uint64, error) {
	return g.conn.BlockNumber(ctx)
}

func (g geth) prepareTransactor(ctx context.Context) (*bind.TransactOpts, error) {
	privateKeyECDSA, err := g.privateKey()
	if err != nil {
		return nil, err
	}

	chainId, err := g.ChainId(ctx)
	if err != nil {
		return nil, err
	}

	nonce, err := g.Nonce(ctx)
	if err != nil {
		return nil, err
	}

	gasPrice, err := g.EstimatedGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	transactor, err := bind.NewKeyedTransactorWithChainID(privateKeyECDSA, chainId)
	if err != nil {
		return nil, err
	}
	transactor.Nonce = big.NewInt(int64(nonce))
	transactor.Value = big.NewInt(0)
	//transactor.GasLimit = g.GasLimit
	transactor.GasPrice = gasPrice

	return transactor, nil
}

func (g geth) prepareCallOpts(ctx context.Context) (*bind.CallOpts, error) {
	caller, err := g.Address()
	if err != nil {
		return nil, err
	}

	callOpts := bind.CallOpts{
		From:        caller,
		Context:     ctx,
	}

	return &callOpts, nil
}
