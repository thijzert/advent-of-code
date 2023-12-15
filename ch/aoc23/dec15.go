package aoc23

import (
	"strings"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec15a(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2023/dec15.txt")
	if err != nil {
		return nil, err
	}
	elems := strings.Split(lines[0], ",")

	answer := 0
	for _, elem := range elems {
		h := byte(0)
		for _, c := range elem {
			h = 17 * (h + byte(c))
		}
		answer += int(h)
	}

	return answer, nil
}

var Dec15b ch.AdventFunc = nil

// func Dec15b(ctx ch.AOContext) (interface{}, error) {
// 	return nil, errNotImplemented
// }
