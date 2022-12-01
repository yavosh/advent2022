package advent2022

import (
	"os"
)

func ReadFile(file string) (string, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
