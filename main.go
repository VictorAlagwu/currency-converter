package main

import (
	"converter/routes"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	port, exist := os.LookupEnv("PORT")

	if !exist {
		log.Fatal("PORT not set in .env")
	}
	fmt.Println("Server Started on Port :"+port)
	log.Fatal(http.ListenAndServe(":"+port, routes.Init()))
}

