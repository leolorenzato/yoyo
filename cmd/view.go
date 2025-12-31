package main

import (
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/ansi"
	"github.com/charmbracelet/x/term"
)

const defaultWidth int = 80

func (m Model) View() string {
	physicalWidth, _, err := term.GetSize(os.Stdout.Fd())
	if err != nil {
		physicalWidth = defaultWidth
		log.Printf(
			"failed to get terminal size, use default size of %d px",
			physicalWidth,
		)

	}

	if physicalWidth < defaultWidth {
		return "terminal size too small"
	}

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(m.palette.Base0E).
		Align(lipgloss.Center).
		Padding(0, 2)

	searchStyle := lipgloss.NewStyle().
		Foreground(m.palette.Base05).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(m.palette.Base03).
		Padding(0, 2)

	menuStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(m.palette.Base03).
		Padding(0, 2)

	footerStyle := lipgloss.NewStyle().
		Foreground(m.palette.Base04).
		Align(lipgloss.Center).
		Padding(0, 2)

	borderStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(m.palette.Base0D).
		Padding(0, 2)

	borderHorizSize := borderStyle.GetBorderLeftSize() +
		borderStyle.GetBorderRightSize()
	borderHorizPadding := borderStyle.GetPaddingLeft() +
		borderStyle.GetPaddingRight()
	viewWidth := physicalWidth - borderHorizSize - borderHorizPadding

	titleWidth := viewWidth -
		titleStyle.GetBorderLeftSize() -
		titleStyle.GetBorderRightSize() -
		titleStyle.GetPaddingLeft() -
		titleStyle.GetPaddingRight()

	searchWidth := viewWidth -
		searchStyle.GetBorderLeftSize() -
		searchStyle.GetBorderRightSize() -
		searchStyle.GetPaddingLeft() -
		searchStyle.GetPaddingRight()

	menuWidth := viewWidth -
		menuStyle.GetBorderLeftSize() -
		menuStyle.GetBorderRightSize() -
		menuStyle.GetPaddingLeft() -
		menuStyle.GetPaddingRight()

	footerWidth := viewWidth -
		footerStyle.GetBorderLeftSize() -
		footerStyle.GetBorderRightSize() -
		footerStyle.GetPaddingLeft() -
		footerStyle.GetPaddingRight()

	titleText := m.title
	searchText := "🔍 " + m.search
	menuText := lipgloss.JoinVertical(lipgloss.Left, m.renderMenu(menuWidth)...)
	footerText := "• ↑/↓ to navigate • enter to select • ctrl+c to quit"

	title := titleStyle.Width(titleWidth).Render(titleText)
	search := searchStyle.Width(searchWidth).Render(searchText)
	menu := menuStyle.Width(menuWidth).Render(menuText)
	footer := footerStyle.Width(footerWidth).Render(footerText)

	var content string
	if m.searchEnabled {
		content = lipgloss.JoinVertical(
			lipgloss.Left,
			title,
			"",
			search,
			menu,
			"",
			footer,
		)
	} else {
		content = lipgloss.JoinVertical(
			lipgloss.Left,
			title,
			"",
			menu,
			"",
			footer,
		)
	}

	border := borderStyle.Width(viewWidth).Render(content)

	return lipgloss.PlaceHorizontal(
		physicalWidth,
		lipgloss.Center,
		border,
	) + "\n"
}

func (m Model) renderMenu(menuWidth int) []string {
	selectedStyle := m.selectedStyle
	normalStyle := m.normalStyle

	items := make([]string, len(m.filteredCmds))
	for i, choice := range m.filteredCmds {
		line := choice.icon + " " + choice.name
		line = ansi.Truncate(line, menuWidth, "...")
		if i == m.cursor {
			items[i] = selectedStyle.Render(line)
		} else {
			items[i] = normalStyle.Render(line)
		}
	}
	return items
}

func filterCmds(cmds []Cmd, query string) []Cmd {
	if query == "" {
		return cmds
	}

	var filtered []Cmd
	for _, cmd := range cmds {
		if strings.Contains(strings.ToLower(cmd.name), strings.ToLower(query)) {
			filtered = append(filtered, cmd)
		}
	}

	return filtered
}
