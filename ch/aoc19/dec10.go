package aoc19

import (
	"github.com/thijzert/advent-of-code/ch"
	"github.com/thijzert/advent-of-code/lib/image"
)

func Dec10a(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2019/dec10.txt")
	if err != nil {
		return nil, err
	}
	lines = []string{"......#.#.", "#..#.#....", "..#######.", ".#.#.###..", ".#..#.....", "..#....#.#", "#..#....#.", ".##.#..###", "##...#..#.", ".#....####"}
	img := image.ReadImage(lines, image.Octothorpe)

	ctx.Printf("Asteroid belt:\n%s", img)
	bestX, bestY, bestAst := 0, 0, 0
	for y := range lines {
		for x := range lines[0] {
			for b := range lines {
				for a := range lines[0] {
					dx, dy := a-x, b-y
					if (dx == 0 && dy == 0) || img.At(a, b) == 0 {
						continue
					}
					for c := range lines[2:] {
						aa, bb := x+(c+2)*dx, y+(c+2)*dy
						img.Set(aa, bb, -img.At(aa, bb))
					}
				}
			}
			ast := 0
			for b := range lines {
				for a := range lines[0] {
					if a == x && b == y {
						continue
					}
					v := img.At(a, b)
					if v == -1 {
						img.Set(a, b, 1)
					} else {
						ast += v
					}
				}
			}
			if ast > bestAst {
				bestX, bestY, bestAst = x, y, ast
			}
		}
	}

	ctx.Printf("Best location is %d,%d with %d asteroids detected", bestX, bestY, bestAst)

	return bestAst, nil
}

var Dec10b ch.AdventFunc = nil

// func Dec10b(ctx ch.AOContext) (interface{}, error) {
// 	return nil, errNotImplemented
// }
