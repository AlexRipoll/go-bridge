package evm

import (
	"context"
	"github.com/AlexRipoll/go-bridge/blockchain/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
)

type Custodian interface {
	RetainNFT(ctx context.Context, tokenId *big.Int) (*Tx, error)
	ReleaseNFT(ctx context.Context, wallet string, tokenId *big.Int) (*Tx, error)
	UpdateOwner(tokenId *big.Int, newHolder common.Address) (*Tx, error)
	Withdraw() (*Tx, error)
	EmergencyDelete(tokenId *big.Int) (*Tx, error)

	HoldCustody(idx *big.Int) (*Custody, error)
	Deployer
}

type Bridger interface {
	Mint(ctx context.Context, to string, tokenId *big.Int) (*Tx, error)
	Burn(ctx context.Context, tokenId *big.Int) (*Tx, error)
	Account
	Deployer
}

type Account interface {
	TokensOf(wallet string) ([]*big.Int, error)
}

type Tx struct {
	Hash    string
	Size    float64
	ChainID *big.Int
	//Data      []byte
	Gas       uint64
	GasPrice  *big.Int
	GasTipCap *big.Int
	GasFeeCap *big.Int
	Value     *big.Int
	Nonce     uint64
	//To        string
}

type Custody struct {
	TokenId *big.Int
	Holder  string
}

func (c custodian) RetainNFT(ctx context.Context, tokenId *big.Int) (*Tx, error) {
	t, err := c.prepareTransactor(ctx)
	if err != nil {
		return nil, err
	}

	gethTx, err := c.Contract.RetainNFT(t, tokenId)
	if err != nil {
		return nil, err
	}
	tx := transactionToTx(gethTx)

	return &tx, nil
}

func (c custodian) ReleaseNFT(ctx context.Context, wallet string, tokenId *big.Int) (*Tx, error) {
	t, err := c.prepareTransactor(ctx)
	if err != nil {
		return nil, err
	}

	walletAddress := common.HexToAddress(wallet)
	gethTx, err := c.Contract.ReleaseNFT(t, tokenId, walletAddress)
	if err != nil {
		return nil, err
	}
	tx := transactionToTx(gethTx)

	return &tx, nil
}

func (c custodian) UpdateOwner(tokenId *big.Int, newHolder common.Address) (*Tx, error) {
	panic("implement me")
}

func (c custodian) Withdraw() (*Tx, error) {
	panic("implement me")
}

func (c custodian) EmergencyDelete(tokenId *big.Int) (*Tx, error) {
	panic("implement me")
}

func (c custodian) HoldCustody(idx *big.Int) (*Custody, error) {
	panic("implement me")
}

func (b bridger) Mint(ctx context.Context, wallet string, tokenId *big.Int) (*Tx, error) {
	t, err := b.prepareTransactor(ctx)
	if err != nil {
		return nil, err
	}

	walletAddress := common.HexToAddress(wallet)
	gethTx, err := b.Contract.BridgeMint(t, walletAddress, tokenId)
	if err != nil {
		return nil, err
	}
	tx := transactionToTx(gethTx)

	return &tx, nil
}

func (b bridger) Burn(ctx context.Context, tokenId *big.Int) (*Tx, error) {
	t, err := b.prepareTransactor(ctx)
	if err != nil {
		return nil, err
	}

	gethTx, err := b.Contract.BridgeBurn(t, tokenId)
	if err != nil {
		return nil, err
	}
	tx := transactionToTx(gethTx)

	return &tx, nil
}

func (b bridger) TokensOf(wallet string) ([]*big.Int, error) {
	address := common.HexToAddress(wallet)
	tokenIds, err := b.Contract.WalletOfOwner(&bind.CallOpts{}, address)
	if err != nil {
		return nil, err
	}

	return tokenIds, err
	// TODO fetch each token data in IPFS and  return ERC721s' data
}

type DeployRx struct {
	Address common.Address
	Tx      Tx
}

func (c custodian) Deploy(ctx context.Context) ([]DeployRx, error) {
	transactor, err := c.prepareTransactor(ctx)
	if err != nil {
		return nil, err
	}

	log.Printf("deploying contract on %s network", c.network)

	var deployRxs []DeployRx
	address, tx, _, err := contract.DeployNFT(transactor, c.conn)
	if err != nil {
		return nil, err
	}
	deployRxs = append(deployRxs, DeployRx{
		Address: address,
		Tx:      transactionToTx(tx),
	})
	log.Printf("deployed contract address: %v", address)

	nonce, err := c.Nonce(ctx)
	if err != nil {
		return nil, err
	}
	transactor.Nonce = big.NewInt(int64(nonce))

	address, tx, _, err = contract.DeployCustosialVault(transactor, c.conn, address)
	if err != nil {
		return nil, err
	}
	deployRxs = append(deployRxs, DeployRx{
		Address: address,
		Tx:      transactionToTx(tx),
	})
	log.Printf("deployed contract address: %v", address)

	return deployRxs, err
}

func (b bridger) Deploy(ctx context.Context) ([]DeployRx, error) {
	transactor, err := b.prepareTransactor(ctx)
	if err != nil {
		return nil, err
	}

	log.Printf("deploying contract on %s network", b.network)

	var deployRxs []DeployRx
	address, tx, _, err := contract.DeployNFT(transactor, b.conn)
	if err != nil {
		return nil, err
	}

	deployRxs = append(deployRxs, DeployRx{
		Address: address,
		Tx:      transactionToTx(tx),
	})
	log.Printf("deployed contract address: %v", address)

	return deployRxs, err
}

func transactionToTx(transaction *types.Transaction) Tx {
	return Tx{
		Hash: transaction.Hash().Hex(),
		Size: float64(transaction.Size()),
		//From:    transaction.,
		ChainID: transaction.ChainId(),
		//Data:      transaction.Data(),
		Gas:       transaction.Gas(),
		GasPrice:  transaction.GasPrice(),
		GasTipCap: transaction.GasTipCap(),
		GasFeeCap: transaction.GasFeeCap(),
		Value:     transaction.Value(),
		Nonce:     transaction.Nonce(),
		//To:        transaction.To().Hex(),
	}
}
