package jsonHandler

import (
	"encoding/json"
	"fmt"
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
	Name     string `json:"name"`
	Filename string `json:"filename"`
}

type SourceFile struct {
	Source []Source `json:"sources"`
}

// structs for the planes.json
type PlaneData struct {
	Name        string `json:"name"`
	Category    string `json:"category"`
	SubCategory string `json:"sub-category,omitempty"`
}

type Plane struct {
	Plane []PlaneData `json:"planes"`
}

// struct for creatures
type CreatureData struct {
	Name        string   `json:"name"`
	Type        string   `json:"type"`
	Exp         int      `json:"Exp"`
	Environment []string `json:"Environment"`
	Plane       []string `json:"plane"`
	Book        string   `json:"book"`
	Page        int      `json:"page"`
}

type Creature struct {
	Creature []CreatureData `json:"creatures"`
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

func LoadCreatures(cfgSources []string) (Creature, error) {
	// set the main folder structure
	dir := "json/creatures/"

	// load the sources.json for loopup purposes
	sources, err := LoadSources()
	if err != nil {
		return Creature{}, err
	}

	// make a slice to store the filepath for the loop
	var filePath []string

	// loop through and store the full paths in the slice
	// If cfgSources is empty, include all source files
	if len(cfgSources) == 0 {
		for _, source := range sources.Source {
			filePath = append(filePath, dir+source.Filename+".json")
		}
	} else {
		for _, source := range sources.Source {
			for _, cfgSource := range cfgSources {
				if source.Name == cfgSource {
					path := dir + source.Filename + ".json"
					filePath = append(filePath, path)
				}
			}
		}
	}

	// check if the filePath is empty, if it is throw and error
	if len(filePath) == 0 {
		return Creature{}, fmt.Errorf("no source files matched your selection: %v", cfgSources)
	}

	// make the var that stores the creatures
	var creatures []CreatureData

	// loop through the fileParh slice and load the creatures into a temp var
	for _, path := range filePath {

		// read the file
		data, err := os.ReadFile(path)
		if err != nil {
			return Creature{}, err
		}

		// make a temp var to store the creatures
		var tempCreatures Creature

		// unmarshal the data
		err = json.Unmarshal(data, &tempCreatures)
		if err != nil {
			return Creature{}, err
		}

		// store the data into the return var
		creatures = append(creatures, tempCreatures.Creature...)
	}

	return Creature{Creature: creatures}, nil
}
