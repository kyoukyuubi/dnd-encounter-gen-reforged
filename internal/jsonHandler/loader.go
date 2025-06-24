package jsonHandler

import (
	"encoding/json"
	"io"
	"os"
)

type Type struct {
	Name string `json:"name"`
}

type TypesFile struct {
	Types []Type `json:"types"`
}

func LoadExpTable() (map[string]map[string]int, error) {
	// set the path
	path := "json/experience_table.json"

	// opent he file handling errors
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	// make sure the file closes
	defer file.Close()

	// read the file
	readFile, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// create an empty map
	fileContent := make(map[string]map[string]int)

	// unmarshal json data into the map
	err = json.Unmarshal(readFile, &fileContent)
	if err != nil {
		return nil, err
	}

	return fileContent, nil
}

func LoadTypes() ([]string, error) {
	// set the path
	path := "json/types.json"

	// open the file, handling errors
	file, err := os.Open(path)
	if err != nil {
		return []string{}, err
	}

	// make sure the file closes once function returns
	defer file.Close()

	// read the file
	readFile, err := io.ReadAll(file)
	if err != nil {
		return []string{}, err
	}

	// create an empty struct
	var fileContent TypesFile 

	// unmarshal json data into the struct
	err = json.Unmarshal(readFile, &fileContent)
	if err != nil {
		return []string{}, err
	}

	// create an empty slice
	var returnSlice []string

	// loop through the struct and put the data into the newly created slice
	for _, t := range fileContent.Types {
		returnSlice = append(returnSlice, t.Name)
	}

	return returnSlice, nil
}