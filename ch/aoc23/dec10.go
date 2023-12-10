package aoc23

import (
	"github.com/thijzert/advent-of-code/ch"
	"github.com/thijzert/advent-of-code/lib/image"
)

type pipeConnection struct {
	Direction [2]int
	Pipe      byte
}

func Dec10(ctx ch.AOContext) (int, int, error) {
	pipemap := make(map[pipeConnection][2]int)
	n := -1
	pipemap[pipeConnection{[2]int{0, 1}, '|'}] = [2]int{0, 1}
	pipemap[pipeConnection{[2]int{0, n}, '|'}] = [2]int{0, n}
	pipemap[pipeConnection{[2]int{1, 0}, '-'}] = [2]int{1, 0}
	pipemap[pipeConnection{[2]int{n, 0}, '-'}] = [2]int{n, 0}

	pipemap[pipeConnection{[2]int{0, 1}, 'L'}] = [2]int{1, 0}
	pipemap[pipeConnection{[2]int{n, 0}, 'L'}] = [2]int{0, n}
	pipemap[pipeConnection{[2]int{0, 1}, 'J'}] = [2]int{n, 0}
	pipemap[pipeConnection{[2]int{1, 0}, 'J'}] = [2]int{0, n}
	pipemap[pipeConnection{[2]int{0, n}, '7'}] = [2]int{n, 0}
	pipemap[pipeConnection{[2]int{1, 0}, '7'}] = [2]int{0, 1}
	pipemap[pipeConnection{[2]int{0, n}, 'F'}] = [2]int{1, 0}
	pipemap[pipeConnection{[2]int{n, 0}, 'F'}] = [2]int{0, 1}
	pmc := func(dir [2]int, ch byte) ([2]int, bool) {
		rv, ok := pipemap[pipeConnection{dir, ch}]
		return rv, ok
	}

	var directions [4][2]int = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	pipeType := make(map[[4]int]rune)
	pipeType[[4]int{1, 0, 1, 0}] = '|'
	pipeType[[4]int{0, 1, 0, 1}] = '-'
	pipeType[[4]int{1, 1, 0, 0}] = 'F'
	pipeType[[4]int{0, 1, 1, 0}] = 'L'
	pipeType[[4]int{0, 0, 1, 1}] = 'J'
	pipeType[[4]int{1, 0, 0, 1}] = '7'

	pipes, err := ctx.DataLines("inputs/2023/dec10.txt")
	if err != nil {
		return 0, 0, err
	}
	//pipes = []string{"7-F7-", ".FJ|7", "SJLL7", "|F--J", "LJ.LJ"}
	//pipes = []string{
	//	"FF7FSF7F7F7F7F7F---7",
	//	"L|LJ||||||||||||F--J",
	//	"FL-7LJLJ||||||LJL-77",
	//	"F--JF--7||LJLJ7F7FJ-",
	//	"L---JF-JLJ.||-FJLJJ7",
	//	"|F|F-JF---7F7-L7L|7|",
	//	"|FFJF7L7F-JF7|JL---7",
	//	"7-L-JL7||F7|L7F-7F7|",
	//	"L.L7LFJ|||||FJL7||LJ",
	//	"L7JLJL-JLJLJL--JLJ.L",
	//}

	var position, dir [2]int

	// Find starting position
	for y, line := range pipes {
		for x, c := range line {
			if c == 'S' {
				position[0], position[1] = x, y
				ctx.Printf("Found starting position %d,%d", x+1, y+1)
			}
		}
	}

	// Find starting direction by finding a pipe connected to S
	var startConnected [4]int
	var startType rune
	for i, d := range directions {
		x, y := position[0]+d[0], position[1]+d[1]
		if x < 0 || x >= len(pipes[0]) || y < 0 || y >= len(pipes) {
			continue
		}
		if _, ok := pmc(d, pipes[y][x]); ok {
			dir = d
			startConnected[i] = 1
		}
	}
	startType = pipeType[startConnected]
	ctx.Printf("Starting point is a '%c' segment", startType)

	loop := [][2]int{}
	var ch byte
	for {
		position[0] += dir[0]
		position[1] += dir[1]
		loop = append(loop, position)
		ch = pipes[position[1]][position[0]]
		if ch == 'S' {
			break
		}
		ndir, ok := pmc(dir, ch)
		if !ok {
			ctx.Printf("Failed to connect from %d,%d via %c going %d,%d", position[0]-dir[0], position[1]-dir[1], ch, dir[0], dir[1])
			return 0, 0, errFailed
		}
		dir = ndir
	}
	ctx.Printf("Total loop length: %d", len(loop))

	img := image.NewImage(len(pipes[0]), len(pipes), func(i1, i2 int) int {
		return 0
	})
	for _, pos := range loop {
		img.Set(pos[0], pos[1], 5)
	}
	enclosedArea := 0
	for y, line := range pipes {
		inside := false
		lastch := rune(0)
		for x, ch := range line {
			if img.At(x, y) == 5 {
				if ch == 'S' {
					ch = startType
				}
				if ch == '-' {
				} else if ch == '|' {
					inside = !inside
				} else if (lastch == 'F' && ch == 'J') || (lastch == 'L' && ch == '7') {
					inside = !inside
					lastch = 0
				} else {
					lastch = ch
				}
			} else if inside {
				enclosedArea++
				img.Set(x, y, 1)
			}
		}
	}
	ctx.Printf("Enclosed area: %d", enclosedArea)
	ctx.Printf("Loop:\n%s", img)

	return len(loop), enclosedArea, nil
}

func Dec10a(ctx ch.AOContext) (interface{}, error) {
	length, _, err := Dec10(ctx)
	return length / 2, err
}

func Dec10b(ctx ch.AOContext) (interface{}, error) {
	_, enclosedArea, err := Dec10(ctx)
	return enclosedArea, err
}
