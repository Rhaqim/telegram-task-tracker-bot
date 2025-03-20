package config

import (
	"os"
	"sync"

	"github.com/joho/godotenv"

	"github.com/Rhaqim/trackdegens/pkg/logger"
)

type AppConfig struct {
	TelegramBotToken string
	BackupFilePath   string
}

type Config struct {
	config *AppConfig
	once   sync.Once
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) LoadConfig() *AppConfig {
	c.once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			logger.ErrorLogger.Printf("Error loading .env file")
		}

		c.config = &AppConfig{
			TelegramBotToken: c.env("TELEGRAM_BOT_TOKEN", ""),
			BackupFilePath:   c.env("BACKUP_FILE_PATH", ""),
		}
	})

	return c.config
}

func (c *Config) env(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// func LoadConfig() *AppConfig {
// 	once.Do(func() {
// 		err := godotenv.Load()
// 		if err != nil {
// 			logger.ErrorLogger.Printf("Error loading .env file")
// 		}

// 		Config = &AppConfig{
// 			TelegramBotToken: Env("TELEGRAM_BOT_TOKEN", ""),
// 		}
// 	})

// }
