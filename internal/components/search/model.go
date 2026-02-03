package search

import (
	"yoyo/internal/components/types"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const defaultIcon string = "🔍"

type Model struct {
	Icon          string
	SearchText    string
	AvailableSize types.Size
	Style         lipgloss.Style
}

func NewModel(style lipgloss.Style) Model {
	return Model{
		Icon:  defaultIcon,
		Style: style,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
