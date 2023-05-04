package main

import (
	"fmt"
	"log"

	"github.com/alec-w/adventofcode/2022/day6/internal"
	"github.com/alec-w/adventofcode/2022/helpers"
)

func main() {
	lines, err := helpers.FileToStrings("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("chars: %d\n", internal.CharactersBeforeXDistinct(lines[0], 14))
}
