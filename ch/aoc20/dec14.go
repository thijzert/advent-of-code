package aoc20

import (
	"fmt"

	"github.com/thijzert/advent-of-code/ch"
)

const U36BIT uint64 = 0xfffffffff

func Dec14a(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2020/dec14.txt")
	if err != nil {
		return nil, err
	}

	var maskOn, maskOff uint64
	var mem []uint64

	for _, l := range lines {
		if len(l) > 7 && l[0:7] == "mask = " {
			maskOn, maskOff = 0, 0
			for _, c := range l[7:] {
				maskOn, maskOff = maskOn<<1, maskOff<<1
				if c == '1' {
					maskOn = maskOn | 1
				} else if c == '0' {
					maskOff = maskOff | 1
				}
			}

			//ctx.Printf("Mask value: % 36b (%d)", maskOn, maskOn)
			//ctx.Printf("            % 36b (%d)", maskOff, maskOff)
			maskOff = U36BIT ^ maskOff
		} else if len(l) > 4 && l[:4] == "mem[" {
			var i int
			var v uint64
			fmt.Sscanf(l, "mem[%d] = %d", &i, &v)
			for len(mem) <= i {
				mem = append(mem, 0)
			}
			mem[i] = (v & maskOff) | maskOn
		}
	}

	rv := uint64(0)
	for _, v := range mem {
		rv += v
	}
	return rv, nil
}

func Dec14b(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2020/dec14.txt")
	if err != nil {
		return nil, err
	}

	var maskOn, maskOff uint64
	var maskFloat uint64
	var enumerated []uint64

	mem := make(map[uint64]uint64)

	for _, l := range lines {
		if len(l) > 7 && l[0:7] == "mask = " {
			maskOn, maskOff = 0, 0
			for _, c := range l[7:] {
				maskOn, maskOff = maskOn<<1, maskOff<<1
				if c == '1' {
					maskOn = maskOn | 1
				} else if c == '0' {
					maskOff = maskOff | 1
				}
			}

			maskOff = maskOff & U36BIT
			maskOn = maskOn & U36BIT

			//ctx.Printf("Mask value: % 36b (%d)", maskOn, maskOn)
			maskOff = U36BIT ^ maskOff
			//ctx.Printf("            % 36b (%d)", maskOff, maskOff)

			maskFloat, enumerated = floatyBits(maskOn, maskOff)

			//ctx.Printf("Floaty bits:% 36b (%d)", U36BIT^maskFloat, U36BIT^maskFloat)
		} else if len(l) > 4 && l[:4] == "mem[" {
			var i, v uint64
			fmt.Sscanf(l, "mem[%d] = %d", &i, &v)

			i = (i | maskOn) & maskFloat
			for _, m := range enumerated {
				mem[i|m] = v
			}
		}
	}

	rv := uint64(0)
	for _, v := range mem {
		rv += v
	}
	return rv, nil
}

func floatyBits(maskOn, maskOff uint64) (uint64, []uint64) {
	maskFl := (maskOn | (U36BIT ^ maskOff))
	return maskFl, enumerateFloatyBits(U36BIT ^ maskFl)
}

func enumerateFloatyBits(floaty uint64) []uint64 {
	if floaty == 0 {
		return []uint64{0}
	}
	downstream := enumerateFloatyBits(floaty >> 1)
	if floaty%2 == 0 {
		for i := range downstream {
			downstream[i] <<= 1
		}
		return downstream
	} else {
		rv := make([]uint64, len(downstream)*2)
		for i := range rv {
			rv[i] = downstream[i/2]<<1 | uint64(i&1)
		}
		return rv
	}
}
