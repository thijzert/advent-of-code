package aoc20

import (
	"fmt"
	"strings"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec02a(ctx ch.AOContext) error {
	lines, err := ctx.DataLines("inputs/2020/dec02.txt")
	if err != nil {
		return err
	}

	validPasswords := 0
	for _, line := range lines {
		pts := strings.Split(line, ": ")
		if len(pts) != 2 {
			continue
		}

		var ch rune
		var rmin, rmax int
		fmt.Sscanf(pts[0], "%d-%d %c", &rmin, &rmax, &ch)

		count := 0
		for _, c := range pts[1] {
			if c == ch {
				count++
			}
		}
		if count >= rmin && count <= rmax {
			validPasswords++
		}
	}

	ctx.FinalAnswer.Print(validPasswords)
	return nil
}

func Dec02b(ctx ch.AOContext) error {
	lines, err := ctx.DataLines("inputs/2020/dec02.txt")
	if err != nil {
		return err
	}

	validPasswords := 0
	for _, line := range lines {
		pts := strings.Split(line, ": ")
		if len(pts) != 2 {
			continue
		}

		var ch byte
		var idx [2]int
		fmt.Sscanf(pts[0], "%d-%d %c", &idx[0], &idx[1], &ch)

		count := 0
		for _, i := range idx {
			if pts[1][i-1] == ch {
				count++
			}
		}
		if count == 1 {
			validPasswords++
		}
	}

	ctx.FinalAnswer.Print(validPasswords)
	return errNotImplemented
}
