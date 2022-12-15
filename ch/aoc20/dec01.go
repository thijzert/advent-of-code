package aoc20

import (
	"errors"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec01a(ctx ch.AOContext) (interface{}, error) {
	expenses, err := ctx.DataAsInts("inputs/2020/dec01.txt")
	if err != nil {
		return nil, err
	}
	for i, a := range expenses {
		for _, b := range expenses[i+1:] {
			if a+b == 2020 {
				return a * b, nil
			}
		}
	}

	return nil, errors.New("failed to find the answer")
}

func Dec01b(ctx ch.AOContext) (interface{}, error) {
	expenses, err := ctx.DataAsInts("inputs/2020/dec01.txt")
	if err != nil {
		return nil, err
	}
	for i, a := range expenses {
		for j, b := range expenses[i+1:] {
			for _, c := range expenses[i+j+2:] {
				if a+b+c == 2020 {
					return a * b * c, nil
				}
			}
		}
	}

	return nil, errors.New("failed to find the answer")
}
