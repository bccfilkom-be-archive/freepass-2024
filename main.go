package main

import (
	"bcc-be-freepass-2024/database"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error when loading .env file: " + err.Error())
	}

	db, err := database.MakeConnection(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	if err != nil {
		log.Fatalln("Error when connecting to database: " + err.Error())
	}
	log.Println(db)

}
