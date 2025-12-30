package main

type GeneralConfig struct {
	title string `toml:"title"`
}

type CmdsConfig struct {
	name string `toml:"name"`
	icon string `toml:"icon"`
	cmd  string `toml:"cmd"`
}

type Config struct {
	general GeneralConfig `toml:"general"`
	cmds    []CmdsConfig  `toml:"cmds"`
}

type ColorsStyle struct {
	base00 string `toml:"base00"`
	base01 string `toml:"base01"`
	base02 string `toml:"base02"`
	base03 string `toml:"base03"`
	base04 string `toml:"base04"`
	base05 string `toml:"base05"`
	base06 string `toml:"base06"`
	base07 string `toml:"base07"`
	base08 string `toml:"base08"`
	base09 string `toml:"base09"`
	base0A string `toml:"base0A"`
	base0B string `toml:"base0B"`
	base0C string `toml:"base0C"`
	base0D string `toml:"base0D"`
	base0E string `toml:"base0E"`
	base0F string `toml:"base0F"`
}

type Style struct {
	colors ColorsStyle `toml:"colors"`
}
