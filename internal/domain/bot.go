package domain

type DegenerousTelegramBot interface {
	SendMessage(chatID int64, message string) error
	Start()
}
