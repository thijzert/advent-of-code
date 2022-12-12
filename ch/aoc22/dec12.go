package aoc22

import (
	"github.com/thijzert/advent-of-code/ch"
	"github.com/thijzert/advent-of-code/lib/dijkstra"
)

func Dec12a(ctx ch.AOContext) error {
	lines, err := ctx.DataLines("inputs/2022/dec12.txt")
	if err != nil {
		return err
	}

	valid := func(x0, y0, x1, y1, cost int) bool {
		if y1 < 0 || y1 >= len(lines) {
			return false
		}
		if x1 < 0 || x1 >= len(lines[y1]) {
			return false
		}

		if lines[y1][x1] == 'a' && lines[y0][x0] == 'S' {
			return true
		}
		if lines[y1][x1] == 'E' && lines[y0][x0] < 'y' {
			return false
		}
		return lines[y1][x1] <= (lines[y0][x0] + 1)
	}
	final := func(x, y int) bool {
		return lines[y][x] == 'E'
	}

	start := []int{}
	for y, line := range lines {
		for x, c := range line {
			if c == 'S' {
				start = append(start, x, y)
			}
		}
	}
	b := positionList{dijkstra.GridWalkerEx(valid, final, start...)}

	_, totalCost, err := dijkstra.ShortestPath(b)
	if err != nil {
		return err
	}

	ctx.FinalAnswer.Print(totalCost)
	return nil
}

type positionList struct {
	start []dijkstra.Position
}

func (p positionList) StartingPositions() []dijkstra.Position {
	return p.start
}

func Dec12b(ctx ch.AOContext) error {
	return errNotImplemented
}
