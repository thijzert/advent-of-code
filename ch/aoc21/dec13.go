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
			if img.Get(x, y) == 1 {
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
		Width:         w,
		Height:        h,
		DisplayWidth:  w,
		DisplayHeight: h,
		Contents:      make([]int, w*h),
	}

	for _, l := range dotCoords {
		var x, y int
		fmt.Sscanf(l, "%d,%d", &x, &y)
		img.Contents[w*y+x] = 1
	}

	//ctx.Printf("Initial image:\n%s", img)

	for _, fi := range foldInstrs {
		var axis rune
		var coord int

		fmt.Sscanf(fi, "fold along %c=%d", &axis, &coord)
		if axis == 'x' {
			img.DisplayWidth = coord
		} else if axis == 'y' {
			img.DisplayHeight = coord
		}

		for y := 0; y < h; y++ {
			b := y
			if y > img.DisplayHeight {
				b = 2*img.DisplayHeight - b
			}
			for x := 0; x < w; x++ {
				a := x
				if x > img.DisplayWidth {
					a = 2*img.DisplayWidth - a
				}
				if x != a || y != b {
					if img.Get(x, y) == 1 {
						img.Set(a, b, 1)
						img.Set(x, y, 0)
					}
				}
			}
		}

		//ctx.Printf("folded along %c = %d:\n%s", axis, coord, img)
	}

	return img
}

type image struct {
	Width, Height               int
	DisplayWidth, DisplayHeight int
	Contents                    []int
}

func readImage(lines []string, one rune) *image {
	rv := &image{
		Height:        len(lines),
		Width:         len(lines[0]),
		DisplayHeight: len(lines),
		DisplayWidth:  len(lines[0]),
		Contents:      make([]int, len(lines)*len(lines[0])),
	}

	for y, l := range lines {
		for x, c := range l {
			if c == one {
				rv.Contents[rv.Width*y+x] = 1
			}
		}
	}

	return rv
}

func (i *image) String() string {
	w, h := i.Width, i.Height
	if i.DisplayHeight != 0 {
		h = i.DisplayHeight
	}
	if i.DisplayWidth != 0 {
		w = i.DisplayWidth
	}

	rv := ""
	for y := 0; y < h; y++ {
		if y > 0 {
			rv += "\n"
		}
		for x := 0; x < w; x++ {
			c := i.Contents[i.Width*y+x]
			if c == 0 {
				rv += "."
			} else {
				rv += "#"
			}
		}
	}
	return rv
}

func (i *image) Get(x, y int) int {
	if x < 0 || x >= i.Width || y < 0 || y > i.Height {
		return 0
	}

	return i.Contents[i.Width*y+x]
}

func (i *image) Set(x, y, v int) {
	if x < 0 || x >= i.Width || y < 0 || y > i.Height {
		return
	}

	i.Contents[i.Width*y+x] = v
}
