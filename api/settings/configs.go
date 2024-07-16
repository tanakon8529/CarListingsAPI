package settings

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// AppSettings holds all the application settings
type AppSettings struct {
	ApiVersion      string
	ApiPath         string
	ApiDoc          string
	Host            string
	HostGateway     string
	Port            string
	UsernameGateWay string
	PasswordGateWay string
	RedisHost       string
	RedisPort       string
	RedisPassword   string
}

// LoadEnv loads environment variables
func LoadEnv() AppSettings {
	err := godotenv.Load("/app/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return AppSettings{
		ApiVersion:      os.Getenv("API_VERSION"),
		ApiPath:         os.Getenv("API_PATH"),
		ApiDoc:          os.Getenv("API_DOC"),
		Host:            os.Getenv("HOST"),
		HostGateway:     os.Getenv("HOST_GATEWAY"),
		Port:            os.Getenv("PORT_GINAPI"),
		UsernameGateWay: os.Getenv("USERNAME_GATEWAY"),
		PasswordGateWay: os.Getenv("PASSWORD_GATEWAY"),
		RedisHost:       os.Getenv("HOST_REDIS"),
		RedisPort:       os.Getenv("PORT_REDIS"),
		RedisPassword:   os.Getenv("REDIS_PASSWORD"),
	}
}
