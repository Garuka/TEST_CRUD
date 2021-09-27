package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/garuda/go/crud/mocks"
	"github.com/garuda/go/crud/models"
	"github.com/gorilla/mux"
)

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	body, err := ioutil.ReadAll(r.Body)

	defer r.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var updateCustomer models.Customer
	json.Unmarshal(body, &updateCustomer)

	for index, customer := range mocks.Customers {
		if customer.Id == id {
			customer.Name = updateCustomer.Name
			customer.Lname = updateCustomer.Lname
			customer.Password = updateCustomer.Password
			customer.Email = updateCustomer.Email
			customer.Img = updateCustomer.Img

			mocks.Customers[index] = customer
			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode("Update")
		}
	}

}
