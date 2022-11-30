package aoc19

import (
	"github.com/thijzert/advent-of-code/ch"
)

func Dec01a(ctx ch.AOContext) error {
	masses, err := ctx.DataAsInts("inputs/2019/dec01a.txt")
	if err != nil {
		return err
	}

	s := 0
	for _, m := range masses {
		s += (m / 3) - 2
	}

	ctx.FinalAnswer.Print(s)
	return nil
}

func Dec01b(ctx ch.AOContext) error {
	masses, err := ctx.DataAsInts("inputs/2019/dec01a.txt")
	if err != nil {
		return err
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

	ctx.FinalAnswer.Print(s)
	return errNotImplemented
}
