package config

import (
	"log"

	"github.com/joho/godotenv"

)

func SetupConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
