package aoc22

import (
	"fmt"
	"strings"

	"github.com/thijzert/advent-of-code/ch"
	"github.com/thijzert/advent-of-code/lib/cube"
	"github.com/thijzert/advent-of-code/lib/image"
)

func Dec14a(ctx ch.AOContext) (interface{}, error) {
	return simulateSandCave(ctx, false)
}

func Dec14b(ctx ch.AOContext) (interface{}, error) {
	return simulateSandCave(ctx, true)
}

func simulateSandCave(ctx ch.AOContext, withFloor bool) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2022/dec14.txt")
	if err != nil {
		return nil, err
	}
	//lines = []string{"498,4 -> 498,6 -> 496,6", "503,4 -> 502,4 -> 502,9 -> 494,9"}

	offsetX := 330
	ymax := 2
	cave := image.NewImage(360, 200, func(int, int) int { return 0 })

	for _, line := range lines {
		coords := strings.Split(line, " -> ")
		start := cube.Point{}
		for i, coord := range coords {
			var p cube.Point
			fmt.Sscanf(coord, "%d,%d", &p.X, &p.Y)
			p.X -= offsetX

			if p.Y > ymax {
				ymax = p.Y
			}

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

	if withFloor {
		for x := 0; x < cave.Width; x++ {
			cave.Set(x, ymax+2, 1)
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
		if grain.Y == 0 {
			break
		}
	}

	ctx.Printf("cave: \n%s", cave)
	return grains, nil
}
