package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/garuda/go/crud/mocks"
)

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Customers)
	// http.ServeFile(w, r, "index.html")
}
