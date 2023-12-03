package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	IntType    = "int"
	SymbolType = "symbol"
	Period     = "period"
)

func main() {
	f, _ := os.ReadFile("./input.txt")
	resOne := partOne(string(f))
	fmt.Println(resOne)
	resTwo := partTwo(string(f))
	fmt.Println(resTwo)
}

func partOne(input string) int {
	sum := 0
	nums := getNumsOne(parse(input))
	for _, num := range nums {
		sum += num
	}
	return sum
}

func partTwo(input string) int {
	sum := 0
	nums := getNumsTwo(parse(input))
	for _, num := range nums {
		sum += num
	}
	return sum
}

func parse(input string) [][]string {
	lines := strings.Split(input, "\n")
	output := make([][]string, len(lines))
	for i, line := range lines {
		for _, char := range line {
			output[i] = append(output[i], string(char))
		}
	}
	return output
}

func getNumsOne(lines [][]string) []int {
	maxR := len(lines)
	maxC := len(lines[0])
	nums := []int{}
	visited := map[string]bool{}
	adjacent := [][]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}, {0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	// find symbols
	for r, line := range lines {
		for c, char := range line {
			if getType(char) == SymbolType {
				// greedily iterate through them to build numbers
				for _, delta := range adjacent {
					curR := r + delta[0]
					curC := c + delta[1]
					key := fmt.Sprintf("%d,%d", curR, curC)
					inBounds := curR >= 0 && curR < maxR && curC >= 0 && curC < maxC
					if !visited[key] && inBounds && getType(lines[curR][curC]) == IntType {
						curChar := lines[curR][curC]
						newC := curC + 1
						left := ""
						right := curChar
						// go right
						for newC < maxC {
							newKey := fmt.Sprintf("%d,%d", curR, newC)
							newChar := lines[curR][newC]
							newC++

							if !visited[newKey] && getType(newChar) == IntType {
								right += newChar
								visited[newKey] = true
							} else {
								break
							}

						}
						newC = curC - 1
						// go left
						for newC >= 0 {
							newKey := fmt.Sprintf("%d,%d", curR, newC)
							newChar := lines[curR][newC]
							newC--

							if !visited[newKey] && getType(newChar) == IntType {
								left += newChar
								visited[newKey] = true
							} else {
								break
							}
						}

						num, err := strconv.Atoi(reverse(left) + right)
						if err != nil {
							panic(err)
						}
						nums = append(nums, num)
						visited[key] = true
					}
				}
			}
		}
	}

	return nums
}

func getNumsTwo(lines [][]string) []int {
	maxR := len(lines)
	maxC := len(lines[0])
	nums := []int{}
	visited := map[string]bool{}
	adjacent := [][]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}, {0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	// find symbols
	for r, line := range lines {
		for c, char := range line {
			if getType(char) == SymbolType {
				symbolNums := []int{}
				// greedily iterate through them to build numbers
				for _, delta := range adjacent {
					curR := r + delta[0]
					curC := c + delta[1]
					key := fmt.Sprintf("%d,%d", curR, curC)
					inBounds := curR >= 0 && curR < maxR && curC >= 0 && curC < maxC
					if !visited[key] && inBounds && getType(lines[curR][curC]) == IntType {
						curChar := lines[curR][curC]
						newC := curC + 1
						left := ""
						right := curChar
						// go right
						for newC < maxC {
							newKey := fmt.Sprintf("%d,%d", curR, newC)
							newChar := lines[curR][newC]
							newC++
							if !visited[newKey] && getType(newChar) == IntType {
								right += newChar
								visited[newKey] = true
							} else {
								break
							}

						}
						newC = curC - 1
						// go left
						for newC >= 0 {
							newKey := fmt.Sprintf("%d,%d", curR, newC)
							newChar := lines[curR][newC]
							newC--
							if !visited[newKey] && getType(newChar) == IntType {
								left += newChar
								visited[newKey] = true
							} else {
								break
							}
						}

						num, err := strconv.Atoi(reverse(left) + right)
						if err != nil {
							panic(err)
						}

						visited[key] = true
						symbolNums = append(symbolNums, num)
					}
				}
				if len(symbolNums) >= 2 {
					product := 1
					for _, num := range symbolNums {
						product *= num
					}
					nums = append(nums, product)
				}
			}
		}
	}
	return nums
}

func getType(char string) string {
	if char == "." {
		return Period
	}
	if strings.Contains("0123456789", char) {
		return IntType
	}
	return SymbolType
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
