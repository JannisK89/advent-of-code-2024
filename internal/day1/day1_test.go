package day1

import (
	"testing"
)

func TestDay1Solution(t *testing.T) {
	var input = []string{"3    4", "4    3", "2    5", "1    3", "3    9", "3    3"}
	result := Part1(input)

	if result != 11 {
		t.Fatalf(`Day 1Part 1 = %d, but expected %d `, result, 11)
	}
}
