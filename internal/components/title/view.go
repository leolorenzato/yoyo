package title

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
	availableContentSize, err := layout.GetStyleContentAvailableSize(m.Style, availableSize)
	if err != nil {
		return "", err
	}
	truncText := layout.Truncate(
		layout.StripNonSpaceWhitespace(m.text),
		availableContentSize.Width,
		"",
	)

	return (m.Style.
		Width(contentSize.Width).
		Render(truncText)), nil
}
