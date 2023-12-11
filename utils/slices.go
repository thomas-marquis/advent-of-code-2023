package utils

import "strconv"

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

func IsIn(val int, slice []int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}