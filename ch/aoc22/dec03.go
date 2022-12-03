package aoc22

import (
	"regexp"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec03a(ctx ch.AOContext) error {
	lines, err := ctx.DataLines("inputs/2022/dec03.txt")
	if err != nil {
		return err
	}
	// lines = []string{
	// 	"vJrwpWtwJgWrhcsFMMfFFhFp",
	// 	"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
	// 	"PmmdzqPrVvPwwTWBwg",
	// 	"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
	// 	"ttgJtRGJQctTZtZT",
	// 	"CrZsJsPPZsGzwwsLwLmpwMDw",
	// }

	sumPr := 0

	prio := func(b byte) int {
		if b >= 'a' && b <= 'z' {
			return int(1 + b - 'a')
		} else if b >= 'A' && b <= 'Z' {
			return int(27 + b - 'A')
		}
		return 0
	}

	for _, line := range lines {
		cutoff := len(line) / 2
		front, back := line[:cutoff], line[cutoff:]

		r, err := regexp.Compile("[" + front + "]")
		if err != nil {
			return err
		}
		m := r.FindString(back)
		ctx.Printf("both compartments contain '%s'", m)
		sumPr += prio(m[0])
	}

	ctx.FinalAnswer.Print(sumPr)
	return nil
}

func Dec03b(ctx ch.AOContext) error {
	return errNotImplemented
}
