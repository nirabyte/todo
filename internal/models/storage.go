package models

import (
	"encoding/json"
	"os"
	"time"

	"github.com/nirabyte/todo/internal/config"
)

func LoadData() AppData {
	data, err := os.ReadFile(config.DataFile)

	hints := []Task{
		{ID: 1, Title: "Press 'n' to add a new task", Done: false},
		{ID: 2, Title: "Press 'e' to edit the selected task", Done: false},
		{ID: 3, Title: "Press 'd' to delete a task", Done: false},
		{ID: 4, Title: "Press 'space' to check/uncheck", Done: true},
		{ID: 5, Title: "Press '@' to set a timer notification", Done: false},
		{ID: 6, Title: "Press 's' to cycle sort modes", Done: false},
		{ID: 7, Title: "Press 't' to change the color theme", Done: false},
	}

	defaultData := AppData{
		ThemeIndex: 0,
		SortMode:   SortOff,
		Tasks:      hints,
	}

	if err != nil {
		return defaultData
	}

	var appData AppData
	if err := json.Unmarshal(data, &appData); err == nil {
		for i := range appData.Tasks {
			if appData.Tasks[i].ID == 0 {
				appData.Tasks[i].ID = time.Now().UnixNano() + int64(i)
			}
		}
		return appData
	}
	return defaultData
}

func (m *Model) Save() {
	var validTasks []Task
	for _, t := range m.Tasks {
		if !t.IsDeleting {
			validTasks = append(validTasks, t)
		}
	}
	data := AppData{
		ThemeIndex: m.ThemeIndex,
		SortMode:   m.SortMode,
		Tasks:      validTasks,
	}
	bytes, _ := json.MarshalIndent(data, "", "  ")
	os.WriteFile(config.DataFile, bytes, 0644)
}

