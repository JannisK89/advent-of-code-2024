package day2

import (
	"strconv"
	"strings"
)

func castLineToInt(line string) []int {
	var values []int

	for _, value := range strings.Split(line, " ") {
		value, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		values = append(values, value)
	}
	return values
}

func isIncreasing(values []int) bool {
	bigger := 0
	smaller := 0
	for i := 0; i < len(values)-1; i++ {
		if values[i] > values[i+1] {
			smaller++
		} else {
			bigger++
		}
	}
	return bigger > smaller
}

func checkIfLineIsSafe(values []int, dampener bool) bool {
	isDampenerUsed := !dampener

	if isIncreasing(values) {
		for i := 0; i < len(values)-1; i++ {
			if (values[i+1]-values[i]) > 3 || (values[i+1]-values[i]) < 1 {

				if isDampenerUsed {
					return false
				} else {
					values[i+1] = values[i]
					isDampenerUsed = true
				}
			}
		}
	} else {
		for i := 0; i < len(values)-1; i++ {
			if (values[i]-values[i+1]) > 3 || (values[i]-values[i+1]) < 1 {
				if isDampenerUsed {
					return false
				} else {
					values[i+1] = values[i]
					isDampenerUsed = true
				}

			}
		}
	}
	return true
}

func Part1(input []string) int {
	total := 0

	for _, line := range input {
		values := castLineToInt(line)
		if checkIfLineIsSafe(values, false) {
			total++
		}
	}
	return total
}

func Part2(input []string) int {
	total := 0

	for _, line := range input {
		values := castLineToInt(line)
		if checkIfLineIsSafe(values, true) {
			total++
		}
	}
	return total
}
