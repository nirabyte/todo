package app

import (
	"math/rand"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nirabyte/todo/internal/models"
	"github.com/nirabyte/todo/internal/styles"
	"github.com/nirabyte/todo/internal/themes"
)

type App struct {
	Model *models.Model
}

func New() *App {
	ti := textinput.New()
	ti.CharLimit = 256
	ti.Width = 50
	ti.Prompt = ""

	rand.Seed(time.Now().UnixNano())

	data := models.LoadData()
	model := &models.Model{
		Tasks:      data.Tasks,
		State:      models.StateBrowse,
		SortMode:   data.SortMode,
		ThemeIndex: data.ThemeIndex,
		TextInput:  ti,
	}

	if model.ThemeIndex >= len(themes.All) {
		model.ThemeIndex = 0
	}
	styles.Update(themes.All[model.ThemeIndex])
	model.ApplySort()

	return &App{Model: model}
}

func (a *App) Run() error {
	p := tea.NewProgram(a.Model, tea.WithAltScreen())
	_, err := p.Run()
	return err
}

