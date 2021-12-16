package aoc21

import (
	"fmt"
	"sort"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec09a(ctx ch.AOContext) error {
	heights, err := getHeightmap(ctx, "inputs/2021/dec09.txt")
	if err != nil {
		return err
	}

	rv := 0
	for y := 0; y < heights.Length; y++ {
		for x := 0; x < heights.Width; x++ {
			rv += heights.RiskAt(x, y)
		}
	}

	ctx.FinalAnswer.Print(rv)
	return nil
}

func Dec09b(ctx ch.AOContext) error {
	heights, err := getHeightmap(ctx, "inputs/2021/dec09.txt")
	if err != nil {
		return err
	}

	var basins []int
	for y := 0; y < heights.Length; y++ {
		for x := 0; x < heights.Width; x++ {
			if heights.RiskAt(x, y) > 0 {
				basins = append(basins, heights.BasinSize(x, y))
			}
		}
	}

	sort.Ints(basins)

	ctx.Print(basins)
	rv := 1
	for _, sz := range basins[len(basins)-3:] {
		rv *= sz
	}

	ctx.FinalAnswer.Print(rv)
	return nil
}

type heightmap struct {
	Width, Length int
	Heights       []int
}

func getHeightmap(ctx ch.AOContext, assetName string) (heightmap, error) {
	var rv heightmap
	lines, err := ctx.DataLines(assetName)
	if err != nil {
		return rv, err
	}

	for _, s := range lines {
		if s != "" {
			rv.Length++
			if rv.Width != 0 {
				if rv.Width != len(s) {
					return rv, fmt.Errorf("inconsistent width")
				}
			}
			rv.Width = len(s)
		}
	}

	rv.Heights = make([]int, rv.Width*rv.Length)
	i := 0
	for _, s := range lines {
		if s == "" {
			continue
		}
		for j, h := range s {
			rv.Heights[rv.Width*i+j] = int(h - '0')
		}
		i++
	}

	return rv, nil
}

func (h heightmap) At(x, y int) int {
	if x < 0 || x >= h.Width || y < 0 || y >= h.Length {
		return 10
	}
	return h.Heights[h.Width*y+x]
}

func (h heightmap) RiskAt(x, y int) int {
	self := h.At(x, y)
	if self >= h.At(x-1, y) || self >= h.At(x+1, y) {
		return 0
	}
	if self >= h.At(x, y-1) || self >= h.At(x, y+1) {
		return 0
	}

	return 1 + self
}

func (h heightmap) BasinSize(x, y int) int {
	self := h.At(x, y)
	rv := 1

	if b := h.At(x-1, y); b > self && b < 9 {
		h.Heights[h.Width*y+x-1] = self
		rv += h.BasinSize(x-1, y)
	}
	if b := h.At(x+1, y); b > self && b < 9 {
		h.Heights[h.Width*y+x+1] = self
		rv += h.BasinSize(x+1, y)
	}
	if b := h.At(x, y-1); b > self && b < 9 {
		h.Heights[h.Width*(y-1)+x] = self
		rv += h.BasinSize(x, y-1)
	}
	if b := h.At(x, y+1); b > self && b < 9 {
		h.Heights[h.Width*(y+1)+x] = self
		rv += h.BasinSize(x, y+1)
	}

	return rv
}
