package menu

import (
	"log"
	"yoyo/internal/components/search"
	"yoyo/internal/components/types"

	tea "charm.land/bubbletea/v2"
)

func (m Model) Update(msg tea.Msg) (types.InternalModel, tea.Cmd) {
	switch msg := msg.(type) {
	case search.SearchChangeMsg:
		m.filterItems(msg.Query)
		return m, nil
	case tea.KeyPressMsg:
		switch msg.String() {
		case "up":
			m.decrementCursor()
		case "down":
			m.incrementCursor()
		case "enter":
			item := m.getSelectedItem()
			if m.dryRun {
				log.Printf("dry run: command to launch: %s", item.Cmd)
				return m, nil
			}
			return m, LaunchCmd(item.Cmd)
		}
	}

	return m, nil
}
