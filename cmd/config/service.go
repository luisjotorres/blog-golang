package config

import (
	"github.com/joho/godotenv"
)

func NewLocalClient() {
	godotenv.Load(".env")
}

func GetDatabaseConfig() {
}
