package internal

import (
	"strconv"
	"strings"

	"github.com/alec-w/adventofcode/2022/helpers"
)

func FileToSignalStrengthByCycle(fileName string) ([]int, error) {
	lines, err := helpers.FileToStrings(fileName)
	if err != nil {
		return nil, err
	}

	signalStrengthByCycle := []int{1}
	for _, line := range lines {
		if line == "noop" {
			signalStrengthByCycle = append(signalStrengthByCycle, signalStrengthByCycle[len(signalStrengthByCycle)-1])
			continue
		}
		parts := strings.Split(line, " ")
		val, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}
		signalStrengthByCycle = append(signalStrengthByCycle, signalStrengthByCycle[len(signalStrengthByCycle)-1])
		signalStrengthByCycle = append(signalStrengthByCycle, signalStrengthByCycle[len(signalStrengthByCycle)-1]+val)
	}

	return signalStrengthByCycle, nil
}
