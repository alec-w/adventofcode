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

type OrderingByFirst struct {
	first int
	later []int
}

type OrderingBySecond struct {
	earlier []int
	second  int
}

func app() (int, error) {
	orderingsByFirst, orderingsBySecond, printings, err := parseInput("../input.txt")
	if err != nil {
		return 0, err
	}
	total := 0
	for _, printing := range printings {
		if _, ok := checkOrder(printing, orderingsByFirst); !ok {
			total += middlePage(reorder(printing, orderingsBySecond))
		}
	}
	return total, nil
}

func parseInput(fileName string) (map[int]OrderingByFirst, map[int]OrderingBySecond, []map[int]int, error) {
	lines, err := load.FileToLines(fileName)
	if err != nil {
		return nil, nil, nil, err
	}
	orderingsByFirst := map[int]OrderingByFirst{}
	orderingsBySecond := map[int]OrderingBySecond{}
	printings := []map[int]int{}
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		if strings.Contains(line, "|") {
			first, second, err := parseOrdering(line)
			if err != nil {
				return nil, nil, nil, err
			}
			orderingByFirst, ok := orderingsByFirst[first]
			if !ok {
				orderingByFirst.first = first
				orderingByFirst.later = []int{}
			}
			orderingByFirst.later = append(orderingByFirst.later, second)
			orderingsByFirst[first] = orderingByFirst
			orderingBySecond, ok := orderingsBySecond[second]
			if !ok {
				orderingBySecond.second = second
				orderingBySecond.earlier = []int{}
			}
			orderingBySecond.earlier = append(orderingBySecond.earlier, first)
			orderingsBySecond[second] = orderingBySecond
		} else {
			printing, err := parsePrinting(line)
			if err != nil {
				return nil, nil, nil, err
			}
			printings = append(printings, printing)
		}
	}
	return orderingsByFirst, orderingsBySecond, printings, nil
}

func parseOrdering(line string) (int, int, error) {
	pages := strings.Split(line, "|")
	first, err := strconv.Atoi(pages[0])
	if err != nil {
		return 0, 0, err
	}
	second, err := strconv.Atoi(pages[1])
	if err != nil {
		return 0, 0, err
	}
	return first, second, nil
}

func parsePrinting(line string) (map[int]int, error) {
	rawPages := strings.Split(line, ",")
	pages := make(map[int]int, len(rawPages))
	for i, rawPage := range rawPages {
		page, err := strconv.Atoi(rawPage)
		if err != nil {
			return nil, err
		}
		pages[page] = i
	}
	return pages, nil
}

func checkOrder(printing map[int]int, orderingsByFirst map[int]OrderingByFirst) (OrderingByFirst, bool) {
	for page, pos := range printing {
		if ordering, ok := orderingsByFirst[page]; ok {
			for _, laterPage := range ordering.later {
				if secondPos, ok := printing[laterPage]; ok {
					if secondPos < pos {
						return ordering, false
					}
				}
			}
		}
	}
	return OrderingByFirst{}, true
}

func reorder(originalPrinting map[int]int, orderingsBySecond map[int]OrderingBySecond) []int {
	graph := map[int][]int{}
	for page := range originalPrinting {
		pagesBefore := []int{}
		ordering, ok := orderingsBySecond[page]
		if !ok {
			continue
		}
		for _, earlierPage := range ordering.earlier {
			if _, ok := originalPrinting[earlierPage]; ok {
				pagesBefore = append(pagesBefore, earlierPage)
			}
		}
		graph[page] = pagesBefore
	}
	printing := []int{}
	placed := map[int]struct{}{}
	for len(printing) < len(graph) {
		for page, pagesBefore := range graph {
			if _, ok := placed[page]; ok {
				continue
			}
			unplacedFound := false
			for _, pageToCheck := range pagesBefore {
				if _, ok := placed[pageToCheck]; !ok {
					unplacedFound = true
					break
				}
			}
			if !unplacedFound {
				printing = append(printing, page)
				placed[page] = struct{}{}
			}
		}
	}
	return printing
}

func middlePage(printing []int) int {
	return printing[len(printing)/2]
}
