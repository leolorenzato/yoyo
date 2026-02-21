package title

import (
	"log"
	"yoyo/internal/layout"
)

func (m Model) View() string {
	rendered, err := m.render()
	if err != nil {
		log.Printf("title render error: %v", err)
		return ""
	}

	return rendered
}

func (m Model) render() (string, error) {
	contentSize, err := layout.GetStyleContentSize(m.Style, m.AvailableSize)
	if err != nil {
		return "", err
	}
	availableContentSize, err := layout.GetStyleContentAvailableSize(m.Style, m.AvailableSize)
	if err != nil {
		return "", err
	}
	truncText := layout.Truncate(layout.StripNonSpaceWhitespace(m.text), availableContentSize.Width, "")

	return (m.Style.
		Width(contentSize.Width).
		Render(truncText)), nil
}
