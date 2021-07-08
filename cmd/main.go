package main

import (
	"log"
	"testSellerX/internal/apiserver"

	"github.com/BurntSushi/toml"
)

func main() {
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile("configs/config.toml", config)
	if err != nil {
		log.Fatal(err)
	}
	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}
}
