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

	toAdd := []int{
		20 * signalStrengthByCycle[19],
		60 * signalStrengthByCycle[59],
		100 * signalStrengthByCycle[99],
		140 * signalStrengthByCycle[139],
		180 * signalStrengthByCycle[179],
		220 * signalStrengthByCycle[219],
	}
	total := 0
	for _, val := range toAdd {
		total += val
	}
	fmt.Printf("total: %d\n", total)
}
