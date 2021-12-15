package aoc20

import (
	"sort"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec10a(ctx ch.AOContext) error {
	inputs := []int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}
	d1, d3, _ := joltageDifferences(inputs)
	ctx.Printf("Example data A: %d×%d = %d", d1, d3, d1*d3)

	inputs = []int{28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3}
	d1, d3, _ = joltageDifferences(inputs)
	ctx.Printf("Example data B: %d×%d = %d", d1, d3, d1*d3)

	inputs, err := ctx.DataAsInts("inputs/2020/dec10.txt")
	if err != nil {
		return err
	}

	d1, d3, _ = joltageDifferences(inputs)
	ctx.Printf("Final data: %d×%d = %d, final voltage %d", d1, d3, d1*d3, inputs[len(inputs)-1]+3)
	ctx.FinalAnswer.Print(d1 * d3)
	return nil
}

func Dec10b(ctx ch.AOContext) error {
	inputs := []int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}
	n := numArrangements(inputs, 0, 0, nil)
	ctx.Printf("Example data A: can be arranged %d ways", n)

	inputs = []int{28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3}
	n = numArrangements(inputs, 0, 0, nil)
	ctx.Printf("Example data B: can be arranged %d ways", n)

	inputs, err := ctx.DataAsInts("inputs/2020/dec10.txt")
	if err != nil {
		return err
	}

	n = numArrangements(inputs, 0, 0, nil)
	ctx.Printf("Final data can be arranged %d ways", n)

	ctx.FinalAnswer.Print(n)
	return nil
}

func joltageDifferences(adapters []int) (diff1, diff3, deviceJoltage int) {
	sort.Ints(adapters)
	last := 0
	for _, jolts := range adapters {
		if jolts-last == 1 {
			diff1++
		} else if jolts-last == 3 {
			diff3++
		}
		last = jolts
	}
	diff3++
	deviceJoltage = last + 3
	return
}

func numArrangements(adapters []int, adapterOffset, joltage int, memory []map[int]int) int {
	if memory == nil {
		sort.Ints(adapters)
		memory = make([]map[int]int, len(adapters))
		for i := range memory {
			memory[i] = make(map[int]int)
		}
	}

	if adapterOffset >= len(adapters)-1 {
		return 1
	}

	if v, ok := memory[adapterOffset][joltage]; ok {
		return v
	}

	rv := 0
	defer func() {
		memory[adapterOffset][joltage] = rv
	}()

	for i, v := range adapters[adapterOffset:] {
		if v > joltage+3 {
			break
		}
		rv += numArrangements(adapters, adapterOffset+i+1, v, memory)
	}

	return rv
}
