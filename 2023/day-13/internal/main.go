package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Pattern [][]string

func main() {
	f, _ := os.ReadFile("./input.txt")
	resOne := partOne(string(f), true)
	fmt.Println(resOne)
	resTwo := partTwo(string(f))
	fmt.Println(resTwo)
}

func parse(input string) []Pattern {
	parts := strings.Split(input, "\n\n")
	parsed := make([]Pattern, len(parts))

	for i, part := range parts {
		lines := strings.Split(part, "\n")
		pattern := Pattern{}
		for _, line := range lines {
			row := []string{}
			for _, chr := range line {
				row = append(row, string(chr))
			}
			pattern = append(pattern, row)
		}
		parsed[i] = pattern
	}

	return parsed
}

func partOne(input string, partOne bool) int {
	parsed := parse(input)
	total := 0

	for _, pattern := range parsed {
		h := getHorizontalLine(pattern, false)
		if h != -1 {
			total += 100 * h
		}
		v := getVerticalLine(pattern, false)
		if v != -1 {
			total += v
		}
	}

	return total
}

func diff(a []string, b []string) int {
	res := 0
	for i := range a {
		if a[i] != b[i] {
			res++
		}
	}
	return res
}
func getHorizontalLineOne(pattern Pattern) int {
	for idx := range pattern[:len(pattern)-1] {
		sym := true
		l := idx
		r := idx + 1
		for l >= 0 && r < len(pattern) {
			if !slices.Equal(pattern[l], pattern[r]) {
				sym = false
				break
			}
			l--
			r++
		}

		if sym {
			return (idx + 1)
		}
	}

	return -1
}

func getHorizontalLineTwo(pattern Pattern) int {
	for idx := range pattern[:len(pattern)-1] {
		diffs := 0
		l := idx
		r := idx + 1
		for l >= 0 && r < len(pattern) {
			diffs += diff(pattern[l], pattern[r])
			l--
			r++
		}

		if diffs == 1 {
			return idx + 1
		}

	}

	return -1
}

func getHorizontalLine(pattern Pattern, partTwo bool) int {
	if partTwo {
		return getHorizontalLineTwo(pattern)
	} else {
		return getHorizontalLineOne(pattern)
	}
}

func getVerticalLine(pattern Pattern, partTwo bool) int {
	return getHorizontalLine(transpose(pattern), partTwo)
}

func transpose(slice [][]string) [][]string {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func partTwo(input string) int {
	parsed := parse(input)
	total := 0

	for _, pattern := range parsed {
		h := getHorizontalLine(pattern, true)
		if h != -1 {
			total += 100 * h
		}
		v := getVerticalLine(pattern, true)
		if v != -1 {
			total += v
		}
	}

	return total
}
