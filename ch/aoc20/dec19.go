package aoc20

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec19a(ctx ch.AOContext) error {
	rset, lines, err := readRuleSet(ctx, "inputs/2020/dec19.txt")
	if err != nil {
		return err
	}

	rv := 0
	for i, line := range lines {
		if line == "" {
			continue
		}
		v := rset.Match(line, 0)
		if v {
			rv++
		}
		ctx.Printf("%3d: %5v  %s", i+1, v, line)
	}

	ctx.FinalAnswer.Print(rv)
	return nil
}

func Dec19b(ctx ch.AOContext) error {
	rset, err := parseRuleSet([]string{
		"0: 1 2",
		"1: 2 2 | \"a\"",
		"2: 1 2 | \"b\"",
	})
	if err != nil {
		return err
	}
	ctx.Printf("Rules: (CNF: %v)", rset.IsChomskyNormalForm())
	for id, rule := range rset {
		ctx.Printf("   %3d : %v", id, rule)
	}

	for _, s := range []string{"aabbb", "abb", "bbb", "abba", "aabba", "abbbb"} {
		ctx.Printf("Matches '%s': %v", s, rset.CYKMatch(s, 0))
	}

	return fmt.Errorf("stopping here for now")

	rset, lines, err := readRuleSet(ctx, "inputs/2020/dec19b.txt")
	if err != nil {
		return err
	}

	rset[8] = []grammarRule{
		{grammarLexeme{RuleID: 42}, grammarLexeme{RuleID: 8}},
		{grammarLexeme{RuleID: 42}},
	}
	rset[11] = []grammarRule{
		{grammarLexeme{RuleID: 42}, grammarLexeme{RuleID: 11}, grammarLexeme{RuleID: 31}},
		{grammarLexeme{RuleID: 42}, grammarLexeme{RuleID: 31}},
	}

	rv := 0
	for i, line := range lines {
		if line == "" {
			continue
		}
		if rset.Match(line, 0) {
			ctx.Printf("%3d: match  %s", i+1, line)
			rv++
		}
	}

	ctx.FinalAnswer.Print(rv)
	return errors.New("not implemented")
}

type grammarLexeme struct {
	Literal string
	RuleID  int
}

type grammarRule []grammarLexeme

type grammarRuleSet map[int][]grammarRule

func readRuleSet(ctx ch.AOContext, assetName string) (grammarRuleSet, []string, error) {
	lines, err := ctx.DataLines(assetName)
	if err != nil {
		return nil, nil, err
	}

	rules := lines
	var rest []string = nil

	for i, l := range lines {
		if l == "" {
			rules = lines[:i]
			rest = lines[i+1:]
			break
		}
	}

	rv, err := parseRuleSet(rules)
	return rv, rest, err
}

func parseRuleSet(rules []string) (grammarRuleSet, error) {
	rv := make(grammarRuleSet)

	for i, l := range rules {
		if l == "" {
			continue
		}

		idparts := strings.SplitN(l, ": ", 2)
		if len(idparts) != 2 {
			return nil, fmt.Errorf("error on line %d: '%s'", i+1, l)
		}
		var id int
		var err error
		if id, err = strconv.Atoi(idparts[0]); err != nil {
			return nil, err
		}

		options := strings.Split(idparts[1], " | ")
		for _, opt := range options {
			parts := strings.Split(opt, " ")
			var rule grammarRule
			for _, p := range parts {
				if p[0] == '"' && p[len(p)-1] == '"' {
					rule = append(rule, grammarLexeme{Literal: p[1 : len(p)-1]})
				} else if n, err := strconv.Atoi(p); err == nil {
					rule = append(rule, grammarLexeme{RuleID: n})
				}
			}
			rv[id] = append(rv[id], rule)
		}
	}

	return rv, nil
}

func (rset grammarRuleSet) IsChomskyNormalForm() bool {
	for _, options := range rset {
		for _, opt := range options {
			if len(opt) == 0 {
				return false
			} else if len(opt) == 1 {
				if len(opt[0].Literal) != 1 {
					return false
				}
			} else if len(opt) == 2 {
				if opt[0].Literal != "" || opt[1].Literal != "" {
					return false
				}
				if len(rset[opt[0].RuleID]) == 0 || len(rset[opt[1].RuleID]) == 0 {
					return false
				}
			} else {
				return false
			}
		}
	}
	return true
}

func (rset grammarRuleSet) Match(str string, initial int) bool {
	n, ok := rset.matchRule(str, 0, initial, 0)
	return ok && (n == len(str))
}

func (rset grammarRuleSet) matchRule(str string, offset int, ruleID int, depth int) (int, bool) {
	if depth > len(str) {
		return 0, false
	}

	options := rset[ruleID]
	for _, opt := range options {
		suboff := offset
		matches := true
		for _, lex := range opt {
			if lex.Literal != "" {
				if suboff+len(lex.Literal) > len(str) || str[suboff:suboff+len(lex.Literal)] != lex.Literal {
					matches = false
					break
				}
				suboff += len(lex.Literal)
			} else {
				n, ok := rset.matchRule(str, suboff, lex.RuleID, depth+1)
				if !ok {
					matches = false
					break
				}
				suboff += n
			}
		}
		if matches {
			return suboff - offset, true
		}
	}

	return 0, false
}

func (rset grammarRuleSet) CYKMatch(str string, initial int) bool {
	V := make([][]map[int]bool, len(str))

	for l := range str {
		V[l] = make([]map[int]bool, len(str)-l)
		if l == 0 {
			for i, c := range str {
				V[l][i] = make(map[int]bool)
				for id, opts := range rset {
					for _, option := range opts {
						if len(option) == 1 && len(option[0].Literal) == 1 && rune(option[0].Literal[0]) == c {
							V[l][i][id] = true
						}
					}
				}
			}
		} else {
			for i := range str[l:] {
				V[l][i] = make(map[int]bool)
			}
			for d := 0; d < l; d++ {
				for j := range V[l] {
					m := V[d][j]
					mm := V[l-d-1][j+d+1]
					for left := range m {
						for right := range mm {
							for id, opts := range rset {
								for _, option := range opts {
									if len(option) == 2 && option[0].Literal == "" && option[1].Literal == "" && option[0].RuleID == left && option[1].RuleID == right {
										V[l][j][id] = true
									}
								}
							}
						}
					}
				}
			}
		}

		// fmt.Printf("  [")
		// for _, m := range V[l] {
		// 	fmt.Printf(" {")
		// 	sep := ""
		// 	for id := range m {
		// 		fmt.Printf("%s%d", sep, id)
		// 		sep = " "
		// 	}
		// 	fmt.Printf("}")
		// }
		// fmt.Printf(" ]\n")
	}

	return V[len(str)-1][0][initial]
}
