package aoc22

import (
	"fmt"

	"github.com/thijzert/advent-of-code/ch"
	"github.com/thijzert/advent-of-code/lib/cube"
)

func Dec09a(ctx ch.AOContext) error {
	lines, err := ctx.DataLines("inputs/2022/dec09.txt")
	if err != nil {
		return err
	}
	//lines = []string{"R 4", "U 4", "L 3", "D 1", "R 4", "D 1", "L 5", "R 2"}

	head, tail := cube.Point{0, 0}, cube.Point{0, 0}
	visited := make(map[cube.Point]bool)
	for _, line := range lines {
		var a string
		var l int
		_, err = fmt.Sscanf(line, "%s %d", &a, &l)
		if err != nil {
			return err
		}
		dir, ok := cube.ParseDirection2D(a)
		if !ok {
			return fmt.Errorf("unknown direction '%c'", a)
		}

		for i := 0; i < l; i++ {
			head = head.Add(dir)
			tail = updatePlanckRope(head, tail)

			visited[tail] = true
		}
	}

	ctx.FinalAnswer.Print(len(visited))
	return nil
}

func updatePlanckRope(head, tail cube.Point) cube.Point {
	tailDiff := head.Sub(tail)
	move := cube.Point{0, 0}
	if tailDiff.X*tailDiff.X > 1 {
		move.X = 1
		if tailDiff.X < 0 {
			move.X = -1
		}

		// Move diagonally if necessary
		if tailDiff.Y < 0 {
			move.Y = -1
		} else if tailDiff.Y > 0 {
			move.Y = 1
		}
	} else if tailDiff.Y*tailDiff.Y > 1 {
		move.Y = 1
		if tailDiff.Y < 0 {
			move.Y = -1
		}

		// Move diagonally if necessary
		if tailDiff.X < 0 {
			move.X = -1
		} else if tailDiff.X > 0 {
			move.X = 1
		}
	}

	return tail.Add(move)
}

func Dec09b(ctx ch.AOContext) error {
	lines, err := ctx.DataLines("inputs/2022/dec09.txt")
	if err != nil {
		return err
	}
	//lines = []string{"R 5", "U 8", "L 8", "D 3", "R 17", "D 10", "L 25", "U 20"}
	//bounds := cube.Square{cube.Interval{-11, 14}, cube.Interval{-5, 15}}

	rope := make([]cube.Point, 10)
	visited := make(map[cube.Point]bool)
	for _, line := range lines {
		var a string
		var l int
		_, err = fmt.Sscanf(line, "%s %d", &a, &l)
		if err != nil {
			return err
		}
		dir, ok := cube.ParseDirection2D(a)
		if !ok {
			return fmt.Errorf("unknown direction '%c'", a)
		}

		for i := 0; i < l; i++ {
			rope[0] = rope[0].Add(dir)
			for j := range rope[1:] {
				rope[j+1] = updatePlanckRope(rope[j], rope[j+1])
			}
			visited[rope[len(rope)-1]] = true
		}

		// ctx.Print(printRopeBridge(rope, bounds))
	}

	ctx.FinalAnswer.Print(len(visited))
	return nil
}

func printRopeBridge(rope []cube.Point, bounds cube.Square) string {
	rv := ""
	for y := bounds.Y.B; y >= bounds.Y.A; y-- {
		rv += "\n"
		for x := bounds.X.A; x <= bounds.X.B; x++ {
			p := cube.Point{x, y}
			found := false
			for i, q := range rope {
				if q == p {
					found = true
					if i == 0 {
						rv += "H"
					} else {
						rv += string('0' + rune(i))
					}
					break
				}
			}
			if !found {
				if x == 0 && y == 0 {
					rv += "s"
				} else {
					rv += "."
				}
			}
		}
	}
	return rv
}
