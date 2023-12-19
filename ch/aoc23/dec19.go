package aoc23

import (
	"fmt"
	"strings"

	"github.com/thijzert/advent-of-code/ch"
)

type machinePart struct {
	X, M, A, S int
}

func (pt machinePart) Rating() int {
	return pt.X + pt.M + pt.A + pt.S
}

type machineOp struct {
	Left  byte
	Op    byte
	Right int
	Jump  string
}

func (op machineOp) Match(pt machinePart) string {
	left := 0
	if op.Left == 'x' {
		left = pt.X
	} else if op.Left == 'm' {
		left = pt.M
	} else if op.Left == 'a' {
		left = pt.A
	} else if op.Left == 's' {
		left = pt.S
	} else {
		return ""
	}

	rv := false
	if op.Op == '<' {
		rv = left < op.Right
	} else if op.Op == '>' {
		rv = left > op.Right
	}

	if rv {
		return op.Jump
	}
	return ""
}

type machineRule struct {
	Operations []machineOp
	Default    string
}

type machineRuleSet map[string]machineRule

func Dec19a(ctx ch.AOContext) (interface{}, error) {
	ruleset, parts, err := dec19read(ctx)
	if err != nil {
		return nil, err
	}

	answer := 0
	for _, pt := range parts {
		ruleName := "in"
		jumps := 0
		for jumps < 1000 {
			rule, ok := ruleset[ruleName]
			if !ok {
				break
			}
			jumps++

			ruleName = rule.Default
			for _, op := range rule.Operations {
				if jmp := op.Match(pt); jmp != "" {
					ruleName = jmp
					break
				}
			}
		}
		ctx.Printf("Part %+v: %s", pt, ruleName)
		if ruleName == "A" {
			answer += pt.Rating()
		}
	}

	return answer, nil
}

var Dec19b ch.AdventFunc = nil

// func Dec19b(ctx ch.AOContext) (interface{}, error) {
// 	return nil, errNotImplemented
// }

func dec19read(ctx ch.AOContext) (machineRuleSet, []machinePart, error) {
	sections, err := ctx.DataSections("inputs/2023/dec19.txt")
	if err != nil {
		return nil, nil, err
	}

	ruleset := make(machineRuleSet)
	parts := []machinePart{}

	for _, line := range sections[0] {
		name := strings.Split(line, "{")
		operations := strings.Split(strings.TrimRight(name[1], "}"), ",")
		rule := machineRule{}
		for _, operation := range operations {
			oppts := strings.Split(operation, ":")
			if len(oppts) == 1 {
				rule.Default = oppts[0]
				continue
			}
			mop := machineOp{Jump: oppts[1]}
			_, err = fmt.Sscanf(oppts[0], "%c%c%d", &mop.Left, &mop.Op, &mop.Right)
			if err != nil {
				return nil, nil, fmt.Errorf("error parsing operation '%s': %w", operation, err)
			}
			rule.Operations = append(rule.Operations, mop)
		}
		//ctx.Printf("Rule '%s':  %v", name[0], rule)
		ruleset[name[0]] = rule
	}

	for _, line := range sections[1] {
		pt := machinePart{}
		_, err = fmt.Sscanf(line, "{x=%d,m=%d,a=%d,s=%d}", &pt.X, &pt.M, &pt.A, &pt.S)
		if err != nil {
			return nil, nil, fmt.Errorf("error parsing part '%s': %w", line, err)
		}
		//ctx.Printf("Part: %+v", pt)
		parts = append(parts, pt)
	}

	return ruleset, parts, nil
}
