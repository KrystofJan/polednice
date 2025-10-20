package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func GetPath(folderPath, fileName string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	absolutePath := strings.Split(folderPath, "/")
	absolutePath = append([]string{home}, absolutePath...)
	dbDir := filepath.Join(absolutePath...)
	err = os.MkdirAll(dbDir, 0755)
	if err != nil {
		return "", err
	}
	return filepath.Join(dbDir, fileName), nil
}
