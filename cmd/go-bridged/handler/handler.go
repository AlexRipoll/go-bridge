package handler

import (
	"github.com/AlexRipoll/go-bridge/blockchain/config"
	v1 "github.com/AlexRipoll/go-bridge/cmd/go-bridged/handler/v1"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func ServeHTTP(config config.Config) {
	r := mux.NewRouter()

	v1.Routes(config, r)

	log.Println("running and listening on port 8000")
	if err := http.ListenAndServe("localhost:8000", r); err != nil {
		log.Fatal(err)
	}
}