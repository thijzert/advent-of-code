package aoc21

import (
	"errors"
	"fmt"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec05a(ctx ch.AOContext) error {
	lines, err := ctx.DataLines("inputs/2021/dec05.txt")
	if err != nil {
		return err
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
			return err
		}

		inc.X = signum(b.X - a.X)
		inc.Y = signum(b.Y - a.Y)

		if inc.X != 0 && inc.Y != 0 {
			if (b.X-a.X)*(b.X-a.X) != (b.Y-a.Y)*(b.Y-a.Y) {
				return fmt.Errorf("vent '%s' isn't even diagonal", l)
			}

			continue
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

	ctx.FinalAnswer.Print(rv)
	return nil
}

func Dec05b(ctx ch.AOContext) error {
	return errors.New("not implemented")
}

type point struct {
	X, Y int
}

func (p point) Add(b point) point {
	p.X += b.X
	p.Y += b.Y
	return p
}

func signum(a int) int {
	if a < 0 {
		return -1
	} else if a > 0 {
		return 1

	}
	return 0
}
