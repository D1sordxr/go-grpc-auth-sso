package main

import (
	"fmt"
	"log"
	loadConfig "sso/internal/infrastructure/config"
)

func main() {
	cfg, err := loadConfig.NewConfig()
	if err != nil {
		log.Fatalf("failed init config: %v", err)
	}
	fmt.Println(cfg)

	// TODO: init logger

	// TODO: init app

	// TODO: run gRPC-server
}
