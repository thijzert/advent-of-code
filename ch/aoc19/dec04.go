package aoc19

import (
	"fmt"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec04a(ctx ch.AOContext) (interface{}, error) {
	input := "359282-820401"

	var min, max int
	fmt.Sscanf(input, "%d-%d", &min, &max)

	rv := 0
	for pw := min; pw <= max; pw++ {
		double := false
		l, p := pw%10, pw/10
		for p > 0 {
			if p%10 == l {
				double = true
			}
			l, p = p%10, p/10
		}

		p = pw
		decreasing := false
		for p >= 10 {
			if (p/10)%10 > p%10 {
				decreasing = true
			}
			p = p / 10
		}

		if double && !decreasing {
			rv++
		}
	}

	return rv, nil
}

func Dec04b(ctx ch.AOContext) (interface{}, error) {
	input := "359282-820401"

	var min, max int
	fmt.Sscanf(input, "%d-%d", &min, &max)

	rv := 0
	//for _, pw := range []int{112233, 123444, 111122} {
	for pw := min; pw <= max; pw++ {
		streak := 1
		double := false
		l, p := pw%10, pw/10
		for p > 0 {
			if p%10 == l {
				streak++
			} else {
				if streak == 2 {
					double = true
				}
				streak = 1
			}
			l, p = p%10, p/10
		}
		if streak == 2 {
			double = true
		}

		p = pw
		decreasing := false
		for p >= 10 {
			if (p/10)%10 > p%10 {
				decreasing = true
			}
			p = p / 10
		}

		if double && !decreasing {
			rv++
		}

		//ctx.Printf("%d meets criteria: %v %v", pw, double, !decreasing)
	}

	return rv, nil
}
