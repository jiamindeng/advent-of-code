package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.ReadFile("./input.txt")
	resOne := partOne(string(f))
	fmt.Println(resOne)
	resTwo := partTwo(string(f))
	fmt.Println(resTwo)
}

func parse(input string) [][]int {
	res := [][]int{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		row := []int{}
		fields := strings.Fields(line)
		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				panic(err)
			}
			row = append(row, num)
		}
		res = append(res, row)
	}
	return res
}

func partOne(input string) int {
	lines := parse(input)
	res := 0
	for _, line := range lines {
		res += findLastHistory(line)
	}
	return res

}

func findLastHistory(line []int) int {
	if allZeros(line) {
		return 0
	}
	diffs := []int{}
	for i := 0; i < len(line)-1; i++ {
		diff := line[i+1] - line[i]
		diffs = append(diffs, diff)
	}
	return findLastHistory(diffs) + line[len(line)-1]
}

func findFirstHistory(line []int) int {
	if allZeros(line) {
		return 0
	}
	diffs := []int{}
	for i := 0; i < len(line)-1; i++ {
		diff := line[i+1] - line[i]
		diffs = append(diffs, diff)
	}

	return line[0] - findFirstHistory(diffs)
}

func allZeros(line []int) bool {
	for _, num := range line {
		if num != 0 {
			return false
		}
	}
	return true
}

func partTwo(input string) int {
	lines := parse(input)
	res := 0
	for _, line := range lines {
		res += findFirstHistory(line)
	}
	return res
}
