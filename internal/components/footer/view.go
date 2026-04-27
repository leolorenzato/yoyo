package footer

import (
	"yoyo/internal/layout"
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

	contentSize, err := layout.GetStyleContentSize(m.Style, availableSize)
	if err != nil {
		return "", err
	}

	return (m.Style.
		Width(contentSize.Width).
		Render(m.text)), nil
}
