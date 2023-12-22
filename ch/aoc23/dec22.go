package aoc23

import (
	"fmt"

	"github.com/thijzert/advent-of-code/ch"
	"github.com/thijzert/advent-of-code/lib/cube"
	"github.com/thijzert/advent-of-code/lib/image"
	"github.com/thijzert/advent-of-code/lib/pq"
)

func Dec22a(ctx ch.AOContext) (interface{}, error) {
	supports, supported, err := dec22(ctx)
	if err != nil {
		return nil, err
	}

	answer := 0
	for k, v := range supports {
		if len(v) == 0 {
			answer++
			if len(supports) < 26 {
				ctx.Printf("Brick '%c' supports no bricks", 'A'+rune(k))
			}
		} else {
			if len(supports) < 26 {
				ctx.Printf("Brick '%c' supports bricks %v", 'A'+rune(k), v)
			}
			allOk := true
			for _, br := range v {
				otherSupport := false
				for _, ob := range supported[br] {
					if ob != k {
						otherSupport = true
						if len(supports) < 26 {
							ctx.Printf("   but brick '%c' is also supported by '%c'", 'A'+rune(br), 'A'+rune(ob))
						}
					}
				}
				allOk = allOk && otherSupport
			}
			if allOk {
				answer++
			}
		}
	}

	return answer, nil
}

func Dec22b(ctx ch.AOContext) (interface{}, error) {
	_, supported, err := dec22(ctx)
	if err != nil {
		return nil, err
	}

	answer := 0
	for brid := range supported {
		gone := make(map[int]bool)
		gone[brid] = true
		changed := true
		for changed {
			changed = false
			for ob, sup := range supported {
				if gone[ob] {
					continue
				}

				isSupported := false
				for _, s := range sup {
					isSupported = isSupported || !gone[s]
				}
				if !isSupported {
					gone[ob] = true
					changed = true
				}
			}
		}

		// ctx.Printf("Disintegrating brick %d will cause %d bricks to fall", brid, len(gone)-1)
		answer += len(gone) - 1
	}

	// 1258: too low
	// 10716: too low
	return answer, nil
}

func dec22(ctx ch.AOContext) ([][]int, [][]int, error) {
	lines, err := ctx.DataLines("inputs/2023/dec22.txt")
	if err != nil {
		return nil, nil, err
	}

	var sortedBricks pq.PriorityQueue[int]
	bounds := cube.Square{}
	for i, line := range lines {
		var x0, y0, z0, x1, y1, z1 int
		fmt.Sscanf(line, "%d,%d,%d~%d,%d,%d", &x0, &y0, &z0, &x1, &y1, &z1)
		bounds = bounds.UpdatedBound(cube.Point{x0, y0})
		bounds = bounds.UpdatedBound(cube.Point{x1, y1})
		sortedBricks.Push(i, z0)
	}

	heightmap := image.NewImage(1+bounds.X.B-bounds.X.A, 1+bounds.Y.B-bounds.Y.A, func(x, y int) int {
		return 0
	})
	topBrick := image.NewImage(1+bounds.X.B-bounds.X.A, 1+bounds.Y.B-bounds.Y.A, func(x, y int) int {
		return -1
	})
	supports, supported := make([][]int, len(lines)), make([][]int, len(lines))

	for sortedBricks.Len() > 0 {
		i, _, _ := sortedBricks.Pop()
		line := lines[i]
		supports[i] = nil
		var bricksBelow []int
		var x0, y0, z0, x1, y1, z1 int
		fmt.Sscanf(line, "%d,%d,%d~%d,%d,%d", &x0, &y0, &z0, &x1, &y1, &z1)
		x0, y0 = x0-bounds.X.A, y0-bounds.Y.A
		x1, y1 = x1-bounds.X.A, y1-bounds.Y.A

		zfinal := 0
		for y := y0; y <= y1; y++ {
			for x := x0; x <= x1; x++ {
				zfinal = max(zfinal, heightmap.At(x, y))
			}
		}
		for y := y0; y <= y1; y++ {
			for x := x0; x <= x1; x++ {
				if heightmap.At(x, y) == zfinal {
					tb := topBrick.At(x, y)
					if tb >= 0 {
						sup := supports[tb]
						if len(sup) == 0 || sup[len(sup)-1] != i {
							sup = append(sup, i)
							supports[tb] = sup
						}
					}
					bricksBelow = append(bricksBelow, tb)
				}
				topBrick.Set(x, y, i)
				heightmap.Set(x, y, zfinal+1+max(z0, z1)-min(z0, z1))
			}
		}
		supported[i] = bricksBelow
	}

	return supports, supported, nil
}
