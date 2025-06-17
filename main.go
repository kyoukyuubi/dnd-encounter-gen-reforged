package main

import (
	"fmt"

	"github.com/kyoukyuubi/dnd-encounter-gen-reforged/internal/config"
)

func main() {
    // cfg := &Config{}

	// test the init function
	config.Init()

	config, err := config.Read()
	if err != nil {
		panic(err)
	}

	// checking if it worked
	fmt.Println("Current file:")
	fmt.Println(config)

	cfg := &Config{
		Config: config,
	}

	// update something in the cfg
	cfg.Config.Level = 6
	cfg.Config.MaxCreatures = 50
	cfg.Config.Types = []string{"Human", "Ooze"}

	fmt.Println("FILE UPDATED")

	startRepl(cfg)
}