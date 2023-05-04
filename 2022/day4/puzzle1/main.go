package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/alec-w/adventofcode/2022/day4/internal"
	"github.com/alec-w/adventofcode/2022/helpers"
)

func main() {
	lines, err := helpers.FileToStrings("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	overlaps := 0
	for _, line := range lines {
		inputs := strings.Split(line, ",")
		assignment1, err := internal.AssignmentFromString(inputs[0])
		if err != nil {
			log.Fatal(err)
		}
		assignment2, err := internal.AssignmentFromString(inputs[1])
		if err != nil {
			log.Fatal(err)
		}
		if (assignment1.Min >= assignment2.Min && assignment1.Max <= assignment2.Max) || (assignment2.Min >= assignment1.Min && assignment2.Max <= assignment1.Max) {
			overlaps += 1
		}
	}

	fmt.Printf("overlaps: %d\n", overlaps)
}
