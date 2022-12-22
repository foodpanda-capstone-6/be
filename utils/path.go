package utils

import (
	"io"
	"log"
	"os"
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

func OpenOrCreateFile(path string) (string, io.Writer, error) {

	fullpath, err := GetFullPathOfPath(path)
	log.Printf("[utils::OpenOrCreateFile] opening %s", fullpath)
	if err != nil {
		return "", nil, err
	}

	err = os.MkdirAll(filepath.Dir(fullpath), os.ModePerm)

	if err != nil {
		return "", nil, err
	}

	file, err := os.OpenFile(fullpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	return fullpath, file, nil
}
