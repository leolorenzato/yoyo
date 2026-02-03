package search

import tea "github.com/charmbracelet/bubbletea"

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyBackspace:
			if len(m.SearchText) > 0 {
				m.SearchText = m.SearchText[:len(m.SearchText)-1]
				return m, func() tea.Msg { return SearchChangeMsg{Query: m.SearchText} }
			}
		case tea.KeyRunes:
			m.SearchText += string(msg.Runes)
			return m, func() tea.Msg { return SearchChangeMsg{Query: m.SearchText} }
		}
	}

	return m, nil
}
