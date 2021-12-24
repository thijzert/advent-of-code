package aoc21

import (
	"fmt"

	"github.com/thijzert/advent-of-code/ch"
)

var Dec24b ch.AdventFunc = nil

func Dec24a(ctx ch.AOContext) error {
	program, err := ctx.DataLines("inputs/2021/dec24.txt")
	if err != nil {
		return err
	}

	alu := &monadALU{}
	err = alu.Run(program, []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 9, 9, 9, 9, 9})
	if err != nil {
		return err
	}
	ctx.Print(alu.Z, alu.Z == 0)

	alu.Reset()
	// Worked this out on paper
	modno := []int{9, 9, 9, 1, 1, 9, 9, 3, 9, 4, 9, 6, 8, 4}
	err = alu.Run(program, modno)
	if err != nil {
		return err
	}
	ctx.Print(alu.Z, alu.Z == 0)

	if alu.Z != 0 {
		return errFailed
	}

	rv := 0
	for _, i := range modno {
		rv *= 10
		rv += i
	}
	ctx.FinalAnswer.Print(rv)
	return nil
}

// func Dec24b(ctx ch.AOContext) error {
// 	return errNotImplemented
// }

type monadALU struct {
	EIP        int
	ESP        int
	W, X, Y, Z int
}

func (m *monadALU) Reset() {
	m.EIP, m.ESP = 0, 0
	m.W, m.X, m.Y, m.Z = 0, 0, 0, 0
}

func (m *monadALU) Run(program []string, input []int) error {
	var l string
	for m.EIP, l = range program {
		if l == "" {
			continue
		}
		if len(l) < 5 {
			return fmt.Errorf("undefined opcode on line %d: '%s'", m.EIP, l)
		}
		if l[3] != ' ' || (len(l) > 5 && l[5] != ' ') {
			return fmt.Errorf("invalid format on line %d: '%s'", m.EIP, l)
		}

		op0 := m.operand(l[4:5])
		if op0 == nil {
			return fmt.Errorf("invalid operand on line %d: '%s'", m.EIP, l[4:5])
		}
		var op1 *int
		if len(l) > 5 {
			op1 = m.operand(l[6:])
			if op1 == nil {
				return fmt.Errorf("invalid operand on line %d: '%s'", m.EIP, l[6:])
			}
		}

		if l[0:3] == "inp" {
			if len(input) <= m.ESP {
				return fmt.Errorf("halting for input on line %d: '%s'", m.EIP, l)
			}
			*op0 = input[m.ESP]
			m.ESP++
		} else if l[0:3] == "add" {
			if op1 == nil {
				return fmt.Errorf("incorrect number of operands on line %d: '%s'", m.EIP, l)
			}
			*op0 += *op1
		} else if l[0:3] == "mul" {
			if op1 == nil {
				return fmt.Errorf("incorrect number of operands on line %d: '%s'", m.EIP, l)
			}
			*op0 *= *op1
		} else if l[0:3] == "div" {
			if op1 == nil {
				return fmt.Errorf("incorrect number of operands on line %d: '%s'", m.EIP, l)
			}
			if *op1 == 0 {
				return fmt.Errorf("division by zero on line %d: '%s'", m.EIP, l)
			}
			*op0 /= *op1
		} else if l[0:3] == "mod" {
			if op1 == nil {
				return fmt.Errorf("incorrect number of operands on line %d: '%s'", m.EIP, l)
			}
			if *op1 == 0 {
				return fmt.Errorf("division by zero on line %d: '%s'", m.EIP, l)
			}
			*op0 %= *op1
		} else if l[0:3] == "eql" {
			if op1 == nil {
				return fmt.Errorf("incorrect number of operands on line %d: '%s'", m.EIP, l)
			}
			if *op0 == *op1 {
				*op0 = 1
			} else {
				*op0 = 0
			}
		} else {
			return fmt.Errorf("undefined opcode on line %d: '%s'", m.EIP, l)
		}
	}

	return nil
}

func (m *monadALU) operand(s string) *int {
	if s == "w" {
		return &m.W
	} else if s == "x" {
		return &m.X
	} else if s == "y" {
		return &m.Y
	} else if s == "z" {
		return &m.Z
	}

	var rv int
	if _, err := fmt.Sscanf(s, "%d", &rv); err == nil {
		return &rv
	}
	return nil
}
