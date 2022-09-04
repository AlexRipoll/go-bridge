package evm

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

type Evm struct {
	client *ethclient.Client
	url string
	pvk string
	address string
	Contracts
}

type Contracts interface {
}

func New(url, pvk, address string) *Evm {
	c, err := ethclient.Dial(url)
	if err != nil {
		panic(err)
	}

	return &Evm{
		client:  c,
		url:     url,
		pvk:     pvk,
		address: address,
	}
}
