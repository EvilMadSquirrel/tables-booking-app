package main

import (
	"log"
	tblBkng "tables-booking-app"
	"tables-booking-app/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(tblBkng.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
