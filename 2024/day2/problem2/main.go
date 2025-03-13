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
	totalSafe := 0
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		values := strings.Split(line, " ")
		safe, idx, err := isSafe(values)
		if err != nil {
			return 0, err
		}
		if !safe {
			safe, _, err = isSafe(append(slices.Clone(values[0:idx]), values[idx+1:]...))
			if err != nil {
				return 0, err
			}
			if !safe {
				safe, _, err = isSafe(append(slices.Clone(values[0:idx+1]), values[idx+2:]...))
				if err != nil {
					return 0, err
				}
				if !safe {
					if idx != 1 {
						continue
					}
					safe, _, err = isSafe(values[1:])
					if err != nil {
						return 0, err
					}
					if !safe {
						continue
					}
				}
			}
		}
		totalSafe++
	}
	return totalSafe, nil
}

func isSafe(values []string) (bool, int, error) {
	value1, err := strconv.Atoi(values[0])
	if err != nil {
		return false, 0, err
	}
	value2, err := strconv.Atoi(values[1])
	if err != nil {
		return false, 0, err
	}
	increasing := value1 < value2
	if math.Abs(float64(value1-value2)) > 3 || value1 == value2 {
		return false, 0, nil
	}
	prev := value2
	for i, strValue := range values[2:] {
		value, err := strconv.Atoi(strValue)
		if err != nil {
			return false, 0, err
		}
		if (value > prev) != increasing {
			return false, i + 1, nil
		}
		if math.Abs(float64(value-prev)) > 3 || value == prev {
			return false, i + 1, nil
		}
		prev = value
	}
	return true, 0, nil
}
