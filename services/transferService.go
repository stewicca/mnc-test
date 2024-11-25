package services

import (
	"errors"
	"fmt"
	"mnc-test/models"
	"mnc-test/repositories"
	"time"
)

func Transfer(senderID uint, recipientID uint, amount int) error {
	customers, err := repositories.LoadCustomers()
	if err != nil {
		return fmt.Errorf("failed to load customers: %w", err)
	}

	for i, customer := range customers {
		if customer.ID == recipientID {
			if customers[senderID].Balance < amount {
				return errors.New("insufficient balance")
			}

			customers[senderID].Balance -= amount
			customers[i].Balance += amount

			err = repositories.SaveCustomers(customers)
			if err != nil {
				return fmt.Errorf("failed to save customers: %w", err)
			}

			history := models.History{
				Activity:  "Transfer",
				Sender:    customers[senderID].Username,
				Recipient: customers[i].Username,
				Amount:    amount,
				Time:      time.Now().Local(),
			}
			return repositories.SaveHistories(history)
		}
	}
	return errors.New("recipient not found")
}
