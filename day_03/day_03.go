package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// wire struct contains the direction and distance values of a wire.
type wire struct {
	direction []string
	distance  []int
}

// readInput converts the input file to a slice of wire structs.
func readInput(f string) ([]wire, error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	var wires []wire
	reader := csv.NewReader(bufio.NewReader(file))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		var d []string
		var v []int
		for i, _ := range line {
			d = append(d, string(line[i][0]))
			val, err := strconv.Atoi(line[i][1:])
			if err != nil {
				fmt.Println(err)
			}
			v = append(v, val)
		}
		wires = append(wires, wire{
			direction: d,
			distance:  v,
		})
	}
	return wires, nil
}

func main() {
	w, err := readInput("input.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(w[0].direction)
	fmt.Println(w[0].distance)
}
