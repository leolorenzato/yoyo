package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	tea "github.com/charmbracelet/bubbletea"
)

const appName string = "yoyo"

func main() {
	f, err := tea.LogToFile("./log/debug.log", "debug")
	if err != nil {
		fmt.Println("failed to setup the logger:", err)
		os.Exit(1)
	}
	defer f.Close()

	configFile := flag.String("config", "config_path", "path to config file")
	flag.Parse()

	var config Config
	if _, err := toml.DecodeFile(*configFile, &config); err != nil {
		log.Printf("failed to load config file %s", *configFile)
		os.Exit(1)
	}

	styleFile := flag.String("style", "style_path", "path to style file")
	flag.Parse()

	var style Style
	if _, err := toml.DecodeFile(*styleFile, &style); err != nil {
		log.Printf("failed to load style file %s", *styleFile)
		os.Exit(1)
	}

	p := tea.NewProgram(NewModel(config, style))
	if _, err := p.Run(); err != nil {
		log.Printf("Error: %v", err)
		os.Exit(1)
	}
}
