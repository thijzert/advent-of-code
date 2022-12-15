package aoc22

import (
	"fmt"

	"github.com/thijzert/advent-of-code/ch"
	"github.com/thijzert/advent-of-code/lib/cube"
)

func Dec15a(ctx ch.AOContext) (interface{}, error) {
	sensors, err := readCaveSensors(ctx)
	if err != nil {
		return nil, err
	}

	yy := 2000000
	rv := beaconCoverage(ctx, sensors, yy)
	ctx.Printf("intervalset %s", rv)

	beacons := make(map[int]bool)
	for _, s := range sensors {
		if s.Beacon.Y == yy {
			beacons[s.Beacon.X] = true
		}
	}

	return rv.Length() - len(beacons), nil
}

func Dec15b(ctx ch.AOContext) (interface{}, error) {
	sensors, err := readCaveSensors(ctx)
	if err != nil {
		return nil, err
	}

	fullSpectrum := cube.Interval{0, 4000000}
	for y := 0; y < 4000000; y++ {
		rv := beaconCoverage(ctx, sensors, y)
		if rv.FullyContains(fullSpectrum) {
			continue
		}
		ctx.Printf("intervalset at y=%d: %s", y, rv)
		for _, iv := range rv.I {
			if iv.B >= -1 {
				x := iv.B + 1
				ctx.Printf("Found mystery beacon at x=%d y=%d", x, y)
				return 4000000*x + y, nil
			}
		}
	}

	return nil, errFailed
}

type caveSensor struct {
	Position cube.Point
	Beacon   cube.Point
	Dist     int
}

func readCaveSensors(ctx ch.AOContext) ([]caveSensor, error) {
	lines, err := ctx.DataLines("inputs/2022/dec15.txt")
	if err != nil {
		return nil, err
	}

	var rv []caveSensor
	for _, line := range lines {
		var s caveSensor
		_, err := fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &s.Position.X, &s.Position.Y, &s.Beacon.X, &s.Beacon.Y)
		if err != nil {
			return nil, err
		}
		s.Dist = s.Beacon.Sub(s.Position).Manhattan()
		rv = append(rv, s)
	}
	return rv, nil
}

func beaconCoverage(ctx ch.AOContext, sensors []caveSensor, y int) *cube.IntervalSet {
	ivs := cube.NewIntervalSet()
	for _, s := range sensors {
		dy := y - s.Position.Y
		if dy < 0 {
			dy = -dy
		}
		dx := s.Dist - dy
		if dx < 0 {
			continue
		}
		iv := cube.Interval{s.Position.X - dx, s.Position.X + dx}
		ivs.Add(iv)
	}

	return ivs
}
