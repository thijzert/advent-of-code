package aoc23

import (
	"github.com/thijzert/advent-of-code/ch"
	"github.com/thijzert/advent-of-code/lib/cube"
	"github.com/thijzert/advent-of-code/lib/dijkstra"
	"github.com/thijzert/advent-of-code/lib/image"
)

func Dec17a(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2023/dec17.txt")
	if err != nil {
		return nil, err
	}
	img := image.ReadImage(lines, func(r rune) int {
		return int(r - '0')
	})
	fm := factoryMap{img}
	_, cost, err := dijkstra.ShortestPath(fm)
	return cost, err
}

var Dec17b ch.AdventFunc = nil

// func Dec17b(ctx ch.AOContext) (interface{}, error) {
// 	return nil, errNotImplemented
// }

type factoryMap struct {
	Map *image.Image
}

func (fm factoryMap) StartingPositions() []dijkstra.Position {
	return []dijkstra.Position{
		crucible{
			Position: cube.Point{X: 0, Y: 0},
			LastTurn: cube.Point{X: -1, Y: 0},
		},
	}
}

type crucible struct {
	Position cube.Point
	LastTurn cube.Point
}

func (c crucible) Final(b dijkstra.Board) bool {
	fm, ok := b.(factoryMap)
	if !ok {
		return false
	}
	return c.Position.X == fm.Map.Width-1 && c.Position.Y == fm.Map.Height-1
}

func (c crucible) Adjacent(b dijkstra.Board, totalCost int) dijkstra.AdjacencyIterator {
	fm, ok := b.(factoryMap)
	if !ok {
		return dijkstra.DeadEnd()
	}

	straight := c.Position.Sub(c.LastTurn)
	l := iabs(straight.X) + iabs(straight.Y)
	straight.X, straight.Y = straight.X/l, straight.Y/l
	back := straight.Mul(-1)

	rv := []dijkstra.Adj{}
	for _, dir := range cube.Cardinal2D {
		if dir == back || (dir == straight && l >= 3) {
			continue
		}
		np := c
		np.Position = c.Position.Add(dir)
		if dir != straight {
			np.LastTurn = c.Position
		}
		if fm.Map.Inside(np.Position.X, np.Position.Y) {
			cost := fm.Map.At(np.Position.X, np.Position.Y)
			rv = append(rv, dijkstra.Adj{Position: np, Cost: cost})
		}
	}
	return dijkstra.AdjacencyList(rv)
}

func iabs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
