package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/paymentintent"
	"io"
	"log"
	"math"
	"net/http" // used to access the request and response object of the api
	"strconv"

	// package used to read the .env file
	_ "github.com/lib/pq" // postgres golang driver
)

// response format
type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

type item struct {
	Id string `json:"id"`
}

// GetTest will return all the users
func GetTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	res := response{
		ID:      0,
		Message: "Hey",
	}

	json.NewEncoder(w).Encode(res)
}

func HandleCreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Items  []item `json:"items"`
		Amount string `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Error decoding request body: %v", err)
		return
	}

	dollarAmount, errConver := strconv.ParseFloat(req.Amount, 64)

	if errConver != nil {
		log.Fatalf("Unable to convert the string into int.  %v", errConver)
	}

	centsAmount := int(math.Round(dollarAmount * 100))

	// Create a PaymentIntent with amount and currency
	stripeParams := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(centsAmount)),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		// In the latest version of the API, specifying the `automatic_payment_methods` parameter is optional because Stripe enables its functionality by default.
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	pi, err := paymentintent.New(stripeParams)
	log.Printf("pi.New: %v", pi.ClientSecret)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("pi.New: %v", err)
		return
	}

	writeJSON(w, struct {
		ClientSecret string `json:"clientSecret"`
	}{
		ClientSecret: pi.ClientSecret,
	})
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("json.NewEncoder.Encode: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := io.Copy(w, &buf); err != nil {
		log.Printf("io.Copy: %v", err)
		return
	}
}
