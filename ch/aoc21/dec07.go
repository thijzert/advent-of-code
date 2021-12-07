package aoc21

import (
	"github.com/thijzert/advent-of-code/ch"
)

func Dec07a(ctx ch.AOContext) error {
	exampleCrabs := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	best, dist := minimalMovementMean(exampleCrabs, abs)

	ctx.Printf("Example data: move towards %d; this will cost %d fuel", best, dist)

	crabs, err := ctx.DataAsIntLists("inputs/2021/dec07.txt")
	if err != nil {
		return err
	}

	best, dist = minimalMovementMean(crabs[0], abs)
	ctx.Printf("Actual data: move towards %d; this will cost %d fuel", best, dist)

	ctx.FinalAnswer.Print(dist)
	return nil
}

func minimalMovementMean(crabs []int, f func(int) int) (int, int) {
	iMin := min(crabs...)
	iMax := max(crabs...)

	best, dist := 0, 0
	for i := iMin; i < iMax; i++ {
		d := 0
		for _, v := range crabs {
			d += f(i - v)
		}
		if d < dist || i == 0 {
			dist = d
			best = i
		}
	}

	return best, dist
}

func Dec07b(ctx ch.AOContext) error {
	f := func(n int) int {
		n = abs(n)
		return n * (n + 1) / 2
	}

	exampleCrabs := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	best, dist := minimalMovementMean(exampleCrabs, f)

	ctx.Printf("Example data: move towards %d; this will cost %d fuel", best, dist)

	crabs, err := ctx.DataAsIntLists("inputs/2021/dec07.txt")
	if err != nil {
		return err
	}

	best, dist = minimalMovementMean(crabs[0], f)
	ctx.Printf("Actual data: move towards %d; this will cost %d fuel", best, dist)

	ctx.FinalAnswer.Print(dist)
	return nil
}
