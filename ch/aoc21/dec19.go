package aoc21

import (
	"fmt"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec19a(ctx ch.AOContext) (interface{}, error) {
	scanners, err := readScannerBeacons(ctx, "inputs/2021/dec19.txt")
	if err != nil {
		return nil, err
	}
	scannerPos, err := findScannerPositionsFromCommonBeacons(ctx, scanners)
	if err != nil {
		return nil, err
	}

	allBeacons := make(map[point3]bool)

	for i, points := range scanners {
		sp, ok := scannerPos[i]
		if !ok {
			return nil, fmt.Errorf("failed to find a position for scanner %d", i)
		}

		for _, pt := range points {
			allBeacons[sp.Abs(pt)] = true
		}
	}

	return len(allBeacons), nil
}

func Dec19b(ctx ch.AOContext) (interface{}, error) {
	scanners, err := readScannerBeacons(ctx, "inputs/2021/dec19.txt")
	if err != nil {
		return nil, err
	}
	scannerPos, err := findScannerPositionsFromCommonBeacons(ctx, scanners)
	if err != nil {
		return nil, err
	}

	maxD := 0

	for _, spA := range scannerPos {
		for _, spB := range scannerPos {
			mhd := abs(spA.Position.X-spB.Position.X) + abs(spA.Position.Y-spB.Position.Y) + abs(spA.Position.Z-spB.Position.Z)
			if mhd > maxD {
				maxD = mhd
			}
		}
	}

	return maxD, nil
}

type scannerPosition struct {
	Position    point3
	Orientation orientation
}

func (sp scannerPosition) Abs(relative point3) point3 {
	relative = relative.Tr(sp.Orientation)
	relative = relative.Add(sp.Position)
	return relative
}

func readScannerBeacons(ctx ch.AOContext, assetName string) ([][]point3, error) {
	sections, err := ctx.DataSections(assetName)
	if err != nil {
		return nil, err
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
				return nil, err
			}
		}
	}

	return scanners, nil
}

func findScannerPositionsFromCommonBeacons(ctx ch.AOContext, scanners [][]point3) (map[int]scannerPosition, error) {
	scannerPos := make(map[int]scannerPosition)
	scannerPos[0] = scannerPosition{}

	changed := true
	for changed {
		changed = false
		for i, points := range scanners {
			if _, ok := scannerPos[i]; ok {
				continue
			}
			weGotEm := false

			for j, sp := range scannerPos {
				if weGotEm {
					break
				}
				if j == i {
					continue
				}
				jPoints := scanners[j]

				for _, ptA := range points[:len(points)-12] {
					if weGotEm {
						break
					}
					for o := orientation(0); o < 24; o++ {
						if weGotEm {
							break
						}
						aabs := ptA.Tr(o)
						for _, ptB := range jPoints[:len(jPoints)-12] {
							if weGotEm {
								break
							}
							// Count overlap between scanner i and j,
							// assuming point A in orientation o is point B
							babs := sp.Abs(ptB)
							iPos := scannerPosition{
								Position:    babs.Sub(aabs),
								Orientation: o,
							}

							match := 0
							for k, c := range points {
								if match+len(points)-k < 12 {
									break
								}
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
								weGotEm = true
								break
							}
						}
					}
				}
			}
		}
	}

	for i := range scanners {
		if _, ok := scannerPos[i]; !ok {
			return nil, fmt.Errorf("failed to find a position for scanner %d", i)
		}
	}

	return scannerPos, nil
}
