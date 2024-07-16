package services

import (
	"daveslist-emdpcv/api/settings"
	"fmt"

	"github.com/google/uuid"
)

// GenerateToken generates a new token if client_id and secret_id are correct
func GenerateToken(clientID, secretID string) (string, error) {
	config := settings.LoadEnv()
	if clientID == config.UsernameGateWay && secretID == config.PasswordGateWay {
		token := uuid.New().String()
		err := StoreAccessToken(token, 60) // Token expires in 60 minutes
		if err != nil {
			return "", fmt.Errorf("could not store token: %v", err)
		}
		return token, nil
	}
	return "", fmt.Errorf("invalid client_id or secret_id")
}
