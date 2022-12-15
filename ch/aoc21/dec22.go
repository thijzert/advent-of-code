package aoc21

import (
	"fmt"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec22a(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2021/dec22.txt")
	if err != nil {
		return nil, err
	}

	reactor := make([]bool, 101*101*101)
	for _, l := range lines {
		var on string
		var xmin, xmax, ymin, ymax, zmin, zmax int
		_, err := fmt.Sscanf(l, "%s x=%d..%d,y=%d..%d,z=%d..%d", &on, &xmin, &xmax, &ymin, &ymax, &zmin, &zmax)
		if err != nil {
			return nil, err
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

	return rv, nil
}

func Dec22b(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2021/dec22.txt")
	if err != nil {
		return nil, err
	}

	reactor := &cuboidSet{}

	for _, l := range lines {
		var on string
		var c cuboid
		_, err := fmt.Sscanf(l, "%s x=%d..%d,y=%d..%d,z=%d..%d", &on, &c[0].min, &c[0].max, &c[1].min, &c[1].max, &c[2].min, &c[2].max)
		if err != nil {
			return nil, err
		}

		if on == "on" {
			reactor.Add(c)
		} else {
			reactor.Remove(c)
		}
	}

	return reactor.Size(), nil
}

type interval struct {
	min, max int
}

func (a interval) Intersect(b interval) (int, int, int, int) {
	if a.min < b.min {
		if b.max < a.max {
			return a.min, b.min, b.max, a.max
		} else {
			return a.min, b.min, a.max, b.max
		}
	} else {
		if a.max < b.max {
			return b.min, a.min, a.max, b.max
		} else {
			return b.min, a.min, b.max, a.max
		}
	}
}

func (a interval) Contains(b interval) bool {
	return a.min <= b.min && a.max >= b.max
}

type cuboid [3]interval

func (c cuboid) Size() int {
	rv := 1

	for _, in := range c {
		if in.max < in.min {
			return 0
		}
		rv *= in.max - in.min + 1
	}

	return rv
}

func (c cuboid) String() string {
	return fmt.Sprintf("x=%d..%d,y=%d..%d,z=%d..%d", c[0].min, c[0].max, c[1].min, c[1].max, c[2].min, c[2].max)
}

type cuboidSet struct {
	Cubes []cuboid
}

func (cs *cuboidSet) Add(c cuboid) {
	if c.Size() == 0 {
		return
	}

	for j, b := range cs.Cubes {
		overlap := true
		for i, in := range c {
			if in.max < b[i].min || in.min > b[i].max {
				// Cuboids are completely distinct
				overlap = false
			}
		}
		if !overlap {
			continue
		}

		for i, in := range c {
			if in == b[i] {
				continue
			}

			// Remove the old cuboid
			cs.Cubes = append(cs.Cubes[:j], cs.Cubes[j+1:]...)

			p, q, r, s := b[i].Intersect(in)
			bs := [3]cuboid{b, b, b}
			bs[0][i] = interval{p, q - 1}
			bs[1][i] = interval{q, r}
			bs[2][i] = interval{r + 1, s}

			for _, bb := range bs {
				if b[i].Contains(bb[i]) {
					cs.Add(bb)
				}
			}

			p, q, r, s = in.Intersect(b[i])
			ccs := [3]cuboid{c, c, c}
			ccs[0][i] = interval{p, q - 1}
			ccs[1][i] = interval{q, r}
			ccs[2][i] = interval{r + 1, s}
			for _, cc := range ccs {
				if in.Contains(cc[i]) {
					cs.Add(cc)
				}
			}
			return
		}

		// Cuboids are identical
		return
	}

	cs.Cubes = append(cs.Cubes, c)
}

func (cs *cuboidSet) Remove(c cuboid) {
	for j, b := range cs.Cubes {
		overlap := true
		for i, in := range c {
			if in.max < b[i].min || in.min > b[i].max {
				// Cuboids are completely distinct
				overlap = false
			}
		}
		if !overlap {
			continue
		}

		for i, in := range c {
			if in == b[i] {
				continue
			}

			// Remove the old cuboid
			cs.Cubes = append(cs.Cubes[:j], cs.Cubes[j+1:]...)

			p, q, r, s := b[i].Intersect(in)
			bs := [3]cuboid{b, b, b}
			bs[0][i] = interval{p, q - 1}
			bs[1][i] = interval{q, r}
			bs[2][i] = interval{r + 1, s}

			for _, bb := range bs {
				if b[i].Contains(bb[i]) {
					cs.Add(bb)
				}
			}

			p, q, r, s = in.Intersect(b[i])
			ccs := [3]cuboid{c, c, c}
			ccs[0][i] = interval{p, q - 1}
			ccs[1][i] = interval{q, r}
			ccs[2][i] = interval{r + 1, s}
			for _, cc := range ccs {
				if in.Contains(cc[i]) {
					cs.Remove(cc)
				}
			}
			return
		}

		// Cuboids are identical
		cs.Cubes = append(cs.Cubes[:j], cs.Cubes[j+1:]...)
		return
	}
}

func (cs *cuboidSet) Size() int {
	rv := 0
	for _, c := range cs.Cubes {
		rv += c.Size()
	}
	return rv
}
