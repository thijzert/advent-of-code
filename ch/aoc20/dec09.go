package aoc20

import "github.com/thijzert/advent-of-code/ch"

func Dec09a(ctx ch.AOContext) (interface{}, error) {
	inputs := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}
	inv, _ := decodeXMASprotocol(inputs, 5)
	ctx.Printf("Example data: first invalid is %d", inv)

	inputs, err := ctx.DataAsInts("inputs/2020/dec09.txt")
	if err != nil {
		return nil, err
	}
	inv, _ = decodeXMASprotocol(inputs, 25)
	return inv, nil
}

func Dec09b(ctx ch.AOContext) (interface{}, error) {
	inputs := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}
	_, weak := decodeXMASprotocol(inputs, 5)
	ctx.Printf("Example data: cryptographic weakness: %d", weak)

	inputs, err := ctx.DataAsInts("inputs/2020/dec09.txt")
	if err != nil {
		return nil, err
	}
	_, weak = decodeXMASprotocol(inputs, 25)
	return weak, nil
}

func decodeXMASprotocol(inputs []int, preambleLength int) (firstInvalid, cryptographicWeakness int) {
	previous := make([]int, preambleLength)
	for i, v := range inputs {
		if i < preambleLength {
			previous[i] = v
			continue
		}

		valid := false
		for _, a := range previous {
			for _, b := range previous {
				if a != b && a+b == v {
					valid = true
				}
			}
		}
		if valid {
			previous[i%preambleLength] = v
		} else {
			previous[i%preambleLength] = v
			if firstInvalid == 0 {
				firstInvalid = v
			}
		}
	}

	for i := range inputs {
		sum := 0
		for j, v := range inputs[i:] {
			sum += v
			if sum == firstInvalid {
				cryptographicWeakness = min(inputs[i:i+j+1]...) + max(inputs[i:i+j+1]...)
				return
			} else if sum > firstInvalid {
				break
			}
		}
	}

	return
}
