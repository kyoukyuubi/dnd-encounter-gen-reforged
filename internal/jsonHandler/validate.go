package jsonHandler

import (
	// "fmt"
	"errors"
	"os"
	"path/filepath"
)

func Check() error {
	// set the parent folder
	folder := "json"

	// set the sub folder
	subFolder := "creatures"

	// make a slice that contains all of the files that needs to be present for the filters to work
	files := []string{"environments.json", "experience_table.json", "plane_catagories.json", "planes.json", "sources.json", "types.json"}

	// loop through the slice checking all files
	for _, file := range files {
		// construct the fullpath
		fullpath := filepath.Join(folder, file)

		// check if the file exists
		if _, err := os.Stat(fullpath); os.IsNotExist(err) {
			// errString := file + " is not found"
			return errors.New("file not found")
			// return fmt.Errorf("%s not found", file)
		}
	}

	// read the entries in the subfolder
	path := filepath.Join(folder, subFolder)
	entries, err := os.ReadDir(path)
	if err != nil {
		// return a folder if reading the entries fails, for example folder not found
		return err
	}

	// check if folder is empty, if it is return error
	if len(entries) == 0 {
		return errors.New("creatures folder is empty")
		// return fmt.Errorf("creatures folder is empty")
	}

	return nil
}