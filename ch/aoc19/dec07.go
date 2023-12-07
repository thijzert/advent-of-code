package aoc19

import (
	"github.com/thijzert/advent-of-code/ch"
)

func Dec07a(ctx ch.AOContext) (interface{}, error) {
	programs, err := ctx.DataAsIntLists("inputs/2019/dec07.txt")
	if err != nil {
		return nil, err
	}
	program := programs[0]
	//program = []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}
	//program = []int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}

	phases := [5]int{}
	maxPhase, maxThrust := [5]int{}, 0
	var findMaxThrust func(i int)
	thrusterValue := func() int {
		thrust := 0
		for _, a := range phases {
			input := []int{a, thrust}
			output := []int{0}
			runIntCodeProgram(program, input, output)
			thrust = output[0]
		}
		return thrust
	}
	findMaxThrust = func(i int) {
		if i >= len(phases) {
			thr := thrusterValue()
			if thr > maxThrust {
				maxPhase, maxThrust = phases, thr
			}
		} else {
			for a := 0; a <= 4; a++ {
				phases[i] = a
				ok := true
				for _, b := range phases[:i] {
					ok = ok && a != b
				}
				if ok {
					findMaxThrust(i + 1)
				}
			}
		}
	}

	findMaxThrust(0)
	ctx.Printf("Max thrust value %d, from phase setting %d", maxThrust, maxPhase)

	return maxThrust, nil
}

func Dec07b(ctx ch.AOContext) (interface{}, error) {
	return nil, errNotImplemented
}
