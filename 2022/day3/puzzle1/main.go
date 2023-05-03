package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/alec-w/adventofcode/2022/helpers"
)

func main() {
	lines, err := helpers.FileToStrings("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	total := 0
	for _, line := range lines {
		rucksack1 := line[:len(line)/2]
		rucksack2 := line[len(line)/2:]
		for _, char := range rucksack2 {
			if strings.ContainsRune(rucksack1, char) {
				if char >= 'A' && char <= 'Z' {
					total += int(char-'A') + 27
				} else {
					total += int(char-'a') + 1
				}
				break
			}
		}
	}
	fmt.Printf("total: %d\n", total)
}
