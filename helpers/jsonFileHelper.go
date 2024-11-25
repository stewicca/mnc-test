package helpers

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"os"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func ReadJSONFile[T any](filePath string, output *T) error {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", filePath, err)
	}
	if err := json.Unmarshal(file, output); err != nil {
		return fmt.Errorf("failed to unmarshal JSON from file %s: %w", filePath, err)
	}
	return nil
}

func WriteJSONFile[T any](filePath string, input T) error {
	data, err := json.MarshalIndent(input, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}
	return nil
}
