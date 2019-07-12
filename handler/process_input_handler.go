package handler

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// ProcessInput Process the input file into a slice.
func ProcessInput(filename string) (input [][]string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		line := scanner.Text()
		input = append(input, strings.Split(line, " "))
	}
	return
}
