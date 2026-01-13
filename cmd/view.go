package main

import (
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/ansi"
	"github.com/charmbracelet/x/term"
)

const (
	minTermWidth  int = 80
	minTermHeight int = 60
)

func (m Model) View() string {
	termWidth, termHeight := getTermSizeOrMinSize()

	if termWidth < minTermWidth || termHeight < minTermHeight {
		return "terminal size too small"
	}

	log.Printf("term size %d, %d px", termWidth, termHeight)

	contentMaxWidth := getStyleMaxWidth(m.contentBorderStyle, termWidth)
	contentMaxHeight := getStyleMaxHeight(m.contentBorderStyle, termHeight)

	title := m.renderTitle(contentMaxWidth)
	search := m.renderSearch(contentMaxWidth)
	footer := m.renderFooter(contentMaxWidth)

	var menuAvailableHeight int
	if m.searchEnabled {
		menuAvailableHeight = contentMaxHeight -
			lipgloss.Height(title) -
			lipgloss.Height(search) -
			lipgloss.Height(footer)
	} else {
		menuAvailableHeight = contentMaxHeight -
			lipgloss.Height(title) -
			lipgloss.Height(footer)
	}

	menu := m.renderMenu(contentMaxWidth, menuAvailableHeight)

	var mainContent string
	if m.searchEnabled {
		mainContent = lipgloss.JoinVertical(
			lipgloss.Left,
			title,
			search,
			menu,
		)
	} else {
		mainContent = lipgloss.JoinVertical(
			lipgloss.Left,
			title,
			menu,
		)
	}

	contentText := lipgloss.Place(
		contentMaxWidth,
		contentMaxHeight,
		lipgloss.Left,
		lipgloss.Top,
		lipgloss.JoinVertical(
			lipgloss.Left,
			mainContent,
			lipgloss.Place(
				contentMaxWidth,
				contentMaxHeight-lipgloss.Height(mainContent),
				lipgloss.Left,
				lipgloss.Bottom,
				footer,
			),
		),
	)

	content := m.contentBorderStyle.
		Width(contentMaxWidth).
		Height(contentMaxHeight).
		Render(contentText)

	return lipgloss.PlaceHorizontal(
		termWidth,
		lipgloss.Center,
		content,
	)
}

func (m Model) renderTitle(contentWidth int) string {
	titleWidth := getStyleMaxWidth(m.titleStyle, contentWidth)
	titleText := truncate(m.title, titleWidth, "...")

	return m.titleStyle.Width(titleWidth).Render(titleText)
}

func (m Model) renderSearch(contentWidth int) string {
	searchWidth := getStyleMaxWidth(m.searchStyle, contentWidth)
	searchText := truncate("🔍 "+m.search, searchWidth, "...")

	return m.searchStyle.Width(searchWidth).Render(searchText)
}

func (m Model) renderMenu(maxContentWidth int, maxContentHeight int) string {
	menuMaxWidth := getStyleMaxWidth(m.menuStyle, maxContentWidth)
	menuMaxHeight := getStyleMaxHeight(m.menuStyle, maxContentHeight)

	startItemIndex := 0
	availableHeight := menuMaxHeight
	var items []string
	for i, cmd := range m.filteredCmds[startItemIndex:] {
		item := m.renderMenuItem(cmd, i, menuMaxWidth)
		itemHeight := lipgloss.Height(item)
		if availableHeight < itemHeight {
			break
		}
		items = append(items, item)
		availableHeight -= itemHeight
	}

	menuText := lipgloss.JoinVertical(lipgloss.Left, items...)
	menuText = truncate(menuText, menuMaxWidth, "...")

	return m.menuStyle.Width(menuMaxWidth).Height(menuMaxHeight).Render(menuText)
}

func (m Model) renderMenuItem(item Cmd, itemIndex int, maxWidth int) string {
	line := item.icon + " " + item.name
	line = truncate(line, maxWidth, "...")
	if itemIndex == m.cursor {
		return m.selectedMenuTextStyle.Render(line)
	}
	return m.normalTextMenuStyle.Render(line)
}

func (m Model) renderFooter(contentWidth int) string {
	footerWidth := getStyleMaxWidth(m.footerStyle, contentWidth)
	footerText := truncate(
		"• ↑/↓ to navigate • enter to select • ctrl+c to quit",
		footerWidth,
		"...",
	)

	return m.footerStyle.Width(footerWidth).Render(footerText)
}

func getTermSizeOrMinSize() (int, int) {
	termWidth, termHeight, err := term.GetSize(os.Stdout.Fd())
	if err != nil {
		termWidth = minTermWidth
		termHeight = minTermHeight
		log.Printf(
			"failed to get terminal size, use default size of %d, %d px",
			termWidth,
			termHeight,
		)
	}

	return termWidth, termHeight
}

func getStyleMaxWidth(s lipgloss.Style, outerWidth int) int {
	return outerWidth -
		s.GetBorderLeftSize() -
		s.GetBorderRightSize() -
		s.GetPaddingLeft() -
		s.GetPaddingRight()
}

func getStyleMaxHeight(s lipgloss.Style, outerHeight int) int {
	return outerHeight -
		s.GetBorderTopSize() -
		s.GetBorderBottomSize() -
		s.GetPaddingTop() -
		s.GetPaddingBottom()
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
