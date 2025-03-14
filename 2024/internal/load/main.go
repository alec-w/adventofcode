package load

import (
	"io"
	"os"
	"strings"
)

func FileToLines(fileName string) ([]string, error) {
	input, err := FileToString(fileName)
	if err != nil {
		return nil, err
	}
	return strings.Split(input, "\n"), nil
}

func FileToString(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	var input strings.Builder
	if _, err := io.Copy(&input, file); err != nil {
		return "", err
	}
	return input.String(), nil
}
