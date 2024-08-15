package service

import (
	"fmt"

	"github.com/Rhaqim/trackdegens/config"
	"github.com/Rhaqim/trackdegens/internal/repo"
	"github.com/Rhaqim/trackdegens/pkg/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start() {
	botToken := config.Config.TelegramBotToken

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		logger.ErrorLogger.Fatalf("Failed to create bot: %v", err)
	}

	logger.InfoLogger.Printf("Authorized on account %s", bot.Self.UserName)

	// Define the bot commands
	commands := []tgbotapi.BotCommand{
		{Command: "start", Description: "Start interacting with the bot"},
		{Command: "track", Description: "Track an item or event"},
		{Command: "status", Description: "Check tracking status"},
	}

	// Set the bot commands
	_, err = bot.Request(tgbotapi.NewSetMyCommands(commands...))
	if err != nil {
		logger.ErrorLogger.Fatalf("Failed to set bot commands: %v", err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Map to store user requests
	userRequests := make(map[int64]string)

	for update := range updates {
		if update.Message == nil { // Ignore non-Message updates
			continue
		}

		userID := update.Message.From.ID
		chatID := update.Message.Chat.ID

		// Handle commands
		command := repo.Commands(update.Message.Command())
		if command.IsValid() {
			command.Handle(bot, update, userRequests)
			continue
		}

		// Handle user responses to the tracking prompt
		if status, exists := userRequests[userID]; exists && status == "awaiting_tracking_info" {
			userRequests[userID] = "tracking_info_received"

			// Store the tracking info
			trackingInfo := update.Message.Text

			// Send confirmation
			confirmationMsg := fmt.Sprintf("Tracking '%s' has been set up.", trackingInfo)
			msg := tgbotapi.NewMessage(chatID, confirmationMsg)
			_, err := bot.Send(msg)
			if err != nil {
				logger.ErrorLogger.Printf("Failed to send message: %v", err)
			}

			// Set a reminder
			go repo.SetReminder(trackingInfo, chatID)
		}
	}
}
