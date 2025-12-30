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
		lipgloss.Color(style.colors.base00),
		lipgloss.Color(style.colors.base01),
		lipgloss.Color(style.colors.base02),
		lipgloss.Color(style.colors.base03),
		lipgloss.Color(style.colors.base04),
		lipgloss.Color(style.colors.base05),
		lipgloss.Color(style.colors.base06),
		lipgloss.Color(style.colors.base07),
		lipgloss.Color(style.colors.base08),
		lipgloss.Color(style.colors.base09),
		lipgloss.Color(style.colors.base0A),
		lipgloss.Color(style.colors.base0B),
		lipgloss.Color(style.colors.base0C),
		lipgloss.Color(style.colors.base0D),
		lipgloss.Color(style.colors.base0E),
		lipgloss.Color(style.colors.base0F),
	}

	selectedStyle := lipgloss.NewStyle().
		Foreground(palette.Base0B).
		Bold(true)

	normalStyle := lipgloss.NewStyle().
		Foreground(palette.Base05)

	var cmds []Cmd
	for _, cmd := range config.cmds {
		cmds = append(cmds, Cmd{
			name: cmd.name,
			icon: cmd.icon,
			cmd:  cmd.cmd,
		})
	}

	return Model{
		title:         config.general.title,
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
