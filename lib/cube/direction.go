package cube

// Cardinal2D contains al cardinal directions in 2D space (right, up, left, down)
var Cardinal2D [4]Point = [4]Point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

// A Walker represents a callback function to be passed to Walk(). If the
// Walker returns true, the Walk is aborted before reaching the boundaries
type Walker func(p Point) bool

// Walk moves in the given direction from the start point until you exit the
// boundaries, calling f at every step. If f returns true, the walk is aborted.
// Walk returns the number of steps taken, and a bool indicating if the walk
// was aborted prematurely or ran outside of the given boundaries.
func Walk(start, direction Point, bounds Square, f Walker) (int, bool) {
	p := start
	rv := 0
	for bounds.Contains(p) {
		if f(p) {
			return rv, true
		}
		p = p.Add(direction)
		rv++
	}
	rv--
	return rv, false
}
