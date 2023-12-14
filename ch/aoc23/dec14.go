package aoc23

import (
	"github.com/thijzert/advent-of-code/ch"
	"github.com/thijzert/advent-of-code/lib/image"
)

func Dec14a(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2023/dec14.txt")
	if err != nil {
		return nil, err
	}
	img := image.ReadImage(lines, dec14rocks)

	ctx.Printf("Rocks:\n%s", img)
	answer := dec14RollNorth(ctx, img)
	ctx.Printf("After rolling:\n%s", img)

	return answer, nil
}

var Dec14b ch.AdventFunc = nil

// func Dec14b(ctx ch.AOContext) (interface{}, error) {
// 	return nil, errNotImplemented
// }

func dec14RollNorth(ctx ch.AOContext, img *image.Image) int {
	rv := 0
	for x := 0; x < img.Width; x++ {
		next := 0
		for y := 0; y < img.Height; y++ {
			c := img.At(x, y)
			if c == 1 {
				// Roll this rock north
				img.Set(x, y, 0)
				img.Set(x, next, 1)
				rv += img.Height - next
				next++
			} else if c == 5 {
				// Square rock
				next = y + 1
			}
		}
	}
	return rv
}

func dec14rocks(c rune) int {
	if c == '#' {
		return 5
	} else if c == 'O' {
		return 1
	}
	return 0
}
