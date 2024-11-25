package repositories

import (
	"fmt"
	"mnc-test/helpers"
	"mnc-test/models"
	"sync"
)

var historyMutex sync.Mutex

func LoadHistories() ([]models.History, error) {
	var histories []models.History
	err := helpers.ReadJSONFile("data/historyData.json", &histories)
	return histories, err
}

func SaveHistories(history models.History) error {
	historyMutex.Lock()
	defer historyMutex.Unlock()

	histories, err := LoadHistories()
	if err != nil {
		return fmt.Errorf("failed to load histories: %w", err)
	}

	histories = append(histories, history)

	return helpers.WriteJSONFile("data/historyData.json", histories)
}
