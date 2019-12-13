package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var fuel int64 = 0

// calcFuel recursively calculates the fuel usage.
func calcFuel(n int64) {
	reduced := (n / 3) - 2
	if reduced <= 0 {
		return
	}
	fuel += reduced
	calcFuel(reduced)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.ParseInt(line, 0, 64)
		if err != nil {
			log.Fatal(err)
		}
		calcFuel(num)
	}
	fmt.Println(fuel)
}
