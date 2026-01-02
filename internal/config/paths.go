package config

import (
	"os"
	"path/filepath"
	"runtime"
)

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getDataHomeOrDefault(defaultValue string) string {
	for _, env := range []string{"XDG_DATA_HOME", "LOCALAPPDATA"} {
		if val := os.Getenv(env); val != "" {
			return ensureTodoDir(val)
		}
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return ensureTodoDir(defaultValue)
	}

	var fallback string
	switch runtime.GOOS {
	case "windows":
		fallback = filepath.Join(home, "AppData", "Local")
	case "darwin":
		fallback = filepath.Join(home, "Library", "Application Support")
	default:
		fallback = filepath.Join(home, ".local", "share")
	}

	return ensureTodoDir(fallback)
}

func ensureTodoDir(path string) string {
	if filepath.Base(path) != "todo" {
		path = filepath.Join(path, "todo")
	}
	_ = os.MkdirAll(path, 0700)
	return path
}
