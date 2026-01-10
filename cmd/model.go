package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	defaulHorizPadding int = 2
	defaulVertPadding  int = 0
)

type Model struct {
	title                 string
	cursor                int
	cmds                  []Cmd
	searchEnabled         bool
	search                string
	filteredCmds          []Cmd
	contentBorderStyle    lipgloss.Style
	titleStyle            lipgloss.Style
	searchStyle           lipgloss.Style
	menuStyle             lipgloss.Style
	selectedMenuTextStyle lipgloss.Style
	normalTextMenuStyle   lipgloss.Style
	footerStyle           lipgloss.Style
}

type Cmd struct {
	name string
	icon string
	cmd  string
}

func NewModel(config Config) Model {
	contentBorderStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(config.UI.ContentBorder)).
		Padding(defaulVertPadding, defaulHorizPadding)

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(config.UI.Title)).
		Align(lipgloss.Center).
		Padding(defaulVertPadding, defaulHorizPadding)

	searchStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(config.UI.SearchText)).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(config.UI.SearchBorder)).
		Padding(defaulVertPadding, defaulHorizPadding)

	menuStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(config.UI.MenuBorder)).
		Padding(defaulVertPadding, defaulHorizPadding)

	selectedMenuTextStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(config.UI.SelectedText)).
		Bold(true)

	normalTextMenuStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(config.UI.NormalText))

	footerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(config.UI.Footer)).
		Align(lipgloss.Center).
		Padding(defaulVertPadding, defaulHorizPadding)

	var cmds []Cmd
	for _, cmd := range config.Cmds {
		cmds = append(cmds, Cmd{
			name: cmd.Name,
			icon: cmd.Icon,
			cmd:  cmd.Cmd,
		})
	}

	return Model{
		title:                 config.General.Title,
		cursor:                0,
		cmds:                  cmds,
		searchEnabled:         true,
		search:                "",
		filteredCmds:          cmds,
		contentBorderStyle:    contentBorderStyle,
		titleStyle:            titleStyle,
		searchStyle:           searchStyle,
		menuStyle:             menuStyle,
		selectedMenuTextStyle: selectedMenuTextStyle,
		normalTextMenuStyle:   normalTextMenuStyle,
		footerStyle:           footerStyle,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.SetWindowTitle(appName)
}
