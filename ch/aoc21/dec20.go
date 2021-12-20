package aoc21

import (
	"github.com/thijzert/advent-of-code/ch"
)

var Dec20b ch.AdventFunc = nil

func Dec20a(ctx ch.AOContext) error {
	sections, err := ctx.DataSections("inputs/2021/dec20.txt")
	if err != nil {
		return err
	}

	enhance := sections[0][0]
	ctx.Printf("Enhance algorithm length: %d", len(enhance))

	img := readImage(sections[1], octothorpe)

	for i := 0; i < 2; i++ {
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

	ctx.FinalAnswer.Print(sum(img.Contents...))
	return nil
}

// 5363: too high (?)
// 5826: too high

// func Dec20b(ctx ch.AOContext) error {
// 	return errNotImplemented
// }
