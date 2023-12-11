package aoc19

import (
	"github.com/thijzert/advent-of-code/ch"
)

func Dec09a(ctx ch.AOContext) (interface{}, error) {
	programs, err := ctx.DataAsIntLists("inputs/2019/dec09.txt")
	if err != nil {
		return nil, err
	}
	program := programs[0]
	//program = []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}
	//program = []int{1102, 34915192, 34915192, 7, 4, 7, 99, 0}
	//program = []int{109, 2000, 1102, 34915192, 34915192, 2000, 2, 2000, 2000, 2000, 109, 19, 109, -34, 204, 15, 99, 0}
	//program = []int{1101, 3, 0, 1000, 109, 988, 209, 12, 9, 1000, 99}

	input := []int{1}
	output := make([]int, 100)
	_, _, no, err := runIntCodeProgram(program, input, output)
	if err != nil {
		return nil, err
	}
	ctx.Printf("Output: %d", output[:no])

	// 209: too low
	return nil, errNotImplemented
}

func Dec09b(ctx ch.AOContext) (interface{}, error) {
	programs, err := ctx.DataAsIntLists("inputs/2019/dec09.txt")
	if err != nil {
		return nil, err
	}
	program := programs[0]

	input := []int{2}
	output := make([]int, 100)
	_, _, no, err := runIntCodeProgram(program, input, output)
	if err != nil {
		return nil, err
	}
	ctx.Printf("Output: %d", output[:no])

	return nil, errNotImplemented
}
