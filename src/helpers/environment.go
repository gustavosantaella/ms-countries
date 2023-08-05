package helpers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvGetProperty(key string) (string, error) {

	var envs map[string]string
	arguments := os.Args
	if len(arguments) == 1 {
		return os.Getenv(key), nil
	}
	finalEnv := "./envs/.env." + arguments[1]

	envs, err := godotenv.Read(finalEnv)

	if err != nil {
		log.Fatal("Error loading environment file")
		return "", err
	}

	return envs[key], nil
}
