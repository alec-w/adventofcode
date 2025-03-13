package main

import (
	"fmt"
	"math"
	"os"
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
	totalSafe := 0
lines:
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		values := strings.Split(line, " ")
		value1, err := strconv.Atoi(values[0])
		if err != nil {
			return 0, err
		}
		value2, err := strconv.Atoi(values[1])
		if err != nil {
			return 0, err
		}
		increasing := value1 < value2
		if math.Abs(float64(value1-value2)) > 3 || value1 == value2 {
			continue
		}
		prev := value2
		for _, strValue := range values[2:] {
			value, err := strconv.Atoi(strValue)
			if err != nil {
				return 0, err
			}
			if (value > prev) != increasing {
				continue lines
			}
			if math.Abs(float64(value-prev)) > 3 || value == prev {
				continue lines
			}
			prev = value
		}
		totalSafe++
	}
	return totalSafe, nil
}
