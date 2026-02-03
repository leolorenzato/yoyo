package theme

type Cfg struct {
	Container ContainerCfg `toml:"container"`
	Title     TitleCfg     `toml:"title"`
	Search    SearchCfg    `toml:"search"`
	Menu      MenuCfg      `toml:"menu"`
	Footer    FooterCfg    `toml:"footer"`
}

type ContainerCfg struct {
	BorderColor   string `toml:"borderColor"`
	BorderRounded bool   `toml:"borderRounded"`
}

type TitleCfg struct {
	TextColor string `toml:"textColor"`
}

type SearchCfg struct {
	BorderColor   string `toml:"borderColor"`
	BorderRounded bool   `toml:"borderRounded"`
	TextColor     string `toml:"textColor"`
	HintIcon      string `toml:"hintIcon"`
}

type MenuCfg struct {
	BorderColor           string `toml:"borderColor"`
	BorderRounded         bool   `toml:"borderRounded"`
	TextColor             string `toml:"textColor"`
	SelectedItemTextColor string `toml:"selectedItemTextColor"`
}

type FooterCfg struct {
	TextColor string `toml:"textColor"`
}
