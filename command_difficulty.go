package main

import (
	"fmt"
	"strconv"
	"strings"

	errorhandling "github.com/kyoukyuubi/dnd-encounter-gen-reforged/internal/errorHandling"
	"github.com/kyoukyuubi/dnd-encounter-gen-reforged/internal/jsonHandler"
)

func commandDifficulty(cfg *Config, args ...string) error {
	// if no args are provided display what is in the config
	if len(args) == 0 {
		fmt.Printf("Difficulty is set to: %s", cfg.Config.Difficulty)
		fmt.Println("")
		return nil
	}

	// store the level for later
	level := cfg.Config.Level

	// convert level to a string
	levelString := strconv.Itoa(level)

	// load the level table
	expData, err := jsonHandler.LoadExpTable()
	if err != nil {
		errorhandling.LogError(err, "commandDifficulty/LoadExpTable")
		fmt.Println("Error occurred, check the log in json/logs for details")
		fmt.Println("")
		return nil
	}

	// loop through the expData, only careing about the keys and see if input matches
	// if it does, set it as the new difficulty
	for _, input := range args {
		for key := range expData[levelString] {

			// check if what was inputtet matches what was in the json file
			// uses the value from the json file for error handling etc.
			if strings.Contains(strings.ToLower(key), input) {
				cfg.Config.Difficulty = key
				fmt.Printf("Difficulty set to: %s\n", key)
				fmt.Println("")
				return nil
			}
		}
	}

	fmt.Printf("%s was not found in the determined difficulties\n", args[0])
	fmt.Println("")
	return nil
}