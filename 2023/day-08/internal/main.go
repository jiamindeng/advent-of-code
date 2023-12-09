package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.ReadFile("./input.txt")
	resOne := partOne(string(f))
	fmt.Println(resOne)
	resTwo := partTwo(string(f))
	fmt.Println(resTwo)
}

func parse(input string) (map[string][]string, string) {
	res := map[string][]string{}
	lines := strings.Split(input, "\n")
	dirs := lines[0]
	for _, line := range lines[2:] {
		var left string
		var right string
		var key string
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, "(", "")
		line = strings.ReplaceAll(line, ")", "")

		fmt.Sscanf(line, "%s = %s %s", &key, &left, &right)
		res[key] = []string{left, right}
	}

	return res, dirs
}

func partOne(input string) int {
	// count := 0
	// d := 0
	// adj, dirs := parse(input)
	// curr := "AAA"
	// for curr != "ZZZ" {
	// 	if string(dirs[d]) == "L" {
	// 		curr = adj[curr][0]
	// 	} else if string(dirs[d]) == "R" {
	// 		curr = adj[curr][1]
	// 	}
	// 	d = (d + 1) % len(dirs)
	// 	count++
	// }
	return 0
}

func partTwo(input string) int {
	count := 0
	d := 0
	adj, dirs := parse(input)
	currStarts := []string{}
	for k := range adj {
		char := string(k[len(k)-1])
		if char == "A" {
			currStarts = append(currStarts, k)
		}
	}
	numSteps := make([]int, len(currStarts))
	done := map[string]bool{}
	for len(done) != len(currStarts) {
		for i, currStart := range currStarts {
			if !done[currStart] {
				if string(dirs[d]) == "L" {
					currStart = adj[currStart][0]
				} else if string(dirs[d]) == "R" {
					currStart = adj[currStart][1]
				}
				currStarts[i] = currStart
				if string(currStart[len(currStart)-1]) == "Z" {
					done[currStart] = true
					numSteps[i] = count + 1
				}
			}

		}
		d = (d + 1) % len(dirs)
		count++
	}
	return LCM(numSteps[0], numSteps[1], numSteps[2:]...)
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
