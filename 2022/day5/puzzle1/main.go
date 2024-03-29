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
		for i := 0; i < instruction.Amount; i++ {
			stacks[instruction.To-1] = append(stacks[instruction.To-1], stacks[instruction.From-1][len(stacks[instruction.From-1])-1])
			stacks[instruction.From-1] = stacks[instruction.From-1][:len(stacks[instruction.From-1])-1]
		}
	}

	for _, stack := range stacks {
		fmt.Printf("%s", string(stack[len(stack)-1]))
	}
	fmt.Println("")
}
