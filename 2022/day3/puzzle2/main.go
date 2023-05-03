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
	for i := 0; i < len(lines); i += 3 {
		elf1 := lines[i]
		elf2 := lines[i+1]
		elf3 := lines[i+2]
		for _, char := range elf1 {
			if strings.ContainsRune(elf2, char) && strings.ContainsRune(elf3, char) {
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
