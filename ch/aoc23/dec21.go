package aoc23

import (
	"github.com/thijzert/advent-of-code/ch"
	"github.com/thijzert/advent-of-code/lib/cube"
	"github.com/thijzert/advent-of-code/lib/image"
)

func Dec21a(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2023/dec21.txt")
	if err != nil {
		return nil, err
	}
	img := image.ReadImage(lines, image.Octothorpe)
	img.Default = -1

	reach := []cube.Point{}
	for y, line := range lines {
		for x, c := range line {
			if c == 'S' {
				reach = append(reach, cube.Point{x, y})
				img.Set(x, y, 2)
			}
		}
	}

	for step := 0; step < 64; step++ {
		newReach := []cube.Point{}
		for _, pos := range reach {
			img.Set(pos.X, pos.Y, 0)
			for _, add := range cube.Cardinal2D {
				np := pos.Add(add)
				if img.At(np.X, np.Y) == 0 {
					newReach = append(newReach, cube.Point{np.X, np.Y})
					img.Set(np.X, np.Y, 2)
				}
			}
		}
		reach = newReach
	}

	img.Default = 0
	ctx.Printf("Garden:\n%s", img)
	return len(reach), nil
}

var Dec21b ch.AdventFunc = nil

// func Dec21b(ctx ch.AOContext) (interface{}, error) {
// 	return nil, errNotImplemented
// }
