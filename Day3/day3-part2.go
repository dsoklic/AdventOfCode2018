package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type rectangle struct {
	id int
	x  int
	y  int
	w  int
	h  int
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

	rectangles := []rectangle{}

	for scanner.Scan() {
		line := scanner.Text()
		params := getParams(`#(?P<ID>\d+) @ (?P<x>\d+),(?P<y>\d+): (?P<w>\d+)x(?P<h>\d+)`, line)

		rectangles = append(rectangles, rectangle{params["ID"], params["x"], params["y"], params["w"], params["h"]})
	}

	var recWithNoOverlap int

	// Go through all combinations, checking if we find a rectangle that doesn't overlap with anything.
	for _, rec1 := range rectangles {
		hasOverlap := false
		for _, rec2 := range rectangles {

			if rec1 == rec2 {
				continue
			}

			if overlap(rec1, rec2) {
				hasOverlap = true
				break
			}
		}

		if !hasOverlap {
			recWithNoOverlap = rec1.id
			break
		}
	}

	fmt.Println(recWithNoOverlap)

}

func valueInRange(value int, min int, max int) bool {
	return (value >= min) && (value <= max)
}

/**
 * Checks if the two rectangles overlap.
 */
func overlap(A rectangle, B rectangle) bool {
	xOverlap := valueInRange(A.x, B.x, B.x+B.w) ||
		valueInRange(B.x, A.x, A.x+A.w)

	yOverlap := valueInRange(A.y, B.y, B.y+B.h) ||
		valueInRange(B.y, A.y, A.y+A.h)

	return xOverlap && yOverlap
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
