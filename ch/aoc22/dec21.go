package aoc22

import (
	"fmt"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec21a(ctx ch.AOContext) (interface{}, error) {
	monkeys, err := readMonkeyMath(ctx)
	if err != nil {
		return nil, err
	}

	//for k, mm := range monkeys {
	//	ctx.Printf("%s: %s", k, mm)
	//}
	changed := 1
	for changed > 0 {
		changed, err = applyMonkeyMath(monkeys)
		if err != nil {
			return nil, err
		}
		ctx.Printf("%d monkeys changed", changed)
		//for k, mm := range monkeys {
		//	ctx.Printf("%s: %s", k, mm)
		//}
	}

	if !monkeys["root"].AnswerKnown {
		return nil, errFailed
	}

	return monkeys["root"].Answer, nil
}

func Dec21b(ctx ch.AOContext) (interface{}, error) {
	monkeys, err := readMonkeyMath(ctx)
	if err != nil {
		return nil, err
	}
	mm := monkeys["root"]
	mm.Operation = '='
	monkeys["root"] = mm

	// Always mount a few of these
	scratchMonkeys := make(map[string]monkeyMath)
	for k, mm := range monkeys {
		scratchMonkeys[k] = mm
	}

	finalAnswer, err := solveMonkeyMath(ctx, monkeys, "root", "humn")
	if err != nil {
		return nil, err
	}

	mm = scratchMonkeys["humn"]
	mm.Answer = finalAnswer
	scratchMonkeys["humn"] = mm

	err = applyAllMonkeyMath(scratchMonkeys)
	if err != nil {
		return nil, err
	}

	return finalAnswer, nil
}

type monkeyMath struct {
	AnswerKnown bool
	Answer      int
	Operation   rune
	LHS, RHS    string
}

func (mm monkeyMath) String() string {
	if mm.Operation == 0 {
		return fmt.Sprintf("%d", mm.Answer)
	}

	if mm.AnswerKnown {
		return fmt.Sprintf("%d (%s %c %s)", mm.Answer, mm.LHS, mm.Operation, mm.RHS)
	} else {
		return fmt.Sprintf("?  (%s %c %s)", mm.LHS, mm.Operation, mm.RHS)
	}
}

func readMonkeyMath(ctx ch.AOContext) (map[string]monkeyMath, error) {
	lines, err := ctx.DataLines("inputs/2022/dec21.txt")
	if err != nil {
		return nil, err
	}
	//lines = []string{
	//	"root: pppw + sjmn",
	//	"dbpl: 5",
	//	"cczh: sllz + lgvd",
	//	"zczc: 2",
	//	"ptdq: humn - dvpt",
	//	"dvpt: 3",
	//	"lfqf: 4",
	//	"humn: 5",
	//	"ljgn: 2",
	//	"sjmn: drzm * dbpl",
	//	"sllz: 4",
	//	"pppw: cczh / lfqf",
	//	"lgvd: ljgn * ptdq",
	//	"drzm: hmdt - zczc",
	//	"hmdt: 32",
	//}

	monkeys := make(map[string]monkeyMath)
	for _, line := range lines {
		var mm monkeyMath
		var k string
		_, err := fmt.Sscanf(line, "%4s: %d", &k, &mm.Answer)
		if err == nil {
			mm.AnswerKnown = true
		} else {
			_, err = fmt.Sscanf(line, "%4s: %4s %c %4s", &k, &mm.LHS, &mm.Operation, &mm.RHS)
			if err != nil {
				ctx.Printf("Line: \"%s\"", line)
				return nil, err
			}
		}

		monkeys[k] = mm
	}
	return monkeys, nil
}

func applyAllMonkeyMath(monkeys map[string]monkeyMath) error {
	var err error
	changed := 1
	for changed > 0 {
		changed, err = applyMonkeyMath(monkeys)
		if err != nil {
			return err
		}
	}
	return nil
}

func applyMonkeyMath(monkeys map[string]monkeyMath) (int, error) {
	changed := 0
	for k, mm := range monkeys {
		if mm.AnswerKnown {
			continue
		}

		var lhs, rhs monkeyMath
		var ok bool
		if lhs, ok = monkeys[mm.LHS]; !ok {
			continue
		}
		if rhs, ok = monkeys[mm.RHS]; !ok {
			continue
		}
		if lhs.AnswerKnown && rhs.AnswerKnown {
			mm.AnswerKnown = true
			if mm.Operation == '+' {
				mm.Answer = lhs.Answer + rhs.Answer
			} else if mm.Operation == '-' {
				mm.Answer = lhs.Answer - rhs.Answer
			} else if mm.Operation == '*' {
				mm.Answer = lhs.Answer * rhs.Answer
			} else if mm.Operation == '/' {
				mm.Answer = lhs.Answer / rhs.Answer
			} else if mm.Operation == '=' {
				mm.Answer = 1
				if lhs.Answer != rhs.Answer {
					return changed, fmt.Errorf("Equation failed: %d ≠ %d", lhs.Answer, rhs.Answer)
				}
			} else {
				return changed, fmt.Errorf("Invalid operation '%c' (0x%2x)", mm.Operation, mm.Operation)
			}
			monkeys[k] = mm
			changed++
		}
	}
	return changed, nil
}

func hasMonkeyVariable(monkeys map[string]monkeyMath, expr, forvar string) bool {
	if expr == forvar {
		return true
	}
	mm, ok := monkeys[expr]
	if !ok {
		return false
	}

	return hasMonkeyVariable(monkeys, mm.LHS, forvar) || hasMonkeyVariable(monkeys, mm.RHS, forvar)
}

func solveMonkeyMath(ctx ch.AOContext, monkeys map[string]monkeyMath, eq, forvar string) (int, error) {
	mm, ok := monkeys[eq]
	if !ok {
		return 0, fmt.Errorf("cannot find root expression '%s'", eq)
	}

	if hasMonkeyVariable(monkeys, mm.RHS, forvar) {
		mm.LHS, mm.RHS = mm.RHS, mm.LHS
		monkeys[eq] = mm
	}
	delete(monkeys, forvar)

	err := applyAllMonkeyMath(monkeys)
	if err != nil {
		return 0, err
	}
	for mm.LHS != forvar {
		mmLHS := mm.LHS
		ctx.Print(printMonkeyExpr(monkeys, eq, forvar))
		ctx.Printf("%v %v %v", mm, monkeys[mm.LHS], monkeys[mm.RHS])
		lhs := monkeys[mm.LHS]
		if !hasMonkeyVariable(monkeys, lhs.RHS, forvar) {
			t := mm.RHS
			mm.RHS = mm.LHS
			mm.LHS = lhs.LHS
			lhs.LHS = t
		} else {
			t := mm.RHS
			mm.RHS = mm.LHS
			mm.LHS = lhs.RHS
			lhs.RHS = t
			if lhs.Operation == '*' || lhs.Operation == '+' {
				lhs.LHS, lhs.RHS = lhs.RHS, lhs.LHS
			}
		}

		if lhs.Operation == '+' {
			lhs.Operation = '-'
		} else if lhs.Operation == '-' {
			lhs.Operation = '+'
		} else if lhs.Operation == '*' {
			lhs.Operation = '/'
		} else if lhs.Operation == '/' {
			lhs.Operation = '*'
		} else {
			return 0, errFailed
		}
		monkeys[mmLHS] = lhs
		monkeys[eq] = mm
		//ctx.Printf("%v %v %v", mm, monkeys[mm.LHS], monkeys[mm.RHS])
	}
	ctx.Print(printMonkeyExpr(monkeys, eq, forvar))

	err = applyAllMonkeyMath(monkeys)
	if err != nil {
		return 0, err
	}
	ctx.Print(printMonkeyExpr(monkeys, eq, forvar))

	if !monkeys[monkeys[eq].RHS].AnswerKnown {
		return 0, errFailed
	}

	return monkeys[monkeys[eq].RHS].Answer, nil
}

func printMonkeyExpr(monkeys map[string]monkeyMath, expr, forvar string) string {
	if expr == forvar {
		return "X"
	}
	mm := monkeys[expr]
	if mm.AnswerKnown {
		return fmt.Sprintf("%d", mm.Answer)
	}
	return fmt.Sprintf("(%s %c %s)", printMonkeyExpr(monkeys, mm.LHS, forvar), mm.Operation, printMonkeyExpr(monkeys, mm.RHS, forvar))
}
