package aoc19

import (
	"github.com/thijzert/advent-of-code/ch"
)

func dec07maxThrust(ctx ch.AOContext, program []int, phaseMin int, phaseMax int, thrusterValue func([5]int) int) ([5]int, int) {
	phases := [5]int{}
	maxPhase, maxThrust := [5]int{}, 0
	var findMaxThrust func(i int)
	findMaxThrust = func(i int) {
		if i >= len(phases) {
			thr := thrusterValue(phases)
			if thr > maxThrust {
				maxPhase, maxThrust = phases, thr
			}
		} else {
			for a := phaseMin; a <= phaseMax; a++ {
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
	return maxPhase, maxThrust
}

func Dec07a(ctx ch.AOContext) (interface{}, error) {
	programs, err := ctx.DataAsIntLists("inputs/2019/dec07.txt")
	if err != nil {
		return nil, err
	}
	program := programs[0]
	//program = []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}
	//program = []int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}

	maxPhase, maxThrust := dec07maxThrust(ctx, program, 0, 4, func(phases [5]int) int {
		thrust := 0
		for _, a := range phases {
			input := []int{a, thrust}
			output := []int{0}
			runIntCodeProgram(program, input, output)
			thrust = output[0]
		}
		return thrust
	})
	ctx.Printf("Max thrust value %d, from phase setting %d", maxThrust, maxPhase)

	return maxThrust, nil
}

func Dec07b(ctx ch.AOContext) (interface{}, error) {
	programs, err := ctx.DataAsIntLists("inputs/2019/dec07.txt")
	if err != nil {
		return nil, err
	}
	program := programs[0]
	//program = []int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26, 27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5}

	maxPhase, maxThrust := dec07maxThrust(ctx, program, 5, 9, func(phases [5]int) int {
		finished := 0
		thrust := 0
		var amps [5]struct {
			In, Out chan int
		}
		for i, a := range phases {
			amps[i].In = make(chan int)
			amps[i].Out = make(chan int)
			go func() {
				startIntCodeProgram(program, amps[i].In, amps[i].Out)
				finished++
			}()
			amps[i].In <- a
		}

		for finished < 5 {
			for _, amp := range amps {
				amp.In <- thrust
				thrust = <-amp.Out
			}
		}
		for _, amp := range amps {
			close(amp.In)
			close(amp.Out)
		}

		return thrust
	})
	ctx.Printf("Max thrust value %d, from phase setting %d", maxThrust, maxPhase)

	return maxThrust, nil
}
