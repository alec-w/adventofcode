package main

import (
	"fmt"
	"os"
	"regexp"
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
	input, err := load.FileToString("../input.txt")
	if err != nil {
		return 0, err
	}
	total := 0
	for _, match := range regexp.MustCompile(`mul\(\d+,\d+\)`).FindAllString(input, -1) {
		values := strings.Split(strings.TrimSuffix(strings.TrimPrefix(match, "mul("), ")"), ",")
		value1, err := strconv.Atoi(values[0])
		if err != nil {
			return 0, err
		}
		value2, err := strconv.Atoi(values[1])
		if err != nil {
			return 0, err
		}
		total += value1 * value2
	}
	return total, nil
}
