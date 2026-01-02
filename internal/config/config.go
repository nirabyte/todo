package config

import (
	"os"
	"path/filepath"
	"runtime"
	"time"
)

const (
	CheckAnimDuration  = 290 * time.Millisecond
	DeleteAnimDuration = 200 * time.Millisecond
	FPS                = 60
)

var (
	// Data file configuration
	DataPath    = getEnvOrDefault("DATA_PATH", getDataHomeOrDefault("data"))
	DataFile    = getEnvOrDefault("DATA_FILE", "todos.json")
	StorageType = "file" // file, s3, mongodb, postgres

	// Encryption
	EncryptionKey = "" // 64 hex chars (32 bytes)

	// S3 configuration
	S3Bucket = ""
	S3Region = "us-east-1"

	// MongoDB configuration
	MongoURI        = "mongodb://localhost:27017"
	MongoDB         = "todo"
	MongoCollection = "tasks"

	// PostgreSQL configuration
	PostgresDSN   = "postgres://user:pass@localhost/todo?sslmode=disable"
	PostgresTable = "tasks"
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
