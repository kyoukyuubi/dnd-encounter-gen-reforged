package jsonHandler

import (
	"encoding/json"
	"io"
	"os"
)

// structs for the types.json
type Type struct {
	Name string `json:"name"`
}

type TypesFile struct {
	Types []Type `json:"types"`
}

// structs for the environments.json
type Environment struct {
	Name string `json:"name"`
}

type EnvironmentFile struct {
	Environments []Environment `json:"environments"`
}

// structs for the sources.json
type Source struct {
	Name string `json:"name"`
	Filename string `json:"filename"`
}

type SourceFile struct {
	Source []Source `json:"sources"`
}

// structs for the planes.json
type PlaneData struct {
	Name string `json:"name"`
	Category string `json:"category"`
	SubCategory string `json:"sub-category,omitempty"`
}

type Plane struct {
	Plane []PlaneData `json:"planes"`
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

func LoadEnvirnments() ([]string, error) {
	// set the patch
	path := "json/environments.json"

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

	// create an empty struct for the data
	var fileContent EnvironmentFile

	// unmarshal json data into the struct
	err = json.Unmarshal(readFile, &fileContent)
	if err != nil {
		return []string{}, err
	}

	// create an empty slice
	var returnSlice []string

	// loop through the struct and put the data into the newly created slice
	for _, t := range fileContent.Environments {
		returnSlice = append(returnSlice, t.Name)
	}

	return returnSlice, nil
}

func LoadSources() (SourceFile, error) {
	// set the patch
	path := "json/sources.json"

	// open the file, handling erros
	file, err := os.Open(path)
	if err != nil {
		return SourceFile{}, err
	}

	// close the file when function closes
	defer file.Close()

	// read the file
	readFile, err := io.ReadAll(file)
	if err != nil {
		return SourceFile{}, err
	}

	// create an empty struct for the data
	var fileContent SourceFile

	// unmarshal json data into the struct
	err = json.Unmarshal(readFile, &fileContent)
	if err != nil {
		return SourceFile{}, err
	}

	return fileContent, nil
}

func LoadPlanes() (Plane, error) {
	// set the path to the file
	path := "json/planes.json"

	// open the file and handle the errors
	file, err := os.Open(path)
	if err != nil {
		return Plane{}, err
	}

	// close the file when the function closes
	defer file.Close()

	// read the file and handle the errors
	readFile, err := io.ReadAll(file)
	if err != nil {
		return Plane{}, err
	}

	// make an empty struct for the json data
	var fileContent Plane

	// unmarshal json data into the struct
	err = json.Unmarshal(readFile, &fileContent)
	if err != nil {
		return Plane{}, err
	}
	
	return fileContent, err
}