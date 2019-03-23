package folder

import (
	"os"
)

// CreateFolders untuk membuat folder recursive
func CreateFolders(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
