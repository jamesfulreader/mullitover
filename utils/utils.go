package utils

import (
	"os"
	"strings"
)

func ReadLines(filepath string) (string, error) {

	contentBytes, err := os.ReadFile(filepath)
	if err != nil {
		return "error reading file", err
	}

	fullCorruptMem := strings.TrimSpace(string(contentBytes))

	return fullCorruptMem, nil
}
