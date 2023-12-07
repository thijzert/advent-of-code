package aoc19

import (
	"strings"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec6OrbitMap(ctx ch.AOContext) (parent map[string]string, level map[string]int, err error) {
	lines, err := ctx.DataLines("inputs/2019/dec06.txt")
	if err != nil {
		return nil, nil, err
	}

	parent = make(map[string]string)
	for _, line := range lines {
		lpts := strings.Split(line, ")")
		parent[lpts[1]] = lpts[0]
	}

	totalIndirectOrbits := 0
	level = make(map[string]int)
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

	return parent, level, nil
}

func Dec06a(ctx ch.AOContext) (interface{}, error) {
	_, level, err := Dec6OrbitMap(ctx)
	if err != nil {
		return nil, err
	}

	totalIndirectOrbits := 0
	for _, l := range level {
		totalIndirectOrbits += l
	}

	return totalIndirectOrbits, nil
}

func Dec06b(ctx ch.AOContext) (interface{}, error) {
	parent, level, err := Dec6OrbitMap(ctx)
	if err != nil {
		return nil, err
	}

	// Find common ancestor
	santas := make(map[string]bool)
	s := "SAN"
	for s != "COM" {
		santas[s] = true
		s = parent[s]
	}
	s = "YOU"
	for s != "COM" {
		if santas[s] {
			break
		}
		s = parent[s]
	}

	return level["SAN"] + level["YOU"] - 2*level[s] - 2, nil
}
