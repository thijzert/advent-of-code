package aoc19

import (
	"github.com/thijzert/advent-of-code/ch"
)

func Dec05a(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataAsIntLists("inputs/2019/dec05.txt")
	if err != nil {
		return nil, err
	}
	program := lines[0]
	//program = []int{3, 0, 4, 0, 99}
	//program = []int{1002, 4, 3, 4, 33}
	//program = []int{1101, 100, -1, 4, 0}

	in, out := []int{1}, make([]int, 100)
	ans, _, outptr, err := runIntCodeProgram(program, in, out)
	if err != nil {
		return nil, err
	}
	ctx.Printf("exit status: %d", ans)
	ctx.Printf("output: %d", out[:outptr])
	if outptr == 0 {
		return nil, errFailed
	}

	return out[outptr-1], nil
}

func Dec05b(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataAsIntLists("inputs/2019/dec05.txt")
	if err != nil {
		return nil, err
	}
	program := lines[0]

	in, out := []int{5}, make([]int, 100)
	ans, _, outptr, err := runIntCodeProgram(program, in, out)
	if err != nil {
		return nil, err
	}
	ctx.Printf("exit status: %d", ans)
	ctx.Printf("output: %d", out[:outptr])
	if outptr == 0 {
		return nil, errFailed
	}

	return out[outptr-1], nil
}
