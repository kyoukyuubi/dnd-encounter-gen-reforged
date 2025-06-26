package main

import (
	"fmt"
	"strings"

	errorhandling "github.com/kyoukyuubi/dnd-encounter-gen-reforged/internal/errorHandling"
	"github.com/kyoukyuubi/dnd-encounter-gen-reforged/internal/jsonHandler"
)

func commandSource(cfg *Config, args ...string) error {
	// if no args are provided display what is in the config
	if len(args) == 0 {
		// check if the stored sources are empty
		if len(cfg.Config.Sources) == 0 {
			fmt.Println("Sources are set to any")
		} else {
			fmt.Println("Sources are set to: ")

			// loop through the stored sources and print them in an unordered list
			for _, s := range cfg.Config.Sources {
				fmt.Printf("* %s\n", s)
			}
		}

		fmt.Println("")
		return nil
	}

	// if args is set to 'Any' reset the config to an empty slice
	if args[0] == "any" {
		fmt.Println("Sources has been reset to default value")
		fmt.Println("")
		cfg.Config.Types = []string{}
		return nil
	}

	// load the sources from the json file
	sources, err := jsonHandler.LoadSources()
	if err != nil {
		errorhandling.LogError(err, "commandSource/LoadSources")
		fmt.Println("Error occurred, check the log in json/logs for details")
		fmt.Println("")
		return nil
	}

	// create an empty slice for putting the found values into
	var sliceFound []string

	// set a value to false, to be used after the loop to see if formatting needs to be printed
	errorOccurred := false

	// loop through the user inputs and see if they match the types from the json
	for _, input := range args {
		matchFound := false
		for _, t := range sources.Source {
			// check if what was inputtet matches what was in the json file
			// uses the value from the json file for error handling etc.
			if strings.Contains(strings.ToLower(t.Name), input) {
				sliceFound = append(sliceFound, t.Name)
				matchFound = true
				continue
			}
		}

		// if no matches is found tell the user
		if !matchFound {
			fmt.Printf("%s was not found in the determined sources\n", input)
			errorOccurred = true
		}
	}

	// check if error occurred, if so print a line break for the user to make things easier to read
	if errorOccurred {
		fmt.Println("-------------------------------")
	}

	// update the config with the new slice if the new slice list is not empty
	if len(sliceFound) != 0 {
		cfg.Config.Sources = sliceFound
		fmt.Println("The sources are now set to:")

		// loop through the created slice and print it in an unorederd list format
		for _, s := range sliceFound {
			fmt.Printf("* %s\n", s)
		}
		fmt.Println("")
	} else {
		fmt.Println("No matches were found in the determined types")
		fmt.Println("Sources list was not altered")
		fmt.Println("")
	}

	return nil
}