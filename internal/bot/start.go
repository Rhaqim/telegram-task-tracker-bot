package bot

import (
	"fmt"

	"github.com/Rhaqim/trackdegens/internal/domain"
	"github.com/Rhaqim/trackdegens/internal/model"
	"github.com/Rhaqim/trackdegens/pkg/logger"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramDegenBot struct {
	bot *tg.BotAPI
}

func Initialize(botToken string) domain.TelegramDegenBot {
	bot, err := tg.NewBotAPI(botToken)
	if err != nil {
		logger.ErrorLogger.Fatalf("Failed to create bot: %v", err)
	}

	logger.InfoLogger.Printf("Authorized on account %s", bot.Self.UserName)

	return &TelegramDegenBot{bot: bot}
}

func (t *TelegramDegenBot) Start() {
	t.loadCommands()
	t.startUpdates()
}

// SendMessage implements domain.TelegramDegenBot.
func (t *TelegramDegenBot) SendMessage(chatID int64, message string) error {
	panic("unimplemented")
}

func (t *TelegramDegenBot) loadCommands() {
	commands := []tg.BotCommand{
		{Command: "start", Description: "Start interacting with the bot"},
		{Command: "track", Description: "Track an item or event"},
		{Command: "status", Description: "Check tracking status"},
		{Command: "list", Description: "List all tracked items"},
		{Command: "done", Description: "Stop tracking an item"},
	}

	_, err := t.bot.Request(tg.NewSetMyCommands(commands...))
	if err != nil {
		logger.ErrorLogger.Fatalf("Failed to set bot commands: %v", err)
	}
}

func (t *TelegramDegenBot) startUpdates() {
	u := tg.NewUpdate(0)
	u.Timeout = 60

	updates := t.bot.GetUpdatesChan(u)

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
			command.Handle(t.bot, update, userRequests, userEntries)
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
			_, err := t.bot.Send(msg)
			if err != nil {
				logger.ErrorLogger.Printf("Failed to send message: %v", err)
			}

			// Set a reminder
			go model.SetReminder(trackingInfo, chatID)
		}
	}
}
