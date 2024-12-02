package utils

import (
	"os"
	"strings"
)

func ReadPuzzleInput(path string) (input []string) {
	file, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	fileString := string(file)
	fileLines := strings.Split(fileString, "\n")

	return fileLines[:len(fileLines)-1]
}
