package aoc23

import (
	"fmt"

	"github.com/thijzert/advent-of-code/ch"
	"github.com/thijzert/advent-of-code/lib/cube"
	"github.com/thijzert/advent-of-code/lib/image"
)

func Dec18a(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2023/dec18.txt")
	if err != nil {
		return nil, err
	}

	bounds := cube.Square{}
	dig := cube.Point{}
	for _, l := range lines {
		dir, dist := "", 0
		fmt.Sscanf(l, "%s %d ", &dir, &dist)
		du, _ := cube.ParseDirection2D(dir)
		for i := 0; i < dist; i++ {
			dig = dig.Add(du)
			bounds = bounds.UpdatedBound(dig)
		}
	}
	ctx.Printf("Bounds: %v", bounds)

	img := image.NewImage(bounds.X.B-bounds.X.A+2, bounds.Y.B-bounds.Y.A+2, func(x, y int) int {
		return 0
	})
	dig = cube.Point{X: 1 - bounds.X.A, Y: 1 - bounds.Y.A}
	for _, l := range lines {
		dir, dist := "", 0
		fmt.Sscanf(l, "%s %d ", &dir, &dist)
		du, _ := cube.ParseDirection2D(dir)
		for i := 0; i < dist; i++ {
			dig = dig.Add(du)
			img.Set(dig.X, dig.Y, 1)
		}
	}
	img.FloodFill(2-bounds.X.A, -bounds.Y.A, 5)

	answer := 0
	for y := 0; y < img.Height; y++ {
		for x := 0; x < img.Width; x++ {
			if img.At(x, y) > 0 {
				answer++
			}
		}
	}

	ctx.Printf("Trench: \n%s", img)
	return answer, nil
}

var Dec18b ch.AdventFunc = nil

// func Dec18b(ctx ch.AOContext) (interface{}, error) {
// 	return nil, errNotImplemented
// }
