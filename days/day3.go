package days

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/thomas-marquis/advent-of-code-2023/utils"
)

type engineNumber struct {
	Value string
	lineIndex int
	startIndex int
	endIndex int
	aroundSymbols []string
}

func (n *engineNumber) IsEnginePart() bool {
	symbolsRegexp := regexp.MustCompile(`[^a-zA-Z\d.]`)

	var isEnginePart bool
	for _, symbol := range n.aroundSymbols {
		if symbolsRegexp.MatchString(symbol) {
			isEnginePart = true
			break
		}
	}

	return isEnginePart
}


type matrix struct {
	Content [][]string

	currentLine int
	currentColumn int
}

func (m *matrix) NbLines() int {
	return len(m.Content)
}

func (m *matrix) NbColumns() int {
	return len(m.Content[0])
}

func (m *matrix) MoveRight() {
	m.currentColumn++
}

func (m *matrix) MoveDown() {
	m.currentLine++
}

func (m *matrix) NextLine() {
	m.MoveDown()
	m.currentColumn = 0
}

func (m *matrix) FillAroundSymbols(n *engineNumber) {
	if n.startIndex > 0 {
		n.aroundSymbols = append(n.aroundSymbols, m.Content[n.lineIndex][n.startIndex-1])
	}
	if n.endIndex < m.NbColumns()-1 {
		n.aroundSymbols = append(n.aroundSymbols, m.Content[n.lineIndex][n.endIndex+1])
	}

	start := max(0, n.startIndex-1)
	stop := min(m.NbColumns()-1, n.endIndex+1)
	for i := start; i <= stop; i++ {
		if n.lineIndex > 0 {
			n.aroundSymbols = append(n.aroundSymbols, m.Content[n.lineIndex-1][i])
		}
		if n.lineIndex < m.NbLines()-1 {
			n.aroundSymbols = append(n.aroundSymbols, m.Content[n.lineIndex+1][i])
		}
	}
}

func (m *matrix) NextNumber() (engineNumber, error) {
	if m.currentLine >= m.NbLines() {
		return engineNumber{}, errors.New("End of matrix")
	}
	if m.currentColumn >= m.NbColumns() {
		m.NextLine()
		return m.NextNumber()
	}

	item := m.Content[m.currentLine][m.currentColumn]
	regexp := regexp.MustCompile(`\d`)
	number := regexp.FindString(item)
	if number == "" {
		m.MoveRight()
		return m.NextNumber()
	}
	
	lastNbIndex := m.currentColumn
	detectedNumber := []string{number}
	for i := m.currentColumn+1; i < m.NbColumns(); i++ {
		nextItem := m.Content[m.currentLine][i]
		nextNumber := regexp.FindString(nextItem)
		if nextNumber == "" {
			break
		}
		lastNbIndex = i
		detectedNumber = append(detectedNumber, nextNumber)
	}

	
	n := engineNumber{
		Value: strings.Join(detectedNumber, ""),
		startIndex: m.currentColumn,
		endIndex: lastNbIndex,
		lineIndex: m.currentLine,
	}

	m.currentColumn = lastNbIndex + 1

	m.FillAroundSymbols(&n)
	return n, nil
}

func Day3() {
	scanner, err := utils.ReadFileLines("resources/day3_input")
	if err != nil {
		panic(err)
	}

	var matrix matrix
	for scanner.Scan() {
		lineContent := scanner.Text()
		lineChars := strings.Split(lineContent, "")
		matrix.Content = append(matrix.Content, lineChars)
	}

	var res int
	for {
		n, err := matrix.NextNumber()
		if err != nil {
			break
		}
		fmt.Printf("Number %s : %s => engine part ? %t\n", n.Value, n.aroundSymbols, n.IsEnginePart())
		if n.IsEnginePart() {
			val, _ := strconv.Atoi(n.Value)
			res += val
		}
	}

	fmt.Printf("Result: %d\n", res)
}