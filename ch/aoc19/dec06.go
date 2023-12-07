package aoc19

import (
	"strings"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec06a(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2019/dec06.txt")
	if err != nil {
		return nil, err
	}

	parent := make(map[string]string)
	for _, line := range lines {
		lpts := strings.Split(line, ")")
		parent[lpts[1]] = lpts[0]
	}

	totalIndirectOrbits := 0
	level := make(map[string]int)
	level["COM"] = 0
	done := false
	for !done {
		done = true
		for k, v := range parent {
			if _, ok := level[k]; !ok {
				done = false
				if l, ok := level[v]; ok {
					level[k] = l + 1
					totalIndirectOrbits += l + 1
				}
			}
		}
	}
	ctx.Printf("Levels: %v", level)

	return totalIndirectOrbits, nil
}

var Dec06b ch.AdventFunc = nil

// func Dec06b(ctx ch.AOContext) (interface{}, error) {
// 	return nil, errNotImplemented
// }
