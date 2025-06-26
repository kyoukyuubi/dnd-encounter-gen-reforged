package main

import "fmt"

func commandListFilters(cfg *Config, args ...string) error {

	// planes
	if len(cfg.Config.Planes) == 0 {
		fmt.Println("Planes: Any")
	} else {
		fmt.Println("Planes: ")

		for _, plane := range cfg.Config.Sources {
			fmt.Printf("* %s\n", plane)
		}
	}

	fmt.Println("")

	// types
	if len(cfg.Config.Types) == 0 {
		fmt.Println("Types: Any")
	} else {
		fmt.Println("Type: ")

		for _, env := range cfg.Config.Sources {
			fmt.Printf("* %s\n", env)
		}
	}

	fmt.Println("")

	// sources
	if len(cfg.Config.Sources) == 0 {
		fmt.Println("Sources: Any")
	} else {
		fmt.Println("Sources: ")

		for _, source := range cfg.Config.Sources {
			fmt.Printf("* %s\n", source)
		}
	}

	fmt.Println("")

	// environments
	if len(cfg.Config.Environments) == 0 {
		fmt.Println("Environments: Any")
	} else {
		fmt.Println("Environments: ")

		for _, env := range cfg.Config.Sources {
			fmt.Printf("* %s\n", env)
		}
	}

	fmt.Println("")

	// party size
	fmt.Printf("Party Size: %v\n", cfg.Config.NumPlayers)
	fmt.Println("")

	// level
	fmt.Printf("Level: %v\n", cfg.Config.Level)
	fmt.Println("")

	// max creatures
	fmt.Printf("Max amount of creatures: %v\n", cfg.Config.MaxCreatures)
	fmt.Println("")

	// min exp
	fmt.Printf("Min amount of experience: %v\n", cfg.Config.MinExperience)
	fmt.Println("")

	// difficulty
	fmt.Printf("Difficulty: %s\n", cfg.Config.Difficulty)
	fmt.Println("")

	return nil
}