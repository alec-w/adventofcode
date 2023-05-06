package main

import (
	"fmt"
	"log"

	"github.com/alec-w/adventofcode/2022/day8/internal"
)

func main() {
	trees, err := internal.TreesFromFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	totalVisible := 0
	for i, row := range trees {
		for j, height := range row {
			if i == 0 || j == 0 || i == len(trees)-1 || j == len(row)-1 {
				totalVisible++
				continue
			}
			maxHeight := 0
			for i1 := 0; i1 < i; i1++ {
				if trees[i1][j] > maxHeight {
					maxHeight = trees[i1][j]
				}
			}
			if maxHeight < height {
				totalVisible++
				continue
			}
			maxHeight = 0
			for i1 := len(trees) - 1; i1 > i; i1-- {
				if trees[i1][j] > maxHeight {
					maxHeight = trees[i1][j]
				}
			}
			if maxHeight < height {
				totalVisible++
				continue
			}
			maxHeight = 0
			for j1 := 0; j1 < j; j1++ {
				if trees[i][j1] > maxHeight {
					maxHeight = trees[i][j1]
				}
			}
			if maxHeight < height {
				totalVisible++
				continue
			}
			maxHeight = 0
			for j1 := len(row) - 1; j1 > j; j1-- {
				if trees[i][j1] > maxHeight {
					maxHeight = trees[i][j1]
				}
			}
			if maxHeight < height {
				totalVisible++
				continue
			}
		}
	}

	fmt.Printf("total visible: %d\n", totalVisible)
}
