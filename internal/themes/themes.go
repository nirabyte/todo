package themes

import "github.com/charmbracelet/lipgloss"

type Theme struct {
	Name      string
	Bg        lipgloss.Color
	Fg        lipgloss.Color
	Dim       lipgloss.Color
	Accent    lipgloss.Color
	Secondary lipgloss.Color
	Success   lipgloss.Color
	Warning   lipgloss.Color
}

var All = []Theme{
	{"Catppuccin", "#000000", "#cdd6f4", "#6c7086", "#cba6f7", "#f5c2e7", "#a6e3a1", "#f38ba8"},
	{"Nord", "#2e3440", "#eceff4", "#4c566a", "#88c0d0", "#81a1c1", "#a3be8c", "#bf616a"},
	{"Gruvbox", "#282828", "#ebdbb2", "#928374", "#fabd2f", "#fe8019", "#b8bb26", "#fb4934"},
	{"Dracula", "#282a36", "#f8f8f2", "#6272a4", "#bd93f9", "#ff79c6", "#50fa7b", "#ff5555"},
	{"Tokyo Night", "#1a1b26", "#c0caf5", "#565f89", "#7aa2f7", "#bb9af7", "#9ece6a", "#f7768e"},
	{"Rose Pine", "#191724", "#e0def4", "#6e6a86", "#ebbcba", "#c4a7e7", "#31748f", "#eb6f92"},
	{"Everforest", "#272e33", "#d3c6aa", "#859289", "#a7c080", "#7fbbb3", "#a7c080", "#e67e80"},
	{"One Dark", "#282c34", "#abb2bf", "#5c6370", "#61afef", "#c678dd", "#98c379", "#e06c75"},
	{"Solarized", "#002b36", "#839496", "#586e75", "#268bd2", "#2aa198", "#859900", "#dc322f"},
	{"Kanagawa", "#1f1f28", "#dcd7ba", "#727169", "#7e9cd8", "#957fb8", "#76946a", "#c34043"},
}

