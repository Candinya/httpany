package main

import (
	"httpany/inits"
	"log"
)

func main() {
	log.Print("Service starting...")
	if err := inits.WebEngine().Run(); err != nil {
		log.Fatalf("Failed to start service: %v", err)
	}
}
