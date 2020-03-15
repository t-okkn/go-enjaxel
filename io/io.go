package io

import (
	"os"
)

func IsExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func FileExists(path string) bool {
	stat, err := os.Stat(path)
	if err != nil { return false }

	return !(stat.IsDir())
}

func DirExists(path string) bool {
	stat, err := os.Stat(path)
	if err != nil { return false }

	return stat.IsDir()
}

