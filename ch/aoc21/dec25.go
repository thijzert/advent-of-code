package aoc21

import (
	"github.com/thijzert/advent-of-code/ch"
)

func Dec25a(ctx ch.AOContext) error {
	lines, err := ctx.DataLines("inputs/2021/dec25.txt")
	if err != nil {
		return err
	}
	img := readImage(lines, seaCucumber)
	ctx.Printf("ocean floor: (%d)\n%s", sum(img.Contents...), img)

	n := 1
	rv := 0
	for n > 0 {
		rv++
		n = moveCucumbers(img)
		if rv%100 == 0 {
			ctx.Printf("ocean floor: (%d)\n%s", sum(img.Contents...), img)
		}
	}

	ctx.FinalAnswer.Print(rv)
	return nil
}

func Dec25b(ctx ch.AOContext) error {
	ctx.FinalAnswer.Print("merry christmas")
	return nil
}

const CUCUMBER_SOUTH int = 2
const CUCUMBER_EAST int = 5

func seaCucumber(x rune) int {
	if x == 'v' {
		return CUCUMBER_SOUTH
	} else if x == '>' {
		return CUCUMBER_EAST
	}
	return 0
}

func moveCucumbers(seafloor *image) int {
	rv := 0

	// Pass 1a: for the eastbound herd, mark all that are going to move
	for y := 0; y < seafloor.Height; y++ {
		for x := seafloor.Width - 1; x >= 0; x-- {
			if seafloor.At(x, y) == CUCUMBER_EAST {
				x1 := (x + 1) % seafloor.Width
				if seafloor.At(x1, y) == 0 {
					rv++
					seafloor.Set(x, y, -CUCUMBER_EAST)
				}
			}
		}
	}

	// Pass 1b: move the eastbound herd
	for y := 0; y < seafloor.Height; y++ {
		for x := seafloor.Width - 1; x >= 0; x-- {
			c := seafloor.At(x, y)
			if c == -CUCUMBER_EAST {
				x1 := (x + 1) % seafloor.Width
				seafloor.Set(x1, y, -c)
				seafloor.Set(x, y, 0)
			} else {
				seafloor.Set(x, y, c)
			}
		}
	}

	// Pass 2a: for the southbound herd, mark all that are going to move
	for y := seafloor.Height - 1; y >= 0; y-- {
		for x := 0; x < seafloor.Width; x++ {
			if seafloor.At(x, y) == CUCUMBER_SOUTH {
				y1 := (y + 1) % seafloor.Height
				if seafloor.At(x, y1) == 0 {
					rv++
					seafloor.Set(x, y, -CUCUMBER_SOUTH)
				}
			}
		}
	}

	// Pass 2b: move the southbound herd
	for y := seafloor.Height - 1; y >= 0; y-- {
		for x := 0; x < seafloor.Width; x++ {
			c := seafloor.At(x, y)
			if c == -CUCUMBER_SOUTH {
				y1 := (y + 1) % seafloor.Height
				seafloor.Set(x, y1, -c)
				seafloor.Set(x, y, 0)
			} else {
				seafloor.Set(x, y, c)
			}
		}
	}

	return rv
}
