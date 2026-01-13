package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	tea "github.com/charmbracelet/bubbletea"
)

const appName string = "yoyo"

func main() {
	configFile := flag.String("c", "config_path", "path to config file")
	debug := flag.Bool("debug", false, "enable debug logging")
	flag.Parse()

	if *debug {
		f, err := tea.LogToFile("./log/debug.log", "debug")
		if err != nil {
			fmt.Println("failed to setup the logger:", err)
			os.Exit(1)
		}
		defer f.Close()
	} else {
		log.SetOutput(io.Discard)
	}

	var config Config
	if _, err := toml.DecodeFile(*configFile, &config); err != nil {
		log.Printf("failed to load config file %s", *configFile)
		os.Exit(1)
	}

	p := tea.NewProgram(
		NewModel(config),
		tea.WithAltScreen(),
	)
	if _, err := p.Run(); err != nil {
		log.Printf("Error: %v", err)
		os.Exit(1)
	}
}
