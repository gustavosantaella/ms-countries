package helpers

import (
	"log"

	"github.com/joho/godotenv"
)

func EnvGetProperty(key string) (string, error) {

	var envs map[string]string

	envs, err := godotenv.Read("./envs/.env.local")

	if err != nil {
		log.Fatal("Error loading environment file")
		return "", err
	}

	return envs[key], nil
}
