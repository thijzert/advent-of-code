package aoc21

import "github.com/thijzert/advent-of-code/ch"

var allDirections [4]point = [4]point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func Dec15a(ctx ch.AOContext) (interface{}, error) {
	sections, err := ctx.DataSections("inputs/2021/dec15.txt")
	if err != nil {
		return nil, err
	}
	ctx.Print(len(sections))

	img := readImage(sections[0], func(r rune) int {
		return int(r - '0')
	})

	rv := shortestPath(img)
	return rv, nil
}

func Dec15b(ctx ch.AOContext) (interface{}, error) {
	sections, err := ctx.DataSections("inputs/2021/dec15.txt")
	if err != nil {
		return nil, err
	}
	ctx.Print(len(sections))

	img := readImage(sections[0], func(r rune) int {
		return int(r - '0')
	})

	tiledImage := &image{
		Width:    5 * img.Width,
		Height:   5 * img.Height,
		Contents: make([]int, 25*len(img.Contents)),
	}
	for ty := 0; ty < 5; ty++ {
		for tx := 0; tx < 5; tx++ {
			d := ty + tx
			for y := 0; y < img.Height; y++ {
				for x := 0; x < img.Width; x++ {
					c := 1 + (img.At(x, y)+d-1)%9
					tiledImage.Set(img.Width*tx+x, img.Height*ty+y, c)
				}
			}
		}
	}

	rv := shortestPath(tiledImage)
	return rv, nil
}

func shortestPath(img *image) int {
	MTD := 10 * (img.Width + img.Height)

	previous := &image{
		Width:    img.Width,
		Height:   img.Height,
		Contents: make([]int, len(img.Contents)),
	}
	cumulative := &image{
		Width:    img.Width,
		Height:   img.Height,
		Contents: make([]int, len(img.Contents)),
	}
	for i := range cumulative.Contents {
		cumulative.Contents[i] = MTD
	}

	pqueue := make([]map[point]bool, MTD+1)
	for i := range pqueue {
		pqueue[i] = make(map[point]bool)
	}
	for y := 0; y < img.Height; y++ {
		for x := 0; x < img.Height; x++ {
			pqueue[MTD][point{x, y}] = true
		}
	}

	for i := 0; i < img.Width*img.Height; i++ {
		d := 0
		m := pqueue[MTD]
		pt := point{0, 0}
		if i > 0 {
			for d, m = range pqueue {
				if len(m) == 0 {
					continue
				}
				for pt = range m {
					break
				}
				break
			}
		}

		delete(m, pt)
		cumulative.Set(pt.X, pt.Y, d)

		for j, dir := range allDirections {
			nb := point{pt.X + dir.X, pt.Y + dir.Y}
			cumul := cumulative.At(nb.X, nb.Y)
			ncum := d + img.At(nb.X, nb.Y)
			if ncum < cumul {
				previous.Set(nb.X, nb.Y, 1+((j+2)%4))
				cumulative.Set(nb.X, nb.Y, ncum)
				delete(pqueue[cumul], nb)
				pqueue[ncum][nb] = true
			}
		}
	}

	pt := point{img.Width - 1, img.Height - 1}
	return cumulative.At(pt.X, pt.Y)
}
