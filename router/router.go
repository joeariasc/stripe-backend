package router

import (
	"github.com/gorilla/mux"
	"stripe-backend/middleware"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/test", middleware.GetTest).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/payment", middleware.HandleCreatePaymentIntent).Methods("POST", "OPTION")

	return router
}
