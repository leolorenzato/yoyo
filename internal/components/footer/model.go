package footer

import (
	"yoyo/internal/components/types"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const defaultText string = "• ↑/↓ to navigate • enter to select • ctrl+c to quit"

type Model struct {
	text          string
	AvailableSize types.Size
	Style         lipgloss.Style
}

func NewModel(style lipgloss.Style) Model {
	return Model{
		text:  defaultText,
		Style: style,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
