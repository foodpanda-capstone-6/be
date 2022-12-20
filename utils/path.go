package utils

import (
	"path/filepath"
)

func GetFullPathOfPath(relativeOrAbsolutePath string) (string, error) {
	if filepath.IsAbs(relativeOrAbsolutePath) {
		return relativeOrAbsolutePath, nil
	} else {
		path, err := filepath.Abs(relativeOrAbsolutePath)
		if err != nil {
			return "", err
		}
		return path, nil
	}
}
