package aoc21

import (
	"fmt"

	"github.com/thijzert/advent-of-code/ch"
)

var Dec22b ch.AdventFunc = nil

func Dec22a(ctx ch.AOContext) error {
	lines, err := ctx.DataLines("inputs/2021/dec22.txt")
	if err != nil {
		return err
	}

	reactor := make([]bool, 101*101*101)
	for _, l := range lines {
		var on string
		var xmin, xmax, ymin, ymax, zmin, zmax int
		_, err := fmt.Sscanf(l, "%s x=%d..%d,y=%d..%d,z=%d..%d", &on, &xmin, &xmax, &ymin, &ymax, &zmin, &zmax)
		if err != nil {
			return err
		}

		for z := zmin; z <= zmax; z++ {
			if z < -50 || z > 50 {
				continue
			}
			for y := ymin; y <= ymax; y++ {
				if y < -50 || y > 50 {
					continue
				}
				for x := xmin; x <= xmax; x++ {
					if x < -50 || x > 50 {
						continue
					}
					reactor[(z+50)*101*101+(y+50)*101+(x+50)] = (on == "on")
				}
			}
		}
	}

	rv := 0
	for _, b := range reactor {
		if b {
			rv++
		}
	}

	ctx.FinalAnswer.Print(rv)
	return errNotImplemented
}

// func Dec22b(ctx ch.AOContext) error {
// 	return errNotImplemented
// }
