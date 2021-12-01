package data

import (
	"strconv"
	"strings"
)

func GetInts(assetName string) ([]int, error) {
	buf, err := getAsset(assetName)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(buf), "\n")
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
