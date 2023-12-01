package aoc23

import (
	"github.com/thijzert/advent-of-code/ch"
)

func Dec01a(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2023/dec01.txt")
	if err != nil {
		return nil, err
	}
	//lines = []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}

	answer := 0
	for _, line := range lines {
		found := false
		f, l := 0, 0
		for _, c := range line {
			if c >= '0' && c <= '9' {
				if !found {
					f = int(c - '0')
				}
				found = true
				l = int(c - '0')
			}
		}
		answer += 10*f + l
	}

	return answer, nil
}

func Dec01b(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2023/dec01.txt")
	if err != nil {
		return nil, err
	}
	//lines = []string{"two1nine", "eightwothree", "abcone2threexyz", "xtwone3four", "4nineeightseven2", "zoneight234", "7pqrstsixteen"}
	digits := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	answer := 0
	for _, line := range lines {
		found := false
		f, l := 0, 0
		for i, c := range line {
			if c >= '0' && c <= '9' {
				if !found {
					f = int(c - '0')
				}
				found = true
				l = int(c - '0')
			} else {
				for d, digit := range digits {
					if len(line[i:]) >= len(digit) && line[i:i+len(digit)] == digit {
						if !found {
							f = d
						}
						found = true
						l = d
					}
				}
			}
		}
		ctx.Printf("Calibration value: %d", 10*f+l)
		answer += 10*f + l
	}

	return answer, nil
}

// func Dec01b(ctx ch.AOContext) (interface{}, error) {
// 	return nil, errNotImplemented
// }
