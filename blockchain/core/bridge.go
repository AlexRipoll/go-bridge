package core

import (
	"context"
	"github.com/AlexRipoll/go-bridge/blockchain/core/evm"
	"math/big"
	"unsafe"
)

type Bridge struct {
	contracts evm.Blockchain
}

func NewBridge(blockchain evm.Blockchain) (*Bridge, error) {
	return &Bridge{
		blockchain,
	}, nil
}

type Tx struct {
	Hash string
	Size float64
	//From string
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

type Rx struct {
	hash string
}

func (b Bridge) RetainNFT(ctx context.Context, tokenId *big.Int) (*Tx, error) {
	tx, err := b.contracts.RetainNFT(ctx, tokenId)
	if err != nil {
		return nil, err
	}

	return toTx(tx), nil
}

func toTx(tx *evm.Tx) *Tx {
	return (*Tx)(unsafe.Pointer(tx))
}
