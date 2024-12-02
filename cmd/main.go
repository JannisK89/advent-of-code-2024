package main

import (
	"fmt"

	"github.com/JannisK89/advent-of-code-2024/internal/day1"
	"github.com/JannisK89/advent-of-code-2024/internal/utils"
)

func main() {
	day1Input := utils.ReadPuzzleInput("internal/day1/input.txt")
	fmt.Printf("Day 1 Part 1: %d\n", day1.Part1(day1Input))
	fmt.Printf("Day 1 Part 2: %d\n", day1.Part2(day1Input))
}
