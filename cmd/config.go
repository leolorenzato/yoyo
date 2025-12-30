package main

type GeneralConfig struct {
	Title string `toml:"title"`
}

type CmdsConfig struct {
	Name string `toml:"name"`
	Icon string `toml:"icon"`
	Cmd  string `toml:"cmd"`
}

type Config struct {
	General GeneralConfig `toml:"general"`
	Cmds    []CmdsConfig  `toml:"cmds"`
}

type ColorsStyle struct {
	Base00 string `toml:"base00"`
	Base01 string `toml:"base01"`
	Base02 string `toml:"base02"`
	Base03 string `toml:"base03"`
	Base04 string `toml:"base04"`
	Base05 string `toml:"base05"`
	Base06 string `toml:"base06"`
	Base07 string `toml:"base07"`
	base08 string `toml:"base08"`
	Base09 string `toml:"base09"`
	Base0A string `toml:"base0A"`
	Base0B string `toml:"base0B"`
	Base0C string `toml:"base0C"`
	Base0D string `toml:"base0D"`
	Base0E string `toml:"base0E"`
	Base0F string `toml:"base0F"`
}

type Style struct {
	Colors ColorsStyle `toml:"colors"`
}
