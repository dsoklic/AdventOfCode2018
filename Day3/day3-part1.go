package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type coordinate struct {
	x int
	y int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("Day3/puzzle3.txt")
	check(err)
	scanner := bufio.NewScanner(f)

	coverage := map[coordinate]int{}
	for scanner.Scan() {
		line := scanner.Text()
		params := getParams(`#(?P<ID>\d+) @ (?P<x>\d+),(?P<y>\d+): (?P<w>\d+)x(?P<h>\d+)`, line)

		for x := params["x"]; x < params["x"]+params["w"]; x++ {
			for y := params["y"]; y < params["y"]+params["h"]; y++ {
				coverage[coordinate{x, y}]++
			}
		}
	}

	sum := 0
	for _, numclaims := range coverage {
		if numclaims > 1 {
			sum++
		}
	}
	fmt.Println(sum)
}

/**
 * Parses url with the given regular expression and returns the
 * group values defined in the expression.
 *
 */
func getParams(regEx, url string) (paramsMap map[string]int) {

	var compRegEx = regexp.MustCompile(regEx)
	match := compRegEx.FindStringSubmatch(url)

	paramsMap = make(map[string]int)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			intComponent, err := strconv.Atoi(match[i])
			check(err)
			paramsMap[name] = intComponent
		}
	}
	return
}
