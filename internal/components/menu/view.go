package menu

import (
	"slices"
	"yoyo/internal/layout"

	"charm.land/lipgloss/v2"
)

func (m Model) View() (string, error) {
	rendered, err := m.render()
	if err != nil {
		return "", err
	}

	return rendered, nil
}

func (m Model) render() (string, error) {
	availableSize, err := m.getAvailableSize()
	if err != nil {
		return "", err
	}

	contentSize, err := layout.GetStyleContentSize(m.ContainerStyle, availableSize)
	if err != nil {
		return "", err
	}
	availableContentSize, err := layout.GetStyleContentAvailableSize(
		m.ContainerStyle,
		availableSize,
	)
	if err != nil {
		return "", err
	}

	var items []string

	// If cursor is visible, render from the top
	if m.cursor < availableContentSize.Height && m.cursor <= len(m.filteredItems)-1 {
		// The assumption is that an item has an height of 1
		items_num := min(availableContentSize.Height, len(m.filteredItems))
		for i := range items_num {
			cmd := m.filteredItems[i]
			item, err := m.renderMenuItem(cmd, i)
			if err != nil {
				return "", err
			}
			items = append(items, item)
		}
	}

	// If cursor is not visible, render from cursor backwards
	if m.cursor >= availableContentSize.Height && m.cursor <= len(m.filteredItems)-1 {
		// The assumption is that an item has an height of 1
		for i := m.cursor; i > m.cursor-availableContentSize.Height; i-- {
			cmd := m.filteredItems[i]
			item, err := m.renderMenuItem(cmd, i)
			if err != nil {
				return "", err
			}
			items = append(items, item)
		}
		slices.Reverse(items)
	}

	menuText := lipgloss.JoinVertical(lipgloss.Left, items...)

	return m.ContainerStyle.
		Width(contentSize.Width).
		Height(contentSize.Height).
		Render(menuText), nil
}

func (m Model) renderMenuItem(item Item, itemIndex int) (string, error) {
	availableSize, err := m.getAvailableSize()
	if err != nil {
		return "", err
	}

	availableContentWidth, err := layout.GetStyleContentAvailableWidth(
		m.ContainerStyle,
		availableSize.Width,
	)
	if err != nil {
		return "", err
	}
	text := item.Icon + " " + item.Name
	truncText := layout.Truncate(
		layout.StripNonSpaceWhitespace(text),
		availableContentWidth,
		"...",
	)
	if itemIndex == m.cursor {
		return m.SelectedItemStyle.Render(truncText), nil
	}
	return m.ItemStyle.Render(truncText), nil
}
