package main

import (
	"fmt"

	"github.com/kyoukyuubi/dnd-encounter-gen-reforged/internal/config"
)

func main() {
    cfg := &Config{}

	// test the init function
	config.Init()

	config, err := config.Read()
	if err != nil {
		panic(err)
	}

	// checking if it worked
	fmt.Println(config)

	startRepl(cfg)
}