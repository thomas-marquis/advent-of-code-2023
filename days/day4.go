package days

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/thomas-marquis/advent-of-code-2023/utils"
)


var (
	reg = regexp.MustCompile(`^Card\s*\d*:\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s\|\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)$`)
)

func ToIntSlice(strSlice []string) []int {
	var intSlice []int
	for _, str := range strSlice {
		intVal, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		intSlice = append(intSlice, intVal)
	}
	return intSlice
}

func Day4() {
	scanner, _ := utils.ReadFileLines("resources/day4_input")

	var res int
	for scanner.Scan() {
		lineContent := scanner.Text()
		groups := reg.FindStringSubmatch(lineContent)

		winningNbs := ToIntSlice(groups[1:11])
		userNbs := ToIntSlice(groups[11:])

		score := 0
		var matchingNb []int
		for _, winN := range winningNbs {
			for _, userN := range userNbs {
				if winN == userN {
					if score == 0 {
						score = 1
					} else {
						score *= 2
					}

					matchingNb = append(matchingNb, winN)
				}
			}
		}
		res += score

		fmt.Printf("Score = %d ; Winning numbers: %#v\n", score, matchingNb)
	}

	fmt.Println(res)
}