package aoc23

import (
	"strings"

	"github.com/thijzert/advent-of-code/ch"
	"github.com/thijzert/advent-of-code/data"
)

func dec12Arrangements(ctx ch.AOContext, conditionRecord []byte, checksum []int) int {
	if len(checksum) == 0 {
		for _, b := range conditionRecord {
			if b == '#' {
				return 0
			}
		}
		return 1
	}

	n := checksum[0]
	//ctx.Printf("  try and fit a %d in %s", n, conditionRecord)

	rv := 0
	for i := range conditionRecord {
		// Suppose we put checksum[0] at index i

		if i > 0 && conditionRecord[i-1] == '#' {
			// nvm, they've already started.
			break
		}

		// Check if we can place n octothorpes at all
		if len(conditionRecord[i:]) < n {
			continue
		}
		canPlace := true
		for j := 0; j < n; j++ {
			canPlace = canPlace && conditionRecord[i+j] != '.'
		}
		if !canPlace {
			continue
		}
		if len(conditionRecord[i:]) == n {
			// We are at the end of the record. This is either good or bad news
			if len(checksum) == 1 {
				rv += 1
			}
			continue
		}
		if conditionRecord[i+n] == '#' {
			continue
		}
		rv += dec12Arrangements(ctx, conditionRecord[i+n+1:], checksum[1:])
	}
	//ctx.Printf("   %d ways", rv)

	return rv
}

func Dec12a(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2023/dec12.txt")
	if err != nil {
		return nil, err
	}
	//lines = []string{"???.### 1,1,3", ".??..??...?##. 1,1,3", "?#?#?#?#?#?#?#? 1,3,1,6", "????.#...#... 4,1,1", "????.######..#####. 1,6,5", "?###???????? 3,2,1"}

	answer := 0
	for _, line := range lines {
		parts := strings.Split(line, " ")
		arr := dec12Arrangements(ctx, []byte(parts[0]), data.CSVInts(parts[1:])[0])
		ctx.Printf("%s - %d arrangements", line, arr)
		answer += arr
	}

	// 7760: too low
	return answer, nil
}

func Dec12b(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2023/dec12.txt")
	if err != nil {
		return nil, err
	}
	//lines = []string{"???.### 1,1,3", ".??..??...?##. 1,1,3", "?#?#?#?#?#?#?#? 1,3,1,6", "????.#...#... 4,1,1", "????.######..#####. 1,6,5", "?###???????? 3,2,1"}

	answer := 0
	for _, line := range lines {
		parts := strings.Split(line, " ")
		parts[0] = strings.Repeat("?"+parts[0], 5)[1:]
		parts[1] = strings.Repeat(","+parts[1], 5)[1:]
		arr := dec12Arrangements(ctx, []byte(parts[0]), data.CSVInts(parts[1:])[0])
		ctx.Printf("%s - %d arrangements", line, arr)
		answer += arr
	}

	// 7760: too low
	return answer, nil
}
