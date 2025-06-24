package main

import (
	"fmt"
	"strconv"

	errorhandling "github.com/kyoukyuubi/dnd-encounter-gen-reforged/internal/errorHandling"
	"github.com/kyoukyuubi/dnd-encounter-gen-reforged/internal/jsonHandler"
)

func calcBudget(cfg *Config) (int, error) {
	// store the level, group size and difficulty
	level := cfg.Config.Level
	size := cfg.Config.NumPlayers
	difficulty := cfg.Config.Difficulty

	// convert the level into a string
	levelString := strconv.Itoa(level)

	// load the exp json
	experienceTable, err := jsonHandler.LoadExpTable()
	if err != nil {
		errorhandling.LogError(err, "calcBudget")
		fmt.Println("Error occurred, check the log in json/logs for details")
		fmt.Println("")
		return 0, nil
	}

	// set base exp
	baseExp := experienceTable[levelString][difficulty]

	// calc the budget
	budget := baseExp * size

	return budget, nil
}