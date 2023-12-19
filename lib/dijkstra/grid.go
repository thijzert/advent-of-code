package dijkstra

type bff struct {
	valid func(int, int, int, int, int) bool
	final func(int, int) bool
}

type gridBoard []gridPoint

func (b gridBoard) StartingPositions() []Position {
	rv := make([]Position, len(b))
	for i, p := range b {
		rv[i] = p
	}
	return rv
}

type gridPoint struct {
	X, Y     int
	Diagonal bool
	F        *bff
}

// GridWalker takes a list of integers, pairs them up into (x,y) coordinates,
// and returns a list of Positions on a rectangular grid. From each grid
// Position agents can take steps of length 1 in all 4 cardinal directions,
// checking if each step to take is a valid one using the valid() func provided
// in the first parameter. In the same vein, the final() func in the second
// parameter returns true if the x,y coordinate is a valid end state.
func GridWalker(valid func(x0, y0, x1, y1, cost int) bool, final func(x, y int) bool, xy ...int) Board {
	f := bff{
		valid: valid,
		final: final,
	}
	return getGridWalker(&f, false, xy...)
}

// DiagonalWalker is the same as GridWalker, only diagonal steps are also
// allowed. Diagonal steps also have length 1.
func DiagonalWalker(valid func(x0, y0, x1, y1, cost int) bool, final func(x, y int) bool, xy ...int) Board {
	f := bff{
		valid: valid,
		final: final,
	}
	return getGridWalker(&f, true, xy...)
}

func getGridWalker(f *bff, diagonal bool, xy ...int) Board {
	var rv gridBoard

	for i := 0; (i + 1) < len(xy); i += 2 {
		rv = append(rv, gridPoint{
			X:        xy[i],
			Y:        xy[i+1],
			Diagonal: diagonal,
			F:        f,
		})
	}

	return rv
}

func (p gridPoint) Final(b Board) bool {
	return p.F.final(p.X, p.Y)
}
func (p gridPoint) Adjacent(b Board, totalCost int) AdjacencyIterator {
	var rv []Adj

	if p.Diagonal {
		positions := [8]gridPoint{
			gridPoint{p.X + 1, p.Y, p.Diagonal, p.F},
			gridPoint{p.X, p.Y + 1, p.Diagonal, p.F},
			gridPoint{p.X - 1, p.Y, p.Diagonal, p.F},
			gridPoint{p.X, p.Y - 1, p.Diagonal, p.F},
			gridPoint{p.X + 1, p.Y - 1, p.Diagonal, p.F},
			gridPoint{p.X + 1, p.Y + 1, p.Diagonal, p.F},
			gridPoint{p.X - 1, p.Y - 1, p.Diagonal, p.F},
			gridPoint{p.X - 1, p.Y + 1, p.Diagonal, p.F},
		}
		for _, pos := range positions {
			if p.F.valid(p.X, p.Y, pos.X, pos.Y, totalCost+1) {
				rv = append(rv, Adj{pos, 1})
			}
		}
	} else {
		positions := [4]gridPoint{
			gridPoint{p.X + 1, p.Y, p.Diagonal, p.F},
			gridPoint{p.X, p.Y + 1, p.Diagonal, p.F},
			gridPoint{p.X - 1, p.Y, p.Diagonal, p.F},
			gridPoint{p.X, p.Y - 1, p.Diagonal, p.F},
		}
		for _, pos := range positions {
			if p.F.valid(p.X, p.Y, pos.X, pos.Y, totalCost+1) {
				rv = append(rv, Adj{pos, 1})
			}
		}
	}

	return AdjacencyList(rv)
}

func (p gridPoint) Pack() int {
	return 0
}
