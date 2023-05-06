package main

import (
	"fmt"
	"log"

	"github.com/alec-w/adventofcode/2022/day7/internal"
)

const (
	totalSpace    = 70000000
	requiredSpace = 30000000
)

func main() {
	rootDir, err := internal.DirFromFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sizes := rootDir.Sizes()
	dirSize := rootDir.Size()
	for _, size := range sizes {
		if size >= requiredSpace-(totalSpace-rootDir.Size()) && size < dirSize {
			dirSize = size
		}
	}

	fmt.Printf("total: %d\n", dirSize)
}
