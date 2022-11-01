package core

import (
	"context"
	"github.com/AlexRipoll/go-bridge/core/evm"
	"math/big"
	"unsafe"
)

type Service struct {
	Erc721Token evm.Erc721Token
}

func NewService(erc721Token evm.Erc721Token) *Service {
	return &Service{Erc721Token: erc721Token}
}

type Tx struct {
	Hash      string
	Size      float64
	ChainID   *big.Int
	Gas       uint64
	GasPrice  *big.Int
	GasTipCap *big.Int
	GasFeeCap *big.Int
	Value     *big.Int
	Nonce     uint64
}

type TxData struct {
	Wallet      string   `json:"wallet"`
	TokenId     *big.Int `json:"token_id"`
	Origin      string   `json:"origin"`
	Destination string   `json:"destination"`
}

func (s Service) WalletTokens(ctx context.Context, walletAddress string) ([]*big.Int, error) {
	tokens, err := s.Erc721Token.TokensOf(ctx, walletAddress)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func (s Service) releaseNFT(ctx context.Context, walletAddress string, tokenId *big.Int) (*Tx, error) {
	return nil, nil
}

// TODO
func (s Service) CreateCollection(ctx context.Context, walletAddress string, tokenId *big.Int) (*Tx, error) {
	return nil, nil
}

func (s Service) MintERC721Token(ctx context.Context, walletAddress string, tokenId *big.Int) (*Tx, error) {
	tx, err := s.Erc721Token.Mint(ctx, walletAddress, tokenId)
	if err != nil {
		return nil, err
	}

	return toTx(tx), nil
}

func (s Service) burnERC721Token(ctx context.Context, origin string, tokenId *big.Int) (*Tx, error) {
	tx, err := s.Erc721Token.Burn(ctx, tokenId)
	if err != nil {
		return nil, err
	}

	return toTx(tx), nil
}

func toTx(tx *evm.Tx) *Tx {
	return (*Tx)(unsafe.Pointer(tx))
}
