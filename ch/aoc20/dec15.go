package aoc20

import (
	"github.com/thijzert/advent-of-code/ch"
)

func Dec15a(ctx ch.AOContext) error {
	rmb := newRambunctius(0, 3, 6)
	c := 0
	for rmb.Turn < 10 {
		c = rmb.Step()
		ctx.Printf("Turn %d: %d", rmb.Turn, c)
	}
	for rmb.Turn < 2020 {
		c = rmb.Step()
	}
	ctx.Printf("Turn %d: %d", rmb.Turn, c)

	rmb = newRambunctius(2, 1, 10, 11, 0, 6)
	c = 0
	for rmb.Turn < 2020 {
		c = rmb.Step()
	}
	ctx.FinalAnswer.Print(c)
	return nil
}

func Dec15b(ctx ch.AOContext) error {
	rmb := newRambunctius(1, 3, 2)
	c := 0
	for rmb.Turn < 30000000 {
		c = rmb.Step()
	}
	ctx.Printf("Turn %d: %d", rmb.Turn, c)

	rmb = newRambunctius(2, 1, 10, 11, 0, 6)
	c = 0
	for rmb.Turn < 30000000 {
		c = rmb.Step()
	}
	ctx.FinalAnswer.Print(c)
	return nil
}

type rambunctius struct {
	Turn   int
	Last   int
	Memory map[int]int
}

func newRambunctius(setup ...int) *rambunctius {
	rv := &rambunctius{
		Memory: make(map[int]int),
	}

	for _, c := range setup {
		rv.Turn++
		rv.Memory[rv.Last] = rv.Turn - 1
		rv.Last = c
	}

	return rv
}

func (r *rambunctius) Step() int {

	rv := 0
	lp := r.Memory[r.Last]
	if lp != 0 {
		rv = r.Turn - lp
	}

	r.Turn++
	r.Memory[r.Last] = r.Turn - 1
	r.Last = rv

	return rv
}
