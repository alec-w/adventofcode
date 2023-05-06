package internal

import (
	"strconv"

	"github.com/alec-w/adventofcode/2022/helpers"
)

func TreesFromFile(fileName string) ([][]int, error) {
	lines, err := helpers.FileToStrings(fileName)
	if err != nil {
		return nil, err
	}

	trees := [][]int{}
	for i, line := range lines {
		trees = append(trees, []int{})
		for _, val := range line {
			height, err := strconv.Atoi(string(val))
			if err != nil {
				return nil, err
			}
			trees[i] = append(trees[i], height)
		}
	}

	return trees, nil
}
