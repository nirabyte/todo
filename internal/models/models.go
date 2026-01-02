package models

import (
	"time"

	"github.com/charmbracelet/bubbles/textinput"
)

type AppState int

const (
	StateBrowse AppState = iota
	StateEditing
	StateCreating
	StateSettingTime
)

type SortMode int

const (
	SortOff SortMode = iota
	SortTodoFirst
	SortDoneFirst
)

const (
	AnimSparkle = iota
	AnimMatrix
	AnimWipeRight
	AnimWipeLeft
	AnimRainbow
	AnimWave
	AnimBinary
	AnimDissolve
	AnimFlip
	AnimPulse
	AnimTypewriter
	AnimParticle
	AnimRedact
	AnimChaos
	AnimConverge
	AnimBounce
	AnimSpin
	AnimZipper
	AnimEraser
	AnimGlitch
	AnimMoons
	AnimBraille
	AnimHex
	AnimReverse
	AnimCaseFlip
	AnimWide
	AnimTraffic
	AnimCenterStrike
	AnimLoading
	AnimSlider

	AnimCount = 30
)

type Task struct {
	ID       int64     `json:"id"`
	Title    string    `json:"title"`
	Done     bool      `json:"done"`
	DueAt    time.Time `json:"dueAt"`
	Notified bool      `json:"notified"`

	// Animation States
	IsAnimatingCheck bool      `json:"-"`
	IsDeleting       bool      `json:"-"`
	AnimType         int       `json:"-"`
	AnimStart        time.Time `json:"-"`
}

type AppData struct {
	ThemeIndex int      `json:"themeIndex"`
	SortMode   SortMode `json:"sortMode"`
	Tasks      []Task   `json:"tasks"`
}

type TickMsg struct{}

type Model struct {
	Tasks      []Task
	State      AppState
	SortMode   SortMode
	ThemeIndex int
	LastAnim   int

	Cursor    int
	Width     int
	Height    int
	TextInput textinput.Model
}

