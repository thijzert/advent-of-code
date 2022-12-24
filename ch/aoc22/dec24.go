package aoc22

import (
	"github.com/thijzert/advent-of-code/ch"
	"github.com/thijzert/advent-of-code/lib/cube"
	"github.com/thijzert/advent-of-code/lib/dijkstra"
	"github.com/thijzert/advent-of-code/lib/image"
)

func Dec24a(ctx ch.AOContext) (interface{}, error) {
	snowstorm, err := readSnowstormboard(ctx, "inputs/2022/dec24.txt")
	if err != nil {
		return nil, err
	}

	_, cost, err := dijkstra.ShortestPath(snowstorm)
	if err != nil {
		return nil, err
	}

	return cost, nil
}

var Dec24b ch.AdventFunc = nil

// func Dec24b(ctx ch.AOContext) (interface{}, error) {
// 	return nil, errNotImplemented
// }

type snowstormboard [4]*image.Image

func (snowstormboard) StartingPositions() []dijkstra.Position {
	return []dijkstra.Position{expeditionPosition{0, -1, 0}}
}

func readSnowstormboard(ctx ch.AOContext, name string) (snowstormboard, error) {
	var rv snowstormboard
	lines, err := ctx.DataLines(name)
	if err != nil {
		return rv, err
	}
	// Remove the outer border; we'll do bounds checking another way
	crustsRemoved := make([]string, len(lines)-2)
	for i, s := range lines {
		if i == 0 || i == (len(lines)-1) {
			continue
		}
		crustsRemoved[i-1] = s[1 : len(s)-1]
	}

	filterRune := func(e rune) func(c rune) int {
		return func(c rune) int {
			if c == e {
				return 1
			}
			return 0
		}
	}

	rv[0] = image.ReadImage(crustsRemoved, filterRune('<'))
	rv[1] = image.ReadImage(crustsRemoved, filterRune('>'))
	rv[2] = image.ReadImage(crustsRemoved, filterRune('^'))
	rv[3] = image.ReadImage(crustsRemoved, filterRune('v'))

	ctx.Printf("Snowstorms going left:\n%s\n, going down:\n%s", rv[0], rv[3])
	return rv, nil
}

type expeditionPosition cube.Point3

func (p expeditionPosition) Final(b dijkstra.Board) bool {
	bb, ok := b.(snowstormboard)
	if !ok {
		return false
	}
	return p.X == bb[0].Width-1 && p.Y == bb[0].Height
}

func (p expeditionPosition) Adjacent(b dijkstra.Board, totalCost int) dijkstra.AdjacencyIterator {
	bb, ok := b.(snowstormboard)
	if !ok {
		return dijkstra.DeadEnd()
	}

	W, H := bb[0].Width, bb[0].Height

	rv := []dijkstra.Adj{}
	for _, add := range [5]cube.Point3{{0, 0, 1}, {1, 0, 1}, {0, 1, 1}, {-1, 0, 1}, {0, -1, 1}} {
		pt := expeditionPosition(cube.Point3(p).Add(add))
		if (pt.X == W-1 && pt.Y == H) || (pt.X == 0 && pt.Y == -1) {
			rv = append(rv, dijkstra.Adj{pt, 1})
			continue
		} else if pt.X < 0 || pt.X >= W || pt.Y < 0 || pt.Y >= H {
			continue
		}
		storms := 0
		t := totalCost + 1
		storms += bb[0].At((pt.X+t)%W, pt.Y)
		storms += bb[1].At(((pt.X-t)%W+W)%W, pt.Y)
		storms += bb[2].At(pt.X, (pt.Y+t)%H)
		storms += bb[3].At(pt.X, ((pt.Y-t)%H+H)%H)
		if storms == 0 {
			rv = append(rv, dijkstra.Adj{pt, 1})
		}
	}

	return dijkstra.AdjacencyList(rv)
}
