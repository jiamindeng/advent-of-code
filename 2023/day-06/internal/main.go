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

func parse(input string) ([]int, []int) {
	res := [][]int{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		nums := []int{}
		f := strings.Fields(line)
		for _, str := range f[1:] {
			num, err := strconv.Atoi(str)
			if err != nil {
				panic(err)
			}
			nums = append(nums, num)
		}
		res = append(res, nums)
	}

	return res[0], res[1]
}

func partOne(input string) int {
	times, distances := parse(input)
	ways := make([]int, len(times))

	for t, time := range times {
		for i := 0; i <= time; i++ {
			timeLeft := time - i
			if i*timeLeft > distances[t] {
				ways[t] = time - 2*i + 1
				break
			}
		}
	}
	product := 1
	for _, way := range ways {
		product *= way
	}
	return product
}

func partTwo(input string) int {
	// lol just modify the input
	return 0
}
