package menu

import (
	"yoyo/internal/components/search"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case search.SearchChangeMsg:
		m.filterItems(msg.Query)
		return m, nil
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyUp:
			m.decrementCursor()
		case tea.KeyDown:
			m.incrementCursor()
		case tea.KeyEnter:
			item := m.getSelectedItem()
			return m, LaunchCmd(item.Cmd)
		}
	}

	return m, nil
}
