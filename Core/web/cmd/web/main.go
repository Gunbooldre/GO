package main

import (
	"fmt"
	"log"
	"web/internal/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("config: %v", err)
	}
	//cfg := config.Config{}
	fmt.Println(cfg)
}
