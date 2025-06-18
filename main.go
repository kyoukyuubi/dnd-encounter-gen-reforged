package main

import (
	"fmt"
	"os"

	"github.com/kyoukyuubi/dnd-encounter-gen-reforged/internal/config"
	"github.com/kyoukyuubi/dnd-encounter-gen-reforged/internal/jsonHandler"
)

func main() {
	err := jsonHandler.Check()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	config.Init()

	config, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cfg := &Config{
		Config: config,
	}

	startRepl(cfg)
}