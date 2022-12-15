package aoc20

import (
	"errors"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec25a(ctx ch.AOContext) (interface{}, error) {
	pubs, err := ctx.DataAsInts("inputs/2020/dec25.txt")
	if err != nil {
		return nil, err
	}

	modulus := 20201227

	doorPub, cardPub := pubs[0], pubs[1]

	ctx.Debug.Print(doorPub)
	ctx.Debug.Print(cardPub)

	doorPriv, cardPriv := 0, 0
	v := 1
	for i := 1; i < modulus; i++ {
		v = (v * 7) % modulus
		if v == doorPub {
			doorPriv = i
			if cardPriv != 0 {
				break
			}
		} else if v == cardPub {
			cardPriv = i
			if doorPriv != 0 {
				break
			}
		}
	}

	ctx.Debug.Printf("Door loop size: %d; card: %d", doorPriv, cardPriv)

	a := transform(doorPub, cardPriv, modulus)
	b := transform(cardPub, doorPriv, modulus)
	ctx.Debug.Print(a, b)
	if a == b {
		return a, nil
	}

	return nil, errors.New("this failed")
}

func transform(subject, loopSize, modulus int) int {
	v := 1
	for i := 0; i < loopSize; i++ {
		v = (v * subject) % modulus
	}
	return v
}

func Dec25b(ctx ch.AOContext) (interface{}, error) {
	return "Sorry for the trouble", nil
}
