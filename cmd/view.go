package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/ansi"
)

const (
	minMenuHeight int = 4
)

func (m Model) View() string {
	mainBoxAvailableSize := m.termSize
	mainBoxContentSize, err := GetStyleContentSize(m.mainBoxStyle, mainBoxAvailableSize)
	if err != nil {
		return fmt.Sprintf("terminal size too small: %v", err)
	}
	mainBoxContentAvailableSize, err := GetStyleContentAvailableSize(m.mainBoxStyle, mainBoxAvailableSize)
	if err != nil {
		return fmt.Sprintf("terminal size too small: %v", err)
	}

	renderedTitle, err := m.renderTitle(mainBoxContentAvailableSize)
	if err != nil {
		return fmt.Sprintf("title rendering error: %v", err)
	}

	var renderedSearch string

	renderedFooter, err := m.renderFooter(mainBoxContentAvailableSize)
	if err != nil {
		return fmt.Sprintf("footer rendering error: %v", err)
	}

	var mainBoxContentFreeHeight int
	if m.searchEnabled {
		renderedSearch, err = m.renderSearch(mainBoxContentAvailableSize)
		if err != nil {
			return fmt.Sprintf("search rendering error: %v", err)
		}
		mainBoxContentFreeHeight = mainBoxContentAvailableSize.Height -
			lipgloss.Height(renderedTitle) -
			lipgloss.Height(renderedSearch) -
			lipgloss.Height(renderedFooter)
	} else {
		mainBoxContentFreeHeight = mainBoxContentAvailableSize.Height -
			lipgloss.Height(renderedTitle) -
			lipgloss.Height(renderedFooter)
	}
	mainBoxContentFreeSize := Size{
		Width:  mainBoxContentAvailableSize.Width,
		Height: mainBoxContentFreeHeight,
	}

	menuAvailableContentSize, err := GetStyleContentAvailableSize(m.menuStyle, mainBoxContentFreeSize)
	if err != nil {
		return fmt.Sprintf("terminal size too small: %v", err)
	}
	if menuAvailableContentSize.Height < minMenuHeight {
		err = fmt.Errorf("menu height too small")
		return fmt.Sprintf("terminal size too small: %v", err)
	}
	renderedMenu, err := m.renderMenu(mainBoxContentFreeSize)
	if err != nil {
		return fmt.Sprintf("menu rendering error: %v", err)
	}

	var joinedContent string
	if m.searchEnabled {
		joinedContent = lipgloss.JoinVertical(
			lipgloss.Center,
			renderedTitle,
			renderedSearch,
			renderedMenu,
			renderedFooter,
		)
	} else {
		joinedContent = lipgloss.JoinVertical(
			lipgloss.Center,
			renderedTitle,
			renderedMenu,
			renderedFooter,
		)
	}

	renderedBox := m.mainBoxStyle.
		Width(mainBoxContentSize.Width).
		Height(mainBoxContentSize.Height).
		Render(joinedContent)

	return lipgloss.Place(
		mainBoxAvailableSize.Width,
		mainBoxAvailableSize.Height,
		lipgloss.Center,
		lipgloss.Center,
		renderedBox,
	)
}

func (m Model) renderTitle(availableSize Size) (string, error) {
	contentSize, err := GetStyleContentSize(m.titleStyle, availableSize)
	if err != nil {
		return "", err
	}
	availableContentSize, err := GetStyleContentAvailableSize(m.titleStyle, availableSize)
	if err != nil {
		return "", err
	}
	truncText := truncate(stripNonSpaceWhitespace(m.title), availableContentSize.Width, "")

	return (m.titleStyle.
		Width(contentSize.Width).
		Render(truncText)), nil
}

func (m Model) renderSearch(availableSize Size) (string, error) {
	text := "🔍 " + m.search
	contentSize, err := GetStyleContentSize(m.searchStyle, availableSize)
	if err != nil {
		return "", err
	}
	availableContentSize, err := GetStyleContentAvailableSize(m.searchStyle, availableSize)
	if err != nil {
		return "", err
	}
	truncText := truncate(stripNonSpaceWhitespace(text), availableContentSize.Width, "...")

	return (m.searchStyle.
		Width(contentSize.Width).
		Render(truncText)), nil
}

func (m Model) renderMenu(availableSize Size) (string, error) {
	contentSize, err := GetStyleContentSize(m.menuStyle, availableSize)
	if err != nil {
		return "", err
	}
	availableContentSize, err := GetStyleContentAvailableSize(m.menuStyle, availableSize)
	if err != nil {
		return "", err
	}

	var items []string

	// If cursor is visible, render from the top
	if m.cursor < availableContentSize.Height && m.cursor <= len(m.filteredCmds)-1 {
		// The assumption is that an item has an height of 1
		items_num := min(availableContentSize.Height, len(m.filteredCmds))
		for i := range items_num {
			cmd := m.filteredCmds[i]
			item, err := m.renderMenuItem(cmd, i, availableSize.Width)
			if err != nil {
				return "", err
			}
			items = append(items, item)
		}
	}

	// If cursor is not visible, render from cursor backwards
	if m.cursor >= availableContentSize.Height && m.cursor <= len(m.filteredCmds)-1 {
		// The assumption is that an item has an height of 1
		for i := m.cursor; i > m.cursor-availableContentSize.Height; i-- {
			cmd := m.filteredCmds[i]
			item, err := m.renderMenuItem(cmd, i, availableSize.Width)
			if err != nil {
				return "", err
			}
			items = append(items, item)
		}
		slices.Reverse(items)
	}

	menuText := lipgloss.JoinVertical(lipgloss.Left, items...)

	return m.menuStyle.
		Width(contentSize.Width).
		Height(contentSize.Height).
		Render(menuText), nil
}

func (m Model) renderMenuItem(item Cmd, itemIndex int, availableWidth int) (string, error) {
	availableContentWidth, err := GetStyleContentAvailableWidth(m.menuStyle, availableWidth)
	if err != nil {
		return "", err
	}
	text := item.icon + " " + item.name
	truncText := truncate(stripNonSpaceWhitespace(text), availableContentWidth, "...")
	if itemIndex == m.cursor {
		return m.selectedMenuTextStyle.Render(truncText), nil
	}
	return truncText, nil
}

func (m Model) renderFooter(availableSize Size) (string, error) {
	text := "• ↑/↓ to navigate • enter to select • ctrl+c to quit"

	contentSize, err := GetStyleContentSize(m.footerStyle, availableSize)
	if err != nil {
		return "", err
	}

	return (m.footerStyle.
		Width(contentSize.Width).
		Render(text)), nil
}

func GetStyleContentSize(
	s lipgloss.Style,
	availableSize Size,
) (Size, error) {
	w, err := GetStyleContentWidth(s, availableSize.Width)
	if err != nil {
		return Size{}, err
	}

	h, err := GetStyleContentHeight(s, availableSize.Height)
	if err != nil {
		return Size{}, err
	}

	return Size{
		Width:  w,
		Height: h,
	}, nil
}

func GetStyleContentAvailableSize(
	s lipgloss.Style,
	availableSize Size,
) (Size, error) {
	w, err := GetStyleContentAvailableWidth(s, availableSize.Width)
	if err != nil {
		return Size{}, err
	}

	h, err := GetStyleContentAvailableHeight(s, availableSize.Height)
	if err != nil {
		return Size{}, err
	}

	return Size{
		Width:  w,
		Height: h,
	}, nil
}

func GetStyleContentAvailableWidth(
	s lipgloss.Style,
	availableWidth int,
) (int, error) {
	w, err := GetStyleContentWidth(s, availableWidth)
	if err != nil {
		return 0, err
	}
	aw := w -
		s.GetPaddingLeft() -
		s.GetPaddingRight()
	if aw < 0 {
		return 0, fmt.Errorf("invalid width %d", aw)
	}

	return aw, nil
}

func GetStyleContentWidth(
	s lipgloss.Style,
	availableWidth int,
) (int, error) {
	w := availableWidth -
		s.GetMarginLeft() -
		s.GetMarginRight() -
		s.GetBorderLeftSize() -
		s.GetBorderRightSize()
	if w < 0 {
		return 0, fmt.Errorf("invalid width %d", w)
	}

	return w, nil
}

func GetStyleContentAvailableHeight(
	s lipgloss.Style,
	availableHeight int,
) (int, error) {
	h, err := GetStyleContentHeight(s, availableHeight)
	if err != nil {
		return 0, err
	}
	ah := h -
		s.GetPaddingTop() -
		s.GetPaddingBottom()
	if ah < 0 {
		return 0, fmt.Errorf("invalid height %d", ah)
	}

	return ah, nil
}

func GetStyleContentHeight(
	s lipgloss.Style,
	availableHeight int,
) (int, error) {
	h := availableHeight -
		s.GetMarginTop() -
		s.GetMarginBottom() -
		s.GetBorderTopSize() -
		s.GetBorderBottomSize()
	if h < 0 {
		return 0, fmt.Errorf("invalid height %d", h)
	}

	return h, nil
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

func stripNonSpaceWhitespace(s string) string {
	return strings.Map(func(r rune) rune {
		if r == ' ' {
			return r
		}
		if r == '\n' || r == '\t' || r == '\r' || r == '\f' || r == '\v' {
			return -1
		}
		return r
	}, s)
}
