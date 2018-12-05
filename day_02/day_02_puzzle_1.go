package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Checksum: ", rudimentaryChecksum())
}

func rudimentaryChecksum() int {
	twoTimes := 0
	threeTimes := 0

	f, _ := os.Open("day_02_puzzle_input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		hasTwoTimes, hasThreeTimes := calcCharTimes(line)
		twoTimes += hasTwoTimes
		threeTimes += hasThreeTimes
	}

	return twoTimes * threeTimes
}

func calcCharTimes(s string) (twoTimes int, threeTimes int) {
	for _, rune := range s {
		numberOf := strings.Count(s, string(rune))
		if numberOf == 2 {
			twoTimes = 1
		} else if numberOf == 3 {
			threeTimes = 1
		}
	}
	return
}
