package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type greaterNap struct {
	filename string
	nap
}

type nap struct {
	minute                            int
	second                            int
	guardID                           int
	greaterMinuteGuardNapCountGuardID int
	greaterMinuteGuardNapCountMinute  int
}

func newGreaterNap(filename string) *greaterNap {
	gn := &greaterNap{}
	gn.filename = filename

	return gn
}

func (gn *greaterNap) processData() error {
	data, err := gn.loadData()
	if err != nil {
		fmt.Println(err)
		return err
	}

	var napAmount map[int]int

	lastGuardID := 0
	lastAsleepMinute := 0
	lastWakeupMinute := 0
	napAmount = make(map[int]int)
	greatherNapAmount := 0

	for _, register := range data {
		k := register[:19]
		v := register[19:]
		if v[0] == 'G' {
			tokens := strings.Split(v, " ")
			lastGuardID, _ = strconv.Atoi(tokens[1][1:])
		} else if v[0] == 'f' {
			lastAsleepMinute, _ = strconv.Atoi(k[15:17])
		} else if v[0] == 'w' {
			lastWakeupMinute, _ = strconv.Atoi(k[15:17])
			nap := lastWakeupMinute - lastAsleepMinute
			napAmount[lastGuardID] += nap

			if napAmount[lastGuardID] > greatherNapAmount {
				gn.guardID = lastGuardID
				greatherNapAmount = napAmount[lastGuardID]
			}
		}
	}

	var minuteNap [60]int

	greatherMinuteNap := 0
	greatherNapValue := 0

	var minuteGuardNapCount map[string]int

	minuteGuardNapCount = make(map[string]int)
	greatherMinuteGuardNapCountValue := 0

	lastGuardID = 0
	lastAsleepMinute = 0

	for _, register := range data {
		k := register[:19]
		v := register[19:]
		if v[0] == 'G' {
			tokens := strings.Split(v, " ")
			lastGuardID, _ = strconv.Atoi(tokens[1][1:])
		} else if v[0] == 'f' {
			lastAsleepMinute, _ = strconv.Atoi(k[15:17])
		} else if v[0] == 'w' {
			lastWakeupMinute, _ = strconv.Atoi(k[15:17])
			for i := lastAsleepMinute; i < lastWakeupMinute; i++ {
				key := fmt.Sprintf("%02d-%04d", i, lastGuardID)
				minuteGuardNapCount[key]++

				if minuteGuardNapCount[key] > greatherMinuteGuardNapCountValue {
					greatherMinuteGuardNapCountValue = minuteGuardNapCount[key]
					gn.greaterMinuteGuardNapCountGuardID = lastGuardID
					gn.greaterMinuteGuardNapCountMinute = i
				}

				if lastGuardID != gn.guardID {
					continue
				}

				minuteNap[i]++
				if minuteNap[i] > greatherNapValue {
					greatherNapValue = minuteNap[i]
					greatherMinuteNap = i
				}
			}
		}
	}

	gn.minute = greatherMinuteNap

	return nil
}

func (gn *greaterNap) loadData() (data []string, err error) {
	file, err := os.Open(gn.filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var unorderData map[string]string
	var keys []string

	unorderData = make(map[string]string)

	for scanner.Scan() {
		line := scanner.Text()
		unorderData[line[0:18]] = line
	}

	for k := range unorderData {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		data = append(data, unorderData[k])
	}

	return data, nil
}
