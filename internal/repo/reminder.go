package repo

import (
	"fmt"
	"time"

	"github.com/Rhaqim/trackdegens/pkg/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// setReminder schedules a reminder for the tracking info
func SetReminder(info string, chatID int64) {
	// Example: Remind in 5 minutes
	time.Sleep(5 * time.Minute)

	reminderMsg := fmt.Sprintf("Reminder: Track '%s'.", info)
	bot, _ := tgbotapi.NewBotAPI("YOUR_TELEGRAM_BOT_TOKEN") // Replace with your bot token
	msg := tgbotapi.NewMessage(chatID, reminderMsg)
	_, err := bot.Send(msg)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to send reminder message: %v", err)
	}
}
