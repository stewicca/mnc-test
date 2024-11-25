package repositories

import (
	"mnc-test/helpers"
	"mnc-test/models"
	"sync"
)

var customerMutex sync.Mutex

func LoadCustomers() ([]models.Customer, error) {
	var customers []models.Customer
	err := helpers.ReadJSONFile("data/customerData.json", &customers)
	return customers, err
}

func SaveCustomers(customers []models.Customer) error {
	customerMutex.Lock()
	defer customerMutex.Unlock()
	return helpers.WriteJSONFile("data/customerData.json", customers)
}
