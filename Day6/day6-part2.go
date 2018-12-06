package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputText := readFile("Day6/puzzle6.txt")
	var input [][]int

	for _, line := range inputText {
		parts := strings.Split(line, ", ")
		x, err := strconv.Atoi(parts[0])
		check(err)
		y, err := strconv.Atoi(parts[1])
		check(err)

		input = append(input, []int{x, y})
	}

	maxX, maxY := max(input)

	regionSize := 0
	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			sumOfDistances := 0
			for _, point := range input {
				pointX := point[0]
				pointY := point[1]
				distance := intAbs(pointX-x) + intAbs(pointY-y)
				sumOfDistances += distance
			}

			if sumOfDistances < 10000 {
				regionSize++
			}
		}
	}

	fmt.Println(regionSize)
}

func max(input [][]int) (int, int) {
	maxX := 0
	maxY := 0

	for _, pair := range input {
		x := pair[0]
		y := pair[1]

		if x > maxX {
			maxX = x
		}

		if y > maxY {
			maxY = y
		}
	}

	return maxX, maxY
}

func intAbs(n int) int {
	if n < 0 {
		return -n
	}
	return n
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}
