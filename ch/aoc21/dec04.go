package aoc21

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec04a(ctx ch.AOContext) error {
	lines, err := ctx.DataLines("inputs/2021/dec04.txt")
	if err != nil {
		return err
	}

	drawStr := strings.Split(lines[0], ",")
	draw := make([]int, len(drawStr))
	for i, s := range drawStr {
		draw[i], _ = strconv.Atoi(s)
	}
	lines = lines[2:]

	ctx.Debug.Printf("Draw: %d", draw)

	minMoves := len(draw) + 5
	finalScore := 0

	for len(lines) > 1 {
		var i int
		var l string
		for i, l = range lines {
			if l == "" {
				break
			}
		}

		bs := newSheet(lines[:i])
		moves, score := bs.PlayBingo(draw)

		if moves < minMoves {
			ctx.Debug.Printf("\n%s", bs)
			ctx.Debug.Printf("this sheet won in %d moves with %d points - a new record", moves, score)

			minMoves = moves
			finalScore = score
		}

		lines = lines[i+1:]
	}

	ctx.FinalAnswer.Print(finalScore)
	return nil
}

type bingoSheet struct {
	Size     int
	Contents []int
}

func newSheet(lines []string) bingoSheet {
	S := len(lines)
	rv := bingoSheet{
		Size:     S,
		Contents: make([]int, S*S),
	}

	for i, l := range lines {
		ptrs := make([]interface{}, S)
		for j := 0; j < S; j++ {
			ptrs[j] = &rv.Contents[i*S+j]
		}
		fmt.Sscan(l, ptrs...)
	}

	return rv
}

func (bs bingoSheet) String() string {
	rv := ""
	for i := 0; i < bs.Size; i++ {
		if i > 0 {
			rv += "\n"
		}
		for j := 0; j < bs.Size; j++ {
			rv += fmt.Sprintf("%3d", bs.Contents[bs.Size*i+j])
		}
	}
	return rv
}

func (bs bingoSheet) PlayBingo(draw []int) (moves, score int) {
	rows := make([]int, bs.Size)
	cols := make([]int, bs.Size)
	for i := range rows {
		rows[i] = bs.Size
		cols[i] = bs.Size
	}

	for _, ball := range draw {
		moves++
		// Find the index of this ball
		var i, c int
		for i, c = range bs.Contents {
			if c == ball {
				break
			}
		}
		if c != ball {
			continue
		}

		rowIdx := i / bs.Size
		colIdx := i % bs.Size

		bs.Contents[i] *= -1
		rows[rowIdx]--
		cols[rowIdx]--

		if rows[rowIdx] == 0 || cols[colIdx] == 0 {
			// We won!
			score = 0
			for _, c := range bs.Contents {
				if c >= 0 {
					score += c
				}
			}
			score *= ball
			return
		}
	}

	return 0, 0
}

func Dec04b(ctx ch.AOContext) error {
	return fmt.Errorf("not implemented")
}
