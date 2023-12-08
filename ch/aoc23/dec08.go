package aoc23

import (
	"github.com/thijzert/advent-of-code/ch"
)

func Dec08a(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2023/dec08.txt")
	if err != nil {
		return nil, err
	}

	pattern := lines[0]
	nodes := make(map[string][2]string)
	for _, line := range lines[2:] {
		nodes[line[:3]] = [2]string{line[7:10], line[12:15]}
	}

	i := 0
	n := "AAA"
	for n != "ZZZ" {
		c := pattern[i%len(pattern)]
		i++
		if c == 'L' {
			n = nodes[n][0]
		} else {
			n = nodes[n][1]
		}
	}

	return i, nil
}

func Dec08b(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2023/dec08.txt")
	if err != nil {
		return nil, err
	}

	pattern := lines[0]
	nodes := make(map[string][2]string)
	for _, line := range lines[2:] {
		nodes[line[:3]] = [2]string{line[7:10], line[12:15]}
	}

	answer := 1

	for k := range nodes {
		if k[2] == 'A' {
			n := k
			i, offset, modulus := 0, 0, 0
			first := ""
			for n != first {
				if n[2] == 'Z' {
					if first == "" {
						first = n
						offset = i
					}
				}

				c := pattern[i%len(pattern)]
				i++
				if c == 'L' {
					n = nodes[n][0]
				} else {
					n = nodes[n][1]
				}
			}
			modulus = i - offset
			ctx.Printf("Ghost %s finishes in %d + k·%d steps", k, offset, modulus)
			if offset%modulus != 0 {
				return nil, errFailed
			}

			answer = lcm(answer, modulus)
		}
	}

	return answer, nil
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
