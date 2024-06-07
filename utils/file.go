package utils

import (
	"path"
	"runtime"
)

func CurrentDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

func CurrentFile() string {
	_, filename, _, _ := runtime.Caller(1)
	return filename
}
