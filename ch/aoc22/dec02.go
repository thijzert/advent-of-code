package aoc22

import (
	"github.com/thijzert/advent-of-code/ch"
)

func Dec02a(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2022/dec02.txt")
	if err != nil {
		return nil, err
	}
	//lines = []string{"A Y", "B X", "C Z"}

	totalScore := 0
	for _, l := range lines {
		opp, me := l[0], l[2]
		s := 0

		s += int(me - 'W')
		ctx.Printf("value score: %d", s)

		if (opp - 'A') == (me - 'X') {
			s += 3
		} else if (opp == 'A' && me == 'Y') || (opp == 'B' && me == 'Z') || (opp == 'C' && me == 'X') {
			s += 6
		}

		ctx.Printf("round: %d", s)
		totalScore += s
	}

	return totalScore, nil
}

func Dec02b(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2022/dec02.txt")
	if err != nil {
		return nil, err
	}
	// lines = []string{"A Y", "B X", "C Z"}

	totalScore := 0
	for _, l := range lines {
		s := 0

		s += 3 * int(l[2]-'X')
		ctx.Printf("result score: %d", s)

		if l == "A Y" || l == "B X" || l == "C Z" {
			ctx.Printf("played rock")
			s += 1
		} else if l == "B Y" || l == "C X" || l == "A Z" {
			ctx.Printf("played paper")
			s += 2
		} else if l == "C Y" || l == "A X" || l == "B Z" {
			ctx.Printf("played scissors")
			s += 3
		}

		ctx.Printf("round: %d", s)
		totalScore += s
	}

	return totalScore, nil
}
