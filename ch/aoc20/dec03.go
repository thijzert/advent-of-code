package aoc20

import "github.com/thijzert/advent-of-code/ch"

func Dec03a(ctx ch.AOContext) error {
	sections, err := ctx.DataSections("inputs/2020/dec03.txt")
	if err != nil {
		return err
	}
	img := readImage(sections[0], octothorpe)

	// ctx.Printf("Forest:\n%s", img)

	a := 3
	b := 1
	rv := 0
	for c := 0; b*c < img.Height; c++ {
		x, y := (c*a)%img.Width, (c*b)%img.Height
		if img.At(x, y) == 1 {
			img.Set(x, y, 5)
			rv++
		} else {
			img.Set(x, y, 2)
		}
	}

	ctx.Printf("Route:\n%s", img)

	ctx.FinalAnswer.Print(rv)
	return nil
}

func Dec03b(ctx ch.AOContext) error {
	sections, err := ctx.DataSections("inputs/2020/dec03.txt")
	if err != nil {
		return err
	}
	img := readImage(sections[0], octothorpe)

	rv := 1

	slopes := []struct{ A, B int }{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	for _, sl := range slopes {
		trees := 0
		for c := 0; sl.B*c < img.Height; c++ {
			trees += img.At((c*sl.A)%img.Width, (c*sl.B)%img.Height)
		}
		rv *= trees
	}

	ctx.FinalAnswer.Print(rv)
	return nil
}
