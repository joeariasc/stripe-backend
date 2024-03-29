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

	if err != nil && !os.IsNotExist(err) {
		fmt.Println(err.Error())
		log.Fatalf("Error loading .env file")
	} else {
		fmt.Println("Running without a .env file. Assuming environment variables are set.")
	}

	// This is your test secret API key.
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	r := router.Router()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	fmt.Printf("Starting server on the port %s...\n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
