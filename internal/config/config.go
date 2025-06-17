package config

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

// the structure of the config file with the json fields
type JsonConfig struct {
	Planes []string `json:"planes"`
	Types []string `json:"types"`
	Sources []string `json:"sources"`
	Environments []string `json:"environments"`
	NumPlayers int `json:"numPlayers"`
	Level int `json:"level"`
	MaxCreatures int `json:"maxCreatures"`
	MinExperience int `json:"minExperience"`
	Difficulty string `json:"difficulty"`
}

func Init() {
	// set the folder names and join them together using filepath
	// to make the entire filepath for the config.json file
	parentFolder := "json"
	configFolder := filepath.Join(parentFolder, "settings")
	filename := "config.json"
	fullpath := filepath.Join(configFolder, filename)

	// make the default struct
	defaultConfig := JsonConfig {
		Planes: []string{},
		Types: []string{},
		Sources: []string{},
		Environments: []string{},
		NumPlayers: 4,
		Level: 1,
		MaxCreatures: 0,
		MinExperience: 0,
		Difficulty: "Moderate",
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

func Read() (JsonConfig, error) {
	// set the path to the config file
	fullpath := "json/settings/config.json"

	// read the file and open it
	file, err := os.Open(fullpath)
	if err != nil {
		return JsonConfig{}, err
	}

	// make sure the file closes once function returns
	defer file.Close()

	// read the entire file
	readFile, err := io.ReadAll(file)
	if err != nil {
		return JsonConfig{}, err
	}

	// initiate the struct
	var config JsonConfig

	// unmarshel into the struct
	err = json.Unmarshal(readFile, &config)
	if err != nil {
		return JsonConfig{}, err
	}

	// return the JsonConfig struct
	return config, nil
}

func Update(newConfig JsonConfig) error {
	// set the path to the config file
	fullpath := "json/settings/config.json"

	// prepare the date to be written
	data, err := json.Marshal(newConfig)
	if err != nil {
		return err
	}

	// write to the actual file
	err = os.WriteFile(fullpath, data, 0644)
	if err != nil {
		return err
	}

	// return nil on success
	return nil
}