package aoc22

import (
	"github.com/thijzert/advent-of-code/ch"
	"github.com/thijzert/advent-of-code/lib/cube"
)

func Dec08a(ctx ch.AOContext) error {
	forest, err := ctx.DataLines("inputs/2022/dec08.txt")
	if err != nil {
		return err
	}
	bounds := cube.Square{cube.Interval{0, len(forest[0]) - 1}, cube.Interval{0, len(forest) - 1}}

	visibleTrees := 0
	for y := range forest {
		for x, height := range forest[y] {
			p := cube.Point{x, y}
			visible := false
			for _, dir := range cube.Cardinal2D {
				_, invisible := cube.Walk(p.Add(dir), dir, bounds, func(q cube.Point) bool {
					return rune(forest[q.Y][q.X]) >= height
				})
				visible = visible || (!invisible)
			}

			if visible {
				visibleTrees++
			}
		}
	}

	ctx.FinalAnswer.Print(visibleTrees)
	return nil
}

func Dec08b(ctx ch.AOContext) error {
	forest, err := ctx.DataLines("inputs/2022/dec08.txt")
	if err != nil {
		return err
	}
	//forest = []string{"30373", "25512", "65332", "33549", "35390"}
	bounds := cube.Square{cube.Interval{0, len(forest[0]) - 1}, cube.Interval{0, len(forest) - 1}}

	bestTree := 0

	for y := range forest {
		if y == 0 || y == len(forest)-1 {
			continue
		}
		for x, height := range forest[y] {
			if x == 0 || x == len(forest[y])-1 {
				continue
			}
			p := cube.Point{x, y}
			scenicScore := 1

			for _, dir := range cube.Cardinal2D {
				steps, _ := cube.Walk(p.Add(dir), dir, bounds, func(q cube.Point) bool {
					return rune(forest[q.Y][q.X]) >= height
				})
				scenicScore *= (steps + 1)
			}

			//ctx.Printf("Scenic score: %d", scenicScore)
			if scenicScore > bestTree {
				bestTree = scenicScore
			}
		}
	}

	ctx.FinalAnswer.Print(bestTree)
	return nil
}
