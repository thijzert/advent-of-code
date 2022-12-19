package aoc22

import (
	"fmt"
	"sync"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec19a(ctx ch.AOContext) (interface{}, error) {
	recipes, err := readBotRecipes(ctx, "inputs/2022/dec19.txt")
	if err != nil {
		return nil, err
	}

	ipts, opts := make(chan int), make(chan int)
	go func() {
		rv := 0
		for i := range ipts {
			rv += i
		}
		opts <- rv
	}()

	var wg sync.WaitGroup
	for i, blp := range recipes {
		wg.Add(1)
		go func(i int, blp botRecipe) {
			var startingBots resourceState
			startingBots[ORE] = 1
			m := mostGeodes(blp, resourceState{}, startingBots, 24)
			ctx.Printf("recipe %d: can crack %d geodes; quality=%d", i, m, i*m[GEODE])
			ipts <- i * m[GEODE]
			wg.Done()
		}(i+1, blp)
	}
	wg.Wait()
	close(ipts)

	return <-opts, nil
}

func Dec19b(ctx ch.AOContext) (interface{}, error) {
	recipes, err := readBotRecipes(ctx, "inputs/2022/dec19.txt")
	if err != nil {
		return nil, err
	}
	if len(recipes) > 3 {
		recipes = recipes[:3]
	}

	ipts, opts := make(chan int), make(chan int)
	go func() {
		rv := 1
		for i := range ipts {
			rv *= i
		}
		opts <- rv
	}()

	var wg sync.WaitGroup
	for i, blp := range recipes {
		wg.Add(1)
		go func(i int, blp botRecipe) {
			var startingBots resourceState
			startingBots[ORE] = 1
			m := mostGeodes(blp, resourceState{}, startingBots, 32)
			ctx.Printf("recipe %d: can crack %d geodes", i+1, m)
			ipts <- m[GEODE]
			wg.Done()
		}(i, blp)
	}
	wg.Wait()
	close(ipts)

	return <-opts, nil
}

func readBotRecipes(ctx ch.AOContext, name string) ([]botRecipe, error) {
	lines, err := ctx.DataLines(name)
	if err != nil {
		return nil, err
	}

	rv := []botRecipe{}
	for _, line := range lines {
		var i int
		var blp botRecipe
		_, err := fmt.Sscanf(line, "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.", &i, &blp[ORE][ORE], &blp[CLAY][ORE], &blp[OBSIDIAN][ORE], &blp[OBSIDIAN][CLAY], &blp[GEODE][ORE], &blp[GEODE][OBSIDIAN])
		if err != nil {
			return nil, err
		}

		for i := 0; i < RESLENGTH; i++ {
			for j := 0; j < RESLENGTH; j++ {
				if blp[RESLENGTH][j] < blp[i][j] && i != j {
					blp[RESLENGTH][j] = blp[i][j]
				}
			}
		}
		ctx.Printf("recipe %d: %v", i, blp)
		rv = append(rv, blp)
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

type botRecipe [RESLENGTH + 1][RESLENGTH]int

func mostGeodes(recipe botRecipe, resources, bots resourceState, timeRemaining int) resourceState {
	nr := resources
	for i, n := range bots {
		nr[i] += n * timeRemaining
	}
	max := nr

	maxBots := recipe[RESLENGTH]
	//for b, cost := range recipe {
	for b := RESLENGTH - 1; b >= 0; b-- {
		if maxBots[b] > 0 && bots[b] >= maxBots[b] {
			continue
		}
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
