package aoc23

import (
	"strings"

	"github.com/thijzert/advent-of-code/ch"
)

func isSymbol(c byte) bool {
	return (c < '0' || c > '9') && (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') && c != '.' && c != ' '
}

func Dec03a(ctx ch.AOContext) (interface{}, error) {
	ilines, err := ctx.DataLines("inputs/2023/dec03.txt")
	if err != nil {
		return nil, err
	}
	lines := []string{strings.Repeat(".", len(ilines[0])+2)}
	for _, s := range ilines {
		lines = append(lines, "."+s+".")
	}
	lines = append(lines, lines[0])

	answer := 0
	for y, line := range lines {
		xst := 0
		partno := 0
		for x, c := range line {
			if c >= '0' && c <= '9' {
				if partno == 0 {
					xst = x - 1
				}
				partno = 10*partno + int(c-'0')
			} else if partno != 0 {
				ispart := isSymbol(line[xst]) || isSymbol(line[x])
				for i := xst; i <= x; i++ {
					ispart = ispart || isSymbol(lines[y-1][i]) || isSymbol(lines[y+1][i])
				}
				if ispart {
					ctx.Printf("Found part number %d", partno)
					answer += partno
				} else {
					ctx.Printf("  number %d is unconnected", partno)
				}
				partno = 0
			}
		}
	}

	return answer, nil
}

func Dec03b(ctx ch.AOContext) (interface{}, error) {
	ilines, err := ctx.DataLines("inputs/2023/dec03.txt")
	if err != nil {
		return nil, err
	}
	lines := []string{strings.Repeat(".", len(ilines[0])+2)}
	for _, s := range ilines {
		lines = append(lines, "."+s+".")
	}
	lines = append(lines, lines[0])

	gears := make(map[[2]int][]int)
	addGear := func(x, y, ratio int) {
		if lines[y][x] == '*' {
			k := [2]int{y, x}
			gears[k] = append(gears[k], ratio)
		}
	}

	for y, line := range lines {
		xst := 0
		teeth := 0
		for x, c := range line {
			if c >= '0' && c <= '9' {
				if teeth == 0 {
					xst = x - 1
				}
				teeth = 10*teeth + int(c-'0')
			} else if teeth != 0 {
				addGear(xst, y, teeth)
				addGear(x, y, teeth)
				for i := xst; i <= x; i++ {
					addGear(i, y-1, teeth)
					addGear(i, y+1, teeth)
				}
				teeth = 0
			}
		}
	}

	answer := 0
	for pos, ratios := range gears {
		if len(ratios) != 2 {
			ctx.Printf("Asterisk at %d,%d only has %d part numbers attached", pos[0], pos[1], len(ratios))
			continue
		}
		ctx.Printf("Gear at %d,%d has ratio %d×%d = %d", pos[0], pos[1], ratios[0], ratios[1], ratios[0]*ratios[1])
		answer += ratios[0] * ratios[1]
	}

	return answer, nil
}
