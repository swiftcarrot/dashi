package generate

import (
	"os"
	"path/filepath"
)

func getName() (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	name := filepath.Base(pwd)

	return name, nil
}
