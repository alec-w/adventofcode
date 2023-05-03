package internal

import (
	"bufio"
	"io"
	"sort"
	"strconv"
)

func InputToElves(input io.Reader) ([][]uint64, error) {
	var elves [][]uint64
	var elf []uint64
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			elves = append(elves, elf)
			elf = []uint64{}
			continue
		}
		meal, err := strconv.ParseUint(line, 10, 64)
		if err != nil {
			return nil, err
		}
		elf = append(elf, meal)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return elves, nil
}

func ElvesToOrderedTotals(elves [][]uint64) []uint64 {
	var totals []uint64
	for _, elf := range elves {
		var total uint64
		for _, val := range elf {
			total += val
		}
		totals = append(totals, total)
	}
	sort.Slice(totals, func(i, j int) bool { return totals[i] > totals[j] })
	return totals
}
