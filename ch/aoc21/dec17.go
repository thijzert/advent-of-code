package aoc21

import (
	"github.com/thijzert/advent-of-code/ch"
)

func Dec17a(ctx ch.AOContext) error {
	exampleTarget := rect{point{20, -10}, point{30, -5}}

	for _, vel := range []point{{7, 2}, {6, 3}, {17, -4}, {6, 9}} {
		hits, height := fireProbe(vel, exampleTarget)
		ctx.Printf("  - %d: hits: %5v; height %d", vel, hits, height)
	}

	_, my := fireProbeWithStyle(exampleTarget)
	ctx.Printf("Example data: max height %d", my)

	actualTarget := rect{point{138, -125}, point{184, -71}}
	_, my = fireProbeWithStyle(actualTarget)
	ctx.FinalAnswer.Print(my)

	return nil
}

func Dec17b(ctx ch.AOContext) error {
	exampleTarget := rect{point{20, -10}, point{30, -5}}

	nh, _ := fireProbeWithStyle(exampleTarget)
	ctx.Printf("Example data: trajectories that hit: %d", nh)

	actualTarget := rect{point{138, -125}, point{184, -71}}
	nh, _ = fireProbeWithStyle(actualTarget)
	ctx.FinalAnswer.Print(nh)

	return nil
}

func fireProbe(velocity point, target rect) (bool, int) {
	ymax := 0
	pos := point{0, 0}
	for (velocity.Y >= 0 || pos.Y >= target.Min.Y) && pos.X <= target.Max.X {
		pos = pos.Add(velocity)

		ymax = max(ymax, pos.Y)
		if target.Contains(pos) {
			return true, ymax
		}

		velocity.X = max(velocity.X-1, 0)
		velocity.Y -= 1
	}

	return false, 0
}

func fireProbeWithStyle(target rect) (int, int) {
	nHits := 0
	maxHeight := 0
	for vy := target.Min.Y; vy <= target.Max.X; vy++ {
		for vx := 0; vx <= target.Max.X; vx++ {
			hits, height := fireProbe(point{vx, vy}, target)
			if hits {
				nHits++
				maxHeight = max(maxHeight, height)
			}
		}
	}
	return nHits, maxHeight
}
