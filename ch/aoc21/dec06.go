package aoc21

import (
	"github.com/thijzert/advent-of-code/ch"
)

func Dec06a(ctx ch.AOContext) (interface{}, error) {
	lanternfish, err := ctx.DataAsIntLists("inputs/2021/dec06.txt")
	if err != nil {
		return nil, err
	}

	_, ex1 := fibonacciSimulate([]int{3, 4, 3, 1, 2}, 8, 6, 18)
	_, ex2 := fibonacciSimulate([]int{3, 4, 3, 1, 2}, 8, 6, 80)
	ctx.Printf("In the example, after 18 days, there are a total of %d fish. After 80 days, there would be a total of %d.", ex1, ex2)

	histogram, rv := fibonacciSimulate(lanternfish[0], 8, 6, 80)
	ctx.Print(histogram)
	return rv, nil
}

func Dec06b(ctx ch.AOContext) (interface{}, error) {
	lanternfish, err := ctx.DataAsIntLists("inputs/2021/dec06.txt")
	if err != nil {
		return nil, err
	}

	_, ex := fibonacciSimulate([]int{3, 4, 3, 1, 2}, 8, 6, 256)
	ctx.Printf("Example data: %d", ex)

	histogram, rv := fibonacciSimulate(lanternfish[0], 8, 6, 256)
	ctx.Print(histogram)
	return rv, nil
}

func fibonacciSimulate(ages []int, initialAge, reproductionInterval, duration int) ([]int, int) {
	histogram := make([]int, initialAge+1)
	for _, timer := range ages {
		histogram[timer]++
	}

	for i := 0; i < duration; i++ {
		multiplying := histogram[0]
		copy(histogram, histogram[1:])
		histogram[initialAge] = multiplying
		histogram[reproductionInterval] += multiplying
	}

	s := 0
	for _, n := range histogram {
		s += n
	}

	return histogram, s
}
