package main

import (
	"github.com/sashindionicus/rest-api"
	"github.com/sashindionicus/rest-api/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(rest.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server %s", err.Error())
	}

}
