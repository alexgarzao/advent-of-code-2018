package main

import (
	"fmt"
)

func main() {
	testSample()

	processInput()
}

func testSample() {
	greaterNap := newGreaterNap("input_sample.txt")
	greaterNap.processData()

	if greaterNap.guardID*greaterNap.minute != 240 {
		fmt.Printf("#### %d != 240 ####\n", greaterNap.guardID*greaterNap.minute)
	}

	if greaterNap.greaterMinuteGuardNapCountGuardID*greaterNap.greaterMinuteGuardNapCountMinute != 4455 {
		fmt.Printf("#### %d != 4455 ####\n", greaterNap.greaterMinuteGuardNapCountGuardID*greaterNap.greaterMinuteGuardNapCountMinute)
	}
}

func processInput() {
	greaterNap := newGreaterNap("input.txt")
	greaterNap.processData()

	fmt.Println("R1:", greaterNap.guardID*greaterNap.minute)
	fmt.Println("R2:", greaterNap.greaterMinuteGuardNapCountGuardID*greaterNap.greaterMinuteGuardNapCountMinute)
}
