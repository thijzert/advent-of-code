package aoc21

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/thijzert/advent-of-code/ch"
)

func Dec02a(ctx ch.AOContext) (interface{}, error) {
	depth := 0
	hPosition := 0

	movement, err := ctx.DataLines("inputs/2021/dec02.txt")
	if err != nil {
		return nil, err
	}

	for _, l := range movement {
		if l == "" {
			continue
		}
		var direction string
		var dist int
		_, err = fmt.Sscanf(l, "%s %d", &direction, &dist)
		if err != nil {
			return nil, errors.Wrapf(err, "invalid line '%s'", l)
		}

		if direction == "forward" {
			hPosition += dist
		} else if direction == "up" {
			depth -= dist
		} else if direction == "down" {
			depth += dist
		} else {
			return nil, fmt.Errorf("invalid line '%s'", l)
		}
	}

	ctx.Debug.Printf("Depth:  %d", depth)
	ctx.Debug.Printf("H. pos: %d", hPosition)

	return depth * hPosition, nil
}

func Dec02b(ctx ch.AOContext) (interface{}, error) {
	aim := 0
	depth := 0
	hPosition := 0

	movement, err := ctx.DataLines("inputs/2021/dec02.txt")
	if err != nil {
		return nil, err
	}

	for _, l := range movement {
		if l == "" {
			continue
		}
		var direction string
		var dist int
		_, err = fmt.Sscanf(l, "%s %d", &direction, &dist)
		if err != nil {
			return nil, errors.Wrapf(err, "invalid line '%s'", l)
		}

		if direction == "forward" {
			hPosition += dist
			depth += aim * dist
		} else if direction == "up" {
			aim -= dist
		} else if direction == "down" {
			aim += dist
		} else {
			return nil, fmt.Errorf("invalid line '%s'", l)
		}
	}

	ctx.Debug.Printf("Depth:  %d", depth)
	ctx.Debug.Printf("H. pos: %d", hPosition)

	return depth * hPosition, nil
}
