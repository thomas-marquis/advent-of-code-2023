package days_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thomas-marquis/advent-of-code-2023/days"
)


func TestGetAllPossiblePlacement_4_2(t *testing.T) {
	// Given
	expected := [][]int{
		{1, 1, 0, 0},
		{1, 0, 1, 0},
		{1, 0, 0, 1},
		{0, 1, 1, 0},
		{0, 1, 0, 1},
		{0, 0, 1, 1},
	}

	// When
	res := days.GetAllPossiblePlacement(4, 2)
	
	// Then
	assert.Equal(t, expected, res)
}

func TestGetAllPossiblePlacement_6_3(t *testing.T) {
	// Given
	expected := [][]int{
		{1, 1, 1, 0, 0, 0},
		{1, 1, 0, 1, 0, 0},
		{1, 1, 0, 0, 1, 0},
		{1, 1, 0, 0, 0, 1},
		{1, 0, 1, 1, 0, 0},
		{1, 0, 1, 0, 1, 0},
		{1, 0, 1, 0, 0, 1},
		{1, 0, 0, 1, 1, 0},
		{1, 0, 0, 1, 0, 1},
		{1, 0, 0, 0, 1, 1},
		{0, 1, 1, 1, 0, 0},
		{0, 1, 1, 0, 1, 0},
		{0, 1, 1, 0, 0, 1},
		{0, 1, 0, 1, 1, 0},
		{0, 1, 0, 1, 0, 1},
		{0, 1, 0, 0, 1, 1},
		{0, 0, 1, 1, 1, 0},
		{0, 0, 1, 1, 0, 1},
		{0, 0, 1, 0, 1, 1},
		{0, 0, 0, 1, 1, 1},
	}

	// When
	res := days.GetAllPossiblePlacement(6, 3)
	
	// Then
	assert.ElementsMatch(t, expected, res)
}

func TestG_ArrayEqual_ShouldBeEqual(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}

	assert.True(t, reflect.DeepEqual(a, b))
}

func TestG_ArrayEqual_ShouldBeNotEqual(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 3, 2}

	assert.False(t, reflect.DeepEqual(a, b))
}