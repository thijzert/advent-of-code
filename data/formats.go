package data

import (
	"strconv"
	"strings"
)

func GetLines(assetName string) ([]string, error) {
	buf, err := getAsset(assetName)
	if err != nil {
		return nil, err
	}

	rv := strings.Split(string(buf), "\n")

	// Remove trailing newline
	if len(rv) > 0 && rv[len(rv)-1] == "" {
		rv = rv[:len(rv)-1]
	}

	return rv, nil
}

func Transpose(lines []string) []string {
	l := 0
	for _, line := range lines {
		l = max(l, len(line))
	}
	rv := make([]string, l)
	for i := 0; i < l; i++ {
		buf := make([]byte, len(lines))
		for j, line := range lines {
			if len(line) > i {
				buf[j] = line[i]
			}
		}
		rv[i] = string(buf)
	}
	return rv
}

func GetInts(assetName string) ([]int, error) {
	lines, err := GetLines(assetName)
	if err != nil {
		return nil, err
	}

	rv := make([]int, 0, len(lines))
	for _, l := range lines {
		if l == "" {
			continue
		}
		if i, err := strconv.Atoi(l); err == nil {
			rv = append(rv, i)
		}
	}

	return rv, nil
}

func CSVInts(lines []string) [][]int {
	var rv [][]int
	for _, line := range lines {
		var iln []int
		for _, s := range strings.Split(line, ",") {
			n, _ := strconv.Atoi(s)
			iln = append(iln, n)
		}
		rv = append(rv, iln)
	}

	return rv
}
