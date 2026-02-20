package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/debug"
	"yoyo/internal/app"
	"yoyo/internal/theme"

	"github.com/BurntSushi/toml"
	tea "github.com/charmbracelet/bubbletea"
)

const appName string = "yoyo"

func main() {
	version := flag.Bool("version", false, "print version information")
	cfgPath := flag.String("c", "", "path to configuration file (required)")
	debug := flag.Bool("debug", false, "debug mode")
	flag.Parse()

	if *version {
		printVersion()
		os.Exit(0)
	}

	requireFlag(cfgPath, "c")

	if *debug {
		f, err := getLogFile()
		if err != nil {
			fmt.Println("failed to setup the logger: ", err)
			os.Exit(1)
		}
		defer f.Close()
	} else {
		log.SetOutput(io.Discard)
	}

	var cfg Cfg
	if _, err := toml.DecodeFile(*cfgPath, &cfg); err != nil {
		log.Printf("failed to load configuration file %s", *cfgPath)
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
		log.Printf("error: %v", err)
		os.Exit(1)
	}
}

func requireFlag(value *string, name string) {
	if *value == "" {
		fmt.Printf("%s is required\n", name)
		os.Exit(1)
	}
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
