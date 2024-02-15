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

	r := router.Router()

	fmt.Println("Starting server on the port 3000...")

	log.Fatal(http.ListenAndServe(":3000", r))
}
