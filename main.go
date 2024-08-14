package main

import (
	"github.com/Rhaqim/trackdegens/config"
	"github.com/Rhaqim/trackdegens/internal/service"
)

func main() {

	config.LoadConfig()

	service.Start()
}
