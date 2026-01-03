package config

import (
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
