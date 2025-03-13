package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
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
	file, err := os.Open("../input.txt")
	if err != nil {
		return 0, err
	}
	var input strings.Builder
	if _, err := io.Copy(&input, file); err != nil {
		return 0, err
	}
	lines := strings.Split(input.String(), "\n")
	list1, list2 := make([]int, len(lines)), make([]int, len(lines))
	for i, line := range strings.Split(input.String(), "\n") {
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
