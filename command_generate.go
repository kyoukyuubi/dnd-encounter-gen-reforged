package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/kyoukyuubi/dnd-encounter-gen-reforged/internal/jsonHandler"
)

type EncounterSelection struct {
	Count int
	Data  jsonHandler.CreatureData
}

func commandGenerate(cfg *Config, args ...string) error {
	// load the creatures according to the selected sources in the cfg/config
	creatures, err := jsonHandler.LoadCreatures(cfg.Config.Sources)
	if err != nil {
		return err
	}

	// check if the envirments, plae, types and MinExperience filters are set and call the helper functions to modify
	// the creatues slice

	// environment Filter
	if len(cfg.Config.Environments) != 0 {
		creatures = filterEnvironment(cfg.Config.Environments, creatures)
	}

	// plane Filter
	if len(cfg.Config.Planes) != 0 {
		creatures = filterPlane(cfg.Config.Planes, creatures)
	}

	// type filter
	if len(cfg.Config.Types) != 0 {
		creatures = filterType(cfg.Config.Types, creatures)
	}

	// min experience filter
	if cfg.Config.MinExperience != 0 {
		creatures = filterMinExp(cfg.Config.MinExperience, creatures)
	}

	// calculated the budget
	budget, err := calcBudget(cfg)
	if err != nil {
		return err
	}

	// create an empty struct that contains the CreatureDate and the amount of creatures selected of set type
	// var selected selectedCreature
	selectedCreatures := make(map[string]*EncounterSelection)
	var totalCreatures int
	var totalExp int

	// shuffle the slice of creatures
	rand.Shuffle(len(creatures.Creature), func(i, j int) {
		creatures.Creature[i], creatures.Creature[j] = creatures.Creature[j], creatures.Creature[i]
	})

	// loop through it and add it to the list
	// making sure to stay within budget and max creatures
	for {
		// check iof the totalcreate exceeds the max creatures
		if totalCreatures >= cfg.Config.MaxCreatures && cfg.Config.MaxCreatures != 0 {
			break
		}

		// see if we can afford any more creatures
		canAfford := false
		for _, c := range creatures.Creature {
			// increase the total creatures by 1 and the totalexp by 1 creature (it's exp)
			projectedCreatures := totalCreatures + 1
			projectedExp := totalExp + c.Exp

			// get the multiplier and calculate the new total exp
			multiplier := getMultiplier(projectedCreatures)
			totalCost := int(math.Round(float64(projectedExp) * multiplier))

			// see if it fits in the budget and break out of the loop if true
			if totalCost <= budget {
				canAfford = true
				break
			}
		}

		// if it cant find anything that fits in the budget break out of the loop
		if !canAfford {
			break
		}

		// select a random creature and store it
		sliceLength := len(creatures.Creature)
		randomIndex := rand.Intn(sliceLength)
		randomCreature := creatures.Creature[randomIndex]

		// Project what would happen if we add the random creature
		projectedCreatures := totalCreatures + 1
		projectedExp := totalExp + randomCreature.Exp
		mult := getMultiplier(projectedCreatures)
		totalCost := int(math.Round(float64(projectedExp) * mult))

		// check if the random creature fits in the budget, if not move on
		if totalCost > budget {
			continue
		}

		// if it fits in the budget add it
		// if it is already in the map, increase the count
		if selected, exists := selectedCreatures[randomCreature.Name]; exists {
			selected.Count++
		} else {
			selectedCreatures[randomCreature.Name] = &EncounterSelection{Count: 1, Data: randomCreature}
		}
		// increase the total creatures and total exp
		totalCreatures++
		totalExp += randomCreature.Exp
	}

	// if selected creatures are empty, tell the user
	if len(selectedCreatures) == 0 {
		fmt.Println("No creatures found in generation")
		fmt.Println("This could be because no creatures matched your filters and/or your level and difficulty.")
		fmt.Println("Please look at your filters using 'list-filters' and change them around, then try again")
		fmt.Println("")
		return nil
	}

	fmt.Println("Encounter Creatures:")
	fmt.Println("===================")

	// loop through and print the result to the user
	for _, creature := range selectedCreatures {
		fmt.Printf("Name:   %s\n", creature.Data.Name)
		fmt.Printf("Type:   %s\n", creature.Data.Type)
		fmt.Printf("Book:   %s (page %d)\n", creature.Data.Book, creature.Data.Page)
		fmt.Printf("Amount: %d\n", creature.Count)
		fmt.Println("--------------------------------")
	}

	displayBudget, err := calcBudget(cfg)
	if err != nil {
		return err
	}

	fmt.Printf("Total Creatures: %d\n", totalCreatures)
	fmt.Printf("Total XP Budget Used: %d out of: %d\n", totalExp, displayBudget)
	fmt.Println("")
	return nil
}

func filterEnvironment(envirments []string, creatures jsonHandler.Creature) jsonHandler.Creature {
	// build a set of environments for quick lookup
	envirmentsSet := make(map[string]struct{})
	for _, env := range envirments {
		// using struct{}{} since it takes less memory than a bool
		envirmentsSet[env] = struct{}{}
	}

	// make the return struct
	var filtered []jsonHandler.CreatureData

	// loop through and add creatures that passes the filter to the return struct
	for _, creature := range creatures.Creature {
		for _, env := range creature.Environment {

			// check to see if the environments from the creature are found in the map that was created
			if _, ok := envirmentsSet[env]; ok {
				filtered = append(filtered, creature)

				// break to avoid duplicates
				break
			}
		}

		// check if the creatures plane slice is empty, if it as add it. Since that means any
		if len(creature.Environment) == 0 {
			filtered = append(filtered, creature)
		}
	}

	return jsonHandler.Creature{Creature: filtered}
}

func filterPlane(planes []string, creatures jsonHandler.Creature) jsonHandler.Creature {
	// build a set of environments for quick lookup
	planeSet := make(map[string]struct{})
	for _, plane := range planes {
		// using struct{}{} since it takes less memory than a bool
		planeSet[plane] = struct{}{}
	}

	// make the return struct
	var filtered []jsonHandler.CreatureData

	// loop through and add creatures that passes the filter to the return struct
	for _, creature := range creatures.Creature {
		for _, plane := range creature.Plane {

			// check to see if the environments from the creature are found in the map that was created
			if _, ok := planeSet[plane]; ok {
				filtered = append(filtered, creature)

				// break to avoid duplicates
				break
			}
		}

		// check if the creatures plane slice is empty, if it as add it. Since that means any
		if len(creature.Plane) == 0 {
			filtered = append(filtered, creature)
		}
	}

	return jsonHandler.Creature{Creature: filtered}
}

func filterType(types []string, creatures jsonHandler.Creature) jsonHandler.Creature {
	// make the return struct
	var filtered []jsonHandler.CreatureData

	// loop through and add creatures that passes the filter to the return struct
	for _, creature := range creatures.Creature {
		for _, t := range types {
			if creature.Type == t {
				filtered = append(filtered, creature)

				// break to avoid duplicates
				break
			}
		}
	}

	return jsonHandler.Creature{Creature: filtered}
}

func filterMinExp(minExp int, creatures jsonHandler.Creature) jsonHandler.Creature {
	// make the return struct
	var filtered []jsonHandler.CreatureData

	// loop through the creatures and compare them to the inputtet minExp
	// add them to filtered if are above or equal to
	for _, creature := range creatures.Creature {
		if creature.Exp >= minExp {
			filtered = append(filtered, creature)
		}
	}

	return jsonHandler.Creature{Creature: filtered}
}

func getMultiplier(totalCreaturs int) float64 {
	// init the multiplier
	var multiplier float64

	// set the correct multiplier
	switch {
	case totalCreaturs == 1:
		multiplier = 1.0
	case totalCreaturs == 2:
		multiplier = 1.5
	case totalCreaturs >= 3 && totalCreaturs <= 6:
		multiplier = 2.0
	case totalCreaturs >= 7 && totalCreaturs <= 10:
		multiplier = 2.5
	case totalCreaturs >= 11 && totalCreaturs <= 14:
		multiplier = 3.0
	case totalCreaturs >= 15:
		multiplier = 4.0
	default:
		multiplier = 1.0
	}

	return multiplier
}
