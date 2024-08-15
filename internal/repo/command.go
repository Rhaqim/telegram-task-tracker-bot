package repo

import (
	"github.com/Rhaqim/trackdegens/pkg/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Commands string

const (
	Track     Commands = "track"
	Status    Commands = "status"
	List      Commands = "list"
	CloseList Commands = "close"
)

func (c Commands) String() string {
	return string(c)
}

func (c Commands) IsValid() bool {
	switch c {
	case Track, Status, List:
		return true
	}
	return false
}

func (c Commands) Handle(bot *tgbotapi.BotAPI, update tgbotapi.Update, userRequests map[int64]string) {
	var numericKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("1"),
			tgbotapi.NewKeyboardButton("2"),
			tgbotapi.NewKeyboardButton("3"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("4"),
			tgbotapi.NewKeyboardButton("5"),
			tgbotapi.NewKeyboardButton("6"),
		),
	)

	switch c {
	case Track:
		// Ask the user what to track
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "What would you like to track?")
		_, err := bot.Send(msg)
		if err != nil {
			logger.ErrorLogger.Printf("Failed to send message: %v", err)
		}

		// Store the tracking request
		userRequests[update.Message.From.ID] = "awaiting_tracking_info"

	case Status:
		// Provide the status of tracking
		// Implement status checking here if needed
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "No tracking info stored.")
		_, err := bot.Send(msg)
		if err != nil {
			logger.ErrorLogger.Printf("Failed to send message: %v", err)
		}
	case List:
		// Provide the list of tracking
		// Implement list checking here if needed
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "No tracking info stored.")
		msg.ReplyMarkup = numericKeyboard
		_, err := bot.Send(msg)
		if err != nil {
			logger.ErrorLogger.Printf("Failed to send message: %v", err)
		}
	case CloseList:
		// Close the list of tracking
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "List closed.")
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		_, err := bot.Send(msg)
		if err != nil {
			logger.ErrorLogger.Printf("Failed to send message: %v", err)
		}
	}
}
