package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Duplicated frequency: ", duplicatedFrequency())
}

func duplicatedFrequency() (frequency int64) {
	frequencies := map[int64]bool{}
	frequencies[0] = true

	for {
		f, _ := os.Open("day_01_puzzle_input.txt")
		defer f.Close()
		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			line := scanner.Text()
			i, _ := strconv.ParseInt(line, 10, 64)
			frequency += i
			_, ok := frequencies[frequency]
			if ok {
				fmt.Println("Frequency: ", frequency)
				os.Exit(0)
			}
			frequencies[frequency] = true
		}
	}
}
