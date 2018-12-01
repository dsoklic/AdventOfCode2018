package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("puzzle1.txt")
	check(err)

	scanner := bufio.NewScanner(f)

	var set []int
	var lines []int
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		i, err := strconv.Atoi(line)
		check(err)
		lines = append(lines, i)
	}

	found := false

	for !found {
		for _, i := range lines {
			check(err)

			sum += i

			if freqInSlice(sum, set) {
				fmt.Println("The freq " + strconv.Itoa(sum) + " has repeated")
				found = true
				break
			}

			set = append(set, sum)
		}
	}

}

func freqInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
