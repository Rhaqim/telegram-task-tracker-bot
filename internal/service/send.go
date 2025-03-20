package service

import (
	"os"

	"github.com/Rhaqim/trackdegens/pkg/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendBackupFile(bot *tgbotapi.BotAPI, chatID int64, filePath string) {

	file, err := GetFile(filePath)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to get file: %v", err)
	}

	docConfig := tgbotapi.NewDocument(chatID, tgbotapi.FileBytes{
		Name:  "backup.sql",
		Bytes: file,
	})

	_, err = bot.Send(docConfig)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to send backup file: %v", err)
	}
}

func GetFile(filePath string) ([]byte, error) {

	file, err := os.ReadFile(filePath)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to read file: %v", err)
	}
	return file, err
}
