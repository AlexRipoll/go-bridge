package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/AlexRipoll/go-bridge/config"
	"github.com/AlexRipoll/go-bridge/core"
	"github.com/AlexRipoll/go-bridge/core/event"
	evm2 "github.com/AlexRipoll/go-bridge/core/evm"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	"math/big"
	"net/http"
	"strconv"
)

type handler struct {
	bridge *core.Bridge
}

func NewHandler(config config.Config) (*handler, error) {
	var custodian evm2.Custodian
	bridgers := make(map[string]evm2.Erc721Token)

	for name, network := range config.Networks {
		conn, err := ethclient.Dial(network.Http)
		if err != nil {
			return nil, err
		}
		for contract, address := range network.Contracts {
			switch contract {
			case "vault":
				custodian, err = evm2.NewCustodian(conn, address, config, name)
				if err != nil {
					return nil, err
				}
			case "bridge":
				bridger, err := evm2.NewBridger(conn, address, config, name)
				if err != nil {
					return nil, err
				}
				bridgers[name] = bridger
			default:
				return nil, err
			}
		}
	}

	bridgeSrv, err := core.NewBridge(config.MainNetwork, custodian, bridgers)
	if err != nil {
		return nil, err
	}

	return &handler{bridge: bridgeSrv}, nil
}

func (h handler) Transfer(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Destination string `json:"destination"`
		Origin      string `json:"origin"`
	}

	vars := mux.Vars(r)
	wallet := vars["wallet_address"]
	tokenIdStr := vars["token_id"]
	tokenId, err := strconv.Atoi(tokenIdStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	fmt.Println(r.Body)
	fmt.Println(body)
	err = h.bridge.TransferNFT(context.Background(), body.Destination, body.Origin, wallet, big.NewInt(int64(tokenId)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func (h handler) CompleteTransfer(ch chan event.Rx) error {
	return h.bridge.CompleteTransfer(context.Background(), ch)
}
