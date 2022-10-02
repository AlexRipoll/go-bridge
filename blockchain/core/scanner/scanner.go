package scanner

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
)

type EventRx struct {
	TxHash common.Hash
	Action string
}

type EventSubscriber interface {
	ListenEvents(ctx context.Context, ch chan EventRx) error
}
