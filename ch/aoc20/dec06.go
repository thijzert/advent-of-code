package aoc20

import "github.com/thijzert/advent-of-code/ch"

func Dec06a(ctx ch.AOContext) error {
	groups, err := ctx.DataSections("inputs/2020/dec06.txt")
	if err != nil {
		return err
	}

	rv := 0
	for _, group := range groups {
		yes := make(map[rune]bool)
		for _, traveler := range group {
			for _, c := range traveler {
				yes[c] = true
			}
		}
		rv += len(yes)
	}

	ctx.FinalAnswer.Print(rv)
	return nil
}

func Dec06b(ctx ch.AOContext) error {
	groups, err := ctx.DataSections("inputs/2020/dec06.txt")
	if err != nil {
		return err
	}

	rv := 0
	for _, group := range groups {
		yes := make(map[rune]int)
		for _, traveler := range group {
			for _, c := range traveler {
				yes[c]++
			}
		}
		for _, ct := range yes {
			if ct == len(group) {
				rv++
			}
		}
	}

	ctx.FinalAnswer.Print(rv)
	return errNotImplemented
}
