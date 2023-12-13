package aoc23

import (
	"github.com/thijzert/advent-of-code/ch"
	"github.com/thijzert/advent-of-code/data"
)

func Dec13a(ctx ch.AOContext) (interface{}, error) {
	return dec13(ctx, 0)
}

func Dec13b(ctx ch.AOContext) (interface{}, error) {
	return dec13(ctx, 1)
}

func dec13(ctx ch.AOContext, smudge int) (interface{}, error) {
	sections, err := ctx.DataSections("inputs/2023/dec13.txt")
	if err != nil {
		return nil, err
	}

	answer := 0
	err = nil
	for i, sect := range sections {
		if j, ok := dec13findAxis(sect, smudge); ok {
			ctx.Printf("section %d reflects in the %dth row", i+1, j)
			answer += 100 * j
		} else if j, ok := dec13findAxis(data.Transpose(sect), smudge); ok {
			ctx.Printf("section %d reflects in the %dth column", i+1, j)
			answer += j
		} else {
			ctx.Printf("section %d does not reflect at all", i+1)
			err = errFailed
		}
	}

	return answer, err
}

func dec13findAxis(lines []string, smudgeFactor int) (int, bool) {
	for i := range lines {
		if i == len(lines)-1 {
			continue
		}
		l := min(i, len(lines)-i-2)

		dist := 0
		for j := 0; j <= l; j++ {
			dist += dec13stringdist(lines[i-j], lines[i+j+1])
		}
		if dist == smudgeFactor {
			return i + 1, true
		}
	}

	return 0, false
}

func dec13stringdist(a, b string) int {
	rv := 0
	l := len(a)
	for i := 0; i < l; i++ {
		if a[i] != b[i] {
			rv++
		}
	}
	return rv
}
