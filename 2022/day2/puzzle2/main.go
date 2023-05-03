package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/alec-w/adventofcode/2022/day2/internal"
)

func RoundFromString(input string) internal.Round {
	result := internal.Round{}
	switch input[0] {
	case 'A':
		result.Opponent = internal.OPTION_ROCK
	case 'B':
		result.Opponent = internal.OPTION_PAPER
	case 'C':
		result.Opponent = internal.OPTION_SCISSORS
	}
	switch input[2] {
	case 'X':
		result.You = (result.Opponent + 2) % 3
	case 'Y':
		result.You = result.Opponent
	case 'Z':
		result.You = (result.Opponent + 1) % 3
	}
	return result
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		total += RoundFromString(scanner.Text()).Score()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("total: %d\n", total)
}
