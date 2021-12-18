package aoc21

import (
	"fmt"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec18a(ctx ch.AOContext) error {
	lines := []string{
		"[1,2]",
		"[[1,2],3]",
		"[9,[8,7]]",
		"[[1,9],[8,5]]",
		"[[[[1,2],[3,4]],[[5,6],[7,8]]],9]",
		"[[[9,[3,8]],[[0,9],6]],[[[3,7],[4,9]],3]]",
		"[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]",
	}

	for _, line := range lines {
		sn, n, err := parseSnailfishNumber(line)
		if err != nil {
			return err
		}
		if n != len(line) || sn == nil {
			ctx.Printf("Error: only parsed %d characters instead of %d", n, len(line))
		}
		if sn.String() != line {
			ctx.Printf(" - in:  %s", line)
			ctx.Printf(" - out: %s", sn)
		}
	}

	tcc := []struct {
		Inp, Exp string
	}{
		{"[[[[[9,8],1],2],3],4]", "[[[[0,9],2],3],4]"},
		{"[7,[6,[5,[4,[3,2]]]]]", "[7,[6,[5,[7,0]]]]"},
		{"[[6,[5,[4,[3,2]]]],1]", "[[6,[5,[7,0]]],3]"},
		{"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[7,0]]]]"},
	}
	for _, tc := range tcc {
		sn, _, _ := parseSnailfishNumber(tc.Inp)
		sn.Reduce()
		if sn.String() != tc.Exp {
			ctx.Printf("%s becomes %s", tc.Inp, sn)
			ctx.Printf(" !! was expecting %s", tc.Exp)
		}
	}

	a, _, _ := parseSnailfishNumber("[[[[4,3],4],4],[7,[[8,4],9]]]")
	b, _, _ := parseSnailfishNumber("[1,1]")
	a.Add(b)
	ctx.Printf("a + b = %s", a)

	lines, err := ctx.DataLines("inputs/2021/dec18.txt")
	if err != nil {
		return err
	}

	var sn *snailfish
	for i, line := range lines {
		if i == 0 {
			sn, _, err = parseSnailfishNumber(line)
			if err != nil {
				return err
			}
			ctx.Printf("Step %d: %s", i+1, sn)
			continue
		} else if line == "" {
			continue
		}

		b, _, err = parseSnailfishNumber(line)
		if err != nil {
			return err
		}
		sn.Add(b)
		ctx.Printf("Step %d: %s", i+1, sn)
	}

	ctx.Printf("snailfish sum: %s", sn)
	ctx.FinalAnswer.Print(sn.Magnitude())
	return nil
}

func Dec18b(ctx ch.AOContext) error {
	lines, err := ctx.DataLines("inputs/2021/dec18.txt")
	if err != nil {
		return err
	}

	mmax := 0

	for _, left := range lines {
		for _, right := range lines {
			if left == right || left == "" || right == "" {
				continue
			}

			a, _, _ := parseSnailfishNumber(left)
			b, _, _ := parseSnailfishNumber(right)

			a.Add(b)
			m := a.Magnitude()
			if m > mmax {
				mmax = m
			}
		}
	}

	ctx.FinalAnswer.Print(mmax)
	return nil
}

type snailfish struct {
	N    int
	L, R *snailfish
}

func parseSnailfishNumber(str string) (*snailfish, int, error) {
	if str == "" {
		return nil, 0, fmt.Errorf("unexpected EOF")
	}
	rv := &snailfish{}
	if str[0] == '[' {
		n := 1
		s, i, err := parseSnailfishNumber(str[n:])
		if err != nil {
			return nil, 0, err
		}
		rv.L = s
		n += i
		if str[n:] == "" || str[n] != ',' {
			return nil, 0, fmt.Errorf("unexpected EOF - expecting ','")
		}
		n++

		s, i, err = parseSnailfishNumber(str[n:])
		if err != nil {
			return nil, 0, err
		}
		rv.R = s
		n += i
		if str[n:] == "" || str[n] != ']' {
			return nil, 0, fmt.Errorf("unexpected EOF - expecting ']'")
		}
		return rv, n + 1, nil
	} else {
		rv.N = int(str[0] - '0')
		return rv, 1, nil
	}
}

func (s *snailfish) String() string {
	if s.L == nil || s.R == nil {
		return fmt.Sprintf("%d", s.N)
	}
	return fmt.Sprintf("[%s,%s]", s.L, s.R)
}

func (s *snailfish) Reduce() {
	changed := true
	for changed {
		changed = false

		if _, _, hasexp := s.explode(0); hasexp {
			changed = true
		} else if s.split() {
			changed = true
		}
	}
}

func (s *snailfish) explode(depth int) (int, int, bool) {
	if s.L == nil || s.R == nil {
		return 0, 0, false
	}

	if depth >= 4 {
		a, b := s.L.N, s.R.N
		s.L = nil
		s.R = nil
		s.N = 0
		return a, b, true
	}

	a, b, hasexp := s.L.explode(depth + 1)
	if hasexp {
		b = s.R.addL(b)
		return a, b, true
	}

	a, b, hasexp = s.R.explode(depth + 1)
	if hasexp {
		a = s.L.addR(a)
		return a, b, true
	}

	return 0, 0, false
}

func (s *snailfish) addL(n int) int {
	if n == 0 {
		return 0
	}
	if s.L == nil {
		s.N += n
		return 0
	}
	return s.L.addL(n)
}
func (s *snailfish) addR(n int) int {
	if n == 0 {
		return 0
	}
	if s.R == nil {
		s.N += n
		return 0
	}
	return s.R.addR(n)
}

func (s *snailfish) split() bool {
	if s.L == nil && s.R == nil {
		if s.N >= 10 {
			s.L = &snailfish{N: s.N / 2}
			s.R = &snailfish{N: (s.N + 1) / 2}
			s.N = 0
			return true
		}
		return false
	}

	if s.L.split() {
		return true
	} else if s.R.split() {
		return true
	}
	return false
}

func (s *snailfish) Add(b *snailfish) {
	left := &snailfish{
		N: s.N,
		L: s.L,
		R: s.R,
	}
	s.N = 0
	s.L = left
	s.R = b
	s.Reduce()
}

func (s *snailfish) Magnitude() int {
	if s.L == nil && s.R == nil {
		return s.N
	}
	return 3*s.L.Magnitude() + 2*s.R.Magnitude()
}
