package service

import (
	"fmt"

	"github.com/Rhaqim/trackdegens/internal/model"
	"github.com/Rhaqim/trackdegens/pkg/logger"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start(bot *tg.BotAPI) {

	// Define the bot commands
	commands := []tg.BotCommand{
		{Command: model.Start.String(), Description: "Start interacting with the bot"},
		{Command: model.Track.String(), Description: "Track an item or event"},
		{Command: model.Status.String(), Description: "Check tracking status"},
		{Command: model.List.String(), Description: "List all tracked items"},
		{Command: model.Done.String(), Description: "Stop tracking an item"},
	}

	// Set the bot commands
	_, err := bot.Request(tg.NewSetMyCommands(commands...))
	if err != nil {
		logger.ErrorLogger.Fatalf("Failed to set bot commands: %v", err)
	}

	u := tg.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Map to store user requests
	userRequests := make(map[int64]string)

	userEntries := make(map[int64][]string)

	for update := range updates {
		if update.Message == nil { // Ignore non-Message updates
			continue
		}

		userID := update.Message.From.ID
		chatID := update.Message.Chat.ID

		// Handle commands
		command := model.Commands(update.Message.Command())
		if command.IsValid() {
			command.Handle(bot, update, userRequests, userEntries)
			continue
		}

		// Handle user responses to the tracking prompt
		if status, exists := userRequests[userID]; exists && status == "awaiting_tracking_info" {
			userRequests[userID] = "tracking_info_received"

			// Store the tracking info
			trackingInfo := update.Message.Text

			// Send confirmation
			confirmationMsg := fmt.Sprintf("Tracking '%s' has been set up.", trackingInfo)
			msg := tg.NewMessage(chatID, confirmationMsg)
			_, err := bot.Send(msg)
			if err != nil {
				logger.ErrorLogger.Printf("Failed to send message: %v", err)
			}

			// Set a reminder
			go model.SetReminder(trackingInfo, chatID)
		}
	}
}
