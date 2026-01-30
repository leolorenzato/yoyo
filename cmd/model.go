package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	defaulHorizPadding int = 2
	defaulVertPadding  int = 0
)

type Size struct {
	Width  int
	Height int
}

type Model struct {
	title                 string
	cursor                int
	cmds                  []Cmd
	searchEnabled         bool
	search                string
	filteredCmds          []Cmd
	termSize              Size
	mainBoxStyle          lipgloss.Style
	titleStyle            lipgloss.Style
	searchStyle           lipgloss.Style
	menuStyle             lipgloss.Style
	selectedMenuTextStyle lipgloss.Style
	footerStyle           lipgloss.Style
}

type Cmd struct {
	name string
	icon string
	cmd  string
}

func NewModel(config Config) Model {
	mainBoxStyle := lipgloss.NewStyle().
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
		Padding(defaulVertPadding, defaulHorizPadding).
		Foreground(lipgloss.Color(config.UI.NormalText))

	selectedMenuTextStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(config.UI.SelectedText)).
		Bold(true)

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
		termSize:              Size{},
		filteredCmds:          cmds,
		mainBoxStyle:          mainBoxStyle,
		titleStyle:            titleStyle,
		searchStyle:           searchStyle,
		menuStyle:             menuStyle,
		selectedMenuTextStyle: selectedMenuTextStyle,
		footerStyle:           footerStyle,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.SetWindowTitle(appName)
}
