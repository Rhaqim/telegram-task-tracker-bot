package model

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/Rhaqim/trackdegens/pkg/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// "podman", "exec", "your_container_name", "pg_dumpall", "-U", "your_username", "-f", "/pg_backups/backup.sql" // OLD BACKUP COMMAND
func backupDatabase() {
	containerTool := "podman"
	containerName := "your_container_name"
	username := "your_username"
	backupDir := "/pg_backups"
	backupFile := "backup.sql"

	cmd := exec.Command(containerTool, "exec", containerName, "pg_dump", "-U", username, "-f", backupDir+"/"+backupFile)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error backing up database: %v\n", err)
	} else {
		fmt.Println("Database backup completed successfully.")
	}
}

// func backup() {
// 	for {
// 		backupDatabase()
// 		time.Sleep(48 * time.Hour)
// 	}
// }

func sendBackupFile(bot *tgbotapi.BotAPI, chatID int64, filePath string) {

	file, err := GetFile(filePath)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to get file: %v", err)
	}

	docConfig := tgbotapi.NewDocument(chatID, tgbotapi.FileBytes{
		Name:  "backup.sql",
		Bytes: file,
	})

	_, err = bot.Send(docConfig)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to send backup file: %v", err)
	}
}

func GetFile(filePath string) ([]byte, error) {

	file, err := os.ReadFile(filePath)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to read file: %v", err)
	}
	return file, err
}

func listBackups() string {
	cmd := exec.Command("podman", "exec", "your_container_name", "ls", "/pg_backups")
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error listing backups: %v\n", err)
	}
	backupFileNames := strings.Split(string(out), "\n")

	var backups string
	for _, backup := range backupFileNames {
		backups += backup + "\n"
	}
	return backups
}

func restoreBackup(backupFile string) {
	cmd := exec.Command("podman", "exec", "your_container_name", "psql", "-U", "your_username", "-f", "/pg_backups/"+backupFile)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error restoring database: %v\n", err)
	} else {
		fmt.Println("Database restored successfully.")
	}
}

func deleteBackup(backupFile string) {
	cmd := exec.Command("podman", "exec", "your_container_name", "rm", "/pg_backups/"+backupFile)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error deleting backup: %v\n", err)
	} else {
		fmt.Println("Backup deleted successfully.")
	}
}
