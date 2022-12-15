package aoc22

import (
	"github.com/thijzert/advent-of-code/ch"
	"github.com/thijzert/advent-of-code/lib/dijkstra"
)

func Dec12a(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2022/dec12.txt")
	if err != nil {
		return nil, err
	}

	valid := func(x0, y0, x1, y1, cost int) bool {
		return validHikeStep(lines, x0, y0, x1, y1)
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

	_, totalCost, err := dijkstra.ShortestPath(dijkstra.GridWalker(valid, final, start...))
	if err != nil {
		return nil, err
	}

	return totalCost, nil
}

func validHikeStep(lines []string, x0, y0, x1, y1 int) bool {
	if y1 < 0 || y1 >= len(lines) {
		return false
	}
	if x1 < 0 || x1 >= len(lines[y1]) {
		return false
	}

	f := lines[y0][x0]
	t := lines[y1][x1]
	if f == 'S' {
		f = 'a'
	}
	if t == 'E' {
		t = 'z'
	}

	return t <= (f + 1)
}

func Dec12b(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2022/dec12.txt")
	if err != nil {
		return nil, err
	}

	valid := func(x0, y0, x1, y1, cost int) bool {
		return validHikeStep(lines, x0, y0, x1, y1)
	}
	final := func(x, y int) bool {
		return lines[y][x] == 'E'
	}

	start := []int{}
	for y, line := range lines {
		for x, c := range line {
			if c == 'S' || c == 'a' {
				start = append(start, x, y)
			}
		}
	}

	_, totalCost, err := dijkstra.ShortestPath(dijkstra.GridWalker(valid, final, start...))
	if err != nil {
		return nil, err
	}

	return totalCost, nil
}
