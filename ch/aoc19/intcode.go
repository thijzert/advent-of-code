package aoc19

import (
	"fmt"
)

func runIntCodeProgram(program []int, input []int, output []int) (int, int, int, error) {
	memory := make([]int, len(program))
	copy(memory, program)

	inputPtr, outputPtr := 0, 0

	get := func(pc, mode int) int {
		if mode == 0 {
			return memory[memory[pc]]
		} else if mode == 1 {
			return memory[pc]
		}
		panic("invalid mode")
	}
	set := func(pc, mode, value int) {
		if mode == 0 {
			memory[memory[pc]] = value
		} else if mode == 1 {
			panic("cannot set an immediate value")
		} else {
			panic("invalid mode")
		}
	}

	pc := 0
	for pc >= 0 && pc <= len(memory) && memory[pc] != 99 {
		opcode := memory[pc] % 100
		modeA, modeB, modeC := (memory[pc]/100)%10, (memory[pc]/1000)%10, (memory[pc]/10000)%10

		if opcode == 1 {
			set(pc+3, modeC, get(pc+1, modeA)+get(pc+2, modeB))
			pc += 4
		} else if opcode == 2 {
			set(pc+3, modeC, get(pc+1, modeA)*get(pc+2, modeB))
			pc += 4
		} else if opcode == 3 {
			set(pc+1, modeB, input[inputPtr])
			inputPtr++
			pc += 2
		} else if opcode == 4 {
			output[outputPtr] = get(pc+1, modeB)
			outputPtr++
			pc += 2
		} else {
			return 0, inputPtr, outputPtr, fmt.Errorf("Undefined opcode %d", opcode)
		}
	}

	return memory[0], inputPtr, outputPtr, nil
}
