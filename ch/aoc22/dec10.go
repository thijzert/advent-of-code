package aoc22

import (
	"fmt"

	"github.com/thijzert/advent-of-code/ch"
	"github.com/thijzert/advent-of-code/lib/image"
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

	img := image.NewImage(40, 6, nil)

	x := 1
	cycleCount := 0
	tick := func() {
		i := cycleCount % 40
		j := cycleCount / 40
		cycleCount++
		if x >= i-1 && x <= i+1 {
			img.Set(i, j, 1)
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

	ctx.Printf("\n%s", img)
	ctx.FinalAnswer.Print("(Maybe I should implement OCR at some point)")
	return nil
}
