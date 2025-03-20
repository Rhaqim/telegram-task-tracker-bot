package domain

type TelegramDegenBot interface {
	SendMessage(chatID int64, message string) error
	Start()
}
