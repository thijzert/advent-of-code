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
