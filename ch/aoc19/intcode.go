package aoc19

import (
	"fmt"
	"math/big"
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
	memory := make([]*big.Int, len(program))
	for i, n := range program {
		memory[i] = big.NewInt(int64(n))
	}

	base := 0

	get := func(pc, mode int) *big.Int {
		if mode == 0 {
			addr := int(memory[pc].Int64())
			if addr >= len(memory) {
				return big.NewInt(0)
			}
			return memory[addr]
		} else if mode == 1 {
			return memory[pc]
		} else if mode == 2 {
			addr := base + int(memory[pc].Int64())
			if addr >= len(memory) {
				return big.NewInt(0)
			}
			return memory[addr]
		}
		panic("invalid mode")
	}
	set := func(pc, mode int, value *big.Int) {
		if mode == 0 {
			addr := int(memory[pc].Int64())
			if addr > 0 && addr < 0x1000000 {
				for len(memory) <= addr {
					memory = append(memory, big.NewInt(0))
				}
			}
			memory[addr].Set(value)
		} else if mode == 1 {
			panic("cannot set an immediate value")
		} else if mode == 2 {
			addr := base + int(memory[pc].Int64())
			if addr > 0 && addr < 0x1000000 {
				for len(memory) <= addr {
					memory = append(memory, big.NewInt(0))
				}
			}
			memory[addr].Set(value)
		} else {
			panic("invalid mode")
		}
	}

	pc := 0
	for pc >= 0 && pc <= len(memory) {
		instr := int(memory[pc].Int64())
		opcode := instr % 100
		modeA, modeB, modeC := (instr/100)%10, (instr/1000)%10, (instr/10000)%10

		if opcode == OP_ADD {
			v := big.NewInt(0)
			v.Add(get(pc+1, modeA), get(pc+2, modeB))
			set(pc+3, modeC, v)
			pc += 4
		} else if opcode == OP_MULT {
			v := big.NewInt(0)
			v.Mul(get(pc+1, modeA), get(pc+2, modeB))
			set(pc+3, modeC, v)
			pc += 4
		} else if opcode == OP_READ {
			set(pc+1, modeA, big.NewInt(int64(<-input)))
			pc += 2
		} else if opcode == OP_WRITE {
			v := get(pc+1, modeA)
			output <- int(v.Int64())
			pc += 2
		} else if opcode == OP_JTRUE {
			v := get(pc+1, modeA)
			if !v.IsInt64() || v.Int64() != 0 {
				pc = int(get(pc+2, modeB).Int64())
			} else {
				pc += 3
			}
		} else if opcode == OP_JFALSE {
			v := get(pc+1, modeA)
			if v.IsInt64() && v.Int64() == 0 {
				pc = int(get(pc+2, modeB).Int64())
			} else {
				pc += 3
			}
		} else if opcode == OP_LESS {
			v := big.NewInt(0)
			if get(pc+1, modeA).Cmp(get(pc+2, modeB)) < 0 {
				v.SetInt64(1)
			}
			set(pc+3, modeC, v)
			pc += 4
		} else if opcode == OP_EQUALS {
			v := big.NewInt(0)
			if get(pc+1, modeA).Cmp(get(pc+2, modeB)) == 0 {
				v.SetInt64(1)
			}
			set(pc+3, modeC, v)
			pc += 4
		} else if opcode == OP_CHRB {
			base += int(get(pc+1, modeA).Int64())
			pc += 2
		} else if opcode == OP_HALT {
			break
		} else {
			return 0, fmt.Errorf("Undefined opcode %d", opcode)
		}
	}

	close(output)
	return int(memory[0].Int64()), nil
}
