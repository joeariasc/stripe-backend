package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v76"
	"log"
	"net/http"
	"os"
	"stripe-backend/router"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// This is your test secret API key.
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	hostAddress := os.Getenv("HOST_ADDRESS")
	if hostAddress == "" {
		// Provide a default host address if none is specified
		hostAddress = "127.0.0.1"
	}

	r := router.Router()

	fmt.Printf("Starting server on the host %s port 3000...\n", hostAddress)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:3000", hostAddress), r))
}
