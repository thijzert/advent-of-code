package aoc20

import "github.com/thijzert/advent-of-code/ch"

func Dec05a(ctx ch.AOContext) error {
	toTest := []struct {
		BoardingPass string
		SeatID       int
	}{
		{"FBFBBFFRLR", 357},
		{"BFFFBBFRRR", 567},
		{"FFFBBBFRRR", 119},
		{"BBFFBBFRLL", 820},
	}

	for _, tc := range toTest {
		i := decodeSeatNumber(tc.BoardingPass)
		ctx.Printf("Seat '%s' → %d (== %d: %v)", tc.BoardingPass, i, tc.SeatID, i == tc.SeatID)
	}

	lines, err := ctx.DataLines("inputs/2020/dec05.txt")
	if err != nil {
		return err
	}
	max := 0
	for _, seat := range lines {
		id := decodeSeatNumber(seat)
		if id > max {
			max = id
		}
	}
	ctx.FinalAnswer.Print(max)
	return nil
}

func Dec05b(ctx ch.AOContext) error {
	lines, err := ctx.DataLines("inputs/2020/dec05.txt")
	if err != nil {
		return err
	}
	seen := make([]bool, 128*8)
	for _, seat := range lines {
		id := decodeSeatNumber(seat)
		seen[id] = true
	}

	rv := 0
	for id, s := range seen {
		if id < 1 || id >= len(seen)-2 {
			continue
		}
		if !s && seen[id-1] && seen[id+1] {
			rv = id
		}
	}

	ctx.FinalAnswer.Print(rv)
	return nil
}

func decodeSeatNumber(boardingPass string) int {
	x, y := 0, 0
	dx, dy := 4, 64

	for _, c := range boardingPass {
		if c == 'B' {
			y += dy
		} else if c == 'R' {
			x += dx
		}
		if c == 'B' || c == 'F' {
			dy >>= 1
		} else if c == 'R' || c == 'L' {
			dx >>= 1
		}
	}

	return 8*y + x
}
