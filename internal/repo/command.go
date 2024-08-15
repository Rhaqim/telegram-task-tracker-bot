package repo

import (
	"fmt"
	"strings"

	"github.com/Rhaqim/trackdegens/pkg/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Commands string

const (
	Start  Commands = "start"
	Track  Commands = "track"
	Status Commands = "status"
	List   Commands = "list"
	Done   Commands = "done"
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

func (c Commands) Handle(bot *tgbotapi.BotAPI, update tgbotapi.Update, userRequests map[int64]string, userEntries map[int64][]string) {
	// var numericKeyboard = tgbotapi.NewReplyKeyboard(
	// 	tgbotapi.NewKeyboardButtonRow(
	// 		tgbotapi.NewKeyboardButton("1"),
	// 		tgbotapi.NewKeyboardButton("2"),
	// 		tgbotapi.NewKeyboardButton("3"),
	// 	),
	// 	tgbotapi.NewKeyboardButtonRow(
	// 		tgbotapi.NewKeyboardButton("4"),
	// 		tgbotapi.NewKeyboardButton("5"),
	// 		tgbotapi.NewKeyboardButton("6"),
	// 	),
	// )

	switch c {
	case Start:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Welcome to the bot! You can start tracking items or events by using the /track command.")
		_, err := bot.Send(msg)
		if err != nil {
			logger.ErrorLogger.Printf("Failed to send message: %v", err)
		}

	case Track:
		// Ask the user what to track
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please enter the title for your tracking:")
		_, err := bot.Send(msg)
		if err != nil {
			logger.ErrorLogger.Printf("Failed to send message: %v", err)
		}
		userRequests[update.Message.From.ID] = "awaiting_title"

	case Status:
		// Provide the status of tracking
		// Implement status checking here if needed
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "No tracking info stored.")
		_, err := bot.Send(msg)
		if err != nil {
			logger.ErrorLogger.Printf("Failed to send message: %v", err)
		}

	case List:
		trackedActivities := getTrackedActivities(update.Message.From.ID)
		if len(trackedActivities) == 0 {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "No tracked activities found.")
			_, err := bot.Send(msg)
			if err != nil {
				logger.ErrorLogger.Printf("Failed to send message: %v", err)
			}
			return
		}

		// Display the tracked activities back to the user
		msgText := fmt.Sprintf("Tracked activities for '%s':\n", trackedActivities[0])
		for _, entry := range trackedActivities[1:] {
			msgText += fmt.Sprintf("- %s\n", entry)
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
		_, err := bot.Send(msg)
		if err != nil {
			logger.ErrorLogger.Printf("Failed to send message: %v", err)
		}

	case Done:
		entries, exists := userEntries[update.Message.From.ID]
		if exists && len(entries) > 0 {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Your tracking entries have been saved:\n"+strings.Join(entries, "\n"))
			_, err := bot.Send(msg)
			if err != nil {
				logger.ErrorLogger.Printf("Failed to send message: %v", err)
			}
			// Optionally, clear the entries after done
			delete(userEntries, update.Message.From.ID)
			delete(userRequests, update.Message.From.ID)
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "No entries were saved.")
			_, err := bot.Send(msg)
			if err != nil {
				logger.ErrorLogger.Printf("Failed to send message: %v", err)
			}
		}

	default:
		if status, exists := userRequests[update.Message.From.ID]; exists {
			if status == "awaiting_title" {
				userRequests[update.Message.From.ID] = "awaiting_entries"
				userEntries[update.Message.From.ID] = []string{update.Message.Text}
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Now, enter each item you want to track, and send /done when you're finished:")
				_, err := bot.Send(msg)
				if err != nil {
					logger.ErrorLogger.Printf("Failed to send message: %v", err)
				}
			} else if status == "awaiting_entries" {
				userEntries[update.Message.From.ID] = append(userEntries[update.Message.From.ID], update.Message.Text)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Entry added. Send /done when finished or keep adding entries.")
				_, err := bot.Send(msg)
				if err != nil {
					logger.ErrorLogger.Printf("Failed to send message: %v", err)
				}
			}
		}
	}
}

func getTrackedActivities(userID int64) []string {
	// Here, we're just returning a sample set of activities.
	return []string{
		"Sample Title",
		"Activity 1: Description",
		"Activity 2: Description",
		"Activity 3: Description",
	}
}
