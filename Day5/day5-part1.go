package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	input := readFile("Day5/puzzle5.txt")

	inputLetters := []rune(input)
	// fmt.Println(string(inputLetters))

	for {
		lettersRemoved := false

		for i := 0; i < len(inputLetters)-1; i++ {
			rune1 := inputLetters[i]
			rune2 := inputLetters[i+1]

			if cancelingRunes(rune1, rune2) {
				// Don't add them to the new slice.
				// fmt.Printf("Removing %s and %s\n", string(rune1), string(rune2))
				inputLetters = removeTwoElements(inputLetters, i)
				lettersRemoved = true
				break
			}
		}

		// fmt.Println(string(inputLetters))

		if !lettersRemoved {
			break
		}
	}

	// fmt.Println(string(inputLetters))
	fmt.Println(len(inputLetters))
}

func cancelingRunes(rune1, rune2 rune) bool {
	return rune1 != rune2 && unicode.ToLower(rune1) == unicode.ToLower(rune2)
}

func removeTwoElements(a []rune, i int) []rune {
	return a[:i+copy(a[i:], a[i+2:])]
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(path string) string {
	f, err := os.Open(path)
	check(err)
	scanner := bufio.NewScanner(f)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return strings.Join(lines, "")
}
