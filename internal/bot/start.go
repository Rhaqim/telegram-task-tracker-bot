package bot

import (
	"fmt"

	"github.com/Rhaqim/trackdegens/internal/domain"
	"github.com/Rhaqim/trackdegens/internal/model"
	"github.com/Rhaqim/trackdegens/pkg/logger"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type DegenerousTelegramBot struct {
	bot *tg.BotAPI
}

func Initialize(botToken string) domain.DegenerousTelegramBot {
	bot, err := tg.NewBotAPI(botToken)
	if err != nil {
		logger.ErrorLogger.Fatalf("Failed to create bot: %v", err)
	}

	logger.InfoLogger.Printf("Authorized on account %s", bot.Self.UserName)

	return &DegenerousTelegramBot{bot: bot}
}

func (t *DegenerousTelegramBot) Start() {
	t.loadCommands()
	t.startUpdates()
}

// SendMessage implements domain.DegenerousTelegramBot.
func (t *DegenerousTelegramBot) SendMessage(chatID int64, message string) error {
	panic("unimplemented")
}

func (t *DegenerousTelegramBot) loadCommands() {
	commands := []tg.BotCommand{
		{Command: model.Start.String(), Description: "Start interacting with the bot"},
		{Command: model.Track.String(), Description: "Track an item or event"},
		{Command: model.Status.String(), Description: "Check tracking status"},
		{Command: model.List.String(), Description: "List all tracked items"},
		{Command: model.Done.String(), Description: "Stop tracking an item"},
		{Command: model.Backups.String(), Description: "Manage database backups"},
		{Command: model.ListBackup.String(), Description: "List all database backups"},
		{Command: model.RestoreBackup.String(), Description: "Restore a database backup"},
		{Command: model.DeleteBackup.String(), Description: "Delete a database backup"},
		{Command: model.SendBackup.String(), Description: "Send the latest database backup"},
	}

	_, err := t.bot.Request(tg.NewSetMyCommands(commands...))
	if err != nil {
		logger.ErrorLogger.Fatalf("Failed to set bot commands: %v", err)
	}
}

func (t *DegenerousTelegramBot) startUpdates() {
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

		tgCommand := update.Message.Command()

		fmt.Println("Command:", tgCommand)

		// Handle commands
		command := model.Commands(tgCommand)

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
