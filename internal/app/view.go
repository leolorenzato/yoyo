package app

import (
	"fmt"
	"log"
	"yoyo/internal/components/types"
	"yoyo/internal/layout"
	"yoyo/internal/utils"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (m Model) View() tea.View {
	v := tea.NewView(m.view())
	v.AltScreen = true
	v.WindowTitle = m.appName

	return v
}

func (m Model) view() string {
	containerAvailableSize, err := m.getAvailableSize()
	if err != nil {
		log.Printf("failed to get available size: %v", err)
		return m.viewErr(fmt.Errorf("invalid available size"), containerAvailableSize)
	}

	containerContentSize, err := layout.GetStyleContentSize(
		m.containerStyle,
		containerAvailableSize,
	)
	if err != nil {
		log.Printf("get container content size error: %v", err)
		return m.viewErr(fmt.Errorf("terminal size too small"), containerAvailableSize)
	}

	content, err := m.getContent(containerAvailableSize)
	if err != nil {
		return m.viewErr(err, containerAvailableSize)
	}

	renderedContainer := m.containerStyle.
		Width(containerContentSize.Width).
		Height(containerContentSize.Height).
		Render(content)

	return lipgloss.Place(
		containerAvailableSize.Width,
		containerAvailableSize.Height,
		lipgloss.Center,
		lipgloss.Center,
		renderedContainer,
	)
}

func (m Model) getContent(containerAvailableSize types.Size) (string, error) {
	containerContentAvailableSize, err := layout.GetStyleContentAvailableSize(
		m.containerStyle,
		containerAvailableSize,
	)
	if err != nil {
		log.Printf("get container content available size error: %v", err)
		return "", fmt.Errorf("terminal size too small")
	}

	m.title.AvailableSize = containerContentAvailableSize
	renderedTitle, err := m.title.View()
	if err != nil {
		log.Printf("title rendering error: %v", err)
		return "", fmt.Errorf("title rendering error")
	}

	var renderedSearch string

	m.footer.AvailableSize = containerContentAvailableSize
	renderedFooter, err := m.footer.View()
	if err != nil {
		log.Printf("footer rendering error: %v", err)
		return "", fmt.Errorf("footer rendering error")
	}

	var containerContentFreeHeight int
	if m.search != nil {
		m.search.AvailableSize = containerContentAvailableSize
		renderedSearch, err = m.search.View()
		if err != nil {
			log.Printf("search rendering error: %v", err)
			return "", fmt.Errorf("search rendering error")
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
	renderedMenu, err := m.menu.View()
	if err != nil {
		log.Printf("menu rendering error: %v", err)
		return "", fmt.Errorf("menu rendering error")
	}

	if m.search != nil {
		return lipgloss.JoinVertical(
			lipgloss.Center,
			renderedTitle,
			renderedSearch,
			renderedMenu,
			renderedFooter,
		), nil
	}
	return lipgloss.JoinVertical(
		lipgloss.Center,
		renderedTitle,
		renderedMenu,
		renderedFooter,
	), nil
}

func (m Model) viewErr(err error, size types.Size) string {
	utils.AssertErr(err)
	renderedErr := m.errorStyle.
		Width(size.Width).
		Height(size.Height).
		Align(lipgloss.Center, lipgloss.Center).
		Render(err.Error())

	return lipgloss.Place(
		size.Width,
		size.Height,
		lipgloss.Center,
		lipgloss.Center,
		renderedErr,
	)
}
