package ch

import (
	"strconv"

	"github.com/thijzert/advent-of-code/2021/data"
)

func dataAsInts(args []string, assetName string) ([]int, error) {
	if len(args) == 0 {
		return data.GetInts(assetName)
	}

	var rv []int
	for _, str := range args {
		i, err := strconv.Atoi(str)
		if err != nil {
			return rv, err
		}
		rv = append(rv, i)
	}
	return rv, nil
}
