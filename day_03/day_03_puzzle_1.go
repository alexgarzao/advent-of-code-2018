package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	mappedFabric := mapFabric()
	fmt.Println("Overlaped inches: ", mappedFabric.overlaped())
	fmt.Println("Good claim: ", mappedFabric.smallerGoodClaim())
}

func mapFabric() *fabric {
	file, err := os.Open("day_03_puzzle_input.txt")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	f := newFabric()

	for scanner.Scan() {
		line := scanner.Text()
		claim := lineToClaim(line)
		f.markInches(claim)
	}

	return f
}
