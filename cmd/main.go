package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"yoyo/internal/app"
	"yoyo/internal/components/menu"
	"yoyo/internal/theme"

	"github.com/BurntSushi/toml"
	tea "github.com/charmbracelet/bubbletea"
)

const appName string = "yoyo"

func main() {
	cfgFile := flag.String("c", "cfg_path", "path to configuration file")
	debug := flag.Bool("debug", false, "enable debug logging")
	flag.Parse()

	if *debug {
		logFilePath := "./log/debug.log"
		err := os.Truncate(logFilePath, 0)
		if err != nil && !os.IsNotExist(err) {
			fmt.Println("failed to truncate log file:", err)
			os.Exit(1)
		}

		f, err := tea.LogToFile(logFilePath, "")
		if err != nil {
			fmt.Println("failed to setup the logger:", err)
			os.Exit(1)
		}
		defer f.Close()
	} else {
		log.SetOutput(io.Discard)
	}

	var cfg Cfg
	if _, err := toml.DecodeFile(*cfgFile, &cfg); err != nil {
		log.Printf("failed to load configuration file %s", *cfgFile)
		os.Exit(1)
	}

	styles := theme.Build(cfg.Theme)
	model := app.NewModel(
		appName,
		getItemsFromCfg(cfg.Items),
		styles,
		cfg.App.Title,
		cfg.App.EnableSearch)
	p := tea.NewProgram(
		model,
		tea.WithAltScreen(),
	)
	if _, err := p.Run(); err != nil {
		log.Printf("Error: %v", err)
		os.Exit(1)
	}
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
