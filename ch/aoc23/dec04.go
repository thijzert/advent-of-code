package aoc23

import (
	"strconv"
	"strings"

	"github.com/thijzert/advent-of-code/ch"
)

func atoid(s string, def int) int {
	rv, err := strconv.Atoi(s)
	if err != nil {
		return def
	}
	return rv
}

func Dec04a(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2023/dec04.txt")
	if err != nil {
		return nil, err
	}

	winnings := 0
	for i, line := range lines {
		contents := strings.Split(strings.Split(line, ":")[1], "|")
		winning := make(map[int]bool)
		for _, s := range strings.Split(contents[0], " ") {
			if s != "" {
				winning[atoid(s, -1)] = true
			}
		}

		worth := 1
		for _, s := range strings.Split(contents[1], " ") {
			if s != "" && winning[atoid(s, 0)] {
				worth <<= 1
			}
		}
		ctx.Printf("Card %d is worth %d points", i+1, worth>>1)
		winnings += worth >> 1
	}

	return winnings, nil
}

func Dec04b(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2023/dec04.txt")
	if err != nil {
		return nil, err
	}

	multipliers := make([]int, len(lines))
	for i := range multipliers {
		multipliers[i] = 1
	}

	totalCards := 0
	for i, line := range lines {
		contents := strings.Split(strings.Split(line, ":")[1], "|")
		winning := make(map[int]bool)
		for _, s := range strings.Split(contents[0], " ") {
			if s != "" {
				winning[atoid(s, -1)] = true
			}
		}

		matching := 0
		for _, s := range strings.Split(contents[1], " ") {
			if s != "" && winning[atoid(s, 0)] {
				matching++
			}
		}
		ctx.Printf("Your %d instances of card %d have %d matching numbers", multipliers[i], i+1, matching)
		for j := 0; j < matching; j++ {
			multipliers[i+j+1] += multipliers[i]
		}
		totalCards += multipliers[i]
	}

	return totalCards, nil
}
