package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/alec-w/adventofcode/2024/internal/load"
)

func main() {
	ans, err := app()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(ans)
}

func app() (int, error) {
	lines, err := load.FileToLines("../input.txt")
	if err != nil {
		return 0, err
	}
	wordsearch := [][]rune{}
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		wordSearchLine := make([]rune, len(line))
		for j, char := range line {
			wordSearchLine[j] = char
		}
		wordsearch = append(wordsearch, wordSearchLine)
	}
	total := 0
	for i, wordSearchLine := range wordsearch {
		for j, char := range wordSearchLine {
			if char != 'A' {
				continue
			}
			if i < 1 || i > len(wordsearch)-2 {
				continue
			}
			if j < 1 || j > len(wordSearchLine)-2 {
				continue
			}
			pair1 := []rune{wordsearch[i-1][j-1], wordsearch[i+1][j+1]}
			pair2 := []rune{wordsearch[i+1][j-1], wordsearch[i-1][j+1]}
			slices.Sort(pair1)
			slices.Sort(pair2)
			if string(pair1) == "MS" && string(pair2) == "MS" {
				total++
			}
		}
	}
	return total, nil
}
