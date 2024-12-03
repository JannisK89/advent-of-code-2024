package day2

import (
	"testing"
)

var input = []string{"7 6 4 2 1", "1 2 7 8 9", "9 7 6 2 1", "1 3 2 4 5", "8 6 4 4 1", "1 3 6 7 9", "2 1 3 6 9", "2 2 3 6 9", "8 8 5 4 1", "8 10 5 4 1", "1 6 5 4 1"}

func TestDay2Part1Solution(t *testing.T) {
	result := Part1(input)

	if result != 2 {
		t.Fatalf(`Day 2 Part 1 = %d, but expected %d `, result, 2)
	}
}

func TestDay2Part2Solution(t *testing.T) {
	result := Part2(input)

	if result != 9 {
		t.Fatalf(`Day 2 Part 2 = %d, but expected %d `, result, 9)
	}
}
