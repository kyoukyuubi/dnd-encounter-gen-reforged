package main

import "fmt"

func commandReset(cfg *Config, args ...string) error {
	// set everything to the default state (found in internal/config.go)
	cfg.Config.Planes = []string{}
	cfg.Config.Types = []string{}
	cfg.Config.Sources = []string{}
	cfg.Config.Environments = []string{}
	cfg.Config.NumPlayers = 4
	cfg.Config.Level = 1
	cfg.Config.MaxCreatures = 0
	cfg.Config.MinExperience = 0
	cfg.Config.Difficulty = "Moderate"

	fmt.Println("Filters has been reset!")
	fmt.Println("you can use 'list-filters' to see the changes")
	fmt.Println("")

	return nil
}
