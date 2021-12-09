package aoc20

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec19a(ctx ch.AOContext) error {
	rset, lines, err := parseRuleSet(ctx, "inputs/2020/dec19.txt")
	if err != nil {
		return err
	}

	ctx.Printf("Rules:")
	for id, rule := range rset {
		ctx.Printf("   %3d : %v", id, rule)
	}
	ctx.Printf("Strings: %s", lines)

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
	return errors.New("not implemented")
}

type grammarLexeme struct {
	Literal string
	RuleID  int
}

type grammarRule []grammarLexeme

type grammarRuleSet map[int][]grammarRule

func parseRuleSet(ctx ch.AOContext, assetName string) (grammarRuleSet, []string, error) {
	lines, err := ctx.DataLines(assetName)
	if err != nil {
		return nil, nil, err
	}

	rv := make(grammarRuleSet)

	for i, l := range lines {
		if l == "" {
			return rv, lines[i+1:], nil
		}

		idparts := strings.SplitN(l, ": ", 2)
		if len(idparts) != 2 {
			return nil, nil, fmt.Errorf("error on line %d: '%s'", i+1, l)
		}
		var id int
		if id, err = strconv.Atoi(idparts[0]); err != nil {
			return nil, nil, err
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

	return rv, nil, nil
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
