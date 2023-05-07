package main

import (
	"fmt"
	"log"

	"github.com/alec-w/adventofcode/2022/day10/internal"
)

func main() {
	signalStrengthByCycle, err := internal.FileToSignalStrengthByCycle("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	crtRows := []string{}
	for i := 0; i < 6; i++ {
		crtRow := ""
		for j := 0; j < 40; j++ {
			cycle := i*40 + j
			if j >= signalStrengthByCycle[cycle]-1 && j <= signalStrengthByCycle[cycle]+1 {
				crtRow += "#"
			} else {
				crtRow += "."
			}
		}
		crtRows = append(crtRows, crtRow)
	}

	for _, row := range crtRows {
		fmt.Println(row)
	}
}
