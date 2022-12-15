package aoc20

import (
	"fmt"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec08a(ctx ch.AOContext) (interface{}, error) {
	inputs := []string{"nop +0", "acc +1", "jmp +4", "acc +3", "jmp -3", "acc -99", "acc +1", "jmp -4", "acc +6"}
	acc, _ := runGameConsoleProgram(inputs)
	ctx.Printf("Example accumulator: %d", acc)

	inputs, err := ctx.DataLines("inputs/2020/dec08.txt")
	if err != nil {
		return nil, err
	}
	acc, _ = runGameConsoleProgram(inputs)
	return acc, nil
}

func Dec08b(ctx ch.AOContext) (interface{}, error) {
	inputs := []string{"nop +0", "acc +1", "jmp +4", "acc +3", "jmp -3", "acc -99", "acc +1", "jmp -4", "acc +6"}
	i, acc := busyConsoleBeaver(inputs)
	if i >= 0 {
		ctx.Printf("Changing instruction %d to '%s' results in %d", i+1, inputs[i], acc)
	}

	inputs, err := ctx.DataLines("inputs/2020/dec08.txt")
	if err != nil {
		return nil, err
	}
	i, acc = busyConsoleBeaver(inputs)
	ctx.Printf("Changing instruction %d to '%s' results in %d", i+1, inputs[i], acc)
	return acc, nil
}

func runGameConsoleProgram(instrs []string) (int, bool) {
	counts := make([]int, len(instrs))
	acc := 0
	pc := 0
	for pc < len(instrs) {
		counts[pc]++
		if counts[pc] > 1 {
			return acc, false
		}

		var opcode string
		var operand int
		var sign rune
		fmt.Sscanf(instrs[pc], "%3s %c%d", &opcode, &sign, &operand)
		if sign == '-' {
			operand = -operand
		}

		if opcode == "acc" {
			acc += operand
			pc++
		} else if opcode == "jmp" {
			pc += operand
		} else {
			pc++
		}
	}

	return acc, true
}

func busyConsoleBeaver(instrs []string) (int, int) {
	for i, s := range instrs {
		if len(s) < 3 {
			continue
		} else if s[0:3] == "nop" {
			instrs[i] = "jmp" + s[3:]
		} else if s[0:3] == "jmp" {
			instrs[i] = "nop" + s[3:]
		} else {
			continue
		}

		acc, exitsNormally := runGameConsoleProgram(instrs)
		if exitsNormally {
			return i, acc
		}
		instrs[i] = s
	}

	return -1, 0
}
