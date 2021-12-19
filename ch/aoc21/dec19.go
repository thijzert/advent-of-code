package aoc21

import (
	"fmt"

	"github.com/thijzert/advent-of-code/ch"
)

var Dec19b ch.AdventFunc = nil

func Dec19a(ctx ch.AOContext) error {
	sections, err := ctx.DataSections("inputs/2021/dec19.txt")
	if err != nil {
		return err
	}

	var scanners [][]point3
	for i, sect := range sections {
		scanners = append(scanners, make([]point3, len(sect)-1))
		for j, l := range sect {
			if j == 0 {
				continue
			}
			_, err := fmt.Sscanf(l, "%d,%d,%d", &scanners[i][j-1].X, &scanners[i][j-1].Y, &scanners[i][j-1].Z)
			if err != nil {
				return err
			}
		}
	}

	scannerPos := make(map[int]scannerPosition)
	scannerPos[0] = scannerPosition{}

	changed := true
	for changed {
		changed = false
		for i, points := range scanners {
			if _, ok := scannerPos[i]; ok {
				continue
			}

			for j, sp := range scannerPos {
				if j == i {
					continue
				}
				jPoints := scanners[j]

				for _, ptA := range points[:len(points)-12] {
					for o := orientation(0); o < 24; o++ {
						aabs := ptA.Tr(o)
						for _, ptB := range jPoints[:len(jPoints)-12] {
							// Count overlap between scanner i and j,
							// assuming point A in orientation o is point B
							babs := sp.Abs(ptB)
							iPos := scannerPosition{
								Position:    babs.Sub(aabs),
								Orientation: o,
							}

							match := 0
							for _, c := range points {
								for _, d := range jPoints {
									if sp.Abs(d) == iPos.Abs(c) {
										match++
									}
								}
							}
							if match >= 12 {
								// ctx.Printf("Scanners %d and %d overlap with %d points", i, j, match)
								ctx.Printf("Scanners %d position: %d; orientation %d", i, iPos.Position, iPos.Orientation)
								scannerPos[i] = iPos
								changed = true
								break
							}
						}
					}
				}
			}
		}
	}

	allBeacons := make(map[point3]bool)

	for i, points := range scanners {
		sp, ok := scannerPos[i]
		if !ok {
			return fmt.Errorf("failed to find a position for scanner %d", i)
		}

		for _, pt := range points {
			allBeacons[sp.Abs(pt)] = true
		}
	}

	ctx.FinalAnswer.Print(len(allBeacons))
	return nil
}

// func Dec19b(ctx ch.AOContext) error {
// 	return errNotImplemented
// }

type scannerPosition struct {
	Position    point3
	Orientation orientation
}

func (sp scannerPosition) Abs(relative point3) point3 {
	relative = relative.Tr(sp.Orientation)
	relative = relative.Add(sp.Position)
	return relative
}
