package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type connection struct {
	from string
	to   string
}

func main() {
	// input := []string{"Step C must be finished before step A can begin.", "Step C must be finished before step F can begin.", "Step A must be finished before step B can begin.", "Step A must be finished before step D can begin.", "Step B must be finished before step E can begin.", "Step D must be finished before step E can begin.", "Step F must be finished before step E can begin."}
	input := readFile("Day7/puzzle7.txt")

	var connections []connection

	for _, line := range input {
		params := getParams(`Step (?P<from>\w) must be finished before step (?P<to>\w) can begin.`, line)
		connections = append(connections, connection{params["from"], params["to"]})
	}

	graphMap := map[string][]string{}
	reverseGraphMap := map[string][]string{}
	for _, pair := range connections {
		graphMap[pair.from] = append(graphMap[pair.from], pair.to)
		reverseGraphMap[pair.to] = append(reverseGraphMap[pair.to], pair.from)
	}

	available := findRootNode(connections)

	sort.Slice(available, func(i, j int) bool {
		return available[i] < available[j]
	})

	order := []string{}

	for len(available) > 0 {
		nextNode := available[0]
		available = available[1:]

		order = append(order, nextNode)

		for _, childOfNextNode := range graphMap[nextNode] {
			// Check if node is available
			if isNodeOpen(childOfNextNode, order, reverseGraphMap) {
				available = append(available, childOfNextNode)
			}
		}

		sort.Slice(available, func(i, j int) bool {
			return available[i] < available[j]
		})
	}

	fmt.Println(strings.Join(order, ""))
}

func isNodeOpen(node string, visited []string, mappings map[string][]string) bool {
	requiredVisited := mappings[node]

	for _, requiredNode := range requiredVisited {
		if !contains(visited, requiredNode) {
			return false
		}
	}

	return true
}

func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

/**
 * Parses url with the given regular expression and returns the
 * group values defined in the expression.
 *
 */
func getParams(regEx, url string) (paramsMap map[string]string) {

	var compRegEx = regexp.MustCompile(regEx)
	match := compRegEx.FindStringSubmatch(url)

	paramsMap = make(map[string]string)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return
}

func findRootNode(connections []connection) []string {
	rootCandidates := []string{}
	disqualified := []string{}

	for _, pair := range connections {
		if !contains(disqualified, pair.to) {
			disqualified = append(disqualified, pair.to)
		}

		if !contains(rootCandidates, pair.from) {
			rootCandidates = append(rootCandidates, pair.from)
		}
	}

	roots := []string{}
	for _, element := range rootCandidates {
		if !contains(disqualified, element) {
			roots = append(roots, element)
		}
	}
	return roots
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
