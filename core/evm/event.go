package evm

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

type TokenCustodyEvent struct {
	TokenId     *big.Int
	Holder             common.Address
	DestinationChainId *big.Int
}

var tokenCustodyEventName = "TokenCustody"
var tokenCustodySig = []byte(fmt.Sprintf("%s(uint256,address,uint256)", tokenCustodyEventName))

func TokenCustodySigHash() string {
	return crypto.Keccak256Hash(tokenCustodySig).Hex()
}
