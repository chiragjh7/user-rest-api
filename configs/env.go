package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func EnvMongoURI() string {
	
	// Load .env file
	err := godotenv.Load()

	// if loading .env file returns an error
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// return the value of MONGOURI
	return os.Getenv("MONGOURI")
}
