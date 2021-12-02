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
	ctx.FinalAnswer.Print(countHexBoard(board))
	return nil
}

func getHexBoard(directions []string, margin int) [][]bool {
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

	board := make([][]bool, maxY-minY+1)
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

func countHexBoard(board [][]bool) int {
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
		iterateHexBoard(bufBoard, board)
		// printBoard(bufBoard)
		board, bufBoard = bufBoard, board
		ctx.Debug.Printf("Day %3d: %d", i+1, countHexBoard(board))
	}

	// Pass 3: count everything
	ctx.FinalAnswer.Print(countHexBoard(board))
	return nil
}

func iterateHexBoard(buf, currentBoard [][]bool) {
	for y, row := range currentBoard {
		for x, black := range row {
			btiles := 0
			if y > 0 {
				if currentBoard[y-1][x] {
					btiles++
				}
				if x > 0 && currentBoard[y-1][x-1] {
					btiles++
				}
			}
			if x > 0 && row[x-1] {
				btiles++
			}
			if x < len(row)-1 && row[x+1] {
				btiles++
			}
			if y < len(currentBoard)-1 {
				if currentBoard[y+1][x] {
					btiles++
				}
				if x < len(row)-1 && currentBoard[y+1][x+1] {
					btiles++
				}
			}

			if black {
				//fmt.Printf("Found a black tile at %d,%d with %d black tiles around it", x, y, btiles)
				if btiles == 0 || btiles > 2 {
					black = false
					//fmt.Printf(" - flipping it to white")
				}
				//fmt.Printf("\n")
			} else {
				if btiles == 2 {
					//fmt.Printf("Flipping tile at %d,%d to black\n", x, y)
					black = true
				}
			}

			buf[y][x] = black
		}
	}
}
