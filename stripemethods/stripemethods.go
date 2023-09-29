package stripemethods

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func DotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}
