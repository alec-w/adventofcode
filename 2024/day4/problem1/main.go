package main

import (
	"fmt"
	"os"
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
			if char != 'X' {
				continue
			}
			canUp := i >= 3
			canDown := i <= len(wordsearch)-4
			canLeft := j >= 3
			canRight := j <= len(wordSearchLine)-4
			toCheck := []string{}
			if canUp {
				toCheck = append(toCheck, string([]rune{char, wordsearch[i-1][j], wordsearch[i-2][j], wordsearch[i-3][j]}))
				if canLeft {
					toCheck = append(toCheck, string([]rune{char, wordsearch[i-1][j-1], wordsearch[i-2][j-2], wordsearch[i-3][j-3]}))
				}
				if canRight {
					toCheck = append(toCheck, string([]rune{char, wordsearch[i-1][j+1], wordsearch[i-2][j+2], wordsearch[i-3][j+3]}))
				}
			}
			if canDown {
				toCheck = append(toCheck, string([]rune{char, wordsearch[i+1][j], wordsearch[i+2][j], wordsearch[i+3][j]}))
				if canLeft {
					toCheck = append(toCheck, string([]rune{char, wordsearch[i+1][j-1], wordsearch[i+2][j-2], wordsearch[i+3][j-3]}))
				}
				if canRight {
					toCheck = append(toCheck, string([]rune{char, wordsearch[i+1][j+1], wordsearch[i+2][j+2], wordsearch[i+3][j+3]}))
				}
			}
			if canLeft {
				toCheck = append(toCheck, string([]rune{char, wordsearch[i][j-1], wordsearch[i][j-2], wordsearch[i][j-3]}))
			}
			if canRight {
				toCheck = append(toCheck, string([]rune{char, wordsearch[i][j+1], wordsearch[i][j+2], wordsearch[i][j+3]}))
			}
			for _, value := range toCheck {
				if value == "XMAS" {
					total++
				}
			}
		}
	}
	return total, nil
}
