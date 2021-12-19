package aoc21

type runemap func(rune) int

type imageFunc func(int, int) int

type image struct {
	Width, Height int
	Contents      []int
}

func newImage(width, height int, f imageFunc) *image {
	rv := &image{
		Height:   height,
		Width:    width,
		Contents: make([]int, height*width),
	}

	if f != nil {
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				rv.Contents[rv.Width*y+x] = f(x, y)
			}
		}
	}

	return rv
}

func readImage(lines []string, f runemap) *image {
	rv := &image{
		Height:   len(lines),
		Width:    len(lines[0]),
		Contents: make([]int, len(lines)*len(lines[0])),
	}

	for y, l := range lines {
		for x, c := range l {
			rv.Contents[rv.Width*y+x] = f(c)
		}
	}

	return rv
}

func octothorpe(r rune) int {
	if r == '#' {
		return 1
	}
	return 0
}

func (i *image) String() string {
	rv := ""

	for y := 0; y < i.Height; y += 2 {
		if y > 0 {
			rv += "\n"
		}
		for x := 0; x < i.Width; x++ {
			rv += blocks(i.At(x, y), i.At(x, 1+y))
		}
	}

	return rv
}

func (i *image) At(x, y int) int {
	if x < 0 || x >= i.Width || y < 0 || y >= i.Height {
		return 0
	}

	return i.Contents[i.Width*y+x]
}

func (i *image) Set(x, y, v int) {
	if x < 0 || x >= i.Width || y < 0 || y >= i.Height {
		return
	}

	i.Contents[i.Width*y+x] = v
}

func blocks(t, b int) string {
	if t == b {
		if t == 0 {
			return " "
		} else if t == 5 {
			return "\x1b[31m\u2588\x1b[0m"
		} else if t == 2 {
			return "\x1b[34m\u2588\x1b[0m"
		} else {
			return "\u2588"
		}
	}

	if t == 0 {
		if b == 5 {
			return "\x1b[31m\u2584\x1b[0m"
		} else if b == 2 {
			return "\x1b[34m\u2584\x1b[0m"
		} else {
			return "\u2584"
		}
	}

	if b == 0 {
		if t == 5 {
			return "\x1b[31m\u2580\x1b[0m"
		} else if t == 2 {
			return "\x1b[34m\u2580\x1b[0m"
		} else {
			return "\u2580"
		}
	}

	if t == 1 {
		if b == 5 {
			return "\x1b[41m\u2580\x1b[0m"
		} else if b == 2 {
			return "\x1b[44m\u2580\x1b[0m"
		}
	}
	if b == 1 {
		if t == 5 {
			return "\x1b[41m\u2584\x1b[0m"
		} else if t == 2 {
			return "\x1b[44m\u2584\x1b[0m"
		}
	}

	if b == 5 && t == 2 {
		return "\x1b[41m\x1b[34m\u2580\x1b[0m"
	} else if b == 2 && t == 5 {
		return "\x1b[41m\x1b[34m\u2584\x1b[0m"
	}

	if t != 0 && b != 0 {
		return "\u2588"
	} else if t != 0 {
		return "\u2580"
	} else if b != 0 {
		return "\u2584"
	} else {
		return " "
	}
}

type point struct {
	X, Y int
}

func (p point) Add(b point) point {
	p.X += b.X
	p.Y += b.Y
	return p
}

type rect struct {
	Min, Max point
}

func (r rect) Contains(p point) bool {
	return p.X >= r.Min.X && p.X <= r.Max.X && p.Y >= r.Min.Y && p.Y <= r.Max.Y
}

type point3 struct {
	X, Y, Z int
}

func (p point3) Add(b point3) point3 {
	p.X += b.X
	p.Y += b.Y
	p.Z += b.Z
	return p
}

func (p point3) Sub(b point3) point3 {
	p.X -= b.X
	p.Y -= b.Y
	p.Z -= b.Z
	return p
}

func (p point3) Tr(o orientation) point3 {
	p.X, p.Y, p.Z = o.Tr(p.X, p.Y, p.Z)
	return p
}

type orientation uint16

func (o orientation) Tr(x, y, z int) (int, int, int) {
	o = o % 24
	switch o {
	// Facing negative Z
	case 1:
		return y, -x, z
	case 2:
		return -x, -y, z
	case 3:
		return -y, x, z
	// Facing positive Z
	case 4:
		return -x, y, -z
	case 5:
		return y, x, -z
	case 6:
		return x, -y, -z
	case 7:
		return -y, -x, -z
	// Facing negative Y
	case 8:
		return -x, z, y
	case 9:
		return -z, -x, y
	case 10:
		return x, -z, y
	case 11:
		return z, x, y
	// Facing positive Y
	case 12:
		return -x, -z, -y
	case 13:
		return z, -x, -y
	case 14:
		return x, z, -y
	case 15:
		return -z, x, -y
	// Facing negative X
	case 16:
		return z, -y, x
	case 17:
		return y, z, x
	case 18:
		return -z, y, x
	case 19:
		return -y, -z, x
	// Facing positive X
	case 20:
		return -z, -y, -x
	case 21:
		return y, -z, -x
	case 22:
		return z, y, -x
	case 23:
		return -y, z, -x
	default:
		return x, y, z
	}
}
