package aoc21

import (
	"fmt"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec13a(ctx ch.AOContext) error {
	sections, err := ctx.DataSections("inputs/2021/dec13.txt")
	if err != nil {
		return err
	}

	img := foldTransparentPaper(ctx, sections[0], sections[1][:1])

	rv := 0
	for y := 0; y < img.Height; y++ {
		for x := 0; x < img.Width; x++ {
			if img.At(x, y) == 1 {
				rv++
			}
		}
	}
	ctx.FinalAnswer.Print(rv)
	return nil
}

func Dec13b(ctx ch.AOContext) error {
	sections, err := ctx.DataSections("inputs/2021/dec13.txt")
	if err != nil {
		return err
	}

	img := foldTransparentPaper(ctx, sections[0], sections[1])
	ctx.FinalAnswer.Printf("Image:  (TODO: OCR)\n%s", img)
	return nil
}

func foldTransparentPaper(ctx ch.AOContext, dotCoords, foldInstrs []string) *image {
	w, h := 1, 1
	for _, l := range dotCoords {
		var x, y int
		fmt.Sscanf(l, "%d,%d", &x, &y)
		w = max(w, x+1)
		h = max(h, y+1)
	}

	img := &image{
		Width:    w,
		Height:   h,
		Contents: make([]int, w*h),
	}

	for _, l := range dotCoords {
		var x, y int
		fmt.Sscanf(l, "%d,%d", &x, &y)
		img.Contents[w*y+x] = 1
	}

	//ctx.Printf("Initial image:\n%s", img)

	newW, newH := w, h

	for _, fi := range foldInstrs {
		var axis rune
		var coord int

		fmt.Sscanf(fi, "fold along %c=%d", &axis, &coord)
		if axis == 'x' {
			newW = coord
		} else if axis == 'y' {
			newH = coord
		}

		for y := 0; y < h; y++ {
			b := y
			if y > newH {
				b = 2*newH - b
			}
			for x := 0; x < w; x++ {
				a := x
				if x > newW {
					a = 2*newW - a
				}
				if x != a || y != b {
					if img.At(x, y) == 1 {
						img.Set(a, b, 1)
						img.Set(x, y, 0)
					}
				}
			}
		}

		//ctx.Printf("folded along %c = %d:\n%s", axis, coord, img)
	}

	rv := &image{
		Width:    newW,
		Height:   newH,
		Contents: make([]int, newW*newH),
	}

	for y := 0; y < newH; y++ {
		for x := 0; x < newW; x++ {
			rv.Set(x, y, img.At(x, y))
		}
	}

	return rv
}
