package main

import (
	"fmt"
	"log"

	"github.com/alec-w/adventofcode/2022/day5/internal"
)

func main() {
	stacks, instructions, err := internal.StacksAndInstructionsFromFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	for _, instruction := range instructions {
		toStack := stacks[instruction.To-1]
		fromStack := stacks[instruction.From-1]
		toStack = append(toStack, fromStack[len(fromStack)-instruction.Amount:]...)
		fromStack = fromStack[:len(fromStack)-instruction.Amount]
		stacks[instruction.To-1] = toStack
		stacks[instruction.From-1] = fromStack
	}

	for _, stack := range stacks {
		fmt.Printf("%s", string(stack[len(stack)-1]))
	}
	fmt.Println("")
}
