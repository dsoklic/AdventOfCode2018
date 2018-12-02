package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	boxesWithTwo := 0
	boxesWithThree := 0

	f, err := os.Open("Day2/puzzle2.txt")
	check(err)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		letterCounter := map[rune]int{}

		// Go through the box name letter by letter and count how many times you see a letter.
		for _, letter := range scanner.Text() {
			letterCounter[letter]++
		}

		// Now check how many letters repeat twice or thrice.
		hasBoxWithTwo := false
		hasBoxWithThree := false
		for _, letterRepetitions := range letterCounter {
			if letterRepetitions == 2 {
				hasBoxWithTwo = true
			} else if letterRepetitions == 3 {
				hasBoxWithThree = true
			}
		}

		if hasBoxWithTwo {
			boxesWithTwo++
		}
		if hasBoxWithThree {
			boxesWithThree++
		}
	}

	fmt.Printf("Appear twice: %d\n", boxesWithTwo)
	fmt.Printf("Appear thrice: %d\n", boxesWithThree)
	fmt.Printf("Checksum is %d\n", boxesWithTwo*boxesWithThree)
}
