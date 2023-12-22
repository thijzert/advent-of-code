package cube

type Point3 struct {
	X, Y, Z int
}

// Manhattan returns the Manhattan length of the vector
func (p Point3) Manhattan() int {
	return iabs(p.X) + iabs(p.Y) + iabs(p.Z)
}
func (p Point3) Add(b Point3) Point3 {
	return Point3{
		p.X + b.X,
		p.Y + b.Y,
		p.Z + b.Z,
	}
}
func (p Point3) Sub(b Point3) Point3 {
	return Point3{
		p.X - b.X,
		p.Y - b.Y,
		p.Z - b.Z,
	}
}
func (p Point3) Mul(n int) Point3 {
	return Point3{
		p.X * n,
		p.Y * n,
		p.Z * n,
	}
}
func (p Point3) Tr(o Orientation) Point3 {
	p.X, p.Y, p.Z = o.Tr(p.X, p.Y, p.Z)
	return p
}

type Cube struct {
	X, Y, Z Interval
}

func NewCube(x0, y0, z0, x1, y1, z1 int) Cube {
	return Cube{
		X: Interval{A: x0, B: x1},
		Y: Interval{A: y0, B: y1},
		Z: Interval{A: z0, B: z1},
	}
}

func (a Cube) Volume() int {
	return a.X.Length() * a.Y.Length() * a.Z.Length()
}

func (a Cube) Contains(p Point3) bool {
	return a.X.Contains(p.X) && a.Y.Contains(p.Y) && a.Z.Contains(p.Z)
}

func (a Cube) Overlap(b Cube) (Cube, bool) {
	xo, ok := a.X.Overlap(b.X)
	if !ok {
		return Cube{}, false
	}

	yo, ok := a.Y.Overlap(b.Y)
	if !ok {
		return Cube{}, false
	}

	zo, ok := a.Z.Overlap(b.Z)
	if !ok {
		return Cube{}, false
	}

	return Cube{xo, yo, zo}, true
}

func (a Cube) FullyContains(b Cube) bool {
	return a.X.FullyContains(b.X) && a.Y.FullyContains(b.Y) && a.Z.FullyContains(b.Z)
}

// UpdatedBound returns an updated interval that contains the original Interval
// as well as the new point
func (a Cube) UpdatedBound(p Point3) Cube {
	a.X = a.X.UpdatedBound(p.X)
	a.Y = a.Y.UpdatedBound(p.Y)
	a.Z = a.Z.UpdatedBound(p.Z)
	return a
}
