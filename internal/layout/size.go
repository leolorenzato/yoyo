package layout

import (
	"fmt"
	"strings"
	"yoyo/internal/components/types"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/ansi"
)

func GetStyleContentSize(
	s lipgloss.Style,
	availableSize types.Size,
) (types.Size, error) {
	w, err := GetStyleContentWidth(s, availableSize.Width)
	if err != nil {
		return types.Size{}, err
	}

	h, err := GetStyleContentHeight(s, availableSize.Height)
	if err != nil {
		return types.Size{}, err
	}

	return types.Size{
		Width:  w,
		Height: h,
	}, nil
}

func GetStyleContentAvailableSize(
	s lipgloss.Style,
	availableSize types.Size,
) (types.Size, error) {
	w, err := GetStyleContentAvailableWidth(s, availableSize.Width)
	if err != nil {
		return types.Size{}, err
	}

	h, err := GetStyleContentAvailableHeight(s, availableSize.Height)
	if err != nil {
		return types.Size{}, err
	}

	return types.Size{
		Width:  w,
		Height: h,
	}, nil
}

func GetStyleContentAvailableWidth(
	s lipgloss.Style,
	availableWidth int,
) (int, error) {
	w, err := GetStyleContentWidth(s, availableWidth)
	if err != nil {
		return 0, err
	}
	aw := w -
		s.GetPaddingLeft() -
		s.GetPaddingRight()
	if aw < 0 {
		return 0, fmt.Errorf("invalid width %d", aw)
	}

	return aw, nil
}

func GetStyleContentWidth(
	s lipgloss.Style,
	availableWidth int,
) (int, error) {
	w := availableWidth -
		s.GetMarginLeft() -
		s.GetMarginRight() -
		s.GetBorderLeftSize() -
		s.GetBorderRightSize()
	if w < 0 {
		return 0, fmt.Errorf("invalid width %d", w)
	}

	return w, nil
}

func GetStyleContentAvailableHeight(
	s lipgloss.Style,
	availableHeight int,
) (int, error) {
	h, err := GetStyleContentHeight(s, availableHeight)
	if err != nil {
		return 0, err
	}
	ah := h -
		s.GetPaddingTop() -
		s.GetPaddingBottom()
	if ah < 0 {
		return 0, fmt.Errorf("invalid height %d", ah)
	}

	return ah, nil
}

func GetStyleContentHeight(
	s lipgloss.Style,
	availableHeight int,
) (int, error) {
	h := availableHeight -
		s.GetMarginTop() -
		s.GetMarginBottom() -
		s.GetBorderTopSize() -
		s.GetBorderBottomSize()
	if h < 0 {
		return 0, fmt.Errorf("invalid height %d", h)
	}

	return h, nil
}

func Truncate(s string, length int, tail string) string {
	return ansi.Truncate(s, length, tail)
}

func StripNonSpaceWhitespace(s string) string {
	return strings.Map(func(r rune) rune {
		if r == ' ' {
			return r
		}
		if r == '\n' || r == '\t' || r == '\r' || r == '\f' || r == '\v' {
			return -1
		}
		return r
	}, s)
}
