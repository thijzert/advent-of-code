package ch

import (
	"fmt"

	"github.com/pkg/errors"
)

func Dec02a(ctx AOContext) error {
	depth := 0
	hPosition := 0

	movement, err := ctx.DataLines("inputs/dec02.txt")
	if err != nil {
		return err
	}

	for _, l := range movement {
		if l == "" {
			continue
		}
		var direction string
		var dist int
		_, err = fmt.Sscanf(l, "%s %d", &direction, &dist)
		if err != nil {
			return errors.Wrapf(err, "invalid line '%s'", l)
		}

		if direction == "forward" {
			hPosition += dist
		} else if direction == "up" {
			depth -= dist
		} else if direction == "down" {
			depth += dist
		} else {
			return fmt.Errorf("invalid line '%s'", l)
		}
	}

	ctx.Debug.Printf("Depth:  %d", depth)
	ctx.Debug.Printf("H. pos: %d", hPosition)

	ctx.FinalAnswer.Print(depth * hPosition)
	return nil
}

func Dec02b(ctx AOContext) error {
	return notImplemented
}
