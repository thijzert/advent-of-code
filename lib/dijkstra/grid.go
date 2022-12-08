package dijkstra

type bff struct {
	valid func(int, int) bool
	final func(int, int, int) bool
}

type gridPoint struct {
	X, Y     int
	Diagonal bool
	F        *bff
}

// GridWalker takes a list of integers, pairs them up into (x,y) coordinates,
// and returns a list of Positions on a rectangular grid. From each grid
// Position agents can take steps of length 1 in all 4 cardinal directions,
// checking if each resulting position is valid using the valid() func provided
// in the first parameter. In the same vein, the final() func in the second
// parameter returns true if the x,y coordinate is a valid end state.
func GridWalker(valid func(x, y int) bool, final func(x, y, cost int) bool, xy ...int) []Position {
	return getGridWalker(valid, final, false, xy...)
}

// DiagonalWalker is the same as GridWalker, only diagonal steps are also
// allowed. Diagonal steps also have length 1.
func DiagonalWalker(valid func(x, y int) bool, final func(x, y, cost int) bool, xy ...int) []Position {
	return getGridWalker(valid, final, true, xy...)
}

func getGridWalker(valid func(x, y int) bool, final func(x, y, cost int) bool, diagonal bool, xy ...int) []Position {
	f := bff{
		valid: valid,
		final: final,
	}

	var rv []Position

	for i := 0; (i + 1) < len(xy); i += 1 {
		rv = append(rv, gridPoint{
			X:        xy[i],
			Y:        xy[i+1],
			Diagonal: diagonal,
			F:        &f,
		})
	}

	return rv
}

func (p gridPoint) Final(b Board, totalCost int) bool {
	return p.F.final(p.X, p.Y, totalCost)
}
func (p gridPoint) Adjacent(b Board) AdjacencyIterator {
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
			if p.F.valid(pos.X, pos.Y) {
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
			if p.F.valid(pos.X, pos.Y) {
				rv = append(rv, Adj{pos, 1})
			}
		}
	}

	return AdjacencyList(rv)
}
