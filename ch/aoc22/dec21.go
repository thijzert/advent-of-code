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

var Dec21b ch.AdventFunc = nil

// func Dec21b(ctx ch.AOContext) (interface{}, error) {
// 	return nil, errNotImplemented
// }

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

func applyMonkeyMath(monkeys map[string]monkeyMath) (int, error) {
	changed := 0
	for k, mm := range monkeys {
		if mm.AnswerKnown {
			continue
		}

		var lhs, rhs monkeyMath
		var ok bool
		if lhs, ok = monkeys[mm.LHS]; !ok {
			return changed, fmt.Errorf("monkey '%s' not found")
		}
		if rhs, ok = monkeys[mm.RHS]; !ok {
			return changed, fmt.Errorf("monkey '%s' not found")
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
			} else {
				return changed, fmt.Errorf("Invalid operation '%c' (0x%2x)", mm.Operation, mm.Operation)
			}
			monkeys[k] = mm
			changed++
		}
	}
	return changed, nil
}
