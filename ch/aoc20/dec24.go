package aoc20

import (
	"github.com/thijzert/advent-of-code/ch"
)

func Dec24a(ctx ch.AOContext) error {
	lines, err := ctx.DataLines("inputs/2020/dec24.txt")
	if err != nil {
		return err
	}

	board := getHexBoard(lines, 0)
	ctx.Debug.Printf("Board dimensions: %d×%d", len(board), len(board[0]))

	// Pass 3: count everything
	ctx.FinalAnswer.Print(board.Count())
	return nil
}

type hexBoard [][]bool

func getHexBoard(directions []string, margin int) hexBoard {
	minX, minY, maxX, maxY := 0, 0, 0, 0

	// Pass 1: determine the max size of the grid
	for _, line := range directions {
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

	minX -= margin
	minY -= margin
	maxX += margin
	maxY += margin

	board := make(hexBoard, maxY-minY+1)
	for i := range board {
		board[i] = make([]bool, maxX-minX+1)
	}

	// Pass 2: flip tiles on the grid
	for _, line := range directions {
		if line == "" {
			continue
		}
		x, y := hexBoardPos(line)

		board[y-minY][x-minX] = !board[y-minY][x-minX]
	}

	return board
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

func (board hexBoard) Count() int {
	rv := 0
	for _, row := range board {
		for _, black := range row {
			if black {
				rv++
			}
		}
	}
	return rv
}

func (board hexBoard) At(x, y int) bool {
	if y < 0 || y >= len(board) {
		return false
	}
	if x < 0 || x >= len(board[y]) {
		return false
	}

	return board[x][y]
}

func (board hexBoard) AdjacentBlack(x, y int) int {
	b := func(x bool) int {
		if x {
			return 1
		} else {
			return 0
		}
	}

	return b(board.At(y+1, x)) +
		b(board.At(y+1, x+1)) +
		b(board.At(y, x-1)) +
		b(board.At(y, x+1)) +
		b(board.At(y-1, x-1)) +
		b(board.At(y-1, x))
}

func Dec24b(ctx ch.AOContext) error {
	lines, err := ctx.DataLines("inputs/2020/dec24.txt")
	if err != nil {
		return err
	}

	board := getHexBoard(lines, 100)

	bufBoard := make([][]bool, len(board))
	for y, row := range board {
		bufBoard[y] = make([]bool, len(row))
	}

	ctx.Debug.Printf("Board dimensions: %d×%d", len(board), len(board[0]))

	for i := 0; i < 100; i++ {
		board.iterate(bufBoard)
		board, bufBoard = bufBoard, board
		ctx.Debug.Printf("Day %3d: %d", i+1, board.Count())
	}

	// Pass 3: count everything
	ctx.FinalAnswer.Print(board.Count())
	return nil
}

func (currentBoard hexBoard) iterate(buf hexBoard) {
	for y, row := range currentBoard {
		for x, black := range row {
			btiles := currentBoard.AdjacentBlack(x, y)

			if black {
				if btiles == 0 || btiles > 2 {
					black = false
				}
			} else {
				if btiles == 2 {
					black = true
				}
			}

			buf[y][x] = black
		}
	}
}
