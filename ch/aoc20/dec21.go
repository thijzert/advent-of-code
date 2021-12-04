package aoc20

import (
	"errors"
	"strings"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec21a(ctx ch.AOContext) error {
	lines, err := ctx.DataLines("inputs/2020/dec21.txt")
	if err != nil {
		return err
	}

	recipes := readRecipeList(lines)

	ingredientOccurrence := make(map[string]int)
	ingredientMap := make(map[string]int)
	allergenMap := make(map[string]int)

	for _, r := range recipes {
		for _, ing := range r.Ingredients {
			ingredientMap[ing] = -1

			oc := ingredientOccurrence[ing]
			ingredientOccurrence[ing] = oc + 1
		}
		for _, alg := range r.Allergens {
			allergenMap[alg] = -1
		}
	}

	allIngredients := make([]string, 0, len(ingredientMap))
	allAllergens := make([]string, 0, len(allergenMap))

	for ing := range ingredientMap {
		ingredientMap[ing] = len(allIngredients)
		allIngredients = append(allIngredients, ing)
	}
	for alg := range allergenMap {
		allergenMap[alg] = len(allAllergens)
		allAllergens = append(allAllergens, alg)
	}
	ctx.Debug.Printf("Allergens: %d; ingredients: %d", len(allAllergens), len(allIngredients))

	transTable := make([][]bool, len(allAllergens))
	for i := range transTable {
		transTable[i] = make([]bool, len(allIngredients))
		for j := range transTable[i] {
			transTable[i][j] = true
		}
	}

	for _, r := range recipes {
		mask := make([]bool, len(allIngredients))
		for _, ing := range r.Ingredients {
			mask[ingredientMap[ing]] = true
		}

		for _, alg := range r.Allergens {
			// ctx.Debug.Printf("Anything that isn't %s cannot possibly contain %s", r.Ingredients, alg)
			i := allergenMap[alg]
			for j := range transTable[i] {
				transTable[i][j] = transTable[i][j] && mask[j]
			}
		}
	}

	var safeIngredients []string
	for j, ing := range allIngredients {
		safe := true
		for i := range allAllergens {
			safe = safe && !transTable[i][j]
		}
		if safe {
			safeIngredients = append(safeIngredients, ing)
		}
	}
	ctx.Debug.Printf("Safe ingredients: %s", safeIngredients)

	totalOcc := 0
	for _, ing := range safeIngredients {
		totalOcc += ingredientOccurrence[ing]
	}

	ctx.FinalAnswer.Print(totalOcc)
	return nil
}

func Dec21b(ctx ch.AOContext) error {
	return errors.New("not implemented")
}

type recipe struct {
	Ingredients []string
	Allergens   []string
}

func readRecipeList(lines []string) []recipe {
	var rv []recipe

	for _, l := range lines {
		if l == "" {
			continue
		}

		sp := strings.Split(l, " (contains ")
		if len(sp) != 2 {
			continue
		}

		r := recipe{
			Ingredients: strings.Split(sp[0], " "),
			Allergens:   strings.Split(sp[1][:len(sp[1])-1], ", "),
		}
		rv = append(rv, r)
	}

	return rv
}
