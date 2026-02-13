package main

import (
	"fmt"
	"os"
	"path/filepath"
	"yoyo/internal/components/menu"

	tea "github.com/charmbracelet/bubbletea"
)

func getLogFile() (*os.File, error) {
	logDir := filepath.Join(os.Getenv("HOME"), ".config", appName, "log")
	os.MkdirAll(logDir, os.ModePerm)
	logFile := filepath.Join(logDir, "debug.log")
	err := os.Truncate(logFile, 0)
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("failed to truncate log file:", err)
		os.Exit(1)
	}

	return tea.LogToFile(logFile, "")
}

func getItemsFromCfg(cfg []ItemCfg) []menu.Item {
	var items []menu.Item
	for _, item := range cfg {
		items = append(items, menu.Item{
			Name: item.Name,
			Icon: item.Icon,
			Cmd:  item.Cmd,
		})
	}

	return items
}
