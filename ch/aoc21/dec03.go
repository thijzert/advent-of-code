package aoc21

import (
	"errors"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec03a(ctx ch.AOContext) error {
	lines, err := ctx.DataLines("inputs/2021/dec03.txt")
	if err != nil {
		return err
	}

	histo := make([][2]int, len(lines[0]))

	for _, l := range lines {
		if l == "" {
			continue
		}

		for i, c := range l {
			if c == '0' {
				histo[i][0]++
			} else if c == '1' {
				histo[i][1]++
			}
		}
	}

	gamma := 0
	epsilon := 0
	for _, h := range histo {
		gamma <<= 1
		epsilon <<= 1
		if h[0] > h[1] {
			epsilon += 1
		} else {
			gamma += 1
		}
	}

	ctx.Debug.Printf("Gamma:   %3d (0b%b)", gamma, gamma)
	ctx.Debug.Printf("Epsilon: %3d (0b%b)", epsilon, epsilon)

	ctx.FinalAnswer.Print(gamma * epsilon)
	return nil
}

func Dec03b(ctx ch.AOContext) error {
	return errors.New("not implemented")
}
