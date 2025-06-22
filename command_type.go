package main

import (
	"fmt"
	"strings"

	errorhandling "github.com/kyoukyuubi/dnd-encounter-gen-reforged/internal/errorHandling"
	"github.com/kyoukyuubi/dnd-encounter-gen-reforged/internal/jsonHandler"
)

func commandType(cfg *Config, args ...string) error {
	// if no args are provided display what is in the config
	if len(args) == 0 {
		fmt.Printf("Types are set to %v\n", cfg.Config.Types)
		return nil
	}

	// if args is set to 'Any' reset the config to an empty slice
	if args[0] == "any" {
		fmt.Println("types has been reset to default value")
		cfg.Config.Types = []string{}
		return nil
	}

	// load the types from the json file
	types, err := jsonHandler.LoadTypes()
	if err != nil {
		errorhandling.LogError(err, "commandType")
		fmt.Println("Error occurred, check the log in json/logs for details")
		return nil
	}

	// create an empty slice for putting the found values into
	var sliceTypes []string

	// loop through the types gotten from the json and check if inputted text matches
	for _, t := range types {
		for _, input := range args {
			// check if what was inputtet matches what was in the json file
			// uses the value from the json file for error handling etc.
			if strings.Contains(strings.ToLower(t), input) {
				sliceTypes= append(sliceTypes, t)
				continue
			}
		}
	}

	// show the user what the types are now set to
	fmt.Printf("The types are now set to %v\n", sliceTypes)

	// update the config with the new slice
	cfg.Config.Types = sliceTypes

	return nil
}