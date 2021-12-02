package aoc21

import "github.com/thijzert/advent-of-code/ch"

func Dec01a(ctx ch.AOContext) error {
	depths, err := ctx.DataAsInts("inputs/2021/dec01.txt")
	if err != nil {
		return err
	}

	rv := 0

	for i, depth := range depths {
		if i == 0 {
			continue
		}
		if depth > depths[i-1] {
			rv++
		}
	}

	ctx.FinalAnswer.Printf("%d\n", rv)
	return nil
}

func Dec01b(ctx ch.AOContext) error {
	depths, err := ctx.DataAsInts("inputs/2021/dec01.txt")
	if err != nil {
		return err
	}

	rv := 0

	avg := 0
	lastAvg := 0

	for i, depth := range depths {
		lastAvg = avg
		avg += depth

		if i < 3 {
			continue
		}
		avg -= depths[i-3]

		if avg > lastAvg {
			rv++
		}
	}

	ctx.FinalAnswer.Printf("%d\n", rv)
	return nil
}
