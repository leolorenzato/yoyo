package types

import tea "charm.land/bubbletea/v2"

type InternalModel interface {
	Init() tea.Cmd

	Update(tea.Msg) (InternalModel, tea.Cmd)

	View() (string, error)
}
