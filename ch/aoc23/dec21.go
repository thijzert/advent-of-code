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
	start := []cube.Point{}
	for y, line := range lines {
		for x, c := range line {
			if c == 'S' {
				start = append(start, cube.Point{x, y})
			}
		}
	}
	answer, img := dec21BFSflood(lines, start, 64)
	ctx.Printf("Garden:\n%s", img)

	return answer, nil
}
func dec21BFSflood(lines []string, start []cube.Point, steps int) (int, *image.Image) {
	img := image.ReadImage(lines, image.Octothorpe)
	img.Default = -1

	reach := make(map[cube.Point]bool)
	for _, pt := range start {
		reach[pt] = true
		img.Set(pt.X, pt.Y, 2)
	}

	for step := 0; step < steps; step++ {
		newReach := make(map[cube.Point]bool)
		for pos := range reach {
			img.Set(pos.X, pos.Y, 0)
			for _, add := range cube.Cardinal2D {
				np := pos.Add(add)
				if img.At((np.X+img.Width)%img.Width, (np.Y+img.Height)%img.Height) == 0 {
					newReach[cube.Point{np.X, np.Y}] = true
					img.Set(np.X, np.Y, 2)
				}
			}
		}
		reach = newReach
	}

	img.Default = 0
	return len(reach), img
}

var Dec21b ch.AdventFunc = nil

// func Dec21b(ctx ch.AOContext) (interface{}, error) {
// 	return nil, errNotImplemented
// }
