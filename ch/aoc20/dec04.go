package aoc20

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec04a(ctx ch.AOContext) error {
	passports, err := readPassports(ctx, "inputs/2020/dec04.txt")
	if err != nil {
		return err
	}

	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	rv := 0
	for _, pp := range passports {
		valid := true
		for _, k := range requiredFields {
			if _, ok := pp[k]; !ok {
				valid = false
			}
		}
		if valid {
			rv++
		}
	}

	ctx.FinalAnswer.Print(rv)
	return nil
}

func Dec04b(ctx ch.AOContext) error {
	passports, err := readPassports(ctx, "inputs/2020/dec04.txt")
	if err != nil {
		return err
	}

	validators := map[string]func(string) bool{
		"byr": func(s string) bool {
			y, _ := strconv.Atoi(s)
			return y >= 1920 && y <= 2002
		},
		"iyr": func(s string) bool {
			y, _ := strconv.Atoi(s)
			return y >= 2010 && y <= 2020
		},
		"eyr": func(s string) bool {
			y, _ := strconv.Atoi(s)
			return y >= 2020 && y <= 2030
		},
		"hgt": func(s string) bool {
			if len(s) < 3 {
				return false
			}
			y, _ := strconv.Atoi(s[:len(s)-2])
			if s[len(s)-2:] == "cm" {
				return y >= 150 && y <= 193
			} else if s[len(s)-2:] == "in" {
				return y >= 59 && y <= 76
			} else {
				return false
			}
		},
		"hcl": func(s string) bool {
			var i int
			_, err := fmt.Sscanf(s+"\n", "#%06x\n", &i)
			return err == nil && len(s) == 7
		},
		"ecl": func(s string) bool {
			return s == "amb" || s == "blu" || s == "brn" || s == "gry" || s == "grn" || s == "hzl" || s == "oth"
		},
		"pid": func(s string) bool {
			var i int
			_, err := fmt.Sscanf(s, "%09d", &i)
			return err == nil && len(s) == 9
		},
	}

	rv := 0
	for _, pp := range passports {
		valid := true
		for k, f := range validators {
			if v, ok := pp[k]; !ok || !f(v) {
				valid = false
			}
		}
		if valid {
			rv++
		}
	}

	ctx.FinalAnswer.Print(rv)
	return nil
}

func readPassports(ctx ch.AOContext, assetName string) ([]map[string]string, error) {
	sections, err := ctx.DataSections(assetName)
	if err != nil {
		return nil, err
	}
	var rv []map[string]string

	for _, passport := range sections {
		pp := make(map[string]string)
		for _, l := range passport {
			for _, cmp := range strings.Split(l, " ") {
				pp[cmp[0:3]] = cmp[4:]
			}
		}
		rv = append(rv, pp)
	}

	return rv, nil
}
