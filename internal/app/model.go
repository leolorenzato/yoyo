package app

import (
	"yoyo/internal/components/footer"
	"yoyo/internal/components/menu"
	"yoyo/internal/components/search"
	"yoyo/internal/components/title"
	"yoyo/internal/components/types"
	"yoyo/internal/theme"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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

	cmds = append(cmds, tea.SetWindowTitle(m.appName))

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
