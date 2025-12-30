package main

import (
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
)

func (m Model) View() string {
	physicalWidth, _, err := term.GetSize(os.Stdout.Fd())
	if err != nil {
		physicalWidth = 80
	}

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(m.palette.Base0E).
		Align(lipgloss.Center)

	menuStyle := lipgloss.NewStyle()

	footerStyle := lipgloss.NewStyle().
		Foreground(m.palette.Base04).
		Align(lipgloss.Center)

	borderStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(m.palette.Base0D).
		Padding(0, 2)

	titleText := m.title
	menuText := lipgloss.JoinVertical(lipgloss.Left, m.renderMenu()...)
	footerText := "• ↑/↓ to navigate • enter to select • ctrl+c to quit"

	contentWidth := maxContentWidth(titleText, menuText, footerText)
	borderHorizSize := borderStyle.GetBorderLeftSize() + borderStyle.GetBorderRightSize()
	borderHorizPadding := borderStyle.GetPaddingLeft() + borderStyle.GetPaddingRight()
	viewWidth := max(
		physicalWidth-borderHorizSize-borderHorizPadding,
		contentWidth,
	)

	title := titleStyle.Width(viewWidth).Render(titleText)
	menu := menuStyle.Width(viewWidth).Render(menuText)
	footer := footerStyle.Width(viewWidth).Render(footerText)

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		menu,
		"",
		footer,
	)

	border := borderStyle.Width(viewWidth).Render(content)

	return lipgloss.PlaceHorizontal(
		physicalWidth,
		lipgloss.Center,
		border,
	) + "\n"
}

func maxContentWidth(strs ...string) int {
	max := 0
	for _, s := range strs {
		if w := lipgloss.Width(s); w > max {
			max = w
		}
	}
	return max
}

func (m Model) renderMenu() []string {
	items := make([]string, len(m.cmds))
	for i, choice := range m.cmds {
		line := choice.icon + " " + choice.name
		if i == m.cursor {
			items[i] = m.selectedStyle.Render(line)
		} else {
			items[i] = m.normalStyle.Render(line)
		}
	}
	return items
}
