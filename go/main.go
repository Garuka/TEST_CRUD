package main

import (
	"log"
	"net/http"

	"github.com/garuda/go/crud/handlers"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/customers", handlers.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id}", handlers.GetCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers", handlers.AddCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{id}", handlers.UpdateCustomer).Methods(http.MethodPut)
	router.HandleFunc("/customers/{id}", handlers.DeleteCustomer).Methods(http.MethodDelete)

	log.Println("API is Running")
	http.ListenAndServe(":4000", router)
}
