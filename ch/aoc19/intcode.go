package aoc19

import (
	"fmt"
	"sync"
)

const (
	OP_ADD    int = 1
	OP_MULT       = 2
	OP_READ       = 3
	OP_WRITE      = 4
	OP_JTRUE      = 5 // Jump if True
	OP_JFALSE     = 6 // Jump if False
	OP_LESS       = 7
	OP_EQUALS     = 8
	OP_CHRB       = 9 // Change Relative Base
	OP_HALT       = 99
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

	base := 0

	get := func(pc, mode int) int {
		if mode == 0 {
			addr := memory[pc]
			if addr >= len(memory) {
				return 0
			}
			return memory[memory[pc]]
		} else if mode == 1 {
			return memory[pc]
		} else if mode == 2 {
			addr := base + memory[pc]
			if addr >= len(memory) {
				return 0
			}
			return memory[addr]
		}
		panic("invalid mode")
	}
	set := func(pc, mode, value int) {
		if mode == 0 {
			addr := memory[pc]
			if addr > 0 && addr < 0x1000000 {
				for len(memory) <= addr {
					memory = append(memory, 0)
				}
			}
			memory[memory[pc]] = value
		} else if mode == 1 {
			panic("cannot set an immediate value")
		} else if mode == 2 {
			addr := base + memory[pc]
			if addr > 0 && addr < 0x1000000 {
				for len(memory) < addr {
					memory = append(memory, 0)
				}
			}
			memory[addr] = value
		} else {
			panic("invalid mode")
		}
	}

	pc := 0
	for pc >= 0 && pc <= len(memory) && memory[pc] != OP_HALT {
		opcode := memory[pc] % 100
		modeA, modeB, modeC := (memory[pc]/100)%10, (memory[pc]/1000)%10, (memory[pc]/10000)%10

		if opcode == OP_ADD {
			set(pc+3, modeC, get(pc+1, modeA)+get(pc+2, modeB))
			pc += 4
		} else if opcode == OP_MULT {
			set(pc+3, modeC, get(pc+1, modeA)*get(pc+2, modeB))
			pc += 4
		} else if opcode == OP_READ {
			set(pc+1, modeA, <-input)
			pc += 2
		} else if opcode == OP_WRITE {
			output <- get(pc+1, modeA)
			pc += 2
		} else if opcode == OP_JTRUE {
			if get(pc+1, modeA) != 0 {
				pc = get(pc+2, modeB)
			} else {
				pc += 3
			}
		} else if opcode == OP_JFALSE {
			if get(pc+1, modeA) == 0 {
				pc = get(pc+2, modeB)
			} else {
				pc += 3
			}
		} else if opcode == OP_LESS {
			v := 0
			if get(pc+1, modeA) < get(pc+2, modeB) {
				v = 1
			}
			set(pc+3, modeC, v)
			pc += 4
		} else if opcode == OP_EQUALS {
			v := 0
			if get(pc+1, modeA) == get(pc+2, modeB) {
				v = 1
			}
			set(pc+3, modeC, v)
			pc += 4
		} else if opcode == OP_CHRB {
			base += get(pc+1, modeA)
			pc += 2
		} else {
			return 0, fmt.Errorf("Undefined opcode %d", opcode)
		}
	}

	close(output)
	return memory[0], nil
}
