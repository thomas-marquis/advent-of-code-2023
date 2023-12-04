package days

import (
	"strings"

	"github.com/thomas-marquis/advent-of-code-2023/utils"
)


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

func (m *matrix) NextNumber() (string, error) {
	if m.currentLine >= len(m.NbLines()) {
		return "", errors.New("End of matrix")
	}
	if m.currentColumn >= len(m.Content[m.currentLine]) {
		m.currentLine++
		m.currentColumn = 0
		return m.NextNumber()
	}

	item := m.Content[m.currentLine][m.currentColumn]
	
	m.currentColumn++
	return number, nil


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
}