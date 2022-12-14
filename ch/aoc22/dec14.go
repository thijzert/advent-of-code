package aoc22

import (
	"fmt"
	"strings"

	"github.com/thijzert/advent-of-code/ch"
	"github.com/thijzert/advent-of-code/lib/cube"
	"github.com/thijzert/advent-of-code/lib/image"
)

func Dec14a(ctx ch.AOContext) error {
	lines, err := ctx.DataLines("inputs/2022/dec14.txt")
	if err != nil {
		return err
	}
	//lines = []string{"498,4 -> 498,6 -> 496,6", "503,4 -> 502,4 -> 502,9 -> 494,9"}

	offsetX := 440
	cave := image.NewImage(200, 200, func(int, int) int { return 0 })

	for _, line := range lines {
		coords := strings.Split(line, " -> ")
		start := cube.Point{}
		for i, coord := range coords {
			var p cube.Point
			fmt.Sscanf(coord, "%d,%d", &p.X, &p.Y)
			p.X -= offsetX
			if i == 0 {
				start = p
				continue
			}

			xrf, xrt := start.X, p.X
			if start.X > p.X {
				xrf, xrt = p.X, start.X
			}
			for x := xrf; x <= xrt; x++ {
				cave.Set(x, start.Y, 1)
			}

			yrf, yrt := start.Y, p.Y
			if start.Y > p.Y {
				yrf, yrt = p.Y, start.Y
			}
			for y := yrf; y <= yrt; y++ {
				cave.Set(p.X, y, 1)
			}

			start = p
		}
	}

	grains := 0
	for {
		grain := cube.Point{500 - offsetX, 0}
		for {
			if grain.Y > 200 {
				break
			} else if cave.At(grain.X, grain.Y+1) == 0 {
				grain.Y += 1
			} else if cave.At(grain.X-1, grain.Y+1) == 0 {
				grain.X -= 1
				grain.Y += 1
			} else if cave.At(grain.X+1, grain.Y+1) == 0 {
				grain.X += 1
				grain.Y += 1
			} else {
				break
			}
		}
		if grain.Y > 200 {
			break
		}
		grains++
		cave.Set(grain.X, grain.Y, 5)
	}

	ctx.Printf("cave: %s", cave)
	ctx.FinalAnswer.Print(grains)
	return errNotImplemented
}

var Dec14b ch.AdventFunc = nil

// func Dec14b(ctx ch.AOContext) error {
// 	return errNotImplemented
// }
