package app

import (
	"yoyo/internal/components/footer"
	"yoyo/internal/components/menu"
	"yoyo/internal/components/search"
	"yoyo/internal/components/title"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.termSize.Width = msg.Width
		m.termSize.Height = msg.Height
		cmds = append(cmds, tea.ClearScreen)
	}

	updated, cmd := m.title.Update(msg)
	m.title = updated.(title.Model)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}

	updated, cmd = m.menu.Update(msg)
	m.menu = updated.(menu.Model)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}

	updated, cmd = m.footer.Update(msg)
	m.footer = updated.(footer.Model)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}

	if m.search != nil {
		updated, cmd = m.search.Update(msg)
		updatedSearch := updated.(search.Model)
		m.search = &updatedSearch
		if cmd != nil {
			cmds = append(cmds, cmd)
		}
	}

	return m, tea.Batch(cmds...)
}
