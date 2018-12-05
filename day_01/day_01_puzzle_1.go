package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Frequency: ", calcFrequency())
}

func calcFrequency() (frequency int64) {
	f, _ := os.Open("day_01_puzzle_input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		i, _ := strconv.ParseInt(line, 10, 64)
		frequency += i
	}

	return
}
