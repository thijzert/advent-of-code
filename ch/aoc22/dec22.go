package aoc22

import (
	"fmt"
	"strings"

	"github.com/thijzert/advent-of-code/ch"
	"github.com/thijzert/advent-of-code/lib/cube"
	"github.com/thijzert/advent-of-code/lib/image"
)

const (
	TILE int = 2
	WALL     = 5
)

func Dec22a(ctx ch.AOContext) (interface{}, error) {
	sections, err := ctx.DataSections("inputs/2022/dec22.txt")
	if err != nil {
		return nil, err
	}

	board := image.ReadImage(sections[0], func(r rune) int {
		if r == '.' {
			return TILE
		} else if r == '#' {
			return WALL
		}
		return 0
	})

	ctx.Printf("board:\n%s", board)

	instructions := sections[1][0]
	dir := 0
	pos := walkBoard(board, cube.Point{0, 0}, cube.Cardinal2D[dir])
	ctx.Printf("Position: %v, direction %v", pos, dir)

	for len(instructions) > 0 {
		j := strings.IndexAny(instructions, "RL")
		steps := 0
		if j < 0 {
			j = len(instructions)
		}
		_, err = fmt.Sscanf(instructions[:j], "%d", &steps)
		if err != nil {
			return nil, err
		}
		for i := 0; i < steps; i++ {
			pos = walkBoard(board, pos, cube.Cardinal2D[dir])
		}
		if j >= len(instructions) {
			ctx.Printf("final position: %v, direction %v", pos, dir)
			break
		} else if instructions[j] == 'R' {
			dir = (dir + 1) % 4
		} else if instructions[j] == 'L' {
			dir = (dir + 3) % 4
		}
		instructions = instructions[j+1:]
		ctx.Printf("Position: %v, direction %v; Remaining instructions: '%s'", pos, dir, instructions)
	}

	return 1000*(pos.Y+1) + 4*(pos.X+1) + dir, nil
}

var Dec22b ch.AdventFunc = nil

// func Dec22b(ctx ch.AOContext) (interface{}, error) {
// 	return nil, errNotImplemented
// }

func walkBoard(board *image.Image, start, dir cube.Point) cube.Point {
	nx := start
	nx.X = (nx.X + dir.X + 2*board.Width) % board.Width
	nx.Y = (nx.Y + dir.Y + 2*board.Height) % board.Height

	for board.At(nx.X, nx.Y) == 0 {
		nx.X = (nx.X + dir.X + 2*board.Width) % board.Width
		nx.Y = (nx.Y + dir.Y + 2*board.Height) % board.Height
	}

	for board.At(nx.X, nx.Y) == WALL {
		return start
	}
	return nx
}
