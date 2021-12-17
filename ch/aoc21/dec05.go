package aoc21

import (
	"fmt"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec05a(ctx ch.AOContext) error {
	rv, err := dec05(ctx, false)
	if err == nil {
		ctx.FinalAnswer.Print(rv)
	}
	return err
}

func Dec05b(ctx ch.AOContext) error {
	rv, err := dec05(ctx, true)
	if err == nil {
		ctx.FinalAnswer.Print(rv)
	}
	return err
}

func dec05(ctx ch.AOContext, includeDiagonalVents bool) (int, error) {
	lines, err := ctx.DataLines("inputs/2021/dec05.txt")
	if err != nil {
		return 0, err
	}

	pointsSeen := make(map[point]int)
	for _, l := range lines {
		if l == "" {
			continue
		}

		var a, b point
		var inc point
		_, err = fmt.Sscanf(l, "%d,%d -> %d,%d", &a.X, &a.Y, &b.X, &b.Y)
		if err != nil {
			return 0, err
		}

		inc.X = signum(b.X - a.X)
		inc.Y = signum(b.Y - a.Y)

		if inc.X != 0 && inc.Y != 0 {
			if (b.X-a.X)*(b.X-a.X) != (b.Y-a.Y)*(b.Y-a.Y) {
				return 0, fmt.Errorf("vent '%s' isn't even diagonal", l)
			}

			if !includeDiagonalVents {
				continue
			}
		}

		for a != b {
			pointsSeen[a]++
			a = a.Add(inc)
		}
		pointsSeen[a]++
	}

	rv := 0
	for _, count := range pointsSeen {
		if count > 1 {
			rv++
		}
	}

	return rv, nil
}
