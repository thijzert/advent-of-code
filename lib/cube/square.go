package cube

type Point struct {
	X, Y int
}

func Pt(x, y int) Point {
	return Point{X: x, Y: y}
}

// Manhattan returns the Manhattan length of the vector
func (p Point) Manhattan() int {
	return iabs(p.X) + iabs(p.Y)
}
func (p Point) Add(b Point) Point {
	return Point{
		p.X + b.X,
		p.Y + b.Y,
	}
}
func (p Point) Sub(b Point) Point {
	return Point{
		p.X - b.X,
		p.Y - b.Y,
	}
}
func (p Point) Mul(n int) Point {
	return Point{
		p.X * n,
		p.Y * n,
	}
}
func (p Point) Tr(o Orientation) Point {
	p.X, p.Y, _ = o.Tr(p.X, p.Y, 0)
	return p
}

type Square struct {
	X, Y Interval
}

func (a Square) Area() int {
	return a.X.Length() * a.Y.Length()
}

func (a Square) Contains(p Point) bool {
	return a.X.Contains(p.X) && a.Y.Contains(p.Y)
}

func (a Square) Overlap(b Square) (Square, bool) {
	xo, ok := a.X.Overlap(b.X)
	if !ok {
		return Square{}, false
	}

	yo, ok := a.Y.Overlap(b.Y)
	if !ok {
		return Square{}, false
	}

	return Square{xo, yo}, true
}

func (a Square) FullyContains(b Square) bool {
	return a.X.FullyContains(b.X) && a.Y.FullyContains(b.Y)
}

// UpdatedBound returns an updated interval that contains the original Interval
// as well as the new point
func (a Square) UpdatedBound(p Point) Square {
	a.X = a.X.UpdatedBound(p.X)
	a.Y = a.Y.UpdatedBound(p.Y)
	return a
}
