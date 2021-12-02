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

	return strings.Split(string(buf), "\n"), nil
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
