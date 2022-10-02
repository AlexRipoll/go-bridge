package main

import (
	"github.com/AlexRipoll/go-bridge/cmd/go-bridged/handler"
	"github.com/AlexRipoll/go-bridge/config"
	"log"
)

func main() {
	c, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	handler.ServeHTTP(c)
}
