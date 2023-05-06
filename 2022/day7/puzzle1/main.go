package main

import (
	"fmt"
	"log"

	"github.com/alec-w/adventofcode/2022/day7/internal"
)

func main() {
	rootDir, err := internal.DirFromFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sizes := rootDir.Sizes()
	total := 0
	for _, size := range sizes {
		if size <= 100000 {
			total += size
		}
	}

	fmt.Printf("total: %d\n", total)
}
