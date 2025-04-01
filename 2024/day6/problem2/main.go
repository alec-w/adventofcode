package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/alec-w/adventofcode/2024/internal/load"
)

const OBSTACLE = '#'
const CLEAR = '.'
const GUARD = '^'

type Position struct {
	row    int
	column int
}

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
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
	labMap, guardPos := linesToMap(lines)
	firstPos := guardPos
	direction := UP
	positions := map[Position]Direction{}
	offEdge := false
	for !offEdge {
		positions[guardPos] = direction
		guardPos, direction, offEdge = move(labMap, guardPos, direction)
	}
	potentialObstablePos := map[Position]struct{}{}
	for position := range positions {
		if position == firstPos {
			continue
		}
		labMap[position.row][position.column] = true
		positions := map[Position]Direction{}
		guardPos := firstPos
		direction := UP
		offEdge := false
		for !offEdge {
			positions[guardPos] = direction
			guardPos, direction, offEdge = move(labMap, guardPos, direction)
			if previousDirection, ok := positions[guardPos]; ok && previousDirection == direction {
				potentialObstablePos[position] = struct{}{}
				break
			}
		}
		labMap[position.row][position.column] = false
	}
	return len(potentialObstablePos), nil
}

func linesToMap(lines []string) ([][]bool, Position) {
	labMap := [][]bool{}
	var guardPos Position
	for i, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		row, potentialGuardPos := parseLine(line)
		labMap = append(labMap, row)
		if potentialGuardPos != -1 {
			guardPos = Position{row: i, column: potentialGuardPos}
		}
	}
	return labMap, guardPos
}

func parseLine(line string) ([]bool, int) {
	row := make([]bool, len(line))
	guardPos := -1
	for i, char := range line {
		row[i] = char == OBSTACLE
		if char == GUARD {
			guardPos = i
		}
	}
	return row, guardPos
}

func move(labMap [][]bool, guardPos Position, direction Direction) (Position, Direction, bool) {
	var newPos Position
	switch direction {
	case UP:
		newPos = Position{row: guardPos.row - 1, column: guardPos.column}
	case DOWN:
		newPos = Position{row: guardPos.row + 1, column: guardPos.column}
	case LEFT:
		newPos = Position{row: guardPos.row, column: guardPos.column - 1}
	case RIGHT:
		newPos = Position{row: guardPos.row, column: guardPos.column + 1}
	}
	if newPos.row == -1 || newPos.row == len(labMap) || newPos.column == -1 || newPos.column >= len(labMap[newPos.row]) {
		return Position{}, Direction(0), true
	}
	if labMap[newPos.row][newPos.column] {
		var newDirection Direction
		switch direction {
		case UP:
			newDirection = RIGHT
		case RIGHT:
			newDirection = DOWN
		case DOWN:
			newDirection = LEFT
		case LEFT:
			newDirection = UP
		}
		return move(labMap, guardPos, newDirection)
	}
	return newPos, direction, false
}

func checkObstable(labMap [][]bool, positions map[Position]Direction, firstPos, currentPos Position, direction Direction) (Position, bool) {
	var potentialPos Position
	var checkPos Position
	var checkDirection Direction
	switch direction {
	case UP:
		potentialPos = Position{row: currentPos.row - 1, column: currentPos.column}
		checkPos = Position{row: currentPos.row, column: currentPos.column + 1}
		checkDirection = RIGHT
	case DOWN:
		potentialPos = Position{row: currentPos.row + 1, column: currentPos.column}
		checkPos = Position{row: currentPos.row, column: currentPos.column - 1}
		checkDirection = LEFT
	case LEFT:
		potentialPos = Position{row: currentPos.row, column: currentPos.column + 1}
		checkPos = Position{row: currentPos.row - 1, column: currentPos.column}
		checkDirection = UP
	case RIGHT:
		potentialPos = Position{row: currentPos.row, column: currentPos.column - 1}
		checkPos = Position{row: currentPos.row + 1, column: currentPos.column}
		checkDirection = DOWN
	}
	if potentialPos == firstPos {
		return potentialPos, false
	}
	if potentialPos.row < 0 || potentialPos.row >= len(labMap) || potentialPos.column < 0 || potentialPos.column >= len(labMap[potentialPos.row]) {
		return potentialPos, false
	}
	previousDirection, ok := positions[checkPos]
	return potentialPos, ok && previousDirection == checkDirection
}
