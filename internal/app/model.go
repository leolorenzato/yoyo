package app

import (
	"fmt"
	"yoyo/internal/components/footer"
	"yoyo/internal/components/menu"
	"yoyo/internal/components/search"
	"yoyo/internal/components/title"
	"yoyo/internal/components/types"
	"yoyo/internal/theme"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type Model struct {
	appName        string
	termSize       types.Size
	errorStyle     lipgloss.Style
	containerStyle lipgloss.Style
	title          title.Model
	search         *search.Model
	menu           menu.Model
	footer         footer.Model
}

func NewModel(
	appName string,
	items []menu.Item,
	styles theme.Styles,
	titletext string,
	enableSearch bool,
	dryRun bool,
) Model {
	m := Model{
		appName:        appName,
		errorStyle:     styles.Error,
		containerStyle: styles.Container,
		title:          title.NewModel(titletext, styles.Title),
		menu: menu.NewModel(
			items,
			styles.Menu.Container,
			styles.Menu.Item,
			styles.Menu.SelectedItem,
			dryRun,
		),
		footer: footer.NewModel(styles.Footer),
	}

	if enableSearch {
		searchModel := search.NewModel(styles.Search)
		m.search = &searchModel
	}

	return m
}

func (m Model) Init() tea.Cmd {
	var cmds []tea.Cmd

	if cmd := m.title.Init(); cmd != nil {
		cmds = append(cmds, cmd)
	}

	if cmd := m.menu.Init(); cmd != nil {
		cmds = append(cmds, cmd)
	}

	if cmd := m.footer.Init(); cmd != nil {
		cmds = append(cmds, cmd)
	}

	if m.search != nil {
		if cmd := m.search.Init(); cmd != nil {
			cmds = append(cmds, cmd)
		}
	}

	return tea.Batch(cmds...)
}

func (m Model) getAvailableSize() (types.Size, error) {
	if m.termSize.Width <= 0 || m.termSize.Height <= 0 {
		return types.Size{}, fmt.Errorf(
			"invalid available size, width: %d height %d",
			m.termSize.Width,
			m.termSize.Height,
		)
	}

	return m.termSize, nil
}
