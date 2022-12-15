package aoc20

import (
	"fmt"
	"sort"
	"strings"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec21a(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2020/dec21.txt")
	if err != nil {
		return nil, err
	}

	recipes := readRecipeList(lines)

	ingredientOccurrence := make(map[string]int)
	for _, r := range recipes {
		for _, ing := range r.Ingredients {
			oc := ingredientOccurrence[ing]
			ingredientOccurrence[ing] = oc + 1
		}
	}

	allIngredients, allAllergens, allergenIngredients, err := translationTableFromRecipes(recipes)
	if err != nil {
		return nil, err
	}
	ctx.Printf("Allergens: %d; ingredients: %d", len(allAllergens), len(allIngredients))

	unsafe := make(map[string]string)
	for alg, ing := range allergenIngredients {
		unsafe[ing] = alg
	}

	var safeIngredients []string
	for _, ing := range allIngredients {
		if _, ok := unsafe[ing]; !ok {
			safeIngredients = append(safeIngredients, ing)
		}
	}
	ctx.Printf("Safe ingredients: (%d) %s", len(safeIngredients), safeIngredients)

	totalOcc := 0
	for _, ing := range safeIngredients {
		totalOcc += ingredientOccurrence[ing]
	}

	ctx.Printf("%v", allergenIngredients)
	return totalOcc, nil
}

func Dec21b(ctx ch.AOContext) (interface{}, error) {
	lines, err := ctx.DataLines("inputs/2020/dec21.txt")
	if err != nil {
		return nil, err
	}

	recipes := readRecipeList(lines)
	allIngredients, allAllergens, allergenIngredients, err := translationTableFromRecipes(recipes)
	if err != nil {
		return nil, err
	}
	ctx.Printf("Allergens: %d; ingredients: %d", len(allAllergens), len(allIngredients))

	unsafe := make(map[string]string)
	for alg, ing := range allergenIngredients {
		unsafe[ing] = alg
	}

	ctx.Printf("%v", allergenIngredients)

	rv := ""
	for _, alg := range allAllergens {
		rv += "," + allergenIngredients[alg]
	}

	return rv[1:], nil
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

func translationTableFromRecipes(recipes []recipe) (allIngredients []string, allAllergens []string, allergenIngredients map[string]string, err error) {
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

	allAllergens = make([]string, 0, len(allergenMap))
	allIngredients = make([]string, 0, len(ingredientMap))

	for alg := range allergenMap {
		allAllergens = append(allAllergens, alg)
	}
	for ing := range ingredientMap {
		allIngredients = append(allIngredients, ing)
	}

	sort.Strings(allAllergens)
	sort.Strings(allIngredients)

	for i, alg := range allAllergens {
		allergenMap[alg] = i
	}
	for j, ing := range allIngredients {
		ingredientMap[ing] = j
	}

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

	allergenIngredients = make(map[string]string)
	done := false
	for !done {
		done = true
		for i, alg := range allAllergens {
			found := 0
			last := -1
			for j := range allIngredients {
				if transTable[i][j] {
					found++
					last = j
				}
			}

			if found == 1 {
				allergenIngredients[alg] = allIngredients[last]
				// update mask
				for j := range allIngredients {
					transTable[i][j] = j == last
				}
				for ii := range allAllergens {
					transTable[ii][last] = ii == i
				}
			} else if found == 0 {
				return nil, nil, nil, fmt.Errorf("no possible ingredients left for allergen '%s'", alg)
			} else {
				done = false
			}
		}
	}

	return
}
