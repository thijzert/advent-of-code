package aoc19

import (
	"fmt"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec02a(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataAsIntLists("inputs/2019/dec02.txt")
	if err != nil {
		return nil, err
	}
	program := lines[0]
	//program = []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}

	program[1] = 12
	program[2] = 2

	ans, err := runIntCodeProgram(program)
	if err != nil {
		return nil, err
	}

	return ans, nil
}

func Dec02b(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataAsIntLists("inputs/2019/dec02.txt")
	if err != nil {
		return nil, err
	}
	program := lines[0]
	//program = []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}

	target := 19690720

	for verb := 0; verb <= 99; verb++ {
		for noun := 0; noun <= 99; noun++ {
			program[1] = noun
			program[2] = verb

			ans, err := runIntCodeProgram(program)
			if err == nil && ans == target {
				return 100*noun + verb, nil
			}
		}
	}

	return nil, errFailed
}

func runIntCodeProgram(program []int) (int, error) {
	memory := make([]int, len(program))
	copy(memory, program)

	pc := 0
	for pc >= 0 && pc <= len(memory) && memory[pc] != 99 {
		if memory[pc] == 1 {
			memory[memory[pc+3]] = memory[memory[pc+1]] + memory[memory[pc+2]]
			pc += 4
		} else if memory[pc] == 2 {
			memory[memory[pc+3]] = memory[memory[pc+1]] * memory[memory[pc+2]]
			pc += 4
		} else {
			return 0, fmt.Errorf("Undefined opcode %d", memory[pc])
		}
	}

	return memory[0], nil
}
