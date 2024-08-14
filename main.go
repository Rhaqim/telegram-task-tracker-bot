package main

import (
	"github.com/Rhaqim/trackdegens/config"
	"github.com/Rhaqim/trackdegens/internal/service"
	"github.com/Rhaqim/trackdegens/pkg/logger"
)

func main() {

	logger.Init()

	config.LoadConfig()

	service.Start()
}
