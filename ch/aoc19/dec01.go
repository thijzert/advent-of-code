package aoc19

import (
	"github.com/thijzert/advent-of-code/ch"
)

func Dec01a(ctx ch.AOContext) (interface{}, error) {
	masses, err := ctx.DataAsInts("inputs/2019/dec01a.txt")
	if err != nil {
		return nil, err
	}

	s := 0
	for _, m := range masses {
		s += (m / 3) - 2
	}

	return s, nil
}

func Dec01b(ctx ch.AOContext) (interface{}, error) {
	masses, err := ctx.DataAsInts("inputs/2019/dec01a.txt")
	if err != nil {
		return nil, err
	}

	s := 0
	for _, m := range masses {
		for m > 0 {
			m = (m / 3) - 2
			if m > 0 {
				s += m
			}
		}
	}

	return s, nil
}
