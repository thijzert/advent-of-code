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

	rv := 0
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
		m := mostGeodes(blp, resourceState{}, startingBots, 24)
		ctx.Printf("recipe %d: can crack %d geodes; quality=%d", i, m, i*m[GEODE])
		rv += i * m[GEODE]
	}

	return rv, nil
}

const (
	ORE int = iota
	CLAY
	OBSIDIAN
	GEODE
	RESLENGTH
)

type resourceState [RESLENGTH]int

func (a resourceState) more(b resourceState) bool {
	for i := RESLENGTH - 1; i >= 0; i-- {
		if a[i] > b[i] {
			return true
		} else if a[i] < b[i] {
			return false
		}
	}
	return false
}

type botRecipe [RESLENGTH][RESLENGTH]int

func mostGeodes(recipe botRecipe, resources, bots resourceState, timeRemaining int) resourceState {
	nr := resources
	for i, n := range bots {
		nr[i] += n * timeRemaining
	}
	max := nr

	//for b, cost := range recipe {
	for b := RESLENGTH - 1; b >= 0; b-- {
		cost := recipe[b]
		steps := 0
		canBuild := true
		for i, n := range cost {
			n = n - resources[i]
			if n > 0 && bots[i] == 0 {
				canBuild = false
			} else if n > 0 {
				s := (n + bots[i] - 1) / bots[i]
				if s > steps {
					steps = s
				}
			}
		}
		steps++
		if !canBuild || steps > timeRemaining {
			continue
		}
		nr := resources
		for i, n := range cost {
			nr[i] += steps*bots[i] - n
		}
		nb := bots
		nb[b] += 1
		newMax := mostGeodes(recipe, nr, nb, timeRemaining-steps)
		if newMax.more(max) {
			max = newMax
		}
	}
	return max
}

var Dec19b ch.AdventFunc = nil

// func Dec19b(ctx ch.AOContext) (interface{}, error) {
// 	return nil, errNotImplemented
// }
