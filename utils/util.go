package utils

import (
	"os"
	"path/filepath"
	"runtime"
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

func GetDbPath() (string, error) {
	// Get dir
	rootDir, re := GetRootPath()
	if re != nil {
		return "", re
	}
	dbDir := filepath.Join(rootDir, "db")
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		return "", err
	}
	return dbDir, nil
}
