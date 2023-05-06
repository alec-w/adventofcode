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

	topScore := 0
	for i, row := range trees {
		for j, height := range row {
			up := 0
			for i1 := i - 1; i1 >= 0; i1-- {
				up++
				if trees[i1][j] >= height {
					break
				}
			}
			down := 0
			for i1 := i + 1; i1 <= len(trees)-1; i1++ {
				down++
				if trees[i1][j] >= height {
					break
				}
			}
			left := 0
			for j1 := j - 1; j1 >= 0; j1-- {
				left++
				if trees[i][j1] >= height {
					break
				}
			}
			right := 0
			for j1 := j + 1; j1 <= len(trees)-1; j1++ {
				right++
				if trees[i][j1] >= height {
					break
				}
			}
			score := up * down * left * right
			if score > topScore {
				topScore = score
			}
		}
	}

	fmt.Printf("top score: %d\n", topScore)
}
