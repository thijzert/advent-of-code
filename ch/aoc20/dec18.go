package aoc20

import (
	"fmt"
	"log"

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
		a := evaluateHomework(tt.Expr, leftToRight)
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
		a := evaluateHomework(l, leftToRight)
		// ctx.Printf("%s = %d", l, a)
		rv += a
	}

	ctx.FinalAnswer.Print(rv)
	return nil
}

func Dec18b(ctx ch.AOContext) error {
	totest := []struct {
		Expr   string
		Answer int
	}{
		{"1 + 2 * 3 + 4 * 5 + 6", 231},
		{"1 + (2 * 3) + (4 * (5 + 6))", 51},
		{"2 * 3 + (4 * 5)", 46},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 1445},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 669060},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 23340},
	}

	var err error
	for _, tt := range totest {
		a := evaluateHomework(tt.Expr, addBeforeMult)
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
		a := evaluateHomework(l, addBeforeMult)
		// ctx.Printf("%s = %d", l, a)
		rv += a
	}

	ctx.FinalAnswer.Print(rv)
	return nil
}

type homeworkLexeme struct {
	Operator rune
	Operand  int
}

type homeworkEvaluator func(tokens []homeworkLexeme) int

func evaluateHomework(expr string, f homeworkEvaluator) int {
	tokens := make([]rune, 0, len(expr))
	for _, c := range expr {
		if c != ' ' {
			tokens = append(tokens, c)
		}
	}

	return homeworkParens(tokens, f)
}

func homeworkParens(tokens []rune, f homeworkEvaluator) int {
	var rv []homeworkLexeme

	i := 0

	for i < len(tokens) {
		c := tokens[i]
		l := 1

		if c == '+' || c == '*' {
			rv = append(rv, homeworkLexeme{Operator: c})
		} else if c == '(' {
			ct := 1
			var j int
			var c rune
			for j, c = range tokens[i+1:] {
				if c == '(' {
					ct++
				} else if c == ')' {
					ct--
					if ct == 0 {
						break
					}
				}
			}
			rv = append(rv, homeworkLexeme{Operand: homeworkParens(tokens[i+1:i+j+1], f)})
			l = j + 2
		} else {
			rv = append(rv, homeworkLexeme{Operand: int(c - '0')})
		}

		i += l
	}

	return f(rv)
}

func leftToRight(expr []homeworkLexeme) int {
	rv := expr[0].Operand

	for i := 1; i < len(expr); i += 2 {
		if expr[i].Operator == '+' {
			rv += expr[i+1].Operand
		} else if expr[i].Operator == '*' {
			rv *= expr[i+1].Operand
		} else {
			log.Printf("Unknown operator '%c' at index %d - (operand value: %d)", expr[i].Operator, i, expr[i].Operand)
			return -1
		}
	}

	return rv
}

func addBeforeMult(expr []homeworkLexeme) int {
	rv := 1
	current := expr[0].Operand

	for i := 1; i < len(expr); i += 2 {
		if expr[i].Operator == '+' {
			current += expr[i+1].Operand
		} else if expr[i].Operator == '*' {
			rv *= current
			current = expr[i+1].Operand
		} else {
			log.Printf("Unknown operator '%c' at index %d - (operand value: %d)", expr[i].Operator, i, expr[i].Operand)
			return -1
		}
	}

	return rv * current
}
