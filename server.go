package main

import (
	"Blog/config"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"Blog/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port := os.Getenv("PORT_ENV")

	fmt.Println("port : ", port)
	if port == "" {
		port = defaultPort
	}

	//initiate Ent Client
	client, err := config.NewEntClient()

	if err != nil {
		log.Printf("err : %s", err)
	}
	defer client.Close()

	if err != nil {
		log.Println("Fail to initialize client")
	}

	//set the client to the variable defined in package config
	//this will enable the client instance to be accessed anywhere through the accessor which is a function
	//named GetClient
	config.SetClient(client)

	//initiate router and register all the route
	r := gin.Default()
	router.RegisterRouter(r)

	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server started on port " + port)
	log.Fatal(srv.ListenAndServe())
}
