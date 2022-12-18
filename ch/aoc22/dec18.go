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

func Dec18b(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2022/dec18.txt")
	if err != nil {
		return nil, err
	}
	//lines = []string{"2,2,2", "1,2,2", "3,2,2", "2,1,2", "2,3,2", "2,2,1", "2,2,3", "2,2,4", "2,2,6", "1,2,5", "3,2,5", "2,1,5", "2,3,5"}

	const LAVA int = 1
	const STEAM = 2

	var boundingBox cube.Cube
	droplets := make(map[cube.Point3]int)
	for i, line := range lines {
		var pt cube.Point3
		fmt.Sscanf(line, "%d,%d,%d", &pt.X, &pt.Y, &pt.Z)
		if i == 0 {
			boundingBox = cube.Cube{X: cube.Interval{pt.X, pt.X}, Y: cube.Interval{pt.Y, pt.Y}, Z: cube.Interval{pt.Z, pt.Z}}
		}
		droplets[pt] = LAVA
		boundingBox = boundingBox.UpdatedBound(pt)
	}

	queue := make([]cube.Point3, 0, boundingBox.Volume())
	ctx.Printf("Bounding box: %v", boundingBox)
	for x := boundingBox.X.A - 2; x <= boundingBox.X.B+2; x++ {
		for y := boundingBox.Y.A - 2; y <= boundingBox.Y.B+2; y++ {
			for z := boundingBox.Z.A - 2; z <= boundingBox.Z.B+2; z++ {
				pt := cube.Point3{x, y, z}
				if x == boundingBox.X.A-2 || x == boundingBox.X.B+2 ||
					y == boundingBox.Y.A-2 || y == boundingBox.Y.B+2 ||
					z == boundingBox.Z.A-2 || z == boundingBox.Z.B+2 {
					droplets[pt] = STEAM
				} else if x == boundingBox.X.A-1 || x == boundingBox.X.B+1 ||
					y == boundingBox.Y.A-1 || y == boundingBox.Y.B+1 ||
					z == boundingBox.Z.A-1 || z == boundingBox.Z.B+1 {
					queue = append(queue, pt)
				}
			}
		}
	}
	for len(queue) > 0 {
		pt := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if droplets[pt] != 0 {
			continue
		}
		droplets[pt] = STEAM
		for _, dir := range cube.Cardinal3D {
			pt1 := pt.Add(dir)
			if droplets[pt1] == 0 {
				queue = append(queue, pt1)
			}
		}
	}

	rv := 0
	for pt, c := range droplets {
		if c != LAVA {
			continue
		}
		for _, dir := range cube.Cardinal3D {
			if droplets[pt.Add(dir)] == STEAM {
				rv++
			}
		}
	}

	return rv, nil
}
