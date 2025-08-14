package utils

import (
	"errors"
	"fmt"
	"os"
)

func FileExsits(path string) (bool, error) {
	if _, err := os.Stat(path); err == nil {
		return true, nil
	} else if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else {
		return false, fmt.Errorf("An error occured when trying to verify if a file exists")
	}
}
