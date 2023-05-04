package internal

import (
	"strconv"
	"strings"
)

type Assignment struct {
	Min uint64
	Max uint64
}

func AssignmentFromString(input string) (Assignment, error) {
	sections := strings.Split(input, "-")
	min, err := strconv.ParseUint(sections[0], 10, 0)
	if err != nil {
		return Assignment{}, err
	}
	max, err := strconv.ParseUint(sections[1], 10, 0)
	if err != nil {
		return Assignment{}, err
	}
	return Assignment{Min: min, Max: max}, nil
}
