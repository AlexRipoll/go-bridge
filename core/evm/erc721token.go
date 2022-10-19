package evm

import (
	"context"
	"github.com/AlexRipoll/go-bridge/contract"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Erc721Token interface {
	Mint(ctx context.Context, to string, tokenId *big.Int) (*Tx, error)
	Burn(ctx context.Context, tokenId *big.Int) (*Tx, error)
	Exists(ctx context.Context, tokenId *big.Int) (bool, error)
	OwnerOf(ctx context.Context, tokenId *big.Int) (string, error)
	TokensOf(ctx context.Context, wallet string) ([]*big.Int, error)
}

type erc721TokenContract struct {
	contract   *contract.Erc721Token
	transactor geth
	//Client
}

func NewErc721TokenContract(address string, client Client) (Erc721Token, error) {
	contract, err := contract.NewErc721Token(common.HexToAddress(address), client.conn)
	if err != nil {
		return nil, err
	}
	return erc721TokenContract{contract: contract}, nil
}

func (b erc721TokenContract) Mint(ctx context.Context, wallet string, tokenId *big.Int) (*Tx, error) {
	txOps, err := b.transactor.prepareTransactor(ctx)
	if err != nil {
		return nil, err
	}

	walletAddress := common.HexToAddress(wallet)
	gethTx, err := b.contract.Mint(txOps, walletAddress, tokenId)
	tx := transactionToTx(gethTx)

	return &tx, nil
}

func (b erc721TokenContract) Burn(ctx context.Context, tokenId *big.Int) (*Tx, error) {
	txOps, err := b.transactor.prepareTransactor(ctx)
	if err != nil {
		return nil, err
	}

	gethTx, err := b.contract.Burn(txOps, tokenId)
	if err != nil {
		return nil, err
	}
	tx := transactionToTx(gethTx)

	return &tx, nil
}

func (b erc721TokenContract) Exists(ctx context.Context, tokenId *big.Int) (bool, error) {
	callOpts, err := b.transactor.prepareCallOpts(ctx)
	if err != nil {
		return false, err
	}

	return b.contract.Exists(callOpts, tokenId)
}

func (b erc721TokenContract) OwnerOf(ctx context.Context, tokenId *big.Int) (string, error) {
	callOpts, err := b.transactor.prepareCallOpts(ctx)
	if err != nil {
		return "", err
	}

	gethTx, err := b.contract.OwnerOf(callOpts, tokenId)
	if err != nil {
		return "", err
	}

	return gethTx.String(), nil
}

func (b erc721TokenContract) TokensOf(ctx context.Context, wallet string) ([]*big.Int, error) {
	callOpts, err := b.transactor.prepareCallOpts(ctx)
	if err != nil {
		return nil, err
	}

	address := common.HexToAddress(wallet)
	tokenIds, err := b.contract.WalletOfOwner(callOpts, address)
	if err != nil {
		return nil, err
	}

	return tokenIds, err
}
