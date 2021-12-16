package aoc21

import (
	"fmt"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec16a(ctx ch.AOContext) error {
	lines, err := ctx.DataLines("inputs/2021/dec16.txt")
	if err != nil {
		return err
	}

	finalErr := fmt.Errorf("empty input")
	for _, line := range lines {
		if line == "" {
			continue
		}
		packets, err := parsePacketsHex(line)
		for _, p := range packets {
			ctx.Printf("Packet type %d: %d", p.Type, p)
		}
		if err != nil {
			return err
		}

		rv := 0
		for _, p := range packets {
			rv += p.RVersion()
		}
		if finalErr != nil {
			ctx.FinalAnswer.Print(rv)
			finalErr = nil
		} else {
			ctx.Printf("Subsequent answer: %d", rv)
		}
	}

	return finalErr
}

func Dec16b(ctx ch.AOContext) error {
	lines, err := ctx.DataLines("inputs/2021/dec16.txt")
	if err != nil {
		return err
	}

	finalErr := fmt.Errorf("empty input")
	for _, line := range lines {
		if line == "" {
			continue
		}
		packets, err := parsePacketsHex(line)
		if err != nil {
			return err
		}

		for _, p := range packets {
			rv, err := p.Evaluate()
			if err != nil {
				return err
			}
			if finalErr != nil {
				ctx.FinalAnswer.Print(rv)
				finalErr = nil
			} else {
				ctx.Printf("Subsequent answer: %d", rv)
			}
		}
	}

	return finalErr
}

const (
	type_LITERAL uint8 = 4
	type_SUM     uint8 = 0
	type_PRODUCT uint8 = 1
	type_MINIMUM uint8 = 2
	type_MAXIMUM uint8 = 3
	type_GREATER uint8 = 5
	type_LESS    uint8 = 6
	type_EQUAL   uint8 = 7
)

type radioPacket struct {
	Version    uint8
	Type       uint8
	Operand    uint64
	Subpackets []radioPacket
}

func (p radioPacket) RVersion() int {
	rv := int(p.Version)
	for _, q := range p.Subpackets {
		rv += q.RVersion()
	}
	return rv
}

func (p radioPacket) Evaluate() (int, error) {
	var subv []int
	for _, q := range p.Subpackets {
		n, err := q.Evaluate()
		if err != nil {
			return 0, err
		}
		subv = append(subv, n)
	}

	if p.Type == type_GREATER || p.Type == type_LESS || p.Type == type_EQUAL {
		if len(subv) != 2 {
			return 0, fmt.Errorf("invalid number of operands for type %3b", p.Type)
		}
	}

	if p.Type == type_LITERAL {
		return int(p.Operand), nil
	} else if p.Type == type_SUM {
		return sum(subv...), nil
	} else if p.Type == type_PRODUCT {
		rv := 1
		for _, n := range subv {
			rv *= n
		}
		return rv, nil
	} else if p.Type == type_MINIMUM {
		return min(subv...), nil
	} else if p.Type == type_MAXIMUM {
		return max(subv...), nil
	} else if p.Type == type_GREATER {
		if subv[0] > subv[1] {
			return 1, nil
		}
		return 0, nil
	} else if p.Type == type_LESS {
		if subv[0] < subv[1] {
			return 1, nil
		}
		return 0, nil
	} else if p.Type == type_EQUAL {
		if subv[0] == subv[1] {
			return 1, nil
		}
		return 0, nil
	} else {
		return 0, fmt.Errorf("unknown operator %3b (%d)", p.Type, p.Type)
	}
}

func parsePacketsHex(hex string) ([]radioPacket, error) {
	binary := ""
	for i := range hex {
		var j uint8
		fmt.Sscanf(hex[i:i+1], "%1x", &j)
		binary += fmt.Sprintf("%04b", j)
	}
	return parsePacketsBin(binary, 8)
}

func parsePacketsBin(binary string, round int) ([]radioPacket, error) {
	var rv []radioPacket
	i := 0
	for i < len(binary) {
		p, n, err := parsePacket(binary[i:])
		if err != nil {
			return rv, err
		}
		rv = append(rv, p)
		i += n
		if n%round != 0 {
			i += round - (n % round)
		}
	}

	return rv, nil
}

func parsePacket(binary string) (radioPacket, int, error) {
	p := radioPacket{}
	fmt.Sscanf(binary[0:3], "%b", &p.Version)
	fmt.Sscanf(binary[3:6], "%b", &p.Type)

	if p.Type == type_LITERAL {
		n := 6
		for n < len(binary) {
			var v uint64
			fmt.Sscanf(binary[n+1:n+5], "%b", &v)
			p.Operand = p.Operand<<4 | v
			n += 5
			if binary[n-5] == '0' {
				break
			}
		}
		return p, n, nil
	} else {
		if binary[6] == '0' {
			var length int
			fmt.Sscanf(binary[7:22], "%b", &length)
			var err error
			p.Subpackets, err = parsePacketsBin(binary[22:22+length], 1)
			return p, length + 22, err
		} else {
			var nPackets int
			fmt.Sscanf(binary[7:18], "%b", &nPackets)
			n := 18
			for i := 0; i < nPackets; i++ {
				q, j, err := parsePacket(binary[n:])
				if err != nil {
					return p, n, err
				}
				p.Subpackets = append(p.Subpackets, q)
				n += j
			}
			return p, n, nil
		}
	}
}
