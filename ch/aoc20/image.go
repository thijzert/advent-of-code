package aoc20

type image struct {
	Width, Height int
	Contents      []int
}

type runemap func(rune) int

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
