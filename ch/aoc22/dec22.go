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
	return dec22(ctx, "inputs/2022/dec22.txt", walkBoard)
}

func Dec22b(ctx ch.AOContext) (interface{}, error) {
	return dec22(ctx, "inputs/2022/dec22a.txt", walkCubeExample)
}

type boardWalker func(*image.Image, cube.Point, int) (cube.Point, int)

func dec22(ctx ch.AOContext, name string, w boardWalker) (interface{}, error) {
	sections, err := ctx.DataSections(name)
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

	instructions := sections[1][0]
	dir := 0
	pos, dir := walkBoard(board, cube.Point{0, 0}, dir)
	board.Set(pos.X, pos.Y, 1)
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
			pos, dir = w(board, pos, dir)
			board.Set(pos.X, pos.Y, 1)
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
		// ctx.Printf("Position: %v, direction %v; Remaining instructions: '%s'", pos, dir, instructions)
	}
	ctx.Printf("board:\n%s", board)

	return 1000*(pos.Y+1) + 4*(pos.X+1) + dir, nil
}

func walkBoard(board *image.Image, start cube.Point, diridx int) (cube.Point, int) {
	dir := cube.Cardinal2D[diridx]
	nx := start
	nx.X = (nx.X + dir.X + 2*board.Width) % board.Width
	nx.Y = (nx.Y + dir.Y + 2*board.Height) % board.Height

	for board.At(nx.X, nx.Y) == 0 {
		nx.X = (nx.X + dir.X + 2*board.Width) % board.Width
		nx.Y = (nx.Y + dir.Y + 2*board.Height) % board.Height
	}

	for board.At(nx.X, nx.Y) == WALL {
		return start, diridx
	}
	return nx, diridx
}

func walkCubeExample(board *image.Image, start cube.Point, diridx int) (cube.Point, int) {
	dir := cube.Cardinal2D[diridx]
	nx := start
	nx.X = (nx.X + dir.X + 2*board.Width) % board.Width
	nx.Y = (nx.Y + dir.Y + 2*board.Height) % board.Height

	lastDir := diridx
	for board.At(nx.X, nx.Y) == 0 {
		lastDir = diridx
		// This hardcodes the cube net in the example data
		if diridx == 0 {
			if nx.Y < 4 {
				nx.X = 15
				nx.Y = 11 - nx.Y
				diridx = 2
			} else if nx.Y < 8 {
				nx.X = 15 - (nx.Y - 4)
				nx.Y = 8
				diridx = 1
			} else {
				nx.X = 11
				nx.Y = 3 - (nx.Y - 8)
				diridx = 2
			}
		} else if diridx == 1 {
			if nx.X < 4 {
				nx.Y = 11
				nx.X = 11 - nx.X
				diridx = 3
			} else if nx.X < 8 {
				nx.Y = 11 - (nx.X - 4)
				nx.X = 8
				diridx = 0
			} else if nx.X < 12 {
				nx.Y = 7
				nx.X = 3 - (nx.X - 8)
				diridx = 3
			} else {
				nx.Y = 7 - (nx.X - 12)
				nx.X = 0
				diridx = 0
			}
		} else if diridx == 2 {
			if nx.Y < 4 {
				nx.X = nx.Y + 4
				nx.Y = 4
				diridx = 1
			} else if nx.Y < 8 {
				nx.X = 15 - (nx.Y - 4)
				nx.Y = 11
				diridx = 3
			} else {
				nx.X = 7 - (nx.Y - 8)
				nx.Y = 7
				diridx = 3
			}
		} else {
			if nx.X < 4 {
				nx.X = 11 - nx.X
				nx.Y = 0
				diridx = 1
			} else if nx.X < 8 {
				nx.Y = 0 + (nx.X - 4)
				nx.X = 8
				diridx = 0
			} else if nx.X < 12 {
				nx.X = 3 - (nx.X - 8)
				nx.Y = 4
				diridx = 1
			} else {
				nx.Y = 7 - (nx.X - 12)
				nx.X = 11
				diridx = 2
			}
		}
	}

	for board.At(nx.X, nx.Y) == WALL {
		return start, lastDir
	}
	return nx, diridx
}
