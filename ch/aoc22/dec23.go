package aoc22

import (
	"github.com/thijzert/advent-of-code/ch"
	"github.com/thijzert/advent-of-code/lib/cube"
	"github.com/thijzert/advent-of-code/lib/image"
)

func Dec23a(ctx ch.AOContext) (interface{}, error) {
	elves, err := readElfBotanists(ctx, "inputs/2022/dec23.txt")
	if err != nil {
		return nil, err
	}

	ctx.Printf("%d elves in %v:\n%s", len(elves), boundingRect(elves), drawElves(elves))

	for i := 0; i < 10; i++ {
		moveBotanistElves(ctx, elves, i)
		//ctx.Printf("%d elves in %v:\n%s", len(elves), boundingRect(elves), drawElves(elves))
	}

	br := boundingRect(elves)
	rv := (br.X.B-br.X.A+1)*(br.Y.B-br.Y.A+1) - len(elves)
	return rv, nil
}

func Dec23b(ctx ch.AOContext) (interface{}, error) {
	elves, err := readElfBotanists(ctx, "inputs/2022/dec23.txt")
	if err != nil {
		return nil, err
	}

	for i := 0; i < 10000; i++ {
		m := moveBotanistElves(ctx, elves, i)
		if m == 0 {
			return i + 1, nil
		}
	}

	return nil, errFailed
}

func readElfBotanists(ctx ch.AOContext, name string) (map[cube.Point]int, error) {
	lines, err := ctx.DataLines(name)
	if err != nil {
		return nil, err
	}

	elves := make(map[cube.Point]int)
	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				elves[cube.Point{x, y}] = 1
			}
		}
	}
	return elves, nil
}

func boundingRect(elves map[cube.Point]int) cube.Square {
	var rv cube.Square
	first := true
	for pt := range elves {
		if first {
			rv.X.A, rv.X.B = pt.X, pt.X
			rv.Y.A, rv.Y.B = pt.Y, pt.Y
			first = false
		}
		if pt.X < rv.X.A {
			rv.X.A = pt.X
		}
		if pt.X > rv.X.B {
			rv.X.B = pt.X
		}
		if pt.Y < rv.Y.A {
			rv.Y.A = pt.Y
		}
		if pt.Y > rv.Y.B {
			rv.Y.B = pt.Y
		}
	}
	return rv
}

func drawElves(elves map[cube.Point]int) string {
	r := boundingRect(elves)
	img := image.NewImage(r.X.B-r.X.A+1, r.Y.B-r.Y.A+1, func(x, y int) int {
		return elves[cube.Point{x + r.X.A, y + r.Y.A}]
	})
	return img.String()
}

type botanistDirection struct {
	Dir   cube.Point
	Clear [3]cube.Point
}

var botanistDirections [4]botanistDirection = [4]botanistDirection{
	{cube.Point{0, -1}, [3]cube.Point{{-1, -1}, {0, -1}, {1, -1}}},
	{cube.Point{0, 1}, [3]cube.Point{{-1, 1}, {0, 1}, {1, 1}}},
	{cube.Point{-1, 0}, [3]cube.Point{{-1, -1}, {-1, 0}, {-1, 1}}},
	{cube.Point{1, 0}, [3]cube.Point{{1, -1}, {1, 0}, {1, 1}}},
}

func moveBotanistElves(ctx ch.AOContext, elves map[cube.Point]int, round int) int {
	maybeMove := make(map[cube.Point]cube.Point)
	propose := make(map[cube.Point]int)
	for elf := range elves {
		neighbours := 0
		for _, dir := range cube.Cardinal2Diag {
			neighbours += elves[elf.Add(dir)]
		}
		if neighbours == 0 {
			continue
		}
		for i := range botanistDirections {
			bot := botanistDirections[(round+i)%4]
			occupied := 0
			for _, ch := range bot.Clear {
				occupied += elves[elf.Add(ch)]
			}
			if occupied == 0 {
				pr := elf.Add(bot.Dir)
				maybeMove[elf] = pr
				propose[pr] += 1
				break
			}
		}
	}

	rv := 0
	for elf, dest := range maybeMove {
		if propose[dest] == 1 {
			delete(elves, elf)
			elves[dest] = 1
			rv++
		}
	}

	return rv
}
