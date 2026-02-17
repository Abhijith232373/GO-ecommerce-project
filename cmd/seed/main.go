package main

import (
	"e-commerce/internal/config"
	"e-commerce/internal/seed"
	"log"
)

func main() {

	log.Println("loading env...")

	config.LoadEnv()
	log.Println("Connecting database...")

	config.ConnectDB()

	log.Println("Running seed...")

	seed.SeedAll()

	log.Println("Seed finished successfully")
}