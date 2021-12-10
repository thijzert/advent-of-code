package aoc21

import (
	"sort"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec10a(ctx ch.AOContext) error {
	rv, err := checkMatchingPairs(ctx, "inputs/2021/dec10.txt", false)
	if err != nil {
		return err
	}
	ctx.FinalAnswer.Print(rv)
	return nil
}

func Dec10b(ctx ch.AOContext) error {
	rv, err := checkMatchingPairs(ctx, "inputs/2021/dec10.txt", true)
	if err != nil {
		return err
	}

	ctx.FinalAnswer.Print(rv)
	return nil
}

func checkMatchingPairs(ctx ch.AOContext, assetName string, autocomplete bool) (int, error) {
	lines, err := ctx.DataLines(assetName)
	if err != nil {
		return 0, err
	}

	var invalid int
	var incomplete []int

	for i, l := range lines {
		if l == "" {
			continue
		}

		valid := true
		stack := make([]rune, 1, len(l))
		for _, c := range l {
			if c == '(' || c == '[' || c == '{' || c == '<' {
				stack = append(stack, c)
			} else if c == closeC(stack[len(stack)-1]) {
				stack = stack[:len(stack)-1]
			} else {
				ctx.Printf("Line %d: expected '%c' but got '%c'", i+1, closeC(stack[len(stack)-1]), c)

				invalid += cpoints(c)
				valid = false
				break
			}
		}

		if valid && autocomplete {
			score := 0
			for i := range stack {
				c := stack[len(stack)-i-1]
				if c == 0 {
					continue
				}
				score = score*5 + cpoints(c)
			}

			ctx.Printf("Line %d: unfinished %c, for %d points", i+1, stack[1:], score)

			incomplete = append(incomplete, score)
		}
	}

	if autocomplete {
		sort.Ints(incomplete)
		ctx.Printf("Scores: %d", incomplete)
		return incomplete[len(incomplete)/2], nil
	} else {
		return invalid, nil
	}
}

func closeC(c rune) rune {
	if c == '(' {
		return ')'
	} else if c == '[' {
		return ']'
	} else if c == '{' {
		return '}'
	} else if c == '<' {
		return '>'
	}
	return c
}

func cpoints(c rune) int {
	if c == ')' {
		return 3
	} else if c == ']' {
		return 57
	} else if c == '}' {
		return 1197
	} else if c == '>' {
		return 25137
	} else if c == '(' {
		return 1
	} else if c == '[' {
		return 2
	} else if c == '{' {
		return 3
	} else if c == '<' {
		return 4
	}
	return 0
}
