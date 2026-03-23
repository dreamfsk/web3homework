package utils

import (
	"github.com/pkg/errors"
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

//@function: PathExists
//@description: 文件目录是否存在
//@param: path string
//@return: bool, error

func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
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
