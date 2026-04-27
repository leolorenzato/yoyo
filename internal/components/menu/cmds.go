package menu

import (
	"log"
	"yoyo/internal/execx"

	tea "charm.land/bubbletea/v2"
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
