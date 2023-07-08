package fyi

import (
	"os"
)

func ReadFile(filePath string) (*string, error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	text := string(fileContent)
	return &text, nil
}
