package aoc22

import (
	"fmt"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec04a(ctx ch.AOContext) error {
	pairs, err := dataAsIVPs(ctx, "inputs/2022/dec04.txt")
	if err != nil {
		return err
	}

	fullyContain := 0
	for _, ivp := range pairs {
		if ivp.P.Contains(ivp.Q) {
			fullyContain++
		} else if ivp.Q.Contains(ivp.P) {
			fullyContain++
		}
	}

	ctx.FinalAnswer.Print(fullyContain)
	return nil
}

func Dec04b(ctx ch.AOContext) error {
	pairs, err := dataAsIVPs(ctx, "inputs/2022/dec04.txt")
	if err != nil {
		return err
	}

	overlap := 0
	for _, ivp := range pairs {
		if _, ok := ivp.P.Overlap(ivp.Q); ok {
			overlap++
		}
	}

	ctx.FinalAnswer.Print(overlap)
	return errNotImplemented
}

// IVP is an interval pair
type IVP struct {
	P, Q Interval
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

// Interval represent an inclusive integer interval, with A <= B
type Interval struct {
	A, B int
}

func (a Interval) Overlap(b Interval) (Interval, bool) {
	if b.A >= a.A && b.A <= a.B {
		if b.B < a.B {
			return Interval{b.A, b.B}, true
		} else {
			return Interval{b.A, a.B}, true
		}
	} else if b.B >= a.A && b.B <= a.B {
		if a.A < b.A {
			return b, true
		} else {
			return Interval{a.A, b.B}, true
		}
	} else if b.A <= a.A && b.B >= a.B {
		return a, true
	}
	return Interval{}, false
}

func (a Interval) Contains(b Interval) bool {
	return b.A >= a.A && b.A <= a.B && b.B >= a.A && b.B <= a.B
}
