package days

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/thomas-marquis/advent-of-code-2023/utils"
)


type gameSet struct {
	NbRed int
	NbGreen int
	NbBlue int
}

type game struct {
	Sets []gameSet
	GameIndex int
}

func (g *game) GetColorMaxNb(colorName string) int{
	var max int
	for _, set := range g.Sets {
		var nb int
		switch colorName {
		case "red":
			nb = set.NbRed
		case "green":
			nb = set.NbGreen
		case "blue":
			nb = set.NbBlue
		}

		if nb > max {
			max = nb
		}
	}
	return max
}


func parseColor(colorName string, color string, nb *int) {
	var err error
	if strings.Contains(color, colorName) {
		nbRedStr := strings.Replace(color, colorName, "", -1)
		nbRedStr = strings.TrimSpace(nbRedStr)
		*nb, err = strconv.Atoi(nbRedStr)
		if err != nil {
			errMsg := fmt.Sprintf("Error while parsing game color %s: %s", color, err)
			panic(errMsg)
		}
	}
}


func parseLine(line string) (game, error) {
	splitted := strings.Split(line, ":")
	gameindexStr := splitted[0]
	gameSetsStr := splitted[1]

	gameIndexStr := strings.Replace(gameindexStr, "Game ", "", -1)
	gameIndex, err := strconv.Atoi(gameIndexStr)
	if err != nil {
		errMsg := fmt.Sprintf("Error while parsing game line %s: %s", line, err)
		return game{}, errors.New(errMsg)
	}
	g := game{
		GameIndex: gameIndex,
	}

	for _, set := range strings.Split(gameSetsStr, ";") {
		setSplitted := strings.Split(set, ",")

		var nbRed int
		var nbGreen int
		var nbBlue int

		for _, color := range setSplitted {
			parseColor("red", color, &nbRed)
			parseColor("green", color, &nbGreen)
			parseColor("blue", color, &nbBlue)
		}

		gameSet := gameSet{
			NbRed: nbRed,
			NbGreen: nbGreen,
			NbBlue: nbBlue,
		}
		g.Sets = append(g.Sets, gameSet)
	}

	return g, nil
}


func Day2() {
	scanner, err := utils.ReadFileLines("resources/day2_input")
	if err != nil {
		panic(err)
	}

	var scoreMax int
	var scorePower int

	for scanner.Scan() {
		lineContent := scanner.Text()
		g, err := parseLine(lineContent)
		if err != nil {
			panic(err)
		}

		if g.GetColorMaxNb("red") <= 12 && g.GetColorMaxNb("green") <= 13 && g.GetColorMaxNb("blue") <= 14 {
			scoreMax += g.GameIndex
		}

		scorePower += g.GetColorMaxNb("red") * g.GetColorMaxNb("green") * g.GetColorMaxNb("blue")
	}

	println(scoreMax)
	println(scorePower)
}