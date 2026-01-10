package main

type GeneralConfig struct {
	Title string `toml:"title"`
}

type CmdsConfig struct {
	Name string `toml:"name"`
	Icon string `toml:"icon"`
	Cmd  string `toml:"cmd"`
}

type UIConfig struct {
	ContentBorder string `toml:"contentBorder"`
	Title         string `toml:"title"`
	SearchBorder  string `toml:"searchBorder"`
	SearchText    string `toml:"searchText"`
	MenuBorder    string `toml:"menuBorder"`
	SelectedText  string `toml:"selectedText"`
	NormalText    string `toml:"normalText"`
	Footer        string `toml:"footer"`
}

type Config struct {
	General GeneralConfig `toml:"general"`
	Cmds    []CmdsConfig  `toml:"cmds"`
	UI      UIConfig      `toml:"ui"`
}
