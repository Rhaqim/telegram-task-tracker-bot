package main

import (
	"github.com/Rhaqim/trackdegens/config"
	"github.com/Rhaqim/trackdegens/internal/bot"
	"github.com/Rhaqim/trackdegens/pkg/logger"
)

func main() {

	logger.Init()

	cfg := config.NewConfig().LoadConfig()

	bot.Initialize(cfg.TelegramBotToken).Start()

}
