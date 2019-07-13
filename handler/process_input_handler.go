package handler

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

// ProcessInput Process the input file into a slice.
func ProcessInput(handler io.Reader) ([][]string, error) {
	scanner := bufio.NewScanner(handler)
	scanner.Split(bufio.ScanLines)
	var input [][]string
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, strings.Split(line, " "))
	}
	if err := scanner.Err(); err != nil {
		return nil, errors.New("unable to read the input")
	}
	return input, nil
}
