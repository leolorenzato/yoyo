package main

import "yoyo/internal/theme"

type AppCfg struct {
	Title        string `toml:"title"`
	EnableSearch bool   `toml:"enableSearch"`
}

type ItemCfg struct {
	Name string `toml:"name"`
	Icon string `toml:"icon"`
	Cmd  string `toml:"cmd"`
}

type Cfg struct {
	App   AppCfg    `toml:"app"`
	Items []ItemCfg `toml:"items"`
	Theme theme.Cfg `toml:"theme"`
}
