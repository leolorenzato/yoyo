package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"yoyo/internal/app"
	"yoyo/internal/theme"

	tea "charm.land/bubbletea/v2"
	"github.com/BurntSushi/toml"
)

const appName string = "yoyo"

func main() {
	version := flag.Bool("version", false, "print version information")
	cfgPath := flag.String("c", "", "path to configuration file (required)")
	dryRun := flag.Bool("dry-run", false, "dry run mode")
	debug := flag.Bool("debug", false, "debug mode")
	flag.Parse()

	if *version {
		printVersion()
		os.Exit(0)
	}

	requireStringFlag(cfgPath, "c")

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
		cfg.App.EnableSearch,
		*dryRun,
	)
	p := tea.NewProgram(
		model,
	)
	if _, err := p.Run(); err != nil {
		log.Printf("error: %v", err)
		os.Exit(1)
	}
}

func requireStringFlag(value *string, name string) {
	if *value == "" {
		fmt.Fprintf(os.Stderr, "%s is required\n", name)
		os.Exit(1)
	}
}
