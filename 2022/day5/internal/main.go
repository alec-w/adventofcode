package internal

import (
	"strconv"
	"strings"

	"github.com/alec-w/adventofcode/2022/helpers"
)

type Stack []rune

type Instruction struct {
	Amount int
	From   int
	To     int
}

func StacksAndInstructionsFromFile(fileName string) ([]Stack, []Instruction, error) {
	lines, err := helpers.FileToStrings(fileName)
	if err != nil {
		return nil, nil, err
	}

	var rawStacks []string
	var rawInstructions []string
	var totalNumberOfStacks int
	for i, line := range lines {
		if line == "" {
			rawStacks = lines[:i-1]
			rawInstructions = lines[i+1:]
			stackNumbers := strings.Split(strings.TrimSpace(lines[i-1]), "   ")
			totalNumberOfStacks, err = strconv.Atoi(stackNumbers[len(stackNumbers)-1])
			if err != nil {
				return nil, nil, err
			}
			break
		}
	}

	stacks := make([]Stack, totalNumberOfStacks)
	for i := range rawStacks {
		rawStack := strings.TrimRight(rawStacks[len(rawStacks)-1-i], " ")
		stackIndex := 0
		for {
			char := rawStack[0]
			if char == '[' {
				stacks[stackIndex] = append(stacks[stackIndex], rune(rawStack[1]))
				if len(rawStack) == 3 {
					break
				}
			}
			rawStack = rawStack[4:]
			stackIndex++
		}
	}

	var instructions []Instruction
	for _, instruction := range rawInstructions {
		parts := strings.Split(instruction, " ")
		amount, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, err
		}
		from, err := strconv.Atoi(parts[3])
		if err != nil {
			return nil, nil, err
		}
		to, err := strconv.Atoi(parts[5])
		if err != nil {
			return nil, nil, err
		}
		instructions = append(instructions, Instruction{Amount: amount, From: from, To: to})
	}

	return stacks, instructions, nil
}
