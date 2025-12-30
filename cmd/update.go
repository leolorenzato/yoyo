package main

import (
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.cmds)-1 {
				m.cursor++
			}
		case "enter":
			cmd := m.cmds[m.cursor].cmd
			return m, launch(cmd)
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
