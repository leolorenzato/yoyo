package footer

import (
	"yoyo/internal/layout"
)

func (m Model) View() string {
	rendered, err := m.render()
	if err != nil {
		return ""
	}

	return rendered
}

func (m Model) render() (string, error) {
	contentSize, err := layout.GetStyleContentSize(m.Style, m.AvailableSize)
	if err != nil {
		return "", err
	}

	return (m.Style.
		Width(contentSize.Width).
		Render(m.text)), nil
}
