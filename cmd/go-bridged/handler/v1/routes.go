package v1

import (
	"github.com/AlexRipoll/go-bridge/blockchain/config"
	"github.com/gorilla/mux"
	"log"
)

func Routes(config config.Config, router *mux.Router) {
	h, err := NewHandler(config)
	if err != nil {
		log.Fatal(err)
	}

	// endpoints
	router.HandleFunc("/wallets/{wallet_address}/tokens/{token_id}/transfer", h.Transfer).Methods("POST")

}
