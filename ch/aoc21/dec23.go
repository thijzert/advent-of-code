package aoc21

import (
	"fmt"
	"log"

	"github.com/thijzert/advent-of-code/ch"
)

var Dec23b ch.AdventFunc

func Dec23a(ctx ch.AOContext) error {
	// Worked this out on a napkin
	ctx.FinalAnswer.Print(15160)
	return nil
}

func thereWasAnAttemptAtDec23b(ctx ch.AOContext) error {
	lines, err := ctx.DataLines("inputs/2021/dec23.txt")
	if err != nil {
		return err
	}

	cave := readAmphipodCave(lines)
	// cave.Move(6, 1, 8, 0)
	// cave.Move(6, 2, -2, 0)
	// cave.Move(4, 1, -1, 0)
	// cave.Move(4, 2, 7, 0)
	// cave.Move(4, 3, 1, 0)
	// cave.Move(6, 3, 3, 0)

	// Poging 2

	// cave.Move(6, 1, 8, 0)
	// cave.Move(6, 2, -2, 0)
	// cave.Move(6, 3, 7, 0)
	// cave.Move(6, 4, -1, 0)
	// cave.Move(0, 1, 6, 4)
	// cave.Move(0, 2, 6, 3)
	// cave.Move(0, 3, 6, 2)
	//cave.Move(0, 4, 1, 0)
	//cave.Move(-1, 0, 0, 4)
	//cave.Move(-2, 0, 0, 3)
	// cave.Move(1, 0, -2, 0)
	// cave.Move(4, 1, 0, 2)
	// cave.Move(4, 2, -1, 0)
	// cave.Move(4, 3, 0, 1)
	// cave.Move(4, 4, 6, 1)
	// cave.Move(2, 1, 4, 4)
	// cave.Move(2, 2, 4, 3)
	// cave.Move(2, 3, 1, 0)
	// cave.Move(2, 4, 4, 2)
	// cave.Move(7, 0, 4, 1)
	// cave.Move(1, 0, 2, 4)
	// cave.Move(8, 0, 2, 3)
	// cave.Move(-1, 0, 2, 2)
	// cave.Move(-2, 0, 2, 1)

	//43506: too low

	// poging 3
	cave.Move(6, 1, 8, 0)
	cave.Move(6, 2, -2, 0)
	cave.Move(6, 3, 7, 0)
	cave.Move(6, 4, -1, 0)
	cave.Move(0, 1, 6, 4)
	cave.Move(0, 2, 6, 3)
	cave.Move(0, 3, 6, 2)

	cave.Move(0, 4, 5, 0)
	cave.Move(-1, 0, 0, 4)
	cave.Move(-2, 0, 0, 3)
	cave.Move(4, 1, 0, 2)
	cave.Move(4, 2, -2, 0)
	cave.Move(4, 3, 0, 1)

	cave.Move(4, 4, -1, 0)
	cave.Move(2, 1, 4, 4)
	cave.Move(2, 2, 4, 3)
	cave.Move(2, 3, 1, 0)
	cave.Move(2, 4, 4, 2)
	cave.Move(1, 0, 2, 4)
	cave.Move(5, 0, 2, 3)
	cave.Move(-1, 0, 6, 1)
	cave.Move(-2, 0, 2, 2)
	cave.Move(7, 0, 4, 1)
	cave.Move(8, 0, 2, 1)
	// 53526: too high

	ctx.Print(cave)

	return errNotImplemented
}

type amphipodCave struct {
	Width, Height int
	Contents      []rune
	EnergyUsed    int
}

func readAmphipodCave(lines []string) *amphipodCave {
	rv := &amphipodCave{
		Height:   len(lines),
		Width:    len(lines[0]),
		Contents: make([]rune, len(lines)*len(lines[0])),
	}

	for y, l := range lines {
		for x, c := range l {
			rv.Contents[rv.Width*y+x] = c
		}
	}

	return rv
}

func (c *amphipodCave) String() string {
	rv := fmt.Sprintf("Energy used: %d", c.EnergyUsed)

	for y := 0; y < c.Height; y++ {
		rv += "\n"
		for x := 0; x < c.Width; x++ {
			rv += string(c.Contents[c.Width*y+x])
		}
	}

	return rv
}

func (c *amphipodCave) Move(x1, y1, x2, y2 int) {
	dist := abs(x1-x2) + abs(y1) + abs(y2)

	x1 += 3
	x2 += 3
	y1 += 1
	y2 += 1

	for y := y1 - 1; y > 1; y-- {
		if c.Contents[c.Width*y+x1] != '.' {
			panic("starting path not empty")
		}
	}
	for y := y2 - 1; y > 1; y-- {
		if c.Contents[c.Width*y+x2] != '.' {
			panic("destination path not empty")
		}
	}
	XM := max(x1, x2)
	for x := min(x1, x2); x <= XM; x++ {
		if x == x1 && y1 == 1 {
			continue
		}
		if c.Contents[c.Width+x] != '.' {
			log.Printf("x: %d xmin %d, xmax %d", x, min(x1, x2), XM)
			log.Print(c)
			panic("hallway path not empty")
		}
	}

	c.Contents[c.Width*y2+x2], c.Contents[c.Width*y1+x1] = c.Contents[c.Width*y1+x1], '.'

	if c.Contents[c.Width*y2+x2] == 'A' {
		c.EnergyUsed += dist
	} else if c.Contents[c.Width*y2+x2] == 'B' {
		c.EnergyUsed += dist * 10
	} else if c.Contents[c.Width*y2+x2] == 'C' {
		c.EnergyUsed += dist * 100
	} else if c.Contents[c.Width*y2+x2] == 'D' {
		c.EnergyUsed += dist * 1000
	} else {
		panic("Don't know the energy budget for this")
	}
}
