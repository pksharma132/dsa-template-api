package main

import (
	"log"
	"dsa-template-api/internal/api"
)

func main() {
	log.Println("Starting dsa-template-api...")
	api.StartServer()
}
