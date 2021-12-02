package aoc20

import (
	"errors"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec24a(ctx ch.AOContext) error {
	lines, err := ctx.DataLines("inputs/2020/dec24.txt")
	if err != nil {
		return err
	}

	minX, minY, maxX, maxY := 0, 0, 0, 0

	// Pass 1: determine the max size of the grid
	for _, line := range lines {
		if line == "" {
			continue
		}
		x, y := hexBoardPos(line)

		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
	}

	ctx.Debug.Printf("Board dimensions: [%d,%d] - [%d,%d]", minX, minY, maxX, maxY)

	board := make([][]bool, maxY-minY+1)
	for i := range board {
		board[i] = make([]bool, maxX-minX+1)
	}

	// Pass 2: flip tiles on the grid
	for _, line := range lines {
		if line == "" {
			continue
		}
		x, y := hexBoardPos(line)

		board[y-minY][x-minX] = !board[y-minY][x-minX]
	}

	// Pass 3: count everything
	rv := 0
	for _, row := range board {
		for _, cell := range row {
			if cell {
				rv++
			}
		}
	}

	ctx.FinalAnswer.Print(rv)
	return nil
}

func hexBoardPos(directions string) (int, int) {
	x, y := 0, 0
	last := '-'
	for _, c := range directions {
		if c == 'e' {
			if last == 'n' {
				x += 1
				y += 1
			} else if last == 's' {
				y -= 1
			} else {
				x += 1
			}
		} else if c == 'w' {
			if last == 'n' {
				y += 1
			} else if last == 's' {
				x -= 1
				y -= 1
			} else {
				x -= 1
			}
		}
		last = c
	}
	return x, y
}

func Dec24b(ctx ch.AOContext) error {
	return errors.New("not implemented")
}
