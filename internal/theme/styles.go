package theme

import "github.com/charmbracelet/lipgloss"

type MenuStyles struct {
	Container    lipgloss.Style
	Item         lipgloss.Style
	SelectedItem lipgloss.Style
}

type Styles struct {
	Container lipgloss.Style
	Title     lipgloss.Style
	Search    lipgloss.Style
	Menu      MenuStyles
	Footer    lipgloss.Style
}
