package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"
	"yoyo/internal/components/menu"

	tea "charm.land/bubbletea/v2"
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

func printVersion() {
	version := "unknown"
	commit := "unknown"
	buildTime := "unknown"
	goVersion := "unknown"
	if info, ok := debug.ReadBuildInfo(); ok {
		version = info.Main.Version
		goVersion = info.GoVersion
		for _, setting := range info.Settings {
			switch setting.Key {
			case "vcs.revision":
				commit = setting.Value
			case "vcs.time":
				buildTime = setting.Value
			}
		}
	}
	fmt.Printf("Version: %s\n", version)
	fmt.Printf("Commit: %s\n", commit)
	fmt.Printf("Build time: %s\n", buildTime)
	fmt.Printf("Go version: %s\n", goVersion)
}
