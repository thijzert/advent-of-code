package aoc22

import (
	"fmt"

	"github.com/thijzert/advent-of-code/ch"
	"github.com/thijzert/advent-of-code/lib/cube"
)

func Dec04a(ctx ch.AOContext) (interface{}, error) {
	pairs, err := dataAsIVPs(ctx, "inputs/2022/dec04.txt")
	if err != nil {
		return nil, err
	}

	fullyContain := 0
	for _, ivp := range pairs {
		if ivp.P.FullyContains(ivp.Q) {
			fullyContain++
		} else if ivp.Q.FullyContains(ivp.P) {
			fullyContain++
		}
	}

	return fullyContain, nil
}

func Dec04b(ctx ch.AOContext) (interface{}, error) {
	pairs, err := dataAsIVPs(ctx, "inputs/2022/dec04.txt")
	if err != nil {
		return nil, err
	}

	overlap := 0
	for _, ivp := range pairs {
		if _, ok := ivp.P.Overlap(ivp.Q); ok {
			overlap++
		}
	}

	return overlap, nil
}

// IVP is an interval pair
type IVP struct {
	P, Q cube.Interval
}

func dataAsIVPs(ctx ch.AOContext, filename string) ([]IVP, error) {
	lines, err := ctx.DataLines(filename)
	if err != nil {
		return nil, err
	}
	rv := []IVP{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		var ivp IVP
		_, err := fmt.Sscanf(line, "%d-%d,%d-%d", &ivp.P.A, &ivp.P.B, &ivp.Q.A, &ivp.Q.B)
		if err != nil {
			return nil, err
		}
		rv = append(rv, ivp)
	}
	return rv, nil
}
