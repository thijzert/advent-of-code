package aoc22

import (
	"github.com/thijzert/advent-of-code/ch"
)

func Dec08a(ctx ch.AOContext) error {
	forest, err := ctx.DataLines("inputs/2022/dec08.txt")
	if err != nil {
		return err
	}
	//forest = []string{"30373", "25512", "65332", "33549", "35390"}

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
	return errNotImplemented
}

var Dec08b ch.AdventFunc = nil

// func Dec08b(ctx ch.AOContext) error {
// 	return errNotImplemented
// }
