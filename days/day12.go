package days

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/thomas-marquis/advent-of-code-2023/utils"
	"gonum.org/v1/gonum/stat/combin"
)

func Day12() {
	scanner, _ := utils.ReadFileLines("resources/day12_input")
	var res int
	for scanner.Scan() {
		lineContent := scanner.Text()

		splittedLine := strings.Split(lineContent, " ")

		spingsStr := splittedLine[0]
		numbersStr := splittedLine[1]

		springs := strings.Split(spingsStr, "")
		numbers := utils.ToIntSlice(strings.Split(numbersStr, ","))
		expectedKoNb := utils.SumSlice(numbers)

		var knownKoNb int
		var unknownSprings int
		var unknownSpringsIdx []int
		for i, c := range springs {
			if c == "#" {
				knownKoNb++
			}
			if c == "?" {
				unknownSprings++
				unknownSpringsIdx = append(unknownSpringsIdx, i)
			}
		}

		missingKos := expectedKoNb - knownKoNb

		placements := GetAllPossiblePlacement(unknownSprings, missingKos)

		for _, p := range placements {
			var springsCopy []string = make([]string, len(springs), len(springs))
			copy(springsCopy, springs)
			var unIdx int
			for i, s := range springsCopy {
				if s == "?" {
					if p[unIdx] == 1 {
						springsCopy[i] = "#"
					} else {
						springsCopy[i] = "."
					}
					unIdx++
				}
			}
			groups := ComputeGroups(springsCopy)
			if reflect.DeepEqual(groups, numbers) {
				res++
			}
		}

		fmt.Printf("Known ko nb: %d vs. expected ones %d, so missing=%d | Unknown states: %d\n", knownKoNb, expectedKoNb, missingKos, unknownSprings)
	}

	fmt.Printf("\nResult: %d\n", res)
}

func ComputeGroups(springs []string) []int {
	var groups []int
	var currGrpIdx int = -1
	var inGroup bool
	for _, c := range springs {
		if c != "." {
			if !inGroup {
				currGrpIdx++
				inGroup = true
				groups = append(groups, 0)
			}
			groups[currGrpIdx]++
		} else {
			inGroup = false
		}
	}
	return groups
}

func GetAllPossiblePlacement(availablePlacementsLength int, nbItemToPlace int) [][]int {
	res := combin.Combinations(availablePlacementsLength, nbItemToPlace)
	var possibilities [][]int
	for _, comb := range res {
		possibility := make([]int, availablePlacementsLength, availablePlacementsLength)
		for _, idx := range comb {
			possibility[idx] = 1
		}
		possibilities = append(possibilities, possibility)
	
	}
	return possibilities
}