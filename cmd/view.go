package main

import (
	"log"
	"os"
	"slices"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/ansi"
	"github.com/charmbracelet/x/term"
)

const (
	minTermWidth  int = 80
	minTermHeight int = 24
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
			lipgloss.Center,
			title,
			search,
			menu,
		)
	} else {
		mainContent = lipgloss.JoinVertical(
			lipgloss.Center,
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

	return lipgloss.Place(
		termWidth,
		termHeight,
		lipgloss.Center,
		lipgloss.Center,
		content,
	)
}

func (m Model) renderTitle(contentWidth int) string {
	titleWidth := getStyleMaxWidth(m.titleStyle, contentWidth)
	titleText := m.title

	return m.titleStyle.Width(titleWidth).Render(titleText)
}

func (m Model) renderSearch(contentWidth int) string {
	searchWidth := getStyleMaxWidth(m.searchStyle, contentWidth)
	searchText := "🔍 " + m.search

	return m.searchStyle.Width(searchWidth).Render(searchText)
}

func (m Model) renderMenu(maxContentWidth int, maxContentHeight int) string {
	menuMaxWidth := getStyleMaxWidth(m.menuStyle, maxContentWidth)
	menuMaxHeight := getStyleMaxHeight(m.menuStyle, maxContentHeight)

	// Try to render from the top
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

	// If cursor is not visible, render from cursor backwards
	if m.cursor > len(items)-1 && m.cursor <= len(m.filteredCmds)-1 {
		lastItemIndex := m.cursor
		availableHeight = menuMaxHeight
		items = []string{}
		for i := lastItemIndex; i >= 0; i-- {
			cmd := m.filteredCmds[i]
			item := m.renderMenuItem(cmd, i, menuMaxWidth)
			itemHeight := lipgloss.Height(item)
			if availableHeight < itemHeight {
				break
			}
			items = append(items, item)
			availableHeight -= itemHeight
		}
		slices.Reverse(items)
	}

	menuText := lipgloss.JoinVertical(lipgloss.Left, items...)

	return m.menuStyle.Width(menuMaxWidth).Height(menuMaxHeight).Render(menuText)
}

func (m Model) renderMenuItem(item Cmd, itemIndex int, maxWidth int) string {
	line := item.icon + " " + item.name
	if itemIndex == m.cursor {
		return m.selectedMenuTextStyle.Render(line)
	}
	return m.normalTextMenuStyle.Render(line)
}

func (m Model) renderFooter(contentWidth int) string {
	footerWidth := getStyleMaxWidth(m.footerStyle, contentWidth)
	footerText := "• ↑/↓ to navigate • enter to select • ctrl+c to quit"

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

func truncate(s string, length int, tail string) string {
	return ansi.Truncate(s, length, tail)
}
