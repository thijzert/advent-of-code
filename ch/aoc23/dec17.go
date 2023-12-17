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
	fm := factoryMap{
		Map:   img,
		Start: crucible{cube.Point{0, 0}, cube.Point{-1, -1}, 1, 3},
	}
	path, cost, err := dijkstra.ShortestPath(fm)
	drawPath(ctx, fm, path)
	return cost, err
}

func Dec17b(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2023/dec17.txt")
	if err != nil {
		return nil, err
	}
	img := image.ReadImage(lines, func(r rune) int {
		return int(r - '0')
	})
	fm := factoryMap{
		Map:   img,
		Start: crucible{cube.Point{0, 0}, cube.Point{-4, -4}, 4, 10},
	}
	path, cost, err := dijkstra.ShortestPath(fm)
	drawPath(ctx, fm, path)
	return cost, err
}

type factoryMap struct {
	Map   *image.Image
	Start dijkstra.Position
}

func (fm factoryMap) StartingPositions() []dijkstra.Position {
	return []dijkstra.Position{fm.Start}
}

type crucible struct {
	Position cube.Point
	LastTurn cube.Point

	MinStraight, MaxStraight int
}

func (c crucible) Final(b dijkstra.Board) bool {
	fm, ok := b.(factoryMap)
	if !ok {
		return false
	}
	straight := c.Position.Sub(c.LastTurn)
	l := iabs(straight.X) + iabs(straight.Y)
	return c.Position.X == fm.Map.Width-1 && c.Position.Y == fm.Map.Height-1 && l >= c.MinStraight
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

	// HACK: The position of the virtual "last turn" is outside the board.
	if c.LastTurn.X < 0 {
		l--
	}

	rv := []dijkstra.Adj{}
	for _, dir := range cube.Cardinal2D {
		if dir == back || (dir == straight && l >= c.MaxStraight) || (dir != straight && l < c.MinStraight) {
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

func (c crucible) Hashcode() [4]uint64 {
	return [4]uint64{uint64(c.Position.X), uint64(c.Position.Y), uint64(c.LastTurn.X), uint64(c.LastTurn.Y)}
}

func iabs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func drawPath(ctx ch.AOContext, fm factoryMap, path []dijkstra.Position) {
	img := image.NewImage(fm.Map.Width, fm.Map.Height, func(x, y int) int {
		return 0
	})
	for _, pos := range path[1:] {
		if c, ok := pos.(crucible); ok {
			img.Set(c.Position.X, c.Position.Y, 2)
		}
	}

	ctx.Printf("Shortest path:\n%s", img)
}
