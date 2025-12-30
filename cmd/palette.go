package main

import "github.com/charmbracelet/lipgloss"

// Base16Palette defines a base16 color palette
type Base16Palette struct {
	Base00 lipgloss.Color // Background
	Base01 lipgloss.Color // Lighter Background
	Base02 lipgloss.Color // Selection Background
	Base03 lipgloss.Color // Comments, Invisibles, Line Highlight
	Base04 lipgloss.Color // Dark Foreground
	Base05 lipgloss.Color // Default Foreground
	Base06 lipgloss.Color // Light Foreground
	Base07 lipgloss.Color // Light Background
	Base08 lipgloss.Color // Variables, Tags, Markup Link Text
	Base09 lipgloss.Color // Integers, Boolean, Constants
	Base0A lipgloss.Color // Classes, Markup Bold, Search Text Background
	Base0B lipgloss.Color // Strings, Inherited Class, Markup Code
	Base0C lipgloss.Color // Support, Regular Expressions, Escape Characters
	Base0D lipgloss.Color // Functions, Methods, Headings
	Base0E lipgloss.Color // Keywords, Storage, Selector, Markup Italic
	Base0F lipgloss.Color // Deprecated, Embedded Language Tags
}

// CatppuccinMocha returns a Catppuccin Mocha palette in base16 colors
func CatppuccinMocha() Base16Palette {
	return Base16Palette{
		Base00: lipgloss.Color("#1E1E2E"),
		Base01: lipgloss.Color("#181825"),
		Base02: lipgloss.Color("#313244"),
		Base03: lipgloss.Color("#45475A"),
		Base04: lipgloss.Color("#585B70"),
		Base05: lipgloss.Color("#CDD6F4"),
		Base06: lipgloss.Color("#F5E0DC"),
		Base07: lipgloss.Color("#B4BEFE"),
		Base08: lipgloss.Color("#F38BA8"),
		Base09: lipgloss.Color("#FAB387"),
		Base0A: lipgloss.Color("#F9E2AF"),
		Base0B: lipgloss.Color("#A6E3A1"),
		Base0C: lipgloss.Color("#94E2D5"),
		Base0D: lipgloss.Color("#89B4FA"),
		Base0E: lipgloss.Color("#CBA6F7"),
		Base0F: lipgloss.Color("#F2CDCD"),
	}
}
