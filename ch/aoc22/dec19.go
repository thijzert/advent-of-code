package aoc22

import (
	"fmt"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec19a(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2022/dec19.txt")
	if err != nil {
		return nil, err
	}

	for _, line := range lines {
		var i int
		var blp botRecipe
		_, err := fmt.Sscanf(line, "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.", &i, &blp[ORE][ORE], &blp[CLAY][ORE], &blp[OBSIDIAN][ORE], &blp[OBSIDIAN][CLAY], &blp[GEODE][ORE], &blp[GEODE][OBSIDIAN])
		if err != nil {
			return nil, err
		}
		ctx.Printf("recipe %d: %v", i, blp)

		var startingBots resourceState
		startingBots[ORE] = 1
		m := mostGeodes(blp, resourceState{}, startingBots, 14)
		ctx.Printf("recipe %d: can crack %d geodes", i, m)
	}

	return nil, errNotImplemented
}

const (
	ORE int = iota
	CLAY
	OBSIDIAN
	GEODE
	RESLENGTH
)

type resourceState [RESLENGTH]int
type botRecipe [RESLENGTH][RESLENGTH]int

func mostGeodes(recipe botRecipe, resources, bots resourceState, timeRemaining int) int {
	for i, n := range bots {
		resources[i] += n
	}
	if timeRemaining == 0 {
		return resources[GEODE]
	}

	max := mostGeodes(recipe, resources, bots, timeRemaining-1)

	for b, cost := range recipe {
		canBuild := true
		for i, n := range cost {
			if resources[i]-bots[i] < n {
				canBuild = false
			}
		}
		if !canBuild {
			continue
		}
		nr := resources
		for i, n := range cost {
			nr[i] -= n
		}
		nb := bots
		nb[b] += 1
		newMax := mostGeodes(recipe, nr, nb, timeRemaining-1)
		if newMax > max {
			max = newMax
		}
	}
	return max
}

var Dec19b ch.AdventFunc = nil

// func Dec19b(ctx ch.AOContext) (interface{}, error) {
// 	return nil, errNotImplemented
// }
