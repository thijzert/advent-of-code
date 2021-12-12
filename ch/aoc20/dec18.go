package aoc20

import (
	"fmt"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec18a(ctx ch.AOContext) error {
	totest := []struct {
		Expr   string
		Answer int
	}{
		{"1 + 2 * 3 + 4 * 5 + 6", 71},
		{"1 + (2 * 3) + (4 * (5 + 6))", 51},
		{"2 * 3 + (4 * 5)", 26},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 437},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 12240},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632},
	}

	var err error
	for _, tt := range totest {
		a := evaluateHomework(tt.Expr)
		if a == tt.Answer {
			ctx.Printf("%s = %d", tt.Expr, a)
		} else {
			ctx.Printf("math error: %s = %d, got %d", tt.Expr, tt.Answer, a)
			err = fmt.Errorf("math error")
		}
	}
	if err != nil {
		return err
	}

	lines, err := ctx.DataLines("inputs/2020/dec18.txt")
	if err != nil {
		return err
	}

	rv := 0
	for _, l := range lines {
		if l == "" {
			continue
		}
		a := evaluateHomework(l)
		ctx.Printf("%s = %d", l, a)
		rv += a
	}

	ctx.FinalAnswer.Print(rv)
	return nil
}

func Dec18b(ctx ch.AOContext) error {
	return errNotImplemented
}

func evaluateHomework(expr string) int {
	tokens := make([]rune, 0, len(expr))
	for _, c := range expr {
		if c != ' ' {
			tokens = append(tokens, c)
		}
	}
	return evaluateHomeworkTokens(tokens)
}

func evaluateHomeworkTokens(tokens []rune) int {
	rv, i := evaluateOperand(tokens)

	for i < len(tokens) {
		b, l := evaluateOperand(tokens[i+1:])

		if tokens[i] == '+' {
			rv += b
		} else if tokens[i] == '*' {
			rv *= b
		}

		i += l + 1
	}

	return rv
}

func evaluateOperand(tokens []rune) (int, int) {
	if tokens[0] == '(' {
		ct := 1
		var j int
		var c rune
		for j, c = range tokens[1:] {
			if c == '(' {
				ct++
			} else if c == ')' {
				ct--
				if ct == 0 {
					break
				}
			}
		}
		return evaluateHomeworkTokens(tokens[1 : j+1]), j + 2
	} else {
		return int(tokens[0] - '0'), 1
	}
}
