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
	program = []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}

	input := make([]int, 0)
	output := make([]int, 100)
	_, _, no, err := runIntCodeProgram(program, input, output)
	if err != nil {
		return nil, err
	}
	ctx.Printf("Output: %d", output[:no])

	return nil, errNotImplemented
}

var Dec09b ch.AdventFunc = nil

// func Dec09b(ctx ch.AOContext) (interface{}, error) {
// 	return nil, errNotImplemented
// }
