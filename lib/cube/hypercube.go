package cube

type Point4 struct {
	X, Y, Z, W int
}

// Manhattan returns the Manhattan length of the vector
func (p Point4) Manhattan() int {
	return iabs(p.X) + iabs(p.Y) + iabs(p.Z) + iabs(p.W)
}
func (p Point4) Add(b Point4) Point4 {
	return Point4{
		p.X + b.X,
		p.Y + b.Y,
		p.Z + b.Z,
		p.W + b.W,
	}
}
func (p Point4) Sub(b Point4) Point4 {
	return Point4{
		p.X - b.X,
		p.Y - b.Y,
		p.Z - b.Z,
		p.W - b.W,
	}
}
func (p Point4) Mul(n int) Point4 {
	return Point4{
		p.X * n,
		p.Y * n,
		p.Z * n,
		p.W * n,
	}
}

type Hypercube struct {
	X, Y, Z, W Interval
}

func (a Hypercube) Hypervolume() int {
	return a.X.Length() * a.Y.Length() * a.Z.Length() * a.W.Length()
}

func (a Hypercube) Contains(p Point4) bool {
	return a.X.Contains(p.X) && a.Y.Contains(p.Y) && a.Z.Contains(p.Z) && a.W.Contains(p.W)
}

func (a Hypercube) Overlap(b Hypercube) (Hypercube, bool) {
	xo, ok := a.X.Overlap(b.X)
	if !ok {
		return Hypercube{}, false
	}

	yo, ok := a.Y.Overlap(b.Y)
	if !ok {
		return Hypercube{}, false
	}

	zo, ok := a.Z.Overlap(b.Z)
	if !ok {
		return Hypercube{}, false
	}

	wo, ok := a.W.Overlap(b.W)
	if !ok {
		return Hypercube{}, false
	}

	return Hypercube{xo, yo, zo, wo}, true
}

func (a Hypercube) FullyContains(b Hypercube) bool {
	return a.X.FullyContains(b.X) && a.Y.FullyContains(b.Y) && a.Z.FullyContains(b.Z)
}

// UpdatedBound returns an updated interval that contains the original Interval
// as well as the new point
func (a Hypercube) UpdatedBound(p Point4) Hypercube {
	a.X = a.X.UpdatedBound(p.X)
	a.Y = a.Y.UpdatedBound(p.Y)
	a.Z = a.Z.UpdatedBound(p.Z)
	a.W = a.W.UpdatedBound(p.W)
	return a
}
