package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Box: ", mostSimilarBoxes())
}

func mostSimilarBoxes() string {
	file, err := os.Open("day_02_puzzle_input.txt")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i := 0; i < len(lines)-1; i++ {
		for j := i + 1; j < len(lines); j++ {
			if lines[i] == lines[j] {
				continue
			}
			diffNumber, equalChars := calcDiff(lines[i], lines[j])
			if diffNumber == 1 {
				return equalChars
			}
		}
	}

	return ""
}

func calcDiff(s1 string, s2 string) (diffNumber int, equalChars string) {
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diffNumber++
		} else {
			equalChars += string(s1[i])
		}
	}
	return
}
