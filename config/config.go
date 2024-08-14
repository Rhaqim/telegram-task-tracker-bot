package config

import (
	"os"
	"sync"

	"github.com/joho/godotenv"

	"github.com/Rhaqim/trackdegens/pkg/logger"
)

type AppConfig struct {
	TelegramBotToken string
}

var (
	Config *AppConfig
	once   sync.Once
)

func LoadConfig() {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			logger.ErrorLogger.Printf("Error loading .env file")
		}

		Config = &AppConfig{
			TelegramBotToken: Env("TELEGRAM_BOT_TOKEN", ""),
		}
	})

}

func Env(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
