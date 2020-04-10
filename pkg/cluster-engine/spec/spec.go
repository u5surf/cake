package spec

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

const appName = ".capv-bootstrap"

// WriteToDisk writes the files to the hidden dir in the home directory
func WriteToDisk(dirname string, fileName string, specFile []byte, perms os.FileMode) error {
	var err error

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	newpath := filepath.Join(home, appName, dirname)
	os.MkdirAll(newpath, os.ModePerm)
	err = ioutil.WriteFile(filepath.Join(newpath, fileName), specFile, perms)

	return err
}
