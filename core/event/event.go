package event

import (
	"context"
	"math/big"
)

// TODO make TxHash generic
type Rx struct {
	TxHash      [32]byte
	TxBlock     uint64
	TokenId     *big.Int
	Holder      string
	Destination *big.Int
}

type Listener interface {
	Listen(ctx context.Context, ch chan Rx) error
}
