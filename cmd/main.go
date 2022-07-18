package main

import (
	"github.com/spf13/viper"
	"log"
	tblBkng "tables-booking-app"
	"tables-booking-app/pkg/handler"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	handlers := new(handler.Handler)

	srv := new(tblBkng.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
