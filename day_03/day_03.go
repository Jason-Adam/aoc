package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// readInput reads in a txt file and converts to slice of ints.
func readInput(f string) ([]string, error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, ",")
		return s, nil
	}
	return nil, nil
}

func main() {
	f, err := readInput("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f)
}
