package image

import "github.com/thijzert/advent-of-code/lib/cube"

type Runemap func(rune) int

type ImageFunc func(int, int) int

type Imagery interface {
	Size() (int, int)
	At(x, y int) int
	Set(x, y, v int)
}

type Image struct {
	Width, Height    int
	OffsetX, OffsetY int
	Contents         []int
	Default          int
}

func NewImage(width, height int, f ImageFunc) *Image {
	rv := &Image{
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

func ReadImage(lines []string, f Runemap) *Image {
	rv := &Image{
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

func Octothorpe(r rune) int {
	if r == '#' {
		return 1
	}
	return 0
}

func (i *Image) String() string {
	rv := ""

	for y := 0; y < i.Height; y += 2 {
		if y > 0 {
			rv += "\n"
		}
		for x := 0; x < i.Width; x++ {
			rv += blocks(i.At(i.OffsetX+x, i.OffsetY+y), i.At(i.OffsetX+x, i.OffsetY+1+y))
		}
	}

	return rv
}

func (i *Image) Size() (int, int) {
	return i.Width, i.Height
}

func (i *Image) Inside(x, y int) bool {
	return x >= 0 && x < i.Width && y >= 0 && y < i.Height
}

func (i *Image) At(x, y int) int {
	x -= i.OffsetX
	y -= i.OffsetY
	if x < 0 || x >= i.Width || y < 0 || y >= i.Height {
		return i.Default
	}

	return i.Contents[i.Width*y+x]
}

func (i *Image) Set(x, y, v int) {
	x -= i.OffsetX
	y -= i.OffsetY
	if x < 0 || x >= i.Width || y < 0 || y >= i.Height {
		return
	}

	i.Contents[i.Width*y+x] = v
}

func (i *Image) Sprite(sprite, mask *Image, offsetX, offsetY int) {
	for y := 0; y < sprite.Height; y++ {
		for x := 0; x < sprite.Width; x++ {
			m := mask.At(x, y)
			v := m*sprite.At(x, y) + (1-m)*i.At(offsetX+x, offsetY+y)
			i.Set(offsetX+x, offsetY+y, v)
		}
	}
}

func (i *Image) FloodFill(x, y int, v int) int {
	v0 := i.At(x, y)
	if v0 == v {
		return 0
	}

	rv := 0
	queue := make([]cube.Point, 1)
	queue[0].X = x
	queue[0].Y = y
	for len(queue) > 0 {
		pos := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if i.At(pos.X, pos.Y) != v0 {
			continue
		}
		i.Set(pos.X, pos.Y, v)
		rv++
		for _, dir := range cube.Cardinal2D {
			next := pos.Add(dir)
			if i.Inside(next.X, next.Y) {
				queue = append(queue, next)
			}
		}
	}
	return rv
}

func (i *Image) MaskAt(mask *Image, x, y int) int {
	rv := 0
	for b := 0; b < mask.Height; b++ {
		for a := 0; a < mask.Width; a++ {
			rv += mask.At(a, b) * i.At(x+a, y+b)
		}
	}
	return rv
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
