package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func GetPath(path string) (string, error) {
	finalPath := ""
	if len(strings.Split(path, ".")) > 0 && strings.Split(path, ".")[0] == "" {
		// this is a relative path
		curentpath, err := os.Getwd()
		if err != nil {
			return "", err
		}
		finalPath = filepath.Join(curentpath, path)
	} else {
		// this is an absolute path
		currentpath, err := filepath.Abs(path)
		if err != nil {
			return "", err
		}
		finalPath = currentpath
	}
	return finalPath, nil
}
