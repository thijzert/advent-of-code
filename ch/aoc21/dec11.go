package aoc21

import (
	"strconv"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec11a(ctx ch.AOContext) error {
	octo, err := readOctopusImage(ctx, "inputs/2021/dec11.txt")
	if err != nil {
		return err
	}

	ctx.Printf("Initial board: (%d×%d)\n%s", octo.Width, octo.Height, octo)
	rv := 0
	for i := 0; i < 100; i++ {
		rv += octo.stepOctopi()
		if i < 10 || i%10 == 9 {
			ctx.Printf("After step %d:\n%s", i+1, octo)
		}
	}

	ctx.FinalAnswer.Print(rv)
	return nil
}

func Dec11b(ctx ch.AOContext) error {
	octo, err := readOctopusImage(ctx, "inputs/2021/dec11.txt")
	if err != nil {
		return err
	}

	prevprev := "(?)"
	prev := "(?)"

	rv := 0
	for i := 0; i < 10000; i++ {
		n := octo.stepOctopi()
		if n == octo.Width*octo.Height {
			ctx.Printf("After step %d:\n%s", i-1, prevprev)
			ctx.Printf("After step %d:\n%s", i, prev)
			ctx.Printf("After step %d:\n%s", i+1, octo)
			rv = i + 1
			break
		}

		prevprev = prev
		prev = octo.String()
	}

	ctx.FinalAnswer.Print(rv)
	return nil
}

type octopusImage struct {
	Width, Height int
	Octopi        []int
}

func readOctopusImage(ctx ch.AOContext, assetName string) (octopusImage, error) {
	lines, err := ctx.DataLines(assetName)
	if err != nil {
		return octopusImage{}, err
	}
	rv := octopusImage{}
	if len(lines) == 0 {
		return rv, nil
	}

	rv.Width = len(lines[0])
	for _, l := range lines {
		if l != "" {
			rv.Height++
		}
		for _, v := range l {
			rv.Octopi = append(rv.Octopi, int(v-'0'))
		}
	}

	return rv, nil
}

func (o octopusImage) String() string {
	rv := ""
	for i, v := range o.Octopi {
		if i > 0 && i%o.Width == 0 {
			rv += "\n"
		}
		if v == 0 {
			rv += "\x1b[1m0\x1b[0m"
		} else {
			rv += strconv.Itoa(v)
		}
	}
	return rv
}

func (o octopusImage) At(x, y int) int {
	if x < 0 || x >= o.Width {
		return 0
	}
	if y < 0 || y >= o.Height {
		return 0
	}

	return o.Octopi[y*o.Width+x]
}

func (o octopusImage) Set(x, y, v int) {
	if x < 0 || x >= o.Width || y < 0 || y >= o.Height {
		return
	}

	o.Octopi[y*o.Width+x] = v
}

func (o octopusImage) Inc(x, y int) int {
	if x < 0 || x >= o.Width || y < 0 || y >= o.Height {
		return 0
	}

	before := o.Octopi[y*o.Width+x]
	o.Octopi[y*o.Width+x]++

	if before <= 9 && o.Octopi[y*o.Width+x] > 9 {
		return o.flash(x, y)
	}
	return 0
}

func (o octopusImage) stepOctopi() int {
	rv := 0
	for y := 0; y < o.Height; y++ {
		for x := 0; x < o.Width; x++ {
			rv += o.Inc(x, y)
		}
	}

	for i, v := range o.Octopi {
		if v > 9 {
			o.Octopi[i] = 0
		}
	}

	return rv
}

func (o octopusImage) flash(x, y int) int {
	if x < 0 || x >= o.Width || y < 0 || y >= o.Height {
		return 0
	}

	rv := 1
	for a := -1; a <= 1; a++ {
		for b := -1; b <= 1; b++ {
			if a != 0 || b != 0 {
				rv += o.Inc(x+b, y+a)
			}
		}
	}

	return rv
}
