package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/alec-w/adventofcode/2022/helpers"
)

func main() {
	lines, err := helpers.FileToStrings("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var rawStacks []string
	var instructions []string
	var totalNumberOfStacks int
	for i, line := range lines {
		if line == "" {
			rawStacks = lines[:i-1]
			instructions = lines[i+1:]
			stackNumbers := strings.Split(strings.TrimSpace(lines[i-1]), "   ")
			totalNumberOfStacks, err = strconv.Atoi(stackNumbers[len(stackNumbers)-1])
			if err != nil {
				log.Fatal(err)
			}
			break
		}
	}

	stacks := make([][]rune, totalNumberOfStacks)
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

	for i, stack := range stacks {
		fmt.Printf("%d:", i+1)
		for _, val := range stack {
			fmt.Printf(" %s", string(val))
		}
		fmt.Println("")
	}

	for _, instruction := range instructions {
		parts := strings.Split(instruction, " ")
		amount, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}
		from, err := strconv.Atoi(parts[3])
		if err != nil {
			log.Fatal(err)
		}
		to, err := strconv.Atoi(parts[5])
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < amount; i++ {
			stacks[to-1] = append(stacks[to-1], stacks[from-1][len(stacks[from-1])-1])
			stacks[from-1] = stacks[from-1][:len(stacks[from-1])-1]
		}
	}

	for _, stack := range stacks {
		fmt.Printf("%s", string(stack[len(stack)-1]))
	}
	fmt.Println("")
}
