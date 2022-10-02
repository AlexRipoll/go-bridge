package handler

import (
	"fmt"
	v1 "github.com/AlexRipoll/go-bridge/cmd/go-bridged/handler/v1"
	"github.com/AlexRipoll/go-bridge/config"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func ServeHTTP(config config.Config) {
	r := mux.NewRouter()

	v1.Routes(config, r)

	log.Println(fmt.Sprintf("running and listening on port %s", config.Host))
	if err := http.ListenAndServe(config.Host, r); err != nil {
		log.Fatal(err)
	}
}
