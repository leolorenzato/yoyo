package theme

type Cfg struct {
	Container ContainerCfg `toml:"container"`
	Title     TitleCfg     `toml:"title"`
	Search    SearchCfg    `toml:"search"`
	Menu      MenuCfg      `toml:"menu"`
	Footer    FooterCfg    `toml:"footer"`
}

type ContainerCfg struct {
	Border        bool   `toml:"border"`
	BorderColor   string `toml:"borderColor"`
	BorderRounded bool   `toml:"borderRounded"`
}

type TitleCfg struct {
	Border        bool   `toml:"border"`
	BorderColor   string `toml:"borderColor"`
	BorderRounded bool   `toml:"borderRounded"`
	TextColor     string `toml:"textColor"`
}

type SearchCfg struct {
	Border        bool   `toml:"border"`
	BorderColor   string `toml:"borderColor"`
	BorderRounded bool   `toml:"borderRounded"`
	TextColor     string `toml:"textColor"`
	HintIcon      string `toml:"hintIcon"`
}

type MenuCfg struct {
	Border                bool   `toml:"border"`
	BorderColor           string `toml:"borderColor"`
	BorderRounded         bool   `toml:"borderRounded"`
	TextColor             string `toml:"textColor"`
	SelectedItemTextColor string `toml:"selectedItemTextColor"`
}

type FooterCfg struct {
	Border        bool   `toml:"border"`
	BorderColor   string `toml:"borderColor"`
	BorderRounded bool   `toml:"borderRounded"`
	TextColor     string `toml:"textColor"`
}
