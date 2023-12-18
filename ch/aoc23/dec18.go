package aoc23

import (
	"fmt"

	"github.com/thijzert/advent-of-code/ch"
	"github.com/thijzert/advent-of-code/lib/cube"
	"github.com/thijzert/advent-of-code/lib/image"
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
	lines, err := ctx.DataLines("inputs/2023/dec18a.txt")
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
	dig := cube.Point{}
	for _, inst := range instrs {
		for i := 0; i < inst.Distance; i++ {
			dig = dig.Add(inst.Direction)
			bounds = bounds.UpdatedBound(dig)
		}
	}
	ctx.Printf("Bounds: %v", bounds)

	img := image.NewImage(bounds.X.B-bounds.X.A+2, bounds.Y.B-bounds.Y.A+2, func(x, y int) int {
		return 0
	})
	dig = cube.Point{X: 1 - bounds.X.A, Y: 1 - bounds.Y.A}
	for _, inst := range instrs {
		for i := 0; i < inst.Distance; i++ {
			dig = dig.Add(inst.Direction)
			img.Set(dig.X, dig.Y, 1)
		}
	}

	// TODO: this, but better
	img.FloodFill(2-bounds.X.A, -bounds.Y.A, 5)

	answer := 0
	for y := 0; y < img.Height; y++ {
		for x := 0; x < img.Width; x++ {
			if img.At(x, y) > 0 {
				answer++
			}
		}
	}

	ctx.Printf("Trench: \n%s", img)
	return answer, nil
}
