package ui

import (
	"math/rand"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gen2brain/beeep"
	"github.com/nirabyte/todo/internal/config"
	"github.com/nirabyte/todo/internal/models"
	"github.com/nirabyte/todo/internal/storage"
	"github.com/nirabyte/todo/internal/styles"
	"github.com/nirabyte/todo/internal/themes"
)

func (m *models.Model) Init() tea.Cmd {
	return textinput.Blink
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second/time.Duration(config.FPS), func(t time.Time) tea.Msg {
		return models.TickMsg{}
	})
}

func (m *models.Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:

		if m.State == models.StateEditing || m.State == models.StateCreating || m.State == models.StateSettingTime {
			switch msg.String() {
			case "enter":
				val := m.TextInput.Value()

				if m.State == models.StateSettingTime {
					if val != "" {
						dur, err := time.ParseDuration(val)
						if err == nil {
							m.Tasks[m.Cursor].DueAt = time.Now().Add(dur)
							m.Tasks[m.Cursor].Notified = false // Reset notification
						} else {
							m.Tasks[m.Cursor].DueAt = time.Time{}
						}
					} else {
						m.Tasks[m.Cursor].DueAt = time.Time{}
					}
					storage.SaveModel(m)
					m.State = models.StateBrowse
					m.TextInput.Blur()
					// Jumpstart ticker for timer updates
					return m, tickCmd()
				}

				if val == "" {
					m.State = models.StateBrowse
					m.TextInput.Blur()
					return m, nil
				}

				if m.State == models.StateCreating {
					m.Tasks = append(m.Tasks, models.Task{
						ID:    time.Now().UnixNano(),
						Title: val,
					})
					if m.SortMode != models.SortOff {
						m.ApplySort()
					}
					storage.SaveModel(m)
					m.State = models.StateBrowse
					m.TextInput.Blur()
					if m.SortMode == models.SortOff {
						m.Cursor = len(m.Tasks) - 1
					}
					return m, nil
				} else {
					m.Tasks[m.Cursor].Title = val
					storage.SaveModel(m)
					m.State = models.StateBrowse
					m.TextInput.Blur()
					return m, nil
				}

			case "esc":
				m.State = models.StateBrowse
				m.TextInput.Blur()
				return m, nil
			}
			m.TextInput, cmd = m.TextInput.Update(msg)
			return m, cmd
		}

		switch msg.String() {
		case "q", "ctrl+c":
			storage.SaveModel(m)
			return m, tea.Quit

		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down", "j":
			if m.Cursor < len(m.Tasks)-1 {
				m.Cursor++
			}

		case "t":
			m.ThemeIndex = (m.ThemeIndex + 1) % len(themes.All)
			styles.Update(themes.All[m.ThemeIndex])
			storage.SaveModel(m)

		case "s":
			m.SortMode = (m.SortMode + 1) % 3
			m.ApplySort()
			storage.SaveModel(m)

		case "n":
			m.State = models.StateCreating
			m.TextInput.Placeholder = "Task name..."
			m.TextInput.SetValue("")
			m.TextInput.Focus()
			m.Cursor = len(m.Tasks)
			return m, textinput.Blink

		case "e":
			if len(m.Tasks) > 0 {
				m.State = models.StateEditing
				m.TextInput.SetValue(m.Tasks[m.Cursor].Title)
				m.TextInput.Focus()
				m.TextInput.SetCursor(len(m.TextInput.Value()))
				return m, textinput.Blink
			}

		case "@":
			if len(m.Tasks) > 0 {
				m.State = models.StateSettingTime
				m.TextInput.Placeholder = "e.g. 10m, 1h2s, 10s..."
				m.TextInput.SetValue("")
				m.TextInput.Focus()
				return m, textinput.Blink
			}

		case "d":
			if len(m.Tasks) > 0 {
				m.Tasks[m.Cursor].IsDeleting = true
				m.Tasks[m.Cursor].AnimStart = time.Now()
				cmds = append(cmds, tickCmd())
			}

		case " ", "enter":
			if len(m.Tasks) > 0 {
				t := &m.Tasks[m.Cursor]
				t.Done = !t.Done

				if t.Done {
					t.IsAnimatingCheck = true
					t.AnimStart = time.Now()

					// Force Unique Random Animation
					newAnim := rand.Intn(models.AnimCount)
					for newAnim == m.LastAnim {
						newAnim = rand.Intn(models.AnimCount)
					}
					t.AnimType = newAnim
					m.LastAnim = newAnim

					cmds = append(cmds, tickCmd())
				} else {
					t.IsAnimatingCheck = false
				}
				m.ApplySort()
				storage.SaveModel(m)
			}
		}

	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		m.TextInput.Width = msg.Width - 10

	case models.TickMsg:
		needsTick := false
		for i := len(m.Tasks) - 1; i >= 0; i-- {
			t := &m.Tasks[i]

			// Animations
			if t.IsDeleting {
				if time.Since(t.AnimStart) > config.DeleteAnimDuration {
					m.Tasks = append(m.Tasks[:i], m.Tasks[i+1:]...)
					if m.Cursor >= len(m.Tasks) && m.Cursor > 0 {
						m.Cursor--
					}
					storage.SaveModel(m)
				} else {
					needsTick = true
				}
			}
			if t.IsAnimatingCheck {
				if time.Since(t.AnimStart) > config.CheckAnimDuration {
					t.IsAnimatingCheck = false
				} else {
					needsTick = true
				}
			}
			// Timer updates
			if !t.Done && !t.DueAt.IsZero() {
				needsTick = true
				if time.Now().After(t.DueAt) && !t.Notified {
					beeep.Notify("Todo Alert!", t.Title, "")
					t.Notified = true
					storage.SaveModel(m)
				}
			}
		}
		if needsTick {
			cmds = append(cmds, tickCmd())
		}
	}

	return m, tea.Batch(cmds...)
}

