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

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		i, err := strconv.Atoi(line)
		check(err)

		sum += i
	}

	fmt.Println(sum)

}
