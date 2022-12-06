package aoc22

import (
	"fmt"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec05a(ctx ch.AOContext) error {
	stacks, moveinstrs, err := readStacks(ctx, "inputs/2022/dec05.txt")
	if err != nil {
		return err
	}

	//for i, st := range stacks {
	//	ctx.Printf("stack %d: [%c]", i+1, st)
	//}

	for _, inst := range moveinstrs {
		ct, from, to := inst[0], inst[1], inst[2]

		for i := 0; i < ct; i++ {
			c := stacks[from][len(stacks[from])-1]
			stacks[from] = stacks[from][:len(stacks[from])-1]
			stacks[to] = append(stacks[to], c)
		}

		//for i, st := range stacks {
		//	ctx.Printf("stack %d: [%c]", i+1, st)
		//}
	}

	rv := []byte{}
	for _, st := range stacks {
		rv = append(rv, st[len(st)-1])
	}

	ctx.FinalAnswer.Print(string(rv))
	return nil
}

func Dec05b(ctx ch.AOContext) error {
	stacks, moveinstrs, err := readStacks(ctx, "inputs/2022/dec05.txt")
	if err != nil {
		return err
	}

	for _, inst := range moveinstrs {
		ct, from, to := inst[0], inst[1], inst[2]

		c := stacks[from][len(stacks[from])-ct:]
		stacks[to] = append(stacks[to], c...)
		stacks[from] = stacks[from][:len(stacks[from])-ct]
	}

	rv := []byte{}
	for _, st := range stacks {
		rv = append(rv, st[len(st)-1])
	}

	ctx.FinalAnswer.Print(string(rv))
	return nil
}

func readStacks(ctx ch.AOContext, name string) ([][]byte, [][3]int, error) {
	sections, err := ctx.DataSections(name)
	if err != nil {
		return nil, nil, err
	}

	stacks := make([][]byte, (len(sections[0][len(sections[0])-1])+3)/4)
	for i := range stacks {
		stacks[i] = make([]byte, 0, 16)
	}
	for i := range sections[0][1:] {
		line := sections[0][len(sections[0])-i-2]

		for j := 0; (4*j + 2) < len(line); j++ {
			c := line[4*j+1]
			if c != ' ' {
				stacks[j] = append(stacks[j], c)
			}
		}
	}

	moveinstrs := make([][3]int, 0, len(sections[1]))
	for _, line := range sections[1] {
		ct, from, to := 0, 0, 0
		_, err := fmt.Sscanf(line, "move %d from %d to %d", &ct, &from, &to)
		if err != nil {
			return nil, nil, err
		}

		from--
		to--
		moveinstrs = append(moveinstrs, [3]int{ct, from, to})
	}

	return stacks, moveinstrs, nil
}
