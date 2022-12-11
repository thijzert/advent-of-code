package aoc22

import (
	"fmt"
	"math/big"
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/thijzert/advent-of-code/ch"
)

type worryMonkey struct {
	Items        []*big.Int
	Operation    rune
	OperandOld   bool
	Operand      *big.Int
	Divisibility *big.Int
	NextMonkey   [2]int
}

func (m worryMonkey) ApplyOperation(v *big.Int) *big.Int {
	b := m.Operand
	if m.OperandOld {
		b = v
	}

	if m.Operation == '+' {
		return v.Add(v, b)
	} else if m.Operation == '*' {
		return v.Mul(v, b)
	}
	panic("invalid operation")
	return big.NewInt(0)
}

func readWorryMonkeys(ctx ch.AOContext, filename string) ([]worryMonkey, error) {
	sections, err := ctx.DataSections(filename)
	if err != nil {
		return nil, err
	}

	var rv []worryMonkey
	for i, sect := range sections {
		if len(sect) != 6 {
			return nil, fmt.Errorf("invalid format: monkey section expects 6 lines")
		}
		if sect[0] != fmt.Sprintf("Monkey %d:", i) {
			return nil, fmt.Errorf("invalid format: invalid monkey header '%s'", sect[0])
		}

		var monkey worryMonkey
		if len(sect[1]) < 18 || sect[1][:18] != "  Starting items: " {
			return nil, fmt.Errorf("invalid format: starting items '%s'", sect[1])
		}
		items := strings.Split(sect[1][18:], ", ")
		for _, item := range items {
			j, err := strconv.Atoi(item)
			if err != nil {
				return nil, fmt.Errorf("invalid format: starting item '%s'", item)
			}
			monkey.Items = append(monkey.Items, big.NewInt(int64(j)))
		}

		if len(sect[2]) < 26 || sect[2][:23] != "  Operation: new = old " {
			return nil, fmt.Errorf("invalid format: operation '%s'", sect[2])
		}
		monkey.Operation = rune(sect[2][23])
		if monkey.Operation != '+' && monkey.Operation != '*' {
			return nil, fmt.Errorf("invalid format: unknown operation %c '%s'", monkey.Operation, sect[2])
		}
		if sect[2][25:] == "old" {
			monkey.OperandOld = true
		} else {
			j, err := strconv.Atoi(sect[2][25:])
			if err != nil {
				return nil, fmt.Errorf("invalid format: operand '%s'", sect[2][25:])
			}
			monkey.Operand = big.NewInt(int64(j))
		}

		monkey.Divisibility = big.NewInt(1)
		_, err := fmt.Sscanf(sect[3], "  Test: divisible by %d", monkey.Divisibility)
		if err != nil {
			return nil, errors.Wrapf(err, "invalid format: divisibility '%s'", sect[3])
		}

		_, err = fmt.Sscanf(sect[4], "    If true: throw to monkey %d", &monkey.NextMonkey[0])
		if err != nil {
			return nil, errors.Wrapf(err, "invalid format: true monkey '%s'", sect[4])
		}
		_, err = fmt.Sscanf(sect[5], "    If false: throw to monkey %d", &monkey.NextMonkey[1])
		if err != nil {
			return nil, errors.Wrapf(err, "invalid format: false monkey '%s'", sect[5])
		}

		rv = append(rv, monkey)
	}

	return rv, nil
}

func Dec11a(ctx ch.AOContext) error {
	monkeys, err := readWorryMonkeys(ctx, "inputs/2022/dec11.txt")
	if err != nil {
		return err
	}

	three := big.NewInt(3)

	inspectedItem := make([]int, len(monkeys))
	for round := 0; round < 20; round++ {
		for i, monkey := range monkeys {
			inspectedItem[i] += len(monkey.Items)
			for _, item := range monkey.Items {
				item = monkey.ApplyOperation(item)
				item.Div(item, three)
				j := monkey.NextMonkey[1]
				rem := big.NewInt(0).Mod(item, monkey.Divisibility)
				if rem.Int64() == 0 {
					j = monkey.NextMonkey[0]
				}
				monkeys[j].Items = append(monkeys[j].Items, item)
			}
			monkeys[i].Items = monkeys[i].Items[:0]
		}
	}

	ctx.Print(inspectedItem)

	monkeyBusiness := 1
	sort.Ints(inspectedItem)
	for _, mb := range inspectedItem[len(inspectedItem)-2:] {
		monkeyBusiness *= mb
	}

	ctx.FinalAnswer.Print(monkeyBusiness)
	return nil
}

func Dec11b(ctx ch.AOContext) error {
	monkeys, err := readWorryMonkeys(ctx, "inputs/2022/dec11.txt")
	if err != nil {
		return err
	}

	lcm := big.NewInt(1)
	for _, monkey := range monkeys {
		lcm.Mul(lcm, monkey.Divisibility)
	}

	inspectedItem := make([]int, len(monkeys))
	for round := 0; round < 10000; round++ {
		for i, monkey := range monkeys {
			inspectedItem[i] += len(monkey.Items)
			for _, item := range monkey.Items {
				item = monkey.ApplyOperation(item)
				item.Mod(item, lcm)
				j := monkey.NextMonkey[1]
				rem := big.NewInt(0).Mod(item, monkey.Divisibility)
				if rem.Int64() == 0 {
					j = monkey.NextMonkey[0]
				}
				monkeys[j].Items = append(monkeys[j].Items, item)
			}
			monkeys[i].Items = monkeys[i].Items[:0]
		}
	}

	ctx.Print(inspectedItem)

	monkeyBusiness := 1
	sort.Ints(inspectedItem)
	for _, mb := range inspectedItem[len(inspectedItem)-2:] {
		monkeyBusiness *= mb
	}

	ctx.FinalAnswer.Print(monkeyBusiness)
	return errNotImplemented
}
