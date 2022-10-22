package evm

import (
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

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
