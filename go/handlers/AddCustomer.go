package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	"github.com/garuda/go/crud/mocks"
	"github.com/garuda/go/crud/models"
)

func AddCustomer(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var customer models.Customer
	json.Unmarshal(body, &customer)

	customer.Id = rand.Intn(100)
	mocks.Customers = append(mocks.Customers, customer)

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Created")
}
