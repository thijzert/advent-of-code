package aoc23

import (
	"fmt"
	"sort"

	"github.com/thijzert/advent-of-code/ch"
	"github.com/thijzert/advent-of-code/lib/cube"
	"github.com/thijzert/advent-of-code/lib/pq"
)

type digInstruction struct {
	Direction cube.Point
	Distance  int
}

func Dec18a(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2023/dec18.txt")
	if err != nil {
		return nil, err
	}

	instrs := []digInstruction{}
	for _, l := range lines {
		inst := digInstruction{}
		dir := ""
		fmt.Sscanf(l, "%s %d ", &dir, &inst.Distance)
		inst.Direction, _ = cube.ParseDirection2D(dir)
		instrs = append(instrs, inst)
	}
	return dec18(ctx, instrs)
}

func Dec18b(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2023/dec18.txt")
	if err != nil {
		return nil, err
	}

	instrs := []digInstruction{}
	for _, l := range lines {
		oldDir, oldDist, colour := "", 0, 0
		fmt.Sscanf(l, "%s %d (#%06x)", &oldDir, &oldDist, &colour)
		instrs = append(instrs, digInstruction{
			Distance:  colour >> 4,
			Direction: cube.Cardinal2D[colour&0xf],
		})
	}
	return dec18(ctx, instrs)
}

func dec18(ctx ch.AOContext, instrs []digInstruction) (any, error) {
	bounds := cube.Square{}
	corners := pq.PriorityQueue[int]{}
	dig := cube.Point{}
	for _, inst := range instrs {
		dig = dig.Add(inst.Direction.Mul(inst.Distance))
		bounds = bounds.UpdatedBound(dig)
		corners.Push(dig.X, dig.Y)
	}
	ctx.Printf("Bounds: %v", bounds)

	answer := 0
	ylast := bounds.Y.A
	ranges := cube.IntervalSet{}
	for corners.Len() > 0 {
		x0, y0, _ := corners.Pop()
		answer += ranges.Length() * (y0 - ylast - 1)
		xs := []int{x0}
		for corners.Len() > 0 {
			x, y, _ := corners.Pop()
			if y != y0 {
				corners.Push(x, y)
				break
			}
			xs = append(xs, x)
		}
		sort.Ints(xs)
		ctx.Printf("Y %d corner points: %d", y0, xs)
		if len(xs)%2 != 0 {
			return nil, errFailed
		}

		// Add current line, with existing and new horizontal segments
		currentLine := cube.IntervalSet{}
		currentLine.I = append(currentLine.I, ranges.I...)
		for i := 0; i < len(xs); i += 2 {
			currentLine.Add(cube.Interval{xs[i], xs[i+1]})
		}
		answer += currentLine.Length()

		// Update vertcial slice
		for i := 0; i < len(xs); i += 2 {
			found := false
			for j, intv := range ranges.I {
				if !intv.Contains(xs[i]) && !intv.Contains(xs[i+1]) {
					continue
				}
				found = true
				oth := cube.Interval{0, -1}
				if xs[i] == intv.A && xs[i+1] == intv.B {
					// Just delete this entire interval
					intv.A, intv.B = 1, 0
				} else if xs[i] == intv.A {
					intv.A = xs[i+1]
				} else if xs[i+1] == intv.A {
					intv.A = xs[i]
				} else if xs[i] == intv.B {
					intv.B = xs[i+1]
				} else if xs[i+1] == intv.B {
					intv.B = xs[i]
				} else {
					oth.A, oth.B = xs[i+1], intv.B
					intv.B = xs[i]
				}
				ranges.I[j] = intv
				if oth.Length() > 0 {
					ranges.Add(oth)
				}
				break
			}
			if !found {
				ranges.Add(cube.Interval{xs[i], xs[i+1]})
			}
		}
		newRanges := cube.IntervalSet{}
		for _, intv := range ranges.I {
			newRanges.Add(intv)
		}
		ranges = newRanges

		ctx.Printf("vertical: (%d) %v", ranges.Length(), ranges)

		ylast = y0
	}

	return answer, nil
}
