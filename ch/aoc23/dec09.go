package aoc23

import (
	"github.com/thijzert/advent-of-code/ch"
)

func dec9Predict(values []int) int {
	if len(values) == 0 {
		return 0
	}
	next := make([]int, len(values)-1)
	allZero := true
	for i := range next {
		next[i] = values[i+1] - values[i]
		allZero = allZero && next[i] == 0
	}
	if allZero {
		return values[len(values)-1]
	} else {
		return values[len(values)-1] + dec9Predict(next)
	}
}

func Dec09a(ctx ch.AOContext) (interface{}, error) {
	lists, err := ctx.DataAsIntLists("inputs/2023/dec09.txt")
	if err != nil {
		return nil, err
	}
	// lists = [][]int{{0, 3, 6, 9, 12, 15}, {1, 3, 6, 10, 15, 21}, {10, 13, 16, 21, 30, 45}}

	answer := 0
	for _, values := range lists {
		pred := dec9Predict(values)
		//ctx.Printf("for [%d], next value is %d", values, pred)
		answer += pred
	}

	return answer, nil
}

func Dec09b(ctx ch.AOContext) (interface{}, error) {
	lists, err := ctx.DataAsIntLists("inputs/2023/dec09.txt")
	if err != nil {
		return nil, err
	}
	// lists = [][]int{{10, 13, 16, 21, 30, 45}}

	answer := 0
	for _, values := range lists {
		for i := range values[:len(values)/2] {
			values[i], values[len(values)-i-1] = values[len(values)-i-1], values[i]
		}
		pred := dec9Predict(values)
		//ctx.Printf("for [%d], preceding value is %d", values, pred)
		answer += pred
	}

	return answer, nil
}
