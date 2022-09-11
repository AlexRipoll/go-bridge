package evm

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/AlexRipoll/go-bridge/blockchain/contract"
	"github.com/AlexRipoll/go-bridge/blockchain/sys/evm"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type geth struct {
	evm.Config
	client         *ethclient.Client
	custodialVault *contract.CustosialVault
}

func NewGeth(c *ethclient.Client, config evm.Config) (Blockchain, error) {

	// TODO add contract address
	vault, err := contract.NewCustosialVault(common.Address{}, c)
	if err != nil {
		return nil, err
	}

	return geth{
		Config: config,
		client:         c,
		custodialVault: vault,
	}, nil
}

func (geth geth) privateKey() (*ecdsa.PrivateKey, error) {
	return crypto.HexToECDSA(geth.Pvk)
}

func (geth geth) Address() (common.Address, error) {
	privateKey, err := geth.privateKey()
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

func (geth geth) Nonce(ctx context.Context) (uint64, error) {
	address, err := geth.Address()
	if err != nil {
		return 0, err
	}
	return geth.client.PendingNonceAt(ctx, address)
}

func (geth geth) EstimatedGasPrice(ctx context.Context) (*big.Int, error) {
	return geth.client.SuggestGasPrice(ctx)
}

func (geth geth) ChainId(ctx context.Context) (*big.Int, error) {
	return geth.client.ChainID(ctx)
}

func (geth geth) prepareTransactor(ctx context.Context) (*bind.TransactOpts, error) {
	privateKeyECDSA, err := geth.privateKey()
	if err != nil {
		return nil, err
	}

	chainId, err := geth.ChainId(ctx)
	if err != nil {
		return nil, err
	}

	nonce, err := geth.Nonce(ctx)
	if err != nil {
		return nil, err
	}

	gasPrice, err := geth.EstimatedGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	transactor, err := bind.NewKeyedTransactorWithChainID(privateKeyECDSA, chainId)
	if err != nil {
		return nil, err
	}
	transactor.Nonce = big.NewInt(int64(nonce))
	transactor.Value = big.NewInt(0)
	transactor.GasLimit = geth.GasLimit
	transactor.GasPrice = gasPrice

	return transactor, nil
}
