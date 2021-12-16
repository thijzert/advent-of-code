package aoc20

import (
	"fmt"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec12a(ctx ch.AOContext) error {
	lines := []string{"F10", "N3", "F7", "R90", "F11"}
	x, y, _, err := navigateShip(lines)
	if err != nil {
		return err
	}

	ctx.Printf("Ship's position: %d,%d", x, y)
	ctx.Printf("Taxicab distance: %d", abs(x)+abs(y))

	lines, err = ctx.DataLines("inputs/2020/dec12.txt")
	if err != nil {
		return err
	}

	x, y, _, err = navigateShip(lines)
	if err != nil {
		return err
	}

	ctx.Printf("Ship's position: %d,%d", x, y)
	ctx.FinalAnswer.Print(abs(x) + abs(y))
	return nil
}

func Dec12b(ctx ch.AOContext) error {
	lines := []string{"F10", "N3", "F7", "R90", "F11"}
	x, y, err := navigateWaypoint(lines)
	if err != nil {
		return err
	}

	ctx.Printf("Ship's position: %d,%d", x, y)
	ctx.Printf("Taxicab distance: %d", abs(x)+abs(y))

	lines, err = ctx.DataLines("inputs/2020/dec12.txt")
	if err != nil {
		return err
	}

	x, y, err = navigateWaypoint(lines)
	if err != nil {
		return err
	}

	ctx.Printf("Ship's position: %d,%d", x, y)
	ctx.FinalAnswer.Print(abs(x) + abs(y))
	return nil
}

func navigateShip(instructions []string) (x, y, dir int, err error) {
	for _, l := range instructions {
		var c rune
		var dist int
		if _, err = fmt.Sscanf(l, "%c%d", &c, &dist); err != nil {
			continue
		}

		if c == 'N' || (c == 'F' && dir == 1) {
			y += dist
		} else if c == 'S' || (c == 'F' && dir == 3) {
			y -= dist
		} else if c == 'E' || (c == 'F' && dir == 0) {
			x += dist
		} else if c == 'W' || (c == 'F' && dir == 2) {
			x -= dist
		} else if c == 'R' || c == 'L' {
			if dist%90 != 0 {
				return 0, 0, 0, fmt.Errorf("Don't know how to handle '%s'", l)
			}
			if c == 'R' {
				dir = (dir - dist/90) % 4
				if dir < 0 {
					dir += 4
				}
			} else {
				dir = (dir + dist/90) % 4
			}
		} else {
			return 0, 0, 0, fmt.Errorf("Don't know how to handle '%s'", l)
		}
	}

	err = nil
	return
}

func navigateWaypoint(instructions []string) (x, y int, err error) {
	wpX, wpY := 10, 1
	for _, l := range instructions {
		var c rune
		var dist int
		if _, err = fmt.Sscanf(l, "%c%d", &c, &dist); err != nil {
			continue
		}

		if c == 'N' {
			wpY += dist
		} else if c == 'S' {
			wpY -= dist
		} else if c == 'E' {
			wpX += dist
		} else if c == 'W' {
			wpX -= dist
		} else if c == 'R' || c == 'L' {
			if dist%90 != 0 {
				return 0, 0, fmt.Errorf("Don't know how to handle '%s'", l)
			}
			dist = (dist / 90) % 4
			if c == 'R' {
				dist = 4 - dist
			}
			for i := 0; i < dist; i++ {
				wpX, wpY = -wpY, wpX
			}
		} else if c == 'F' {
			x += dist * wpX
			y += dist * wpY
		} else {
			return 0, 0, fmt.Errorf("Don't know how to handle '%s'", l)
		}
	}

	err = nil
	return
}
