package aoc21

import (
	"errors"

	"github.com/thijzert/advent-of-code/ch"
)

var errNotImplemented = errors.New("not implemented")

func ExampleChallengeA(ctx ch.AOContext) error {
	lines, err := ctx.DataLines("inputs/2021/example.txt")
	if err != nil {
		return err
	}
	ctx.Print(len(lines))

	return errNotImplemented
}

func ExampleChallengeB(ctx ch.AOContext) error {
	return errNotImplemented
}
