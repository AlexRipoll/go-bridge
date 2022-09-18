package evm

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/AlexRipoll/go-bridge/blockchain/config"
	"github.com/AlexRipoll/go-bridge/blockchain/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type Client struct {
	config.Config
	conn    *ethclient.Client
	network string
}

func NewClient(conn *ethclient.Client, config config.Config, network string) (*Client, error) {
	return &Client{
		Config:  config,
		conn:    conn,
		network: network,
	}, nil
}

type ContractTransactor interface {
	Address() (common.Address, error)
	Nonce(ctx context.Context) (uint64, error)
	EstimatedGasPrice(ctx context.Context) (*big.Int, error)
	ChainId(ctx context.Context) (*big.Int, error)
}

type Deployer interface {
	Deploy(ctx context.Context) ([]DeployRx, error)
}

type custodian struct {
	Client
	Contract *contract.CustosialVault
}

func NewCustodian(conn *ethclient.Client, contractAddr string, config config.Config, network string) (Custodian, error) {
	geth, err := NewClient(conn, config, network)
	if err != nil {
		return nil, err
	}

	var vault *contract.CustosialVault
	if contractAddr != "" {
		vault, err = contract.NewCustosialVault(common.HexToAddress(contractAddr), conn)
		if err != nil {
			return nil, err
		}
	}

	return custodian{
		Client:   *geth,
		Contract: vault,
	}, nil
}

type bridger struct {
	Client
	Contract *contract.NFT
}

func NewBridger(conn *ethclient.Client, contractAddr string, config config.Config, network string) (Bridger, error) {
	geth, err := NewClient(conn, config, network)
	if err != nil {
		return nil, err
	}

	var nft *contract.NFT
	if contractAddr != "" {
		nft, err = contract.NewNFT(common.HexToAddress(contractAddr), conn)
		if err != nil {
			return nil, err
		}
	}
	return bridger{
		Client:   *geth,
		Contract: nft,
	}, nil
}

func (c Client) privateKey() (*ecdsa.PrivateKey, error) {
	return crypto.HexToECDSA(c.PrivateKey)
}

func (c Client) Address() (common.Address, error) {
	privateKey, err := c.privateKey()
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

func (c Client) Nonce(ctx context.Context) (uint64, error) {
	address, err := c.Address()
	if err != nil {
		return 0, err
	}
	return c.conn.PendingNonceAt(ctx, address)
}

func (c Client) EstimatedGasPrice(ctx context.Context) (*big.Int, error) {
	return c.conn.SuggestGasPrice(ctx)
}

func (c Client) ChainId(ctx context.Context) (*big.Int, error) {
	return c.conn.ChainID(ctx)
}

func (c Client) prepareTransactor(ctx context.Context) (*bind.TransactOpts, error) {
	privateKeyECDSA, err := c.privateKey()
	if err != nil {
		return nil, err
	}

	chainId, err := c.ChainId(ctx)
	if err != nil {
		return nil, err
	}

	nonce, err := c.Nonce(ctx)
	if err != nil {
		return nil, err
	}

	gasPrice, err := c.EstimatedGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	transactor, err := bind.NewKeyedTransactorWithChainID(privateKeyECDSA, chainId)
	if err != nil {
		return nil, err
	}
	transactor.Nonce = big.NewInt(int64(nonce))
	transactor.Value = big.NewInt(0)
	//transactor.GasLimit = c.GasLimit
	transactor.GasPrice = gasPrice

	return transactor, nil
}
