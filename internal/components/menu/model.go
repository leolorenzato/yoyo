package menu

import (
	"strings"
	"yoyo/internal/components/types"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
}

func NewModel(
	items []Item,
	containerStyle lipgloss.Style,
	itemStyle lipgloss.Style,
	selectedItemStyle lipgloss.Style,
) Model {
	return Model{
		items:             items,
		filteredItems:     items,
		ContainerStyle:    containerStyle,
		ItemStyle:         itemStyle,
		SelectedItemStyle: selectedItemStyle,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m *Model) filterItems(query string) {
	m.filteredItems = filterItems(m.items, query)
	m.clipCursor()
}

func (m *Model) incrementCursor() {
	if m.cursor < len(m.filteredItems)-1 {
		m.cursor++
	}
	m.clipCursor()
}

func (m *Model) decrementCursor() {
	if m.cursor > 0 {
		m.cursor--
	}
	m.clipCursor()
}

func (m *Model) clipCursor() {
	if m.cursor >= len(m.filteredItems) {
		if len(m.filteredItems) == 0 {
			m.cursor = 0
		} else {
			m.cursor = len(m.filteredItems) - 1
		}
	}
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
