package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/alec-w/adventofcode/2024/internal/load"
)

func main() {
	ans, err := app()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(ans)
}

func app() (int, error) {
	lines, err := load.FileToLines("../input.txt")
	if err != nil {
		return 0, err
	}
	list1, list2 := make([]int, len(lines)), make([]int, len(lines))
	for i, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		values := strings.Split(line, "   ")
		value, err := strconv.Atoi(strings.TrimSpace(values[0]))
		if err != nil {
			return 0, err
		}
		list1[i] = value
		value, err = strconv.Atoi(strings.TrimSpace(values[1]))
		if err != nil {
			return 0, err
		}
		list2[i] = value
	}
	slices.Sort(list1)
	slices.Sort(list2)
	totalDiffs := 0
	for i, val1 := range list1 {
		totalDiffs += int(math.Abs(float64(val1 - list2[i])))
	}
	return totalDiffs, nil
}
