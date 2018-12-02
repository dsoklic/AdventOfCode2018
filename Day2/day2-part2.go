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
	f, err := os.Open("Day2/puzzle2.txt")
	check(err)

	scanner := bufio.NewScanner(f)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	firstBox, secondBox := findBoxes(lines)
	fmt.Println("Box 1: " + firstBox)
	fmt.Println("Box 2: " + secondBox)

	fmt.Print("Answer: ")
	for i := range firstBox {
		letter1 := firstBox[i]
		letter2 := secondBox[i]

		if letter1 == letter2 {
			fmt.Print(string(letter1))
		}
	}

	fmt.Println()
}

func findBoxes(lines []string) (string, string) {
	for i := 0; i < len(lines)-1; i++ {
		for j := i + 1; j < len(lines); j++ {
			line1 := lines[i]
			line2 := lines[j]

			numDiffering := 0
			for k := 0; k < len(line1); k++ {
				letter1 := line1[k]
				letter2 := line2[k]

				if letter1 != letter2 {
					numDiffering++
				}
			}

			if numDiffering == 1 {
				return line1, line2
			}
		}
	}

	return "", ""
}
