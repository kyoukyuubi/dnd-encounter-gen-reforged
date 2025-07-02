package main

import (
	"fmt"
)

func commandHelp(cfg *Config, args ...string) error {
	// display text welcomming the user
	fmt.Println("Welcome to the D&D Encounter Generator!")
	fmt.Println("Usage:")
	fmt.Println("")

	// loops through the commands and displays info to the user
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
		fmt.Printf("example: %s\n", command.example)
		fmt.Println("")
	}

	fmt.Println("-----------------")
	fmt.Println("example use of the generator")
	fmt.Println("level 5")
	fmt.Println("party-size 5")
	fmt.Println("environment forest")
	fmt.Println("generate")
	fmt.Println("")

	return nil
}