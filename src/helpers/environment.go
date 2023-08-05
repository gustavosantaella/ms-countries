package helpers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvGetProperty(key string) (string, error) {

	var envs map[string]string
	environment := os.Args[1]

	if environment == "production" {
		return os.Getenv(key), nil
	}
	finalEnv := "./envs/.env." + environment

	envs, err := godotenv.Read(finalEnv)

	if err != nil {
		log.Fatal("Error loading environment file")
		return "", err
	}

	return envs[key], nil
}
