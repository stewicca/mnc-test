package repositories

import (
	"fmt"
	"mnc-test/helpers"
	"mnc-test/models"
	"sync"
	"time"
)

var tokenBlacklistMutex sync.Mutex

func LoadTokenBlacklist() ([]models.TokenBlacklist, error) {
	var blacklist []models.TokenBlacklist
	err := helpers.ReadJSONFile("data/tokenBlacklistData.json", &blacklist)
	return blacklist, err
}

func SaveTokenBlacklist(tokenBlacklist models.TokenBlacklist) error {
	tokenBlacklistMutex.Lock()
	defer tokenBlacklistMutex.Unlock()

	blacklist, err := LoadTokenBlacklist()
	if err != nil {
		return fmt.Errorf("failed to load blacklist: %w", err)
	}

	blacklist = append(blacklist, tokenBlacklist)

	return helpers.WriteJSONFile("data/tokenBlacklistData.json", blacklist)
}

func CleanTokenExpired() error {
	tokenBlacklistMutex.Lock()
	defer tokenBlacklistMutex.Unlock()

	blacklist, err := LoadTokenBlacklist()
	if err != nil {
		return fmt.Errorf("failed to load blacklist: %w", err)
	}

	var updatedBlacklist []models.TokenBlacklist
	for _, token := range blacklist {
		if token.ExpiresAt.After(time.Now()) {
			updatedBlacklist = append(updatedBlacklist, token)
		}
	}

	return helpers.WriteJSONFile("data/tokenBlacklistData.json", updatedBlacklist)
}
