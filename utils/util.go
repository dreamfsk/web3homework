package utils

import (
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

type ctxKey string

const CTX_KEY_OP ctxKey = "operator"

var (
	envLoaded bool
	envOnce   sync.Once
)

func GetRootPath() (string, error) {
	// Get dir
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		return "", os.ErrInvalid
	}
	util := filepath.Dir(currentFile)
	rootDir := filepath.Dir(util)
	return rootDir, nil
}

func GetPath() (string, error) {
	// Get dir
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		return "", os.ErrInvalid
	}
	util := filepath.Dir(currentFile)
	rootDir := filepath.Dir(util)
	dbDir := filepath.Join(rootDir, "db")
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		return "", err
	}
	return dbDir, nil
}

// GetEnv get env
func GetEnv(key, defaultValue string) string {
	LoadEnv()
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func LoadEnv() {
	if envLoaded != true {
		envOnce.Do(func() {
			_, currentFile, _, ok := runtime.Caller(0)
			if !ok {
				return
			}
			_utils := filepath.Dir(currentFile)
			_root := filepath.Dir(_utils)

			// Load .env file from examples directory
			envPath := filepath.Join(_root, ".env")
			if err := godotenv.Load(envPath); err != nil {
				// .env file is optional, so we don't fail if it doesn't exist
				// Environment variables can still be set directly via system environment
				return
			}
			envLoaded = true
		})
	}
}
