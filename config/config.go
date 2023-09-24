package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	PGconnStr string
	Port      string
}

func Init() config {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("Please provide a DATABASE_URL env var")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Please provide a PORT env var")
	}

	return config{
		PGconnStr: connStr,
		Port:      port,
	}
}
