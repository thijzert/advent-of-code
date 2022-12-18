package aoc22

import (
	"fmt"

	"github.com/thijzert/advent-of-code/ch"
	"github.com/thijzert/advent-of-code/lib/cube"
)

func Dec18a(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2022/dec18.txt")
	if err != nil {
		return nil, err
	}
	//lines = []string{"2,2,2", "1,2,2", "3,2,2", "2,1,2", "2,3,2", "2,2,1", "2,2,3", "2,2,4", "2,2,6", "1,2,5", "3,2,5", "2,1,5", "2,3,5"}

	droplets := make(map[cube.Point3]bool)
	for _, line := range lines {
		var pt cube.Point3
		fmt.Sscanf(line, "%d,%d,%d", &pt.X, &pt.Y, &pt.Z)
		if droplets[pt] {
			ctx.Printf("duplicate droplet %s", line)
			return nil, errFailed
		}
		droplets[pt] = true
	}

	rv := 6 * len(droplets)
	for pt := range droplets {
		for _, dir := range cube.Cardinal3D {
			if droplets[pt.Add(dir)] {
				rv--
			}
		}
	}

	return rv, nil
}

var Dec18b ch.AdventFunc = nil

// func Dec18b(ctx ch.AOContext) (interface{}, error) {
// 	return nil, errNotImplemented
// }
