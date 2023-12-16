package aoc23

import (
	"strings"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec15a(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2023/dec15.txt")
	if err != nil {
		return nil, err
	}
	elems := strings.Split(lines[0], ",")

	answer := 0
	for _, elem := range elems {
		h := byte(0)
		for _, c := range elem {
			h = 17 * (h + byte(c))
		}
		answer += int(h)
	}

	return answer, nil
}

func Dec15b(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2023/dec15.txt")
	if err != nil {
		return nil, err
	}
	//lines = []string{"rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"}
	elems := strings.Split(lines[0], ",")

	type lens struct {
		Label string
		Focus int
	}

	slots := [256][]lens{}

	for _, elem := range elems {
		h := 0
		for i, c := range elem {
			if c == '-' {
				for j, l := range slots[h] {
					if l.Label == elem[:i] {
						slots[h] = append(slots[h][:j], slots[h][j+1:]...)
						break
					}
				}
				break
			} else if c == '=' {
				focus := atoid(elem[i+1:], 11)
				found := false
				for j, l := range slots[h] {
					if l.Label == elem[:i] {
						found = true
						slots[h][j].Focus = focus
						break
					}
				}
				if !found {
					slots[h] = append(slots[h], lens{
						Label: elem[:i],
						Focus: focus,
					})
				}
				break
			} else {
				h = (17 * (h + int(c))) & 0xff
			}
		}

		//ctx.Printf("After %s:", elem)
		//for h, slot := range slots {
		//	if len(slot) == 0 {
		//		continue
		//	}
		//	ctx.Printf("Slot %d: %v", h, slot)
		//}
	}

	answer := 0
	for i, sl := range slots {
		for j, l := range sl {
			answer += (i + 1) * (j + 1) * l.Focus
		}
	}

	return answer, nil
}
