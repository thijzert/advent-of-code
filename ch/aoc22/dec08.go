package aoc22

import (
	"github.com/thijzert/advent-of-code/ch"
)

func Dec08a(ctx ch.AOContext) error {
	forest, err := ctx.DataLines("inputs/2022/dec08.txt")
	if err != nil {
		return err
	}

	visibleTrees := 0
	for y := range forest {
		for x, height := range forest[y] {
			visTop, visBottom, visLeft, visRight := true, true, true, true
			for i, h := range forest[y] {
				if h >= height {
					if i > x {
						visRight = false
					} else if i < x {
						visLeft = false
					}
				}
			}
			for i := range forest {
				h := rune(forest[i][x])
				if h >= height && i != y {
					if i > y {
						visBottom = false
					} else if i < y {
						visTop = false
					}
				}
			}
			if visTop || visBottom || visLeft || visRight {
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

	bestTree := 0

	for y := range forest {
		if y == 0 || y == len(forest)-1 {
			continue
		}
		for x := range forest[y] {
			if x == 0 || x == len(forest[y])-1 {
				continue
			}
			height := forest[y][x]
			scenicScore := 1

			obstructed := false
			for i := 1; (x - i) >= 0; i++ {
				if forest[y][x-i] >= height {
					obstructed = true
					scenicScore *= i
					break
				}
			}
			if !obstructed {
				scenicScore *= x
			}

			obstructed = false
			for i := 1; (x + i) < len(forest[y]); i++ {
				if forest[y][x+i] >= height {
					obstructed = true
					scenicScore *= i
					break
				}
			}
			if !obstructed {
				scenicScore *= len(forest[y]) - x - 1
			}

			obstructed = false
			for i := 1; (y - i) >= 0; i++ {
				if forest[y-i][x] >= height {
					obstructed = true
					scenicScore *= i
					break
				}
			}
			if !obstructed {
				scenicScore *= y
			}

			obstructed = false
			for i := 1; (y + i) < len(forest); i++ {
				if forest[y+i][x] >= height {
					obstructed = true
					scenicScore *= i
					break
				}
			}
			if !obstructed {
				scenicScore *= len(forest) - y - 1
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
