package v1

import (
	"github.com/AlexRipoll/go-bridge/config"
	"github.com/AlexRipoll/go-bridge/core/event"
	"github.com/gorilla/mux"
	"log"
)

func Routes(config config.Config, router *mux.Router) {
	h, err := NewHandler(config)
	if err != nil {
		log.Fatal(err)
	}

	ch := make(chan event.Rx)

	go h.CompleteTransfer(ch)
	go RunEvmEventListener(config, ch)

	// endpoints
	router.HandleFunc("/wallets/{wallet_address}/tokens/{token_id}/transfer", h.Transfer).Methods("POST")
}
