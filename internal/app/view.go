package app

import (
	"log"
	"yoyo/internal/components/types"
	"yoyo/internal/layout"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	containerAvailableSize := m.termSize
	containerContentSize, err := layout.GetStyleContentSize(
		m.containerStyle,
		containerAvailableSize,
	)
	if err != nil {
		log.Printf("get container content size error: %v", err)
		return m.viewErr("terminal size too small", containerAvailableSize)
	}
	containerContentAvailableSize, err := layout.GetStyleContentAvailableSize(
		m.containerStyle,
		containerAvailableSize,
	)
	if err != nil {
		log.Printf("get container content available size error: %v", err)
		return m.viewErr("terminal size too small", containerAvailableSize)
	}

	m.title.AvailableSize = containerContentAvailableSize
	renderedTitle := m.title.View()
	if renderedTitle == "" {
		return m.viewErr("title rendering error", containerAvailableSize)
	}

	var renderedSearch string

	m.footer.AvailableSize = containerContentAvailableSize
	renderedFooter := m.footer.View()
	if renderedFooter == "" {
		return m.viewErr("footer rendering error", containerAvailableSize)
	}

	var containerContentFreeHeight int
	if m.search != nil {
		m.search.AvailableSize = containerContentAvailableSize
		renderedSearch = m.search.View()
		if renderedSearch == "" {
			return m.viewErr("search rendering error", containerAvailableSize)
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
		return m.viewErr("menu rendering error", containerAvailableSize)
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

func (m Model) viewErr(err string, size types.Size) string {
	renderedErr := m.errorStyle.
		Width(size.Width).
		Render(err)

	return lipgloss.Place(
		size.Width,
		size.Height,
		lipgloss.Center,
		lipgloss.Center,
		renderedErr,
	)
}
