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

func parse(input string) [][]string {
	parts := strings.Split(input, "\n")
	res := make([][]string, len(parts))

	for p, part := range parts {
		res[p] = make([]string, len(parts[0]))
		for i, chr := range part {
			res[p][i] = string(chr)
		}
	}

	return res
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

func partOne(input string) int {
	parsed := transpose(parse(input))
	counts := make([]int, len(parsed[0]))

	for _, col := range parsed {
		newCol := shiftRocks(col)
		for i, chr := range newCol {
			if chr == "O" {
				counts[i]++
			}
		}

	}

	sum := 0
	for i, count := range counts {
		mul := len(counts) - i
		sum += mul * count

	}
	return sum
}

func rotateMatrix(matrix [][]string) [][]string {
	if len(matrix) == 0 {
		return matrix
	}
	for i := 0; i < len(matrix)/2; i++ {
		top := i
		bottom := len(matrix) - 1 - i
		for j := top; j < bottom; j++ {
			temp := matrix[top][j]
			matrix[top][j] = matrix[j][bottom]
			matrix[j][bottom] = matrix[bottom][bottom-(j-top)]
			matrix[bottom][bottom-(j-top)] = matrix[bottom-(j-top)][top]
			matrix[bottom-(j-top)][top] = temp
		}
	}
	return matrix
}

func partTwo(input string) int {
	cache := map[string][]int{}
	curr := transpose(parse(input))
	det := 0
	start := -1
	end := -1
	big := 1000000000
	for c := 0; c < big; c++ {
		key := getKey(curr)
		v, ok := cache[key]
		if ok {
			start = v[1]
			end = c
			det++
			break
		}

		counts := make([]int, len(curr[0]))

		for i := 0; i < 4; i++ {
			next := make([][]string, 0)
			for _, col := range curr {
				newCol := shiftRocks(col)
				next = append(next, newCol)
			}
			curr = rotateMatrix(next)
		}

		for _, col := range curr {
			for i, chr := range col {
				if chr == "O" {
					counts[i]++
				}
			}
		}
		sum := 0
		for i, count := range counts {
			mul := len(counts) - i
			sum += mul * count
		}

		cache[key] = []int{sum, c}
	}
	for _, v := range cache {
		if v[1] == start+(big-end)%(start-end)-1 {
			return v[0]
		}
	}
	return 0
}

func getKey(m [][]string) string {
	res := ""
	for _, r := range m {
		res += strings.Join(r, ",")
	}
	return res
}

func shiftRocks(col []string) []string {
	curr := 0
	l := 0
	for r, chr := range col {
		if chr == "O" {
			col[r] = "."
			curr++
		}
		if chr == "#" || r == len(col)-1 {
			for n := l; n < l+curr; n++ {
				col[n] = "O"
			}
			curr = 0
			l = r + 1
		}
	}
	return col
}
