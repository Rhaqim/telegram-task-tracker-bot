package service

import (
	"fmt"
	"log"
	"time"

	"github.com/Rhaqim/trackdegens/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TrackingRequest struct {
	UserID       int64
	TrackingInfo string
	Timestamp    time.Time
}

func Start() {
	botToken := config.Config.TelegramBotToken

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

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

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "track":
				// Ask the user what to track
				msg := tgbotapi.NewMessage(chatID, "What would you like to track?")
				_, err := bot.Send(msg)
				if err != nil {
					log.Printf("Failed to send message: %v", err)
				}

				// Store the tracking request
				userRequests[userID] = "awaiting_tracking_info"
				continue

			case "status":
				// Provide the status of tracking
				// Implement status checking here if needed
				msg := tgbotapi.NewMessage(chatID, "No tracking info stored.")
				_, err := bot.Send(msg)
				if err != nil {
					log.Printf("Failed to send message: %v", err)
				}

				continue
			}
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
				log.Printf("Failed to send message: %v", err)
			}

			// Set a reminder
			go setReminder(trackingInfo, chatID)
		}
	}
}

// setReminder schedules a reminder for the tracking info
func setReminder(info string, chatID int64) {
	// Example: Remind in 5 minutes
	time.Sleep(5 * time.Minute)

	reminderMsg := fmt.Sprintf("Reminder: Track '%s'.", info)
	bot, _ := tgbotapi.NewBotAPI("YOUR_TELEGRAM_BOT_TOKEN") // Replace with your bot token
	msg := tgbotapi.NewMessage(chatID, reminderMsg)
	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Failed to send reminder message: %v", err)
	}
}

// func Start() {
// 	// Replace with your bot's token
// 	botToken := config.Config.TelegramBotToken

// 	bot, err := tgbotapi.NewBotAPI(botToken)
// 	if err != nil {
// 		log.Fatalf("Failed to create bot: %v", err)
// 	}

// 	log.Printf("Authorized on account %s", bot.Self.UserName)

// 	u := tgbotapi.NewUpdate(0)
// 	u.Timeout = 60

// 	updates := bot.GetUpdatesChan(u)

// 	for update := range updates {
// 		if update.Message == nil {
// 			continue
// 		}

// 		log.Printf("Received message from %s: %s", update.Message.From.UserName, update.Message.Text)

// 		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello, World!")
// 		_, err := bot.Send(msg)
// 		if err != nil {
// 			log.Printf("Failed to send message: %v", err)
// 		}
// 	}
// }
