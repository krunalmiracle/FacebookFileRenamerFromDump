package utils

import (
	"errors"
	"os"
)

func IsFileExist(newPath string) bool {
	if _, err := os.Stat(newPath); err == nil {
		// EXIST
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		// DOES NOT EXIST
		return false
	} else {
		return false
	}
}
