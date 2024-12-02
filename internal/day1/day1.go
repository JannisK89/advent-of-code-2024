package day1

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func processLine(line string) []string {
	values := strings.Fields(line)
	return values
}

func calculateDistance(firstNum int, secondNum int) int {
	abs := math.Abs(float64(firstNum) - float64(secondNum))
	return int(abs)

}

func sortSlice(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	return slice
}

func Part1(input []string) int {
	var totalDistance int
	var leftNumbers, rightNumbers []int

	for i, line := range input {
		values := processLine(line)
		firstValue, firstErr := strconv.Atoi(values[0])
		secondvalue, secondErr := strconv.Atoi(values[1])

		if firstErr != nil || secondErr != nil {
			fmt.Printf("Error in line %d: %s\n", i, line)
		}

		leftNumbers = append(leftNumbers, firstValue)
		rightNumbers = append(rightNumbers, secondvalue)
	}

	leftNumbers = sortSlice(leftNumbers)
	rightNumbers = sortSlice(rightNumbers)

	for i, left := range leftNumbers {
		right := rightNumbers[i]
		distance := calculateDistance(left, right)
		totalDistance += distance
	}
	return totalDistance
}
