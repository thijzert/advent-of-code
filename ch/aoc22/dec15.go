package aoc22

import (
	"fmt"

	"github.com/thijzert/advent-of-code/ch"
	"github.com/thijzert/advent-of-code/lib/cube"
)

func Dec15a(ctx ch.AOContext) error {
	sensors, err := readCaveSensors(ctx)
	if err != nil {
		return err
	}

	rv := beaconFreeSpots(ctx, sensors, 2000000)
	ctx.FinalAnswer.Print(rv)
	return nil
}

var Dec15b ch.AdventFunc = nil

// func Dec15b(ctx ch.AOContext) error {
// 	return errNotImplemented
// }

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

func beaconFreeSpots(ctx ch.AOContext, sensors []caveSensor, y int) int {
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
		// ctx.Printf("add %v", cube.Interval{s.Position.X - dx, s.Position.X + dx})
		ivs.Add(cube.Interval{s.Position.X - dx, s.Position.X + dx})
	}

	ctx.Printf("intervalset %s", ivs)

	rv := 0
	for _, iv := range ivs.I {
		rv += iv.B - iv.A + 1
	}

	beacons := make(map[int]bool)
	for _, s := range sensors {
		if s.Beacon.Y == y {
			beacons[s.Beacon.X] = true
		}
	}
	return rv - len(beacons)
}
