package command

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type Command func(source, destination string) error

// MkDirCommand create a new directory if it does not exist.
func MkDirCommand(_, destination string) error {
	dir, _ := filepath.Split(destination)
	return os.MkdirAll(dir, os.ModePerm)
}

// CopyFile copies the file.
func CopyFile(source, destination string) error {
	input, err := ioutil.ReadFile(source)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(destination, input, 0644)
}
