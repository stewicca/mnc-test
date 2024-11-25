package services

import (
	"errors"
	"fmt"
	"mnc-test/helpers"
	"mnc-test/models"
	"mnc-test/repositories"
	"time"
)

func Login(username string, password string) (string, error) {
	customers, err := repositories.LoadCustomers()
	if err != nil {
		return "", fmt.Errorf("failed to load customers: %w", err)
	}

	for _, customer := range customers {
		if customer.Username == username && helpers.VerifyPassword(password, customer.Password) {
			token := helpers.GenerateToken(customer.ID, customer.Username)

			history := models.History{
				Activity: "Login",
				Sender:   customer.Username,
				Time:     time.Now().Local(),
			}

			err = repositories.SaveHistories(history)

			if err != nil {
				return "", fmt.Errorf("failed to save history: %w", err)
			}

			return token, nil
		} else {
			return "", errors.New("invalid credentials")
		}
	}

	return "", errors.New("customer not found")
}

func Logout(token string, username string, expireAt int64) error {
	history := models.History{
		Activity: "Logout",
		Sender:   username,
		Time:     time.Now().Local(),
	}

	err := repositories.SaveHistories(history)

	if err != nil {
		return fmt.Errorf("failed to save history: %w", err)
	}

	tokenBlacklist := models.TokenBlacklist{
		Token:     token,
		ExpiresAt: time.Unix(expireAt, 0).Local(),
	}

	err = repositories.SaveTokenBlacklist(tokenBlacklist)

	if err != nil {
		return fmt.Errorf("failed to save blacklist: %w", err)
	}

	return nil
}
