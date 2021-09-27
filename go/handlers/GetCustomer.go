package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/garuda/go/crud/mocks"
	"github.com/gorilla/mux"
)

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for _, customer := range mocks.Customers {
		if customer.Id == id {
			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(customer)
			break
		}
	}
}
