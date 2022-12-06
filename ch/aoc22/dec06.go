package aoc22

import (
	"github.com/thijzert/advent-of-code/ch"
)

func Dec06a(ctx ch.AOContext) error {
	lines, err := ctx.DataLines("inputs/2022/dec06.txt")
	if err != nil {
		return err
	}
	//lines = []string{
	//	"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
	//	"bvwbjplbgvbhsrlpgdmjqwftvncz",
	//	"nppdvjthqldpwncqszvftbrmjlhg",
	//	"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
	//	"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
	//}

	for _, line := range lines {
		var i int
		for i = range line[4:] {
			seen := make(map[rune]bool)
			for _, c := range line[i : i+4] {
				seen[c] = true
			}
			if len(seen) == 4 {
				break
			}
			i = -1
		}
		ctx.FinalAnswer.Print(i + 4)
	}

	return nil
}

func Dec06b(ctx ch.AOContext) error {
	lines, err := ctx.DataLines("inputs/2022/dec06.txt")
	if err != nil {
		return err
	}

	for _, line := range lines {
		var i int
		for i = range line[14:] {
			seen := make(map[rune]bool)
			for _, c := range line[i : i+14] {
				seen[c] = true
			}
			if len(seen) == 14 {
				break
			}
			i = -1
		}
		ctx.FinalAnswer.Print(i + 14)
	}

	return nil
}
