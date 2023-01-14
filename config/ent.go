package config

import (
	"Blog/ent"
	"database/sql"
	"log"

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

	client, err := sql.Open("postgres", "host=localhost port=5432 user=postgres dbname=Blog password=1234 sslmode=disable")

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Ping(); err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	return ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=Blog password=1234 sslmode=disable")
}
