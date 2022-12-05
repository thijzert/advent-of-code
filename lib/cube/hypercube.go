package cube

type Point4 struct {
	X, Y, Z, W int
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
