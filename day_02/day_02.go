package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// sliceAtoi converts a list of strings to ints.
func sliceAtoi(sa []string) ([]int, error) {
	var sint = []int{}
	for _, i := range sa {
		j, err := strconv.Atoi(i)
		if err != nil {
			return nil, err
		}
		sint = append(sint, j)
	}
	return sint, nil
}

// prepInput reads in a txt file and converts to slice of ints.
func prepInput(f string) ([]int, error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, ",")
		d, err := sliceAtoi(s)
		if err != nil {
			fmt.Println(err)
		}
		return d, nil
	}
	return nil, nil
}

func main() {
	s, err := prepInput("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	s[1] = 12
	s[2] = 2
	i := 0
	for s[i] != 99 {
		a := s[i+1]
		b := s[i+2]
		c := s[i+3]
		switch s[i] {
		case 1:
			s[c] = s[a] + s[b]
		case 2:
			s[c] = s[a] * s[b]
		}
		i += 4
	}
	fmt.Println(s[0])
}
