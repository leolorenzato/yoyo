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
