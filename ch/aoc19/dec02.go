package aoc19

import (
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

	ans, err := runIntCodeProgram(program, nil, nil)
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

			ans, err := runIntCodeProgram(program, nil, nil)
			if err == nil && ans == target {
				return 100*noun + verb, nil
			}
		}
	}

	return nil, errFailed
}
