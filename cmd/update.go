package main

import (
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		return m, tea.ClearScreen
	case tea.KeyMsg:
		if m.searchEnabled {
			switch msg.Type {

			case tea.KeyBackspace:
				if len(m.search) > 0 {
					m.search = m.search[:len(m.search)-1]
				}

			case tea.KeyRunes:
				m.search += string(msg.Runes)
			}

			m.filteredCmds = filterCmds(m.cmds, m.search)
		}

		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.filteredCmds)-1 {
				m.cursor++
			}
		case "enter":
			cmd := m.filteredCmds[m.cursor].cmd
			return m, launch(cmd)
		}

		if m.cursor >= len(m.filteredCmds) {
			if len(m.filteredCmds) == 0 {
				m.cursor = 0
			} else {
				m.cursor = len(m.filteredCmds) - 1
			}
		}
	}

	return m, nil
}

func launch(cmd string) tea.Cmd {
	return func() tea.Msg {
		shell := os.Getenv("SHELL")
		if shell == "" {
			return nil
		}
		cmd_ := exec.Command(shell, "-lc", cmd)
		cmd_.Stdin = nil
		cmd_.Stdout = nil
		cmd_.Stderr = nil
		_ = cmd_.Start()

		return nil
	}
}
