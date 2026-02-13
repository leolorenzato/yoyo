package menu

import (
	"log"
	"yoyo/internal/execx"

	tea "github.com/charmbracelet/bubbletea"
)

func LaunchCmd(cmdText string) tea.Cmd {
	return func() tea.Msg {
		err := execx.Launch(cmdText)
		if err != nil {
			log.Printf("command launch error: %v", err)
		}

		return tea.Quit()
	}
}
