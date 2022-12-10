package aoc22

import (
	"fmt"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec10a(ctx ch.AOContext) error {
	program, err := ctx.DataLines("inputs/2022/dec10.txt")
	if err != nil {
		return err
	}

	signalStrength := 0
	x := 1
	cycleCount := 0
	tick := func() {
		cycleCount++
		if (cycleCount+20)%40 == 0 {
			ctx.Printf("During the %dth cycle, register X has the value %d, so the signal strength is %d * %d = %d.", cycleCount, x, cycleCount, x, cycleCount*x)
			signalStrength += cycleCount * x
		}
	}

	for _, instr := range program {
		if instr == "noop" {
			tick()
		} else if len(instr) > 5 && instr[:5] == "addx " {
			var operand int
			_, err := fmt.Sscanf(instr, "addx %d", &operand)
			if err != nil {
				return err
			}
			tick()
			tick()
			x += operand
		} else {
			return fmt.Errorf("unknown instruction '%s'", instr)
		}
	}

	ctx.FinalAnswer.Print(signalStrength)
	return nil
}

func Dec10b(ctx ch.AOContext) error {
	program, err := ctx.DataLines("inputs/2022/dec10.txt")
	if err != nil {
		return err
	}

	x := 1
	cycleCount := 0
	scanline := make([]byte, 40)
	tick := func() {
		i := cycleCount % 40
		cycleCount++
		if x >= i-1 && x <= i+1 {
			scanline[i] = '#'
		} else {
			scanline[i] = ' '
		}
		if i == 39 {
			ctx.Printf("%s", scanline)
		}
	}

	for _, instr := range program {
		if instr == "noop" {
			tick()
		} else if len(instr) > 5 && instr[:5] == "addx " {
			var operand int
			fmt.Sscanf(instr, "addx %d", &operand) // We checked the validity in the last challenge
			tick()
			tick()
			x += operand
		} else {
			return fmt.Errorf("unknown instruction '%s'", instr)
		}
	}

	ctx.FinalAnswer.Print("(Maybe I should implement OCR at some point)")
	return nil
}
