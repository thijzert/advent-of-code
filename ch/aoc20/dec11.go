package aoc20

import (
	"github.com/thijzert/advent-of-code/ch"
)

const (
	SEAT_EMPTY    = 1
	SEAT_OCCUPIED = 5
)

func Dec11a(ctx ch.AOContext) error {
	img := readImage([]string{
		"L.LL.LL.LL",
		"LLLLLLL.LL",
		"L.L.L..L..",
		"LLLL.LL.LL",
		"L.LL.LL.LL",
		"L.LLLLL.LL",
		"..L.L.....",
		"LLLLLLLLLL",
		"L.LLLLLL.L",
		"L.LLLLL.LL",
	}, seatmap)

	ctx.Printf("Example seating map:\n%s", img)

	i, n := 0, 1
	for n > 0 {
		i++
		n = fillSeatsStep(img, 1, 4)
	}
	ctx.Printf("After %d steps:\n%s", i, img)

	rv := seatsOccupied(img)
	ctx.Printf("In the example data, %d seats are occupied", rv)

	lines, err := ctx.DataLines("inputs/2020/dec11.txt")
	if err != nil {
		return err
	}
	img = readImage(lines[:len(lines)-1], seatmap)

	i, n = 0, 1
	for n > 0 {
		i++
		n = fillSeatsStep(img, 1, 4)
	}
	ctx.Printf("After %d steps:\n%s", i, img)

	rv = seatsOccupied(img)
	ctx.Printf("In the final data, %d seats are occupied", rv)

	ctx.FinalAnswer.Print(rv)
	return nil
}

func Dec11b(ctx ch.AOContext) error {
	img := readImage([]string{
		"L.LL.LL.LL",
		"LLLLLLL.LL",
		"L.L.L..L..",
		"LLLL.LL.LL",
		"L.LL.LL.LL",
		"L.LLLLL.LL",
		"..L.L.....",
		"LLLLLLLLLL",
		"L.LLLLLL.L",
		"L.LLLLL.LL",
	}, seatmap)

	ctx.Printf("Example seating map:\n%s", img)

	maxDist := max(img.Width, img.Height) + 6
	i, n := 0, 1
	for n > 0 {
		i++
		n = fillSeatsStep(img, maxDist, 5)
	}
	ctx.Printf("After %d steps:\n%s", i, img)

	rv := seatsOccupied(img)
	ctx.Printf("In the example data, %d seats are occupied", rv)

	lines, err := ctx.DataLines("inputs/2020/dec11.txt")
	if err != nil {
		return err
	}
	img = readImage(lines[:len(lines)-1], seatmap)

	i, n = 0, 1
	for n > 0 {
		i++
		n = fillSeatsStep(img, maxDist, 5)
		//n = fillSeatsStep(img, 1, 4)
		//n = 5 - i
	}
	ctx.Printf("After %d steps:\n%s", i, img)

	rv = seatsOccupied(img)
	ctx.Printf("In the final data, %d seats are occupied", rv)

	ctx.FinalAnswer.Print(rv)
	return nil
}

func fillSeatsStep(img *image, maxDist int, occupancyTolerance int) int {
	changed := 0

	for y := 0; y < img.Height; y++ {
		for x := 0; x < img.Width; x++ {
			current := img.At(x, y)
			next := current
			if current == 0 {
				continue
			}

			n := 0
			for b := -1; b <= 1; b++ {
				for a := -1; a <= 1; a++ {
					if a != 0 || b != 0 {
						for c := 1; c <= maxDist; c++ {
							other := img.At(x+c*a, y+c*b) & 0xf
							if other == SEAT_OCCUPIED {
								n++
							}
							if other != 0 {
								break
							}
						}
					}
				}
			}

			if current == SEAT_EMPTY && n == 0 {
				next = SEAT_OCCUPIED
				changed++
			} else if current == SEAT_OCCUPIED && n >= occupancyTolerance {
				next = SEAT_EMPTY
				changed++
			}

			img.Set(x, y, current|(next<<4))
		}
	}

	for y := 0; y < img.Height; y++ {
		for x := 0; x < img.Width; x++ {
			img.Set(x, y, img.At(x, y)>>4)
		}
	}

	return changed
}

func seatmap(r rune) int {
	if r == '#' {
		return SEAT_OCCUPIED
	} else if r == 'L' {
		return SEAT_EMPTY
	}
	return 0
}

func seatsOccupied(img *image) int {
	rv := 0
	for y := 0; y < img.Height; y++ {
		for x := 0; x < img.Width; x++ {
			if img.At(x, y) == SEAT_OCCUPIED {
				rv++
			}
		}
	}
	return rv
}
