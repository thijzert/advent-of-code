package aoc19

import (
	"fmt"
	"sync"
)

func runIntCodeProgram(program []int, input []int, output []int) (int, int, int, error) {
	chin, chout := make(chan int), make(chan int)
	inputPtr, outputPtr := 0, 0

	var wv sync.WaitGroup
	go func() {
		wv.Add(1)
		for v := range chout {
			output[outputPtr] = v
			outputPtr++
		}
		wv.Done()
	}()
	go func() {
		wv.Add(1)
		for _, v := range input {
			chin <- v
			inputPtr++
		}
		close(chin)
		wv.Done()
	}()

	mem0, err := startIntCodeProgram(program, chin, chout)
	wv.Wait()
	return mem0, inputPtr, outputPtr, err
}

func startIntCodeProgram(program []int, input chan int, output chan int) (int, error) {
	memory := make([]int, len(program))
	copy(memory, program)

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
			set(pc+1, modeB, <-input)
			pc += 2
		} else if opcode == 4 {
			output <- get(pc+1, modeB)
			pc += 2
		} else if opcode == 5 {
			if get(pc+1, modeA) != 0 {
				pc = get(pc+2, modeB)
			} else {
				pc += 3
			}
		} else if opcode == 6 {
			if get(pc+1, modeA) == 0 {
				pc = get(pc+2, modeB)
			} else {
				pc += 3
			}
		} else if opcode == 7 {
			v := 0
			if get(pc+1, modeA) < get(pc+2, modeB) {
				v = 1
			}
			set(pc+3, modeC, v)
			pc += 4
		} else if opcode == 8 {
			v := 0
			if get(pc+1, modeA) == get(pc+2, modeB) {
				v = 1
			}
			set(pc+3, modeC, v)
			pc += 4
		} else {
			return 0, fmt.Errorf("Undefined opcode %d", opcode)
		}
	}

	close(output)
	return memory[0], nil
}
