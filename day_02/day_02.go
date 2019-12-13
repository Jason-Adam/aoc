package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// sliceAtoi converts a slice of strings to ints.
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

// makeRange creates a slice of ints from min to max.
func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func main() {
	s, err := prepInput("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	noun := makeRange(0, 99)
	verb := makeRange(0, 99)
	for _, n := range noun {
		for _, v := range verb {
			sc := make([]int, len(s))
			copy(sc, s)
			sc[1] = n
			sc[2] = v
			i := 0
			for sc[i] != 99 {
				a := sc[i+1]
				b := sc[i+2]
				c := sc[i+3]
				switch sc[i] {
				case 1:
					sc[c] = sc[a] + sc[b]
				case 2:
					sc[c] = sc[a] * sc[b]
				}
				i += 4
			}
			if sc[0] == 19690720 {
				fmt.Println(100*n + v)
			}
		}
	}
}
