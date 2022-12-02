package aoc22

import (
	"fmt"
	"sort"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec01a(ctx ch.AOContext) error {
	elves, err := elfCalories(ctx)
	if err != nil {
		return err
	}

	ctx.FinalAnswer.Print(elves[len(elves)-1])
	return nil
}

func Dec01b(ctx ch.AOContext) error {
	elves, err := elfCalories(ctx)
	if err != nil {
		return err
	}

	ctx.FinalAnswer.Print(elves[len(elves)-1] + elves[len(elves)-2] + elves[len(elves)-3])
	return nil
}

func elfCalories(ctx ch.AOContext) ([]int, error) {
	sections, err := ctx.DataSections("inputs/2022/dec01.txt")
	if err != nil {
		return nil, err
	}

	rv := []int{}

	for _, elf := range sections {
		total := 0
		for _, food := range elf {
			calories := 0
			fmt.Sscanf(food, "%d", &calories)
			total += calories
		}
		rv = append(rv, total)
	}

	sort.Ints(rv)
	return rv, nil
}
