package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func main() {
	apiKey, err := generateAPIKey(32)
	if err != nil {
		panic(err)
	}

	fmt.Printf("API_KEY: %s\n", apiKey)
}

func generateAPIKey(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	apiKey := base64.URLEncoding.EncodeToString(bytes)
	return apiKey, nil
}
