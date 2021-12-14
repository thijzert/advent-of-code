package aoc20

import (
	"strconv"
	"strings"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec13a(ctx ch.AOContext) error {
	bus, wait := nextBus(939, []int{7, 13, 0, 0, 59, 0, 31, 19})
	ctx.Printf("Example data: %d", bus*wait)

	lines, err := ctx.DataLines("inputs/2020/dec13.txt")
	if err != nil {
		return err
	}

	offset, _ := strconv.Atoi(lines[0])
	var buses []int
	for _, s := range strings.Split(lines[1], ",") {
		n, _ := strconv.Atoi(s)
		buses = append(buses, n)
	}

	bus, wait = nextBus(offset, buses)
	ctx.Printf("Can take bus %d %d minutes from now", bus, wait)
	ctx.FinalAnswer.Print(bus * wait)
	return nil
}

func Dec13b(ctx ch.AOContext) error {
	buses := []int{7, 13, 0, 0, 59, 0, 31, 19}
	offset := synchroniseDepartures(buses)

	for i, b := range buses {
		if b != 0 {
			ctx.Printf("   - Does bus %d depart at time %d: %v", b, offset+i, (offset+i)%b == 0)
		}
	}
	ctx.Printf("Offset %d seems like a real winner", offset)

	lines, err := ctx.DataLines("inputs/2020/dec13.txt")
	if err != nil {
		return err
	}
	buses = nil
	for _, s := range strings.Split(lines[1], ",") {
		n, _ := strconv.Atoi(s)
		buses = append(buses, n)
	}

	offset = synchroniseDepartures(buses)

	ctx.Printf("Consider offset %d", offset)
	for i, b := range buses {
		if b != 0 {
			ctx.Printf("   - Does bus %d depart at time %d: %v", b, offset+i, (offset+i)%b == 0)
		}
	}

	ctx.FinalAnswer.Print(offset)
	return nil
}

func nextBus(offset int, buses []int) (busline int, waitTime int) {
	found := false
	var i, b int
	for !found && i < 1000000 {
		for _, b = range buses {
			if b > 1 && (offset+i)%b == 0 {
				found = true
				break
			}
		}
		if !found {
			i++
		}
	}
	if !found {
		return 0, 0
	}

	return b, i
}

func synchroniseDepartures(buses []int) int {
	offset := 0
	lcm := 0
	for i, b := range buses {
		if lcm == 0 {
			if b != 0 {
				lcm = b
				offset = (b - i) % b
			}
			continue
		}
		if b == 0 {
			continue
		}

		for (offset+i)%b != 0 {
			offset += lcm
		}
		lcm = lcm * b / gcd(lcm, b)
	}

	return offset
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
