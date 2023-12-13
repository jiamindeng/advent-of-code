package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.ReadFile("./input.txt")
	resOne := partOne(string(f), true)
	fmt.Println(resOne)
	resTwo := partTwo(string(f))
	fmt.Println(resTwo)
}

func parse(input string, partOne bool) ([][]string, [][]int) {
	lines := strings.Split(input, "\n")
	parsed := make([][]string, len(lines))
	cc := [][]int{}
	lim := 5
	if partOne {
		lim = 1
	}

	for r, line := range lines {
		if parsed[r] == nil {
			parsed[r] = []string{}
		}
		groups := strings.Split(line, " ")
		subP := []string{}
		subC := []int{}
		for i := 0; i < lim; i++ {
			for _, chr := range groups[0] {
				subP = append(subP, string(chr))
			}
			counts := strings.Split(groups[1], ",")
			for _, count := range counts {
				c, err := strconv.Atoi(count)
				if err != nil {
					panic(err)
				}
				subC = append(subC, c)
			}
			if i != lim-1 {
				subP = append(subP, "?")
			}
		}
		parsed[r] = subP
		cc = append(cc, subC)
	}
	return parsed, cc
}

var cache = map[string]int{}

func getKey(row []string, want []int) string {
	res := ""

	res += strings.Join(row, "") + "|"
	nums := []string{}
	for _, w := range want {
		nums = append(nums, fmt.Sprintf("%d", w))
	}

	return res + strings.Join(nums, ",")
}

func countCombos(s []string, nums []int) int {
	k := getKey(s, nums)
	if _, ok := cache[k]; ok {
		return cache[k]
	}

	if len(nums) == 0 {
		if !slices.Contains(s, "#") {
			return 1
		} else {
			return 0
		}

	}

	size := nums[0]
	total := 0

	for i := range s {
		if i+size <= len(s) && !slices.Contains(s[i:i+size], ".") && (i == 0 || s[i-1] != "#") && (i+size == len(s) || s[i+size] != "#") {
			if i+size+1 >= len(s) {
				total += countCombos([]string{}, nums[1:])
			} else {
				total += countCombos(s[i+size+1:], nums[1:])
			}
		}

		if s[i] == "#" {
			break
		}
	}

	cache[k] = total
	return total
}

func partOne(input string, partOne bool) int {
	parsed, cc := parse(input, true)
	total := 0
	for i := range parsed {
		total += countCombos(parsed[i], cc[i])

	}
	return total
}

func partTwo(input string) int {
	parsed, cc := parse(input, false)
	total := 0
	for i := range parsed {
		total += countCombos(parsed[i], cc[i])

	}
	return total
}
