package main

import (
	"fmt"
	"os/exec"
	"time"
)

func backupDatabase() {
	cmd := exec.Command("podman", "exec", "your_container_name", "pg_dumpall", "-U", "your_username", "-f", "/pg_backups/backup.sql")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error backing up database: %v\n", err)
	} else {
		fmt.Println("Database backup completed successfully.")
	}
}

func backup() {
	for {
		backupDatabase()
		time.Sleep(48 * time.Hour)
	}
}
