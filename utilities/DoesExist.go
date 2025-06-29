package utilities

import (
	"os"
	"errors"
)

func DoesExist(filepath string) bool {
	_, err := os.Stat(filepath)
	return !errors.Is(err, os.ErrNotExist)
}
