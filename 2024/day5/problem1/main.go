package main

import (
	"fmt"
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
	rawOrderings := []string{}
	rawPrintings := []string{}
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		if strings.Contains(line, "|") {
			rawOrderings = append(rawOrderings, line)
		} else {
			rawPrintings = append(rawPrintings, line)
		}
	}
	orderings := map[int][]int{}
	for _, rawOrdering := range rawOrderings {
		pages := strings.Split(rawOrdering, "|")
		page0, err := strconv.Atoi(pages[0])
		if err != nil {
			return 0, err
		}
		page1, err := strconv.Atoi(pages[1])
		if err != nil {
			return 0, err
		}
		ordering := orderings[page0]
		ordering = append(ordering, page1)
		orderings[page0] = ordering
	}
	total := 0
rawPrintings:
	for _, rawPrinting := range rawPrintings {
		rawPages := strings.Split(rawPrinting, ",")
		pages := make(map[int]int, len(rawPages))
		for i, rawPage := range rawPages {
			page, err := strconv.Atoi(rawPage)
			if err != nil {
				return 0, err
			}
			pages[page] = i
		}
		for page, idx := range pages {
			rules := orderings[page]
			for _, rule := range rules {
				otherIdx, ok := pages[rule]
				if !ok {
					continue
				}
				if otherIdx < idx {
					continue rawPrintings
				}
			}
		}
		middle, err := strconv.Atoi(rawPages[len(rawPages)/2])
		if err != nil {
			return 0, err
		}
		total += middle
	}
	return total, nil
}
