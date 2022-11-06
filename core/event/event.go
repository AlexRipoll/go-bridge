package event

import (
	"context"
	"encoding/json"
	"math/big"
)

// TODO make TxHash generic
type Rx struct {
	TxHash      [32]byte `json:"tx_hash"`
	TxBlock     uint64   `json:"tx_block"`
	TokenId     *big.Int `json:"token_id"`
	Holder      string   `json:"holder"`
	Origin      *big.Int `json:"origin"`
	Destination *big.Int `json:"destination"`
}

func (rx Rx) Bytes() ([]byte, error) {
	buf, err := json.Marshal(rx)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

type Listener interface {
	Listen(ctx context.Context, ch chan Rx) error
}
