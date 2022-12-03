package aoc22

import (
	"regexp"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec03a(ctx ch.AOContext) error {
	lines, err := ctx.DataLines("inputs/2022/dec03.txt")
	if err != nil {
		return err
	}

	sumPr := 0
	for _, line := range lines {
		cutoff := len(line) / 2
		front, back := line[:cutoff], line[cutoff:]

		r, err := regexp.Compile("[" + front + "]")
		if err != nil {
			return err
		}
		m := r.FindString(back)
		ctx.Printf("both compartments contain '%s'", m)
		sumPr += rucksackPriority(m[0])
	}

	ctx.FinalAnswer.Print(sumPr)
	return nil
}

func Dec03b(ctx ch.AOContext) error {
	lines, err := ctx.DataLines("inputs/2022/dec03.txt")
	if err != nil {
		return err
	}

	sumPr := 0

	elfIdx := byte(2)
	seen := [53]byte{}
	for _, line := range lines {
		elfIdx++
		if elfIdx == 3 {
			for i := range seen {
				seen[i] = 0
			}
			elfIdx = 0
		}
		for _, c := range line {
			i := rucksackPriority(byte(c))
			if seen[i] == elfIdx {
				seen[i]++
				if seen[i] == 3 {
					ctx.Printf("All elves carry '%c'", c)
					sumPr += i
				}
			}
		}
	}

	ctx.FinalAnswer.Print(sumPr)
	return nil
}

func rucksackPriority(b byte) int {
	if b >= 'a' && b <= 'z' {
		return int(1 + b - 'a')
	} else if b >= 'A' && b <= 'Z' {
		return int(27 + b - 'A')
	}
	return 0
}
