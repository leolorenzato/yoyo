package title

import (
	"yoyo/internal/components/types"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	text          string
	AvailableSize types.Size
	Style         lipgloss.Style
}

func NewModel(text string, style lipgloss.Style) Model {
	return Model{
		text:  text,
		Style: style,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
