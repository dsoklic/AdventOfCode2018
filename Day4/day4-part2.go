package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type KeyValue struct {
	time      string
	paramsMap map[string]string
}

type KeyValueArray []KeyValue

func (a KeyValueArray) Len() int           { return len(a) }
func (a KeyValueArray) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a KeyValueArray) Less(i, j int) bool { return a[i].time < a[j].time }

func main() {
	input := readFile("Day4/puzzle4.txt")

	var sortedArray KeyValueArray

	regexPattern := `\[(?P<time>(?P<year>\d+)-(?P<month>\d+)-(?P<day>\d+) (?P<h>\d+):(?P<min>\d+))\] (?P<action>[^\n]+)`
	for _, line := range input {
		params := getParams(regexPattern, line)

		sortedArray = append(sortedArray, KeyValue{params["time"], params})
	}

	sort.Sort(KeyValueArray(sortedArray))

	// The array is now in chronological order.

	var currentID int
	var minuteFallingAsleep int
	sleepingTimes := map[int]int{}
	sleepingPeriods := map[int][][]int{}

	for _, data := range sortedArray {
		action := data.paramsMap["action"]

		if strings.HasPrefix(action, "Guard") {
			// Guard assumed duty.
			id, err := strconv.Atoi(getParams(`Guard #(?P<id>\d+) begins shift`, action)["id"])
			check(err)
			currentID = id
		} else if strings.HasPrefix(action, "falls") {
			minute, err := strconv.Atoi(data.paramsMap["min"])
			check(err)
			minuteFallingAsleep = minute
		} else if strings.HasPrefix(action, "wakes") {
			minute, err := strconv.Atoi(data.paramsMap["min"])
			check(err)
			sleepingTimes[currentID] += (minute - minuteFallingAsleep)
			sleepingPeriods[currentID] = append(sleepingPeriods[currentID], []int{minuteFallingAsleep, minute})
		}
	}

	var guardID int
	maxTimes := 0
	var maxMinute int
	for id, times := range sleepingPeriods {
		bestMinutes := map[int]int{}
		for _, time := range times {
			for i := time[0]; i < time[1]; i++ {
				bestMinutes[i]++
			}
		}

		currentMaxMin := findMaxId(bestMinutes)
		currentMaxTimes := bestMinutes[currentMaxMin]

		if currentMaxTimes > maxTimes {
			maxMinute = currentMaxMin
			maxTimes = currentMaxTimes
			guardID = id
		}
	}

	fmt.Printf("Best is guard %d at min %d\n", guardID, maxMinute)
}

/**
 * Parses url with the given regular expression and returns the
 * group values defined in the expression.
 *
 */
func getParams(regEx, url string) (paramsMap map[string]string) {

	var compRegEx = regexp.MustCompile(regEx)
	match := compRegEx.FindStringSubmatch(url)

	paramsMap = make(map[string]string)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func findMaxId(input map[int]int) int {
	var maxId int
	foundMax := 0

	for k, v := range input {
		if v > foundMax {
			maxId = k
			foundMax = v
		}
	}

	return maxId
}

func readFile(path string) []string {
	f, err := os.Open(path)
	check(err)
	scanner := bufio.NewScanner(f)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
