package event

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
)

const (
	MintAction    = "mint"
	BurnAction    = "burn"
	ReleaseAction = "release"
)

// TODO make TxHash generic
type Rx struct {
	TxHash common.Hash
	Action string
}

type Listener interface {
	Listen(ctx context.Context, ch chan Rx) error
}
