package service

import (
	"errors"
	"fmt"
	"os"
)

func CreateFolders(dirs []string) error {
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0750); err != nil {
			return errors.New(fmt.Sprintf("create folder error: %v", err))
		}
	}

	return nil
}

func Exists(serviceName string) bool {
	info, err := os.Stat(serviceName)
	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir() && !os.IsNotExist(err)
}
