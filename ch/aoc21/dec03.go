package aoc21

import (
	"fmt"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec03a(ctx ch.AOContext) error {
	lines, err := ctx.DataLines("inputs/2021/dec03.txt")
	if err != nil {
		return err
	}

	histo, _ := bitHistogram(lines, "")

	gamma := 0
	epsilon := 0
	for _, h := range histo {
		gamma <<= 1
		epsilon <<= 1
		if h[0] > h[1] {
			epsilon += 1
		} else {
			gamma += 1
		}
	}

	ctx.Debug.Printf("Gamma:   %3d (0b%b)", gamma, gamma)
	ctx.Debug.Printf("Epsilon: %3d (0b%b)", epsilon, epsilon)

	ctx.FinalAnswer.Print(gamma * epsilon)
	return nil
}

func bitHistogram(lines []string, prefix string) ([][2]int, int) {
	histo := make([][2]int, len(lines[0]))
	count := 0

	for _, l := range lines {
		if l == "" {
			continue
		}
		if len(l) < len(prefix) || l[:len(prefix)] != prefix {
			continue
		}

		for i, c := range l {
			if c == '0' {
				histo[i][0]++
			} else if c == '1' {
				histo[i][1]++
			}
		}
		count++
	}
	return histo, count
}

func Dec03b(ctx ch.AOContext) error {
	lines, err := ctx.DataLines("inputs/2021/dec03.txt")
	if err != nil {
		return err
	}

	count := 0
	o2 := ""
	co2 := ""

	for i := 0; i < len(lines[0]); i++ {
		histo, _ := bitHistogram(lines, o2)
		if histo[i][0] > histo[i][1] {
			o2 += "0"
		} else {
			o2 += "1"
		}

		histo, count = bitHistogram(lines, co2)
		if count < 2 {
			if histo[i][0] > histo[i][1] {
				co2 += "0"
			} else {
				co2 += "1"
			}
		} else {
			if histo[i][1] >= histo[i][0] {
				co2 += "0"
			} else {
				co2 += "1"
			}
		}
	}

	var a, b int
	_, err = fmt.Sscanf(o2+" "+co2, "%b %b", &a, &b)
	if err != nil {
		return err
	}

	ctx.Debug.Printf("O₂ generator: %s (%d)", o2, a)
	ctx.Debug.Printf("CO₂ scrubber: %s (%d)", co2, b)

	ctx.FinalAnswer.Print(a * b)
	return nil
}
