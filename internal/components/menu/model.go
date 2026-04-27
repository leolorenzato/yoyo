package menu

import (
	"fmt"
	"strings"
	"yoyo/internal/components/types"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type Item struct {
	Name string
	Icon string
	Cmd  string
}

type Model struct {
	cursor            int
	items             []Item
	filteredItems     []Item
	AvailableSize     types.Size
	ContainerStyle    lipgloss.Style
	ItemStyle         lipgloss.Style
	SelectedItemStyle lipgloss.Style
	dryRun            bool
}

func NewModel(
	items []Item,
	containerStyle lipgloss.Style,
	itemStyle lipgloss.Style,
	selectedItemStyle lipgloss.Style,
	dryRun bool,
) Model {
	return Model{
		items:             items,
		filteredItems:     items,
		ContainerStyle:    containerStyle,
		ItemStyle:         itemStyle,
		SelectedItemStyle: selectedItemStyle,
		dryRun:            dryRun,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) getAvailableSize() (types.Size, error) {
	if m.AvailableSize.Width <= 0 || m.AvailableSize.Height <= 0 {
		return types.Size{}, fmt.Errorf(
			"invalid available size, width: %d height %d",
			m.AvailableSize.Width,
			m.AvailableSize.Height,
		)
	}

	return m.AvailableSize, nil
}

func (m *Model) filterItems(query string) {
	m.filteredItems = filterItems(m.items, query)
	m.clipCursor()
}

func (m *Model) incrementCursor() {
	m.cursor++
	m.clipCursor()
}

func (m *Model) decrementCursor() {
	m.cursor--
	m.clipCursor()
}

func (m *Model) clipCursor() {
	n := len(m.filteredItems)
	if n == 0 {
		m.cursor = 0
		return
	}

	m.cursor = (m.cursor%n + n) % n
}

func (m *Model) getSelectedItem() Item {
	return m.filteredItems[m.cursor]
}

func filterItems(items []Item, query string) []Item {
	if query == "" {
		return items
	}

	var filtered []Item
	for _, item := range items {
		if strings.Contains(strings.ToLower(item.Name), strings.ToLower(query)) {
			filtered = append(filtered, item)
		}
	}

	return filtered
}
