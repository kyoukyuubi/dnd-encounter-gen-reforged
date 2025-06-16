package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type JsonConfig struct {
	Planes []string `json:"planes"`
	Types []string `json:"types"`
	Sources []string `json:"sources"`
	Environments []string `json:"environments"`
}

func Init() {
	// set the folder names and join them together using filepath
	// to make the entire filepath for the config.json file
	parentFolder := "json"
	configFolder := filepath.Join(parentFolder, "settings")
	filename := "config.json"
	fullpath := filepath.Join(configFolder, filename)

	// make the defualt struct
	defaultConfig := JsonConfig {
		Planes: []string{},
		Types: []string{},
		Sources: []string{},
		Environments: []string{},
	}

	// check if the folder exists, if not, create it
	// if there is an error, panic to close the software
	if err := os.MkdirAll(configFolder, 0755); err != nil {
		panic(err)
	}

	// check if the file exists, if not, creatite it
	// panic if there is an error to exit the software
	if _, err := os.Stat(fullpath); os.IsNotExist(err) {
		file, err := os.Create(fullpath)
		if err != nil {
			panic(err)
		}

		// make sure the file closes when function returns
		defer file.Close()

		// make the actual file
		if err := json.NewEncoder(file).Encode(defaultConfig); err != nil {
    		panic(err)
		}
	}
}