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

func createSlices(input []string) ([]int, []int) {
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
	return leftNumbers, rightNumbers
}

func Part1(input []string) int {
	var totalDistance int
	leftNumbers, rightNumbers := createSlices(input)

	leftNumbers = sortSlice(leftNumbers)
	rightNumbers = sortSlice(rightNumbers)

	for i, left := range leftNumbers {
		right := rightNumbers[i]
		distance := calculateDistance(left, right)
		totalDistance += distance
	}
	return totalDistance
}

func Part2(input []string) int {
	var totalSimilarity int
	var counter int
	leftNumbers, rightNumbers := createSlices(input)

	for _, left := range leftNumbers {
		counter = 0
		for _, right := range rightNumbers {
			if left == right {
				counter++
			}
		}
		totalSimilarity += counter * left
	}
	return totalSimilarity
}
