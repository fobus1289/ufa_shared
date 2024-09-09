package service

import (
	"fmt"
	"os"
)

func CreateFolders(dirs []string) error {
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0750); err != nil {
			return fmt.Errorf("create folder error: %v", err)
		}
	}

	return nil
}

func Exists(serviceName string) bool {
	info, err := os.Stat(serviceName)
	{
		if os.IsNotExist(err) {
			return false
		}
	}

	return info.IsDir() && !os.IsNotExist(err)
}
