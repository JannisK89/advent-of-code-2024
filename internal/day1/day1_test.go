package day1

import (
	"testing"
)

var input = []string{"3    4", "4    3", "2    5", "1    3", "3    9", "3    3"}

func TestDay1Part1Solution(t *testing.T) {
	result := Part1(input)

	if result != 11 {
		t.Fatalf(`Day 1 Part 1 = %d, but expected %d `, result, 11)
	}
}

func TestDay1Part2Solution(t *testing.T) {
	result := Part2(input)

	if result != 31 {
		t.Fatalf(`Day 1 Part 2 = %d, but expected %d `, result, 31)
	}
}
