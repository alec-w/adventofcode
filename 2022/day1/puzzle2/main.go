package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alec-w/adventofcode/2022/day1/internal"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	elves, err := internal.InputToElves(file)
	if err != nil {
		log.Fatal(err)
	}

	totals := internal.ElvesToOrderedTotals(elves)

	fmt.Printf("Top 3: %d\n", totals[0]+totals[1]+totals[2])
}
