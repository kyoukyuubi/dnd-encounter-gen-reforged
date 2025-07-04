package main

import (
	"fmt"
	"strings"

	errorhandling "github.com/kyoukyuubi/dnd-encounter-gen-reforged/internal/errorHandling"
	"github.com/kyoukyuubi/dnd-encounter-gen-reforged/internal/jsonHandler"
)

func commandList(cfg *Config, args ...string) error {
	// if no args are provided
    if len(args) == 0 {
        return fmt.Errorf("no argument provided")
    }

    // store the user input for easier use
    input := args[0]
    
    // define valid lists
    validList := []string{"environments", "planes", "sources", "types"}
    
    // find matching list using partial matching
    var matchedList string
    for _, list := range validList {
        if strings.Contains(list, input) {
            matchedList = list
            break
        }
    }
    
    // check if we found a match
    if matchedList == "" {
        return fmt.Errorf("unrecognized list: %s", input)
    }
    
    // run function based on what was matched
    switch matchedList {
    case "environments":
        listEnvironments()
    case "planes":
        listPlanes()
    case "sources":
        listSources()
    case "types":
        listTypes()
    }

	return nil
}

func listEnvironments() {
	// load the enviroments.json
	file, err := jsonHandler.LoadEnvirnments()
	if err != nil {
		errorhandling.LogError(err, "listEnvironments")
		fmt.Println("Error occurred, check the log in json/logs for details")
		fmt.Println("")
		return
	}

	// loop through the slice and print what was found:
	fmt.Println("Environments: ")
	for _, env := range file {
		fmt.Printf("* %s \n", env)
	}

	fmt.Println("")
}

func listPlanes() {
	// load the planes.json
	file, err := jsonHandler.LoadPlanes()
	if err != nil {
		errorhandling.LogError(err, "listPlanes")
		fmt.Println("Error occurred, check the log in json/logs for details")
		fmt.Println("")
		return
	}

	// set previous category and subCategory for chekcs to avoid duplicate entries
	previousCategory := ""
	previousSubCategory := ""

	// loop through the struct and print what was found, printing category and sub category
	for _, plane := range file.Plane {
		if previousCategory != plane.Category {
			// set the sub category for the next iteration of the loop
			previousCategory = plane.Category
			fmt.Println("")
			fmt.Printf("%s: \n", plane.Category)
		}

		// do the same check, but for the sub category
		// checking for no sub category too, to avoid printing issues
		if previousSubCategory != plane.SubCategory && plane.SubCategory != "" {
			// set the sub category for the next iteration of the loop
			previousSubCategory = plane.SubCategory
			fmt.Printf("  %s: \n", plane.SubCategory)
		} else if plane.SubCategory == "" {
    		// Reset previousSubCategory when we encounter an item with no subcategory
    		previousSubCategory = ""
}

		if previousSubCategory == "" {
			fmt.Printf("* %s\n", plane.Name)
		} else {
			fmt.Printf("    * %s\n", plane.Name)
		}
	}

	fmt.Println("")
}

func listSources() {
	// load the sources.json
	file, err := jsonHandler.LoadSources()
	if err != nil {
		errorhandling.LogError(err, "listSources")
		fmt.Println("Error occurred, check the log in json/logs for details")
		fmt.Println("")
		return
	}

	// loop through the struct and print the names, not the filenames
	fmt.Println("Sources: ")
	for _, source := range file.Source {
		fmt.Printf("* %s\n", source.Name)
	}

	fmt.Println("")
}

func listTypes() {
	// load the types.json
	file, err := jsonHandler.LoadTypes()
	if err != nil {
		errorhandling.LogError(err, "listTypes")
		fmt.Println("Error occurred, check the log in json/logs for details")
		fmt.Println("")
		return
	}

	// loop through the slice and print it to the user
	fmt.Println("Types: ")
	for _, types := range file {
		fmt.Printf("* %s\n", types)
	}

	fmt.Println("")
}