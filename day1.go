package main

import (
	"regexp"
	"strconv"

	"github.com/thomas-marquis/advent-of-code-2023/utils"
)

var (
	one = "one"
	two = "two"
	three = "three"
	four = "four"
	five = "five"
	six = "six"
	seven = "seven"
	eight = "eight"
	nine = "nine" 
	digits = []string{one, two, three, four, five, six, seven, eight, nine}
	digitsMap = make(map[string]int)
	digitRegex = regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`)
	invertedDigitRegex = regexp.MustCompile(`\d|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin`)
)

func invertString(str string) string {
	var inverted string
	for i := len(str) - 1; i >= 0; i-- {
		inverted += string(str[i])
	}
	return inverted
}

func main() {
	for i, digit := range digits {
		digitsMap[digit] = i + 1
	}
	var numbers []int
	
	scanner, err := utils.ReadFileLines("resources/day1_input")
	if err != nil {
		panic(err)
	}

	for scanner.Scan() {
		lineContent := scanner.Text()

		var firstNumber int
		var lastNumber int

		res := digitRegex.FindAllString(lineContent, 1)
		firstNumberStr := res[0]
		numValFirst, errFirst := strconv.Atoi(firstNumberStr)
		if errFirst == nil {
			firstNumber = numValFirst
		} else {
			firstNumber = digitsMap[firstNumberStr]
		}

		lastNumberStr := invertedDigitRegex.FindAllString(invertString(lineContent), 1)[0]
		numValLast, errLast := strconv.Atoi(lastNumberStr)
		if errLast == nil {
			lastNumber = numValLast
		} else {
			lastNumber = digitsMap[invertString(lastNumberStr)]
		}

		x := firstNumber*10 + lastNumber
		println(x)
		numbers = append(numbers, x)
	}

	var sum int
	for _, number := range numbers {
		sum += number
	}

	println(sum)
}
