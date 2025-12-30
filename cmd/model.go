package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	title         string
	cursor        int
	cmds          []Cmd
	palette       Base16Palette
	selectedStyle lipgloss.Style
	normalStyle   lipgloss.Style
}

type Cmd struct {
	name string
	icon string
	cmd  string
}

func NewModel(config Config, style Style) Model {
	palette := Base16Palette{
		lipgloss.Color(style.Colors.Base00),
		lipgloss.Color(style.Colors.Base01),
		lipgloss.Color(style.Colors.Base02),
		lipgloss.Color(style.Colors.Base03),
		lipgloss.Color(style.Colors.Base04),
		lipgloss.Color(style.Colors.Base05),
		lipgloss.Color(style.Colors.Base06),
		lipgloss.Color(style.Colors.Base07),
		lipgloss.Color(style.Colors.base08),
		lipgloss.Color(style.Colors.Base09),
		lipgloss.Color(style.Colors.Base0A),
		lipgloss.Color(style.Colors.Base0B),
		lipgloss.Color(style.Colors.Base0C),
		lipgloss.Color(style.Colors.Base0D),
		lipgloss.Color(style.Colors.Base0E),
		lipgloss.Color(style.Colors.Base0F),
	}

	selectedStyle := lipgloss.NewStyle().
		Foreground(palette.Base0B).
		Bold(true)

	normalStyle := lipgloss.NewStyle().
		Foreground(palette.Base05)

	var cmds []Cmd
	for _, cmd := range config.Cmds {
		cmds = append(cmds, Cmd{
			name: cmd.Name,
			icon: cmd.Icon,
			cmd:  cmd.Cmd,
		})
	}

	return Model{
		title:         config.General.Title,
		cursor:        0,
		cmds:          cmds,
		palette:       palette,
		selectedStyle: selectedStyle,
		normalStyle:   normalStyle,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.SetWindowTitle(appName)
}
