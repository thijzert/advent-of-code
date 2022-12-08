package cube

// Cardinal2D contains all cardinal directions in 2D space (right, up, left, down)
var Cardinal2D [4]Point = [4]Point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

// Cardinal2Diag contains all unit directions in 2D space, including diagonals
var Cardinal2Diag [8]Point = [8]Point{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}

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

// A Flyer represent a callback function to be passed to Fly().
// Flyer is like a Walker, but in 3D space
type Flyer func(p Point3) bool

// Fly moves in the given direction from the start point until you exit the
// boundaries, calling f at every step. If f returns true, the flight is
// aborted. Fly returns the number of steps taken, and a bool indicating if
// the flight was aborted prematurely or ran outside of the given boundaries.
func Fly(start, direction Point3, bounds Cube, f Flyer) (int, bool) {
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

// A Phaser represent a callback function to be passed to Fly().
// Phaser is like a Walker, but in 4D space
type Phaser func(p Point4) bool

// Phase moves in the given direction from the start point until you exit the
// boundaries, calling f at every step. If f returns true, the phase is
// aborted. Phase returns the number of steps taken, and a bool indicating if
// the phase was aborted prematurely or ran outside of the given boundaries.
func Phase(start, direction Point4, bounds Hypercube, f Phaser) (int, bool) {
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
