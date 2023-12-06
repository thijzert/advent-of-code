package aoc23

import (
	"strings"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec06a(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2023/dec06.txt")
	if err != nil {
		return nil, err
	}
	races := make([][]int, len(lines))
	for i, line := range lines {
		line = strings.Split(line, ":")[1]
		for _, s := range strings.Split(line, " ") {
			if s != "" {
				races[i] = append(races[i], atoid(s, 0))
			}
		}
	}
	//races = [][]int{{7, 15, 30}, {9, 40, 200}}

	answer := 1

	for i, tmax := range races[0] {
		dmin := races[1][i]

		found := false
		f, l := 0, 0

		for t := 1; t < tmax; t++ {
			dist := t * (tmax - t)
			if dist > dmin {
				if !found {
					f = t
				}
				found = true
				l = t
			}
		}
		ctx.Printf("In race %d, you can win by charging at least %d ms and at most %d ms", i+1, f, l)
		answer *= (l - f + 1)
	}

	return answer, nil
}

func Dec06b(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2023/dec06.txt")
	if err != nil {
		return nil, err
	}

	tmax := atoid(strings.ReplaceAll(strings.Split(lines[0], ":")[1], " ", ""), 0)
	dmin := atoid(strings.ReplaceAll(strings.Split(lines[1], ":")[1], " ", ""), 0)

	found := false
	f, l := 0, 0

	for t := 1; t < tmax; t++ {
		dist := t * (tmax - t)
		if dist > dmin {
			if !found {
				f = t
			}
			found = true
			l = t
		}
	}
	ctx.Printf("In this race, you can win by charging at least %d ms and at most %d ms", f, l)
	return (l - f + 1), nil
}
