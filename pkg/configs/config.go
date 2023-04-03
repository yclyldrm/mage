package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port          string
	RedisURL      string
	RedisPassword string
	DBUser        string
	DBPassword    string
	SecretKey     string
}

var Cfg Config

func LoadConfig() {
	err := godotenv.Load("./pkg/configs/envs/.env")
	if err != nil {
		log.Fatalf("Error loading env file,error: %v", err)

	}

	Cfg.Port = os.Getenv("PORT")
	Cfg.RedisURL = os.Getenv("REDIS_URL")
	Cfg.RedisPassword = os.Getenv("REDIS_PASSWORD")
	Cfg.DBUser = os.Getenv("DB_USERNAME")
	Cfg.DBPassword = os.Getenv("DB_PASSWORD")
	Cfg.SecretKey = os.Getenv("SECRET_KEY")
}
