package image

import "github.com/thijzert/advent-of-code/lib/cube"

type Rotated struct {
	Orientation cube.Orientation
	Img         Imagery
}

func (r Rotated) transform(x, y int) (int, int) {
	w, h := r.Img.Size()
	w, h = w-1, h-1

	if r.Orientation == 0 {
		return x, y
	} else if r.Orientation == 1 {
		return y, w - x
	} else if r.Orientation == 2 {
		return w - x, h - y
	} else if r.Orientation == 3 {
		return h - y, x
	}
	return -1, -1
}

func (r Rotated) Size() (int, int) {
	w, h := r.Img.Size()
	if r.Orientation%2 == 1 {
		return h, w
	}
	return w, h
}

func (r Rotated) At(x, y int) int {
	xx, yy := r.transform(x, y)
	return r.Img.At(xx, yy)
}

func (r Rotated) Set(x, y, v int) {
	xx, yy := r.transform(x, y)
	r.Img.Set(xx, yy, v)
}
