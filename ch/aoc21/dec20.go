package aoc21

import (
	"github.com/thijzert/advent-of-code/ch"
)

func Dec20a(ctx ch.AOContext) (interface{}, error) {
	sections, err := ctx.DataSections("inputs/2021/dec20.txt")
	if err != nil {
		return nil, err
	}

	enhance := sections[0][0]
	ctx.Printf("Enhance algorithm length: %d", len(enhance))

	img := readImage(sections[1], octothorpe)

	img = zoomInAndEnhance(img, enhance, 2)

	return sum(img.Contents...), nil
}

func Dec20b(ctx ch.AOContext) (interface{}, error) {
	sections, err := ctx.DataSections("inputs/2021/dec20.txt")
	if err != nil {
		return nil, err
	}

	img := readImage(sections[1], octothorpe)
	img = zoomInAndEnhance(img, sections[0][0], 50)

	return sum(img.Contents...), nil
}

func zoomInAndEnhance(img *image, enhance string, count int) *image {
	for i := 0; i < count; i++ {
		//ctx.Printf("input image:\n%s", img)

		nextImg := newImage(img.Width+2, img.Height+2, nil)
		if enhance[0] == '#' && i%2 == 0 {
			nextImg.Default = 1
		}

		for y := 0; y < nextImg.Height; y++ {
			for x := 0; x < nextImg.Width; x++ {
				idx := 0
				for b := -1; b <= 1; b++ {
					for a := -1; a <= 1; a++ {
						idx = idx<<1 | img.At(x+a-1, y+b-1)
					}
				}
				nextImg.Set(x, y, octothorpe(rune(enhance[idx])))
			}
		}
		img = nextImg
	}

	//ctx.Printf("final image:\n%s", img)
	return img
}
