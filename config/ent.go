package config

import (
	"Blog/ent"
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	client *ent.Client
)

func SetClient(c *ent.Client) {
	client = c
}

func GetClient() *ent.Client {
	return client
}

func NewEntClient() (*ent.Client, error) {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	DB_DATABASE := os.Getenv("DB_DATABASE")
	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")

	client, err := sql.Open("postgres", "host="+DB_HOST+" port="+DB_PORT+" user="+DB_USERNAME+" dbname="+DB_DATABASE+" password="+DB_PASSWORD+" sslmode=disable")

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Ping(); err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	return ent.Open("postgres", "host="+DB_HOST+" port="+DB_PORT+" user="+DB_USERNAME+" dbname="+DB_DATABASE+" password="+DB_PASSWORD+" sslmode=disable")
}
