package main

import (
	"fmt"
	"log"

	config "github.com/joejosephvarghese/message/server/pkg"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	fmt.Printf("Loaded Config: %+v\n", cfg)
}
