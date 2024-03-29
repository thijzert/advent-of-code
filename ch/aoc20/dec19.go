package aoc20

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec19a(ctx ch.AOContext) (interface{}, error) {
	rset, lines, err := readRuleSet(ctx, "inputs/2020/dec19.txt")
	if err != nil {
		return nil, err
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

	return rv, nil
}

func Dec19b(ctx ch.AOContext) (interface{}, error) {
	rset, lines, err := readRuleSet(ctx, "inputs/2020/dec19.txt")
	if err != nil {
		return nil, err
	}
	ctx.Printf("%d rules, %d strings to match", len(rset.Rules), len(lines))

	rset.Rules[8] = []grammarRule{
		{grammarLexeme{RuleID: 42}, grammarLexeme{RuleID: 8}},
		{grammarLexeme{RuleID: 42}},
	}
	rset.Rules[11] = []grammarRule{
		{grammarLexeme{RuleID: 42}, grammarLexeme{RuleID: 11}, grammarLexeme{RuleID: 31}},
		{grammarLexeme{RuleID: 42}, grammarLexeme{RuleID: 31}},
	}
	rset.ChomskyNormalForm()
	ctx.Print(rset)

	rv := 0
	for i, line := range lines {
		if line == "" {
			continue
		}
		if rset.CYKMatch(line, 0) {
			ctx.Printf("%3d: match  %s", i+1, line)
			rv++
		}
	}

	return rv, nil
}

type grammarLexeme struct {
	Literal string
	RuleID  int
}

func (l grammarLexeme) String() string {
	if l.Literal != "" {
		return "\"" + l.Literal + "\""
	}
	return strconv.Itoa(l.RuleID)
}

type grammarRule []grammarLexeme

func (g grammarRule) String() string {
	rv := ""
	sep := ""
	for _, l := range g {
		rv += sep + l.String()
		sep = " "
	}
	return rv
}

type invRule struct {
	Left, Right int
}

type grammarRuleSet struct {
	Rules   map[int][]grammarRule
	inverse map[invRule][]int
}

func newGrammarRuleSet() *grammarRuleSet {
	return &grammarRuleSet{
		Rules: make(map[int][]grammarRule),
	}
}

func readRuleSet(ctx ch.AOContext, assetName string) (*grammarRuleSet, []string, error) {
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

func parseRuleSet(rules []string) (*grammarRuleSet, error) {
	rv := newGrammarRuleSet()

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
			rv.Rules[id] = append(rv.Rules[id], rule)
		}
	}

	rv.recreateInverses()

	return rv, nil
}

func (rset *grammarRuleSet) recreateInverses() int {
	rset.inverse = make(map[invRule][]int)
	for id, opts := range rset.Rules {
		for _, opt := range opts {
			if len(opt) == 2 && opt[0].Literal == "" && opt[1].Literal == "" {
				j := invRule{opt[0].RuleID, opt[1].RuleID}
				rset.inverse[j] = append(rset.inverse[j], id)
			}
		}
	}
	rv := 0
	for _, ids := range rset.inverse {
		if len(ids) > 1 {
			rv++
		}
	}
	return rv
}

func (rset *grammarRuleSet) String() string {
	rv := fmt.Sprintf("Rules: (CNF: %v)", rset.IsChomskyNormalForm())

	for id, rules := range rset.Rules {
		rv += fmt.Sprintf("\n   %3d", id)
		sep := " : "
		for _, rule := range rules {
			rv += sep + rule.String()
			sep = " | "
		}
	}
	return rv
}

func (rset *grammarRuleSet) IsChomskyNormalForm() bool {
	for _, options := range rset.Rules {
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
				if len(rset.Rules[opt[0].RuleID]) == 0 || len(rset.Rules[opt[1].RuleID]) == 0 {
					return false
				}
			} else {
				return false
			}
		}
	}
	return true
}

func (rset *grammarRuleSet) ChomskyNormalForm() error {
	nextID := 0
	for id := range rset.Rules {
		if id >= nextID {
			nextID = id + 1
		}
	}
	nextID += 1000

	changed := true
	for changed {
		changed = false

		for id, options := range rset.Rules {
			for i, opt := range options {
				if len(opt) == 0 {
					return fmt.Errorf("calculating the CNF of non-λ-free languages is not implemented")
				} else if len(opt) == 1 && opt[0].Literal == "" {
					// Pass-through all productions of the other rule
					for j, oopt := range rset.Rules[opt[0].RuleID] {
						if j == 0 {
							rset.Rules[id][i] = oopt
						} else {
							rset.Rules[id] = append(rset.Rules[id], oopt)
						}
					}
					changed = true
				} else if len(opt) > 2 {
					// Split this option into the first lexeme, and a new symbol with 2-n
					newProduction := make(grammarRule, len(opt)-1)
					copy(newProduction, opt[1:])
					rset.Rules[nextID] = []grammarRule{newProduction}
					rset.Rules[id][i] = opt[0:2]
					rset.Rules[id][i][1] = grammarLexeme{RuleID: nextID}
					nextID++
					changed = true
				} else if len(opt) == 2 {
					// Check if both lexemes aren't literals
					for j, l := range opt {
						if l.Literal != "" {
							rset.Rules[nextID] = []grammarRule{{l}}
							rset.Rules[id][i][j] = grammarLexeme{RuleID: nextID}
							nextID++
							changed = true
						}
					}
				}
			}
		}
	}

	rset.recreateInverses()

	return nil
}

func (rset grammarRuleSet) Match(str string, initial int) bool {
	n, ok := rset.matchRule(str, 0, initial, 0)
	return ok && (n == len(str))
}

func (rset grammarRuleSet) matchRule(str string, offset int, ruleID int, depth int) (int, bool) {
	if depth > len(str) {
		return 0, false
	}

	options := rset.Rules[ruleID]
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
				for id, opts := range rset.Rules {
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
							invI := invRule{left, right}
							for _, id := range rset.inverse[invI] {
								V[l][j][id] = true
							}
						}
					}
				}
			}
		}
	}

	return V[len(str)-1][0][initial]
}
