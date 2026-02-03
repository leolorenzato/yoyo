package theme

import "github.com/charmbracelet/lipgloss"

const (
	defaulHorizPadding int = 2
	defaulVertPadding  int = 0
)

func Build(cfg Cfg) Styles {

	return Styles{
		Container: buildContainerStyle(cfg.Container),
		Title:     buildTitleStyle(cfg.Title),
		Search:    buildSearchStyle(cfg.Search),
		Menu:      buildMenuStyle(cfg.Menu),
		Footer:    buildFooterStyle(cfg.Footer),
	}
}

func buildContainerStyle(cfg ContainerCfg) lipgloss.Style {
	var border lipgloss.Border
	if cfg.BorderRounded {
		border = lipgloss.RoundedBorder()
	} else {
		border = lipgloss.NormalBorder()
	}
	style := lipgloss.NewStyle().
		Border(border).
		BorderForeground(lipgloss.Color(cfg.BorderColor)).
		Padding(defaulVertPadding, defaulHorizPadding)

	return style
}

func buildTitleStyle(cfg TitleCfg) lipgloss.Style {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color(cfg.TextColor)).
		Align(lipgloss.Center).
		Bold(true)

	return style
}

func buildSearchStyle(cfg SearchCfg) lipgloss.Style {
	var border lipgloss.Border
	if cfg.BorderRounded {
		border = lipgloss.RoundedBorder()
	} else {
		border = lipgloss.NormalBorder()
	}
	style := lipgloss.NewStyle().
		Border(border).
		BorderForeground(lipgloss.Color(cfg.BorderColor)).
		Foreground(lipgloss.Color(cfg.TextColor)).
		Padding(defaulVertPadding, defaulHorizPadding)

	return style
}

func buildMenuStyle(cfg MenuCfg) MenuStyles {
	var border lipgloss.Border
	if cfg.BorderRounded {
		border = lipgloss.RoundedBorder()
	} else {
		border = lipgloss.NormalBorder()
	}
	containerStyle := lipgloss.NewStyle().
		Border(border).
		BorderForeground(lipgloss.Color(cfg.BorderColor)).
		Padding(defaulVertPadding, defaulHorizPadding)

	itemStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(cfg.TextColor))

	selectedItemsStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(cfg.SelectedItemTextColor)).
		Bold(true)

	return MenuStyles{
		Container:    containerStyle,
		Item:         itemStyle,
		SelectedItem: selectedItemsStyle,
	}
}

func buildFooterStyle(cfg FooterCfg) lipgloss.Style {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color(cfg.TextColor)).
		Align(lipgloss.Center).
		Padding(defaulVertPadding, defaulHorizPadding)

	return style
}
