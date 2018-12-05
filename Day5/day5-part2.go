package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"unicode"
)

func main() {
	input := readFile("Day5/puzzle5.txt")

	minLength := math.MaxInt32
	for i := rune('a'); i <= rune('z'); i++ {
		inputCopy := input
		inputCopy = strings.Replace(inputCopy, string(i), "", -1)
		inputCopy = strings.Replace(inputCopy, string(unicode.ToUpper(i)), "", -1)
		inputCopy = react(inputCopy)

		if len(inputCopy) < minLength {
			minLength = len(inputCopy)
		}
	}

	fmt.Println(minLength)
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

func react(input string) string {
	inputLetters := []rune(input)

	for {
		lettersRemoved := false

		for i := 0; i < len(inputLetters)-1; i++ {
			rune1 := inputLetters[i]
			rune2 := inputLetters[i+1]

			if cancelingRunes(rune1, rune2) {
				// Don't add them to the new slice.
				inputLetters = removeTwoElements(inputLetters, i)
				lettersRemoved = true
				break
			}
		}

		if !lettersRemoved {
			break
		}
	}

	return string(inputLetters)
}
