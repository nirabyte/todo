package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/nirabyte/todo/internal/styles"
	"github.com/nirabyte/todo/internal/themes"
)

func (m *Model) View() string {
	currentTheme := themes.All[m.ThemeIndex]
	var content string

	content = m.viewList(currentTheme)

	header := styles.HeaderStyle.Render("// TODO LIST")

	container := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(currentTheme.Accent).
		Width(min(m.Width-4, 100)).
		Height(m.Height - 7).
		Render(content)

	sortStr := "Off"
	if m.SortMode == SortTodoFirst {
		sortStr = "Todo"
	}
	if m.SortMode == SortDoneFirst {
		sortStr = "Done"
	}

	help := fmt.Sprintf("Theme: %s (t) • Sort: %s (s) • New (n) • Edit (e) • Check (Space) • Notify (@) • Del (d)", currentTheme.Name, sortStr)
	status := styles.HelpStyle.Render(help)

	ui := lipgloss.JoinVertical(lipgloss.Center, header, container, status)
	return lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center, ui)
}

func (m *Model) viewList(t themes.Theme) string {
	if len(m.Tasks) == 0 && m.State != StateCreating {
		return styles.HelpStyle.Padding(2).Render("No tasks.")
	}
	var s strings.Builder

	count := len(m.Tasks)
	if m.State == StateCreating {
		count++
	}
	creatingIndex := len(m.Tasks)

	// Layout Calc: Window - Borders(2) - Number(4) - Icon(3) - Timer(approx 25) - Spacers(6)
	availableWidth := min(m.Width-4, 100)
	textWidth := availableWidth - 40 // Give extra room for timer
	if textWidth < 10 {
		textWidth = 10
	}

	for i := 0; i < count; i++ {
		selected := false
		if m.State == StateCreating {
			if i == creatingIndex {
				selected = true
			}
		} else {
			if m.Cursor == i {
				selected = true
			}
		}

		numberStr := fmt.Sprintf("%d.", i+1)
		var checkIcon string
		var titleContent string
		var dueContent string

		isEditingThis := (m.State == StateEditing && i == m.Cursor)
		isCreatingThis := (m.State == StateCreating && i == creatingIndex)
		isSettingTime := (m.State == StateSettingTime && i == m.Cursor)

		if isEditingThis || isCreatingThis {
			checkIcon = lipgloss.NewStyle().Foreground(t.Accent).Render(">")
			m.TextInput.Width = textWidth
			titleContent = styles.InlineInputStyle.Render(m.TextInput.View())
		} else {
			task := m.Tasks[i]

			if task.Done {
				checkIcon = lipgloss.NewStyle().Foreground(t.Success).Render("[✔]")
			} else {
				checkIcon = lipgloss.NewStyle().Foreground(t.Accent).Render("[ ]")
			}

			var rawTitle string
			if task.IsDeleting {
				rawTitle = renderDeleteAnim(task.Title, t)
			} else if task.IsAnimatingCheck {
				rawTitle = renderCheckAnim(task, t)
			} else if task.Done {
				rawTitle = styles.StrikeStyle.Render(task.Title)
			} else {
				rawTitle = lipgloss.NewStyle().Foreground(t.Fg).Render(task.Title)
			}

			titleContent = lipgloss.NewStyle().Width(textWidth).Render(rawTitle)

			if isSettingTime {
				m.TextInput.Width = 20
				dueContent = styles.InlineInputStyle.Render(m.TextInput.View())
			} else if !task.DueAt.IsZero() && !task.Done {
				timeRemaining := time.Until(task.DueAt)
				if timeRemaining < 0 {
					dueContent = styles.OverdueStyle.Render("[OVERDUE]")
				} else {
					dueContent = styles.DueStyle.Render(shortDur(timeRemaining))
				}
			}
		}

		leftBlock := lipgloss.JoinHorizontal(lipgloss.Top,
			lipgloss.NewStyle().Foreground(t.Dim).Width(4).Align(lipgloss.Right).Render(numberStr),
			" ",
			lipgloss.NewStyle().Width(3).Align(lipgloss.Center).Render(checkIcon),
			" ",
		)

		row := lipgloss.JoinHorizontal(lipgloss.Top,
			leftBlock,
			titleContent,
			"   ",
			dueContent,
		)

		if selected {
			s.WriteString(styles.ListSelectedStyle.Render(row))
		} else {
			s.WriteString(styles.ListItemStyle.Render(row))
		}
		s.WriteString("\n")
	}
	return s.String()
}

func shortDur(d time.Duration) string {
	d = d.Round(time.Second)
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60

	if h > 0 {
		return fmt.Sprintf("%dh%dm%ds", h, m, s)
	}
	if m > 0 {
		return fmt.Sprintf("%dm%ds", m, s)
	}
	return fmt.Sprintf("%ds", s)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

