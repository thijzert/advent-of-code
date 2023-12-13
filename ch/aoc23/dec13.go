package aoc23

import (
	"github.com/thijzert/advent-of-code/ch"
	"github.com/thijzert/advent-of-code/data"
)

func Dec13a(ctx ch.AOContext) (interface{}, error) {
	sections, err := ctx.DataSections("inputs/2023/dec13.txt")
	if err != nil {
		return nil, err
	}

	answer := 0
	err = nil
	for i, sect := range sections {
		if j, ok := dec13findAxis(sect); ok {
			ctx.Printf("section %d reflects in the %dth row", i+1, j)
			answer += 100 * j
		} else if j, ok := dec13findAxis(data.Transpose(sect)); ok {
			ctx.Printf("section %d reflects in the %dth column", i+1, j)
			answer += j
		} else {
			ctx.Printf("section %d does not reflect at all", i+1)
			err = errFailed
		}
	}

	return answer, err
}

var Dec13b ch.AdventFunc = nil

// func Dec13b(ctx ch.AOContext) (interface{}, error) {
// 	return nil, errNotImplemented
// }

func dec13findAxis(lines []string) (int, bool) {
	for i := range lines {
		if i == len(lines)-1 {
			continue
		}
		l := min(i, len(lines)-i-2)

		ok := true
		for j := 0; j <= l; j++ {
			ok = ok && lines[i-j] == lines[i+j+1]
		}
		if ok {
			return i + 1, true
		}
	}

	return 0, false
}
