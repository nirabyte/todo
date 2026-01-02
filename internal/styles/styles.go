package styles

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/nirabyte/todo/internal/themes"
)

var (
	AppStyle          lipgloss.Style
	HeaderStyle       lipgloss.Style
	ListSelectedStyle lipgloss.Style
	ListItemStyle     lipgloss.Style
	InlineInputStyle  lipgloss.Style
	StrikeStyle       lipgloss.Style
	BinaryStyle       lipgloss.Style
	HelpStyle         lipgloss.Style
	DueStyle          lipgloss.Style
	OverdueStyle      lipgloss.Style
)

func Update(t themes.Theme) {
	AppStyle = lipgloss.NewStyle().Padding(1).Background(t.Bg)

	HeaderStyle = lipgloss.NewStyle().
		Foreground(t.Bg).
		Background(t.Accent).
		Bold(true).
		Padding(0, 1).
		MarginBottom(1)

	ListSelectedStyle = lipgloss.NewStyle().
		Border(lipgloss.ThickBorder(), false, false, false, true).
		BorderForeground(t.Accent).
		PaddingLeft(1).
		Foreground(t.Accent).
		Bold(true)

	ListItemStyle = lipgloss.NewStyle().
		PaddingLeft(2).
		Foreground(t.Fg)

	InlineInputStyle = lipgloss.NewStyle().
		Foreground(t.Accent).
		Bold(true)

	StrikeStyle = lipgloss.NewStyle().Foreground(t.Dim).Strikethrough(true)
	BinaryStyle = lipgloss.NewStyle().Foreground(t.Warning).Bold(true)
	HelpStyle = lipgloss.NewStyle().Foreground(t.Dim)

	DueStyle = lipgloss.NewStyle().Foreground(t.Secondary).Italic(true)
	OverdueStyle = lipgloss.NewStyle().Foreground(t.Warning).Bold(true).Blink(true)
}

