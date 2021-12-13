package aoc20

import (
	"fmt"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec17a(ctx ch.AOContext) error {
	layers := [][]string{{".#.", "..#", "###"}}
	cc := readConwayCube(layers)
	ctx.Printf("Initial state: %d", cc.Length())

	for i := 0; i < 6; i++ {
		cc.Iterate()
	}
	ctx.Printf("Example data: After stage 6: %d", cc.Length())

	layers, err := ctx.DataSections("inputs/2020/dec17.txt")
	if err != nil {
		return err
	}

	cc = readConwayCube(layers)
	for i := 0; i < 6; i++ {
		cc.Iterate()
		ctx.Printf("After stage %d: %d", i+1, cc.Length())
	}
	ctx.FinalAnswer.Print(cc.Length())
	return nil
}

func Dec17b(ctx ch.AOContext) error {
	layers := [][]string{{".#.", "..#", "###"}}
	cc := readConwayHypercube(layers)

	for i := 0; i < 6; i++ {
		cc.Iterate()
	}
	ctx.Printf("Example data: After stage 6: %d", cc.Length())

	layers, err := ctx.DataSections("inputs/2020/dec17.txt")
	if err != nil {
		return err
	}

	cc = readConwayHypercube(layers)
	for i := 0; i < 6; i++ {
		cc.Iterate()
		ctx.Printf("After stage %d: %d", i+1, cc.Length())
	}
	ctx.FinalAnswer.Print(cc.Length())
	return nil
}

type point3 struct {
	X, Y, Z int
}

type conwayCube struct {
	Min, Max point3
	Contents map[point3]bool
}

func readConwayCube(layers [][]string) *conwayCube {
	rv := &conwayCube{
		Contents: make(map[point3]bool),
	}

	rv.Max.Z = len(layers)
	for z, layer := range layers {
		rv.Max.Y = max(rv.Max.Y, len(layer))
		for y, line := range layer {
			rv.Max.X = max(rv.Max.X, len(line))
			for x, c := range line {
				if c == '#' {
					rv.Contents[point3{x, y, z}] = true
				}
			}
		}
	}

	return rv
}

func (cc *conwayCube) Iterate() {
	next := make(map[point3]bool)

	for z := cc.Min.Z - 1; z <= cc.Max.Z+1; z++ {
		for y := cc.Min.Y - 1; y <= cc.Max.Y+1; y++ {
			for x := cc.Min.X - 1; x <= cc.Max.X+1; x++ {
				n := cc.Neighbours(x, y, z)
				if n == 3 || (n == 2 && cc.Contents[point3{x, y, z}]) {
					next[point3{x, y, z}] = true
					cc.Min.X = min(cc.Min.X, x)
					cc.Min.Y = min(cc.Min.Y, y)
					cc.Min.Z = min(cc.Min.Z, z)
					cc.Max.X = max(cc.Max.X, x)
					cc.Max.Y = max(cc.Max.Y, y)
					cc.Max.Z = max(cc.Max.Z, z)
				}
			}
		}
	}

	cc.Contents = next
}

func (cc *conwayCube) Neighbours(x, y, z int) int {
	rv := 0
	for c := -1; c <= 1; c++ {
		for b := -1; b <= 1; b++ {
			for a := -1; a <= 1; a++ {
				if a != 0 || b != 0 || c != 0 {
					if cc.Contents[point3{x + a, y + b, z + c}] {
						rv++
					}
				}
			}
		}
	}
	return rv
}

func (cc *conwayCube) Length() int {
	rv := 0
	for _, v := range cc.Contents {
		if v {
			rv++
		}
	}
	return rv
}

type point4 struct {
	X, Y, Z, W int
}

type conwayHypercube struct {
	Min, Max point4
	Contents map[point4]bool
}

func readConwayHypercube(layers [][]string) *conwayHypercube {
	rv := &conwayHypercube{
		Contents: make(map[point4]bool),
	}

	rv.Max.W = 0
	rv.Max.Z = len(layers) - 1
	for z, layer := range layers {
		rv.Max.Y = max(rv.Max.Y, len(layer))
		for y, line := range layer {
			rv.Max.X = max(rv.Max.X, len(line))
			for x, c := range line {
				if c == '#' {
					rv.Contents[point4{x, y, z, 0}] = true
				}
			}
		}
	}

	return rv
}

func (cc *conwayHypercube) Iterate() {
	next := make(map[point4]bool)

	for w := cc.Min.W - 1; w <= cc.Max.W+1; w++ {
		for z := cc.Min.Z - 1; z <= cc.Max.Z+1; z++ {
			for y := cc.Min.Y - 1; y <= cc.Max.Y+1; y++ {
				for x := cc.Min.X - 1; x <= cc.Max.X+1; x++ {
					n := cc.Neighbours(x, y, z, w)
					if n == 3 || (n == 2 && cc.Contents[point4{x, y, z, w}]) {
						next[point4{x, y, z, w}] = true
						cc.Min.X = min(cc.Min.X, x)
						cc.Min.Y = min(cc.Min.Y, y)
						cc.Min.Z = min(cc.Min.Z, z)
						cc.Min.W = min(cc.Min.W, w)
						cc.Max.X = max(cc.Max.X, x)
						cc.Max.Y = max(cc.Max.Y, y)
						cc.Max.Z = max(cc.Max.Z, z)
						cc.Max.W = max(cc.Max.W, w)
					}
				}
			}
		}
	}

	cc.Contents = next
}

func (cc *conwayHypercube) Neighbours(x, y, z, w int) int {
	rv := 0
	for d := -1; d <= 1; d++ {
		for c := -1; c <= 1; c++ {
			for b := -1; b <= 1; b++ {
				for a := -1; a <= 1; a++ {
					if a != 0 || b != 0 || c != 0 || d != 0 {
						if cc.Contents[point4{x + a, y + b, z + c, w + d}] {
							rv++
						}
					}
				}
			}
		}
	}
	return rv
}

func (cc *conwayHypercube) Length() int {
	rv := 0
	for _, v := range cc.Contents {
		if v {
			rv++
		}
	}
	return rv
}

func (cc *conwayHypercube) String() string {
	rv := ""

	for w := cc.Min.W; w <= cc.Max.W; w++ {
		for z := cc.Min.Z; z <= cc.Max.Z; z++ {
			if w != cc.Min.W || z != cc.Min.Z {
				rv += "\n\n"
			}
			rv += fmt.Sprintf("z = %d, w = %d", z, w)
			for y := cc.Min.Y; y <= cc.Max.Y; y++ {
				rv += "\n"
				for x := cc.Min.X; x <= cc.Max.X; x++ {
					if cc.Contents[point4{x, y, z, w}] {
						rv += "#"
					} else {
						rv += "."
					}
				}
			}
		}
	}

	return rv
}
