package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/deckarep/golang-set"
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

	regionSizes := map[int]int{}
	infiniteRegions := mapset.NewSet()

	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			closest, tied := findClosest(x, y, input)

			if !tied {
				if x == 0 || y == 0 || x == maxX || y == maxY {
					// Infinite region, ignore it.
					infiniteRegions.Add(closest)
				}

				if !infiniteRegions.Contains(closest) {
					regionSizes[closest]++
				}
			}
		}
	}

	fmt.Println(regionSizes)

	biggestRegion := 0
	for _, size := range regionSizes {
		if size > biggestRegion {
			biggestRegion = size
		}
	}

	fmt.Println(biggestRegion)
}

func findClosest(x, y int, points [][]int) (int, bool) {
	var closestIndex int
	minDistance := math.MaxInt32

	for i, point := range points {
		pointX := point[0]
		pointY := point[1]

		distance := intAbs(pointX-x) + intAbs(pointY-y)

		if distance < minDistance {
			minDistance = distance
			closestIndex = i
		} else if distance == minDistance {
			// Tied.
			return 0, true
		}
	}

	return closestIndex, false
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
