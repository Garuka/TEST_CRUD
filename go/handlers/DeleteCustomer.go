package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/garuda/go/crud/mocks"
	"github.com/gorilla/mux"
)

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for index, customer := range mocks.Customers {
		if customer.Id == id {
			mocks.Customers = append(mocks.Customers[:index], mocks.Customers[index+1:]...)

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Delete")
			break
		}
	}
}
