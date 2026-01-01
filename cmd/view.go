package main

import (
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/ansi"
	"github.com/charmbracelet/x/term"
)

const minTermWidth int = 80

func (m Model) View() string {
	termWidth := getTermWidth()

	if termWidth < minTermWidth {
		return "terminal size too small"
	}

	contentWidth := getStyleWidth(m.contentBorderStyle, termWidth)

	title := m.renderTitle(contentWidth)
	search := m.renderSearch(contentWidth)
	menu := m.renderMenu(contentWidth)
	footer := m.renderFooter(contentWidth)

	var contentText string
	if m.searchEnabled {
		contentText = lipgloss.JoinVertical(
			lipgloss.Left,
			title,
			"",
			search,
			menu,
			"",
			footer,
		)
	} else {
		contentText = lipgloss.JoinVertical(
			lipgloss.Left,
			title,
			"",
			menu,
			"",
			footer,
		)
	}

	content := m.contentBorderStyle.Width(contentWidth).Render(contentText)

	return lipgloss.PlaceHorizontal(
		termWidth,
		lipgloss.Center,
		content,
	) + "\n"
}

func (m Model) renderTitle(contentWidth int) string {
	titleWidth := getStyleWidth(m.titleStyle, contentWidth)
	titleText := truncate(m.title, titleWidth, "...")

	return m.titleStyle.Width(titleWidth).Render(titleText)
}

func (m Model) renderSearch(contentWidth int) string {
	searchWidth := getStyleWidth(m.searchStyle, contentWidth)
	searchText := truncate("🔍 "+m.search, searchWidth, "...")

	return m.searchStyle.Width(searchWidth).Render(searchText)
}

func (m Model) renderMenu(contentWidth int) string {
	menuWidth := getStyleWidth(m.menuStyle, contentWidth)
	menuText := lipgloss.JoinVertical(lipgloss.Left, m.renderMenuItems(menuWidth)...)
	menuText = truncate(menuText, menuWidth, "...")

	return m.menuStyle.Width(menuWidth).Render(menuText)
}

func (m Model) renderMenuItems(menuWidth int) []string {
	selectedStyle := m.selectedMenuTextStyle
	normalStyle := m.normalTextMenuStyle

	items := make([]string, len(m.filteredCmds))
	for i, choice := range m.filteredCmds {
		line := choice.icon + " " + choice.name
		line = truncate(line, menuWidth, "...")
		if i == m.cursor {
			items[i] = selectedStyle.Render(line)
		} else {
			items[i] = normalStyle.Render(line)
		}
	}
	return items
}

func (m Model) renderFooter(contentWidth int) string {
	footerWidth := getStyleWidth(m.footerStyle, contentWidth)
	footerText := truncate(
		"• ↑/↓ to navigate • enter to select • ctrl+c to quit",
		footerWidth,
		"...",
	)

	return m.footerStyle.Width(footerWidth).Render(footerText)
}

func getTermWidth() int {
	termWidth, _, err := term.GetSize(os.Stdout.Fd())
	if err != nil {
		termWidth = minTermWidth
		log.Printf(
			"failed to get terminal size, use default size of %d px",
			termWidth,
		)
	}

	return termWidth
}

func getStyleWidth(s lipgloss.Style, outerWidth int) int {
	return outerWidth -
		s.GetBorderLeftSize() -
		s.GetBorderRightSize() -
		s.GetPaddingLeft() -
		s.GetPaddingRight()
}

func truncate(s string, length int, tail string) string {
	return ansi.Truncate(s, length, tail)
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
