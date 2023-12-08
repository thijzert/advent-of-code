package aoc19

import (
	"fmt"

	"github.com/thijzert/advent-of-code/ch"
	"github.com/thijzert/advent-of-code/lib/image"
)

func Dec08a(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2019/dec08.txt")
	if err != nil {
		return nil, err
	}
	imageData := lines[0]

	w, h := 25, 6
	layers := len(imageData) / (w * h)
	minZeros, onesTwos := len(imageData), 0
	for layer := 0; layer < layers; layer++ {
		var counts [3]int
		for _, c := range imageData[layer*w*h : (layer+1)*w*h] {
			if c >= '0' && int(c-'0') < len(counts) {
				counts[int(c-'0')]++
			}
		}
		ctx.Printf("Layer %d has %d zeros", layer+1, counts[0])
		if counts[0] < minZeros {
			minZeros = counts[0]
			onesTwos = counts[1] * counts[2]
		}
	}

	return onesTwos, nil
}

func Dec08b(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2019/dec08.txt")
	if err != nil {
		return nil, err
	}
	imageData := lines[0]
	w, h := 25, 6

	img := image.NewImage(w, h, func(x, y int) int {
		for offset := y*w + x; offset < len(imageData); offset += w * h {
			if imageData[offset] == '0' || imageData[offset] == '1' {
				return int(imageData[offset] - '0')
			}
		}
		return 5
	})
	ctx.Printf("Image data:\n%s", img)

	rv, err := img.OCR()
	if err != nil {
		err = fmt.Errorf("error finding text: %w", err)
	}
	return rv, err
}
