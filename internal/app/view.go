package app

import (
	"fmt"
	"yoyo/internal/components/types"
	"yoyo/internal/layout"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	containerAvailableSize := m.termSize
	containerContentSize, err := layout.GetStyleContentSize(m.containerStyle, containerAvailableSize)
	if err != nil {
		return fmt.Sprintf("terminal size too small: %v", err)
	}
	containerContentAvailableSize, err := layout.GetStyleContentAvailableSize(m.containerStyle, containerAvailableSize)
	if err != nil {
		return fmt.Sprintf("terminal size too small: %v", err)
	}

	m.title.AvailableSize = containerContentAvailableSize
	renderedTitle := m.title.View()
	if renderedTitle == "" {
		return fmt.Sprintln("title rendering error")
	}

	var renderedSearch string

	m.footer.AvailableSize = containerContentAvailableSize
	renderedFooter := m.footer.View()
	if renderedFooter == "" {
		return fmt.Sprintln("footer rendering error")
	}

	var containerContentFreeHeight int
	if m.search != nil {
		m.search.AvailableSize = containerContentAvailableSize
		renderedSearch = m.search.View()
		if renderedSearch == "" {
			return fmt.Sprintln("search rendering error")
		}

		containerContentFreeHeight = containerContentAvailableSize.Height -
			lipgloss.Height(renderedTitle) -
			lipgloss.Height(renderedSearch) -
			lipgloss.Height(renderedFooter)
	} else {
		containerContentFreeHeight = containerContentAvailableSize.Height -
			lipgloss.Height(renderedTitle) -
			lipgloss.Height(renderedFooter)
	}
	containerContentFreeSize := types.Size{
		Width:  containerContentAvailableSize.Width,
		Height: containerContentFreeHeight,
	}

	m.menu.AvailableSize = containerContentFreeSize
	renderedMenu := m.menu.View()
	if renderedMenu == "" {
		return fmt.Sprintln("menu rendering error")
	}

	var joinedContent string
	if m.search != nil {
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

	renderedContainer := m.containerStyle.
		Width(containerContentSize.Width).
		Height(containerContentSize.Height).
		Render(joinedContent)

	return lipgloss.Place(
		containerAvailableSize.Width,
		containerAvailableSize.Height,
		lipgloss.Center,
		lipgloss.Center,
		renderedContainer,
	)
}
