package search

import (
	"yoyo/internal/components/types"

	tea "charm.land/bubbletea/v2"
)

func (m Model) Update(msg tea.Msg) (types.InternalModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch {
		case msg.Code == tea.KeyBackspace:
			if len(m.SearchText) > 0 {
				m.SearchText = m.SearchText[:len(m.SearchText)-1]
				return m, func() tea.Msg { return SearchChangeMsg{Query: m.SearchText} }
			}
		case msg.Text != "":
			m.SearchText += string(msg.Text)
			return m, func() tea.Msg { return SearchChangeMsg{Query: m.SearchText} }
		}
	}

	return m, nil
}
