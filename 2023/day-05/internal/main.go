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
	resOne := partOne(string(f))
	fmt.Println(resOne)
	resTwo := partTwo(string(f))
	fmt.Println(resTwo)
}

func parse(input string) ([]int, [][][]int) {
	lines := strings.Split(input, "\n")
	i := strings.Split(lines[0], ": ")
	initial := strings.Split(i[1], " ")
	initNums := []int{}
	for _, s := range initial {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		initNums = append(initNums, num)
	}

	groups := strings.Split(input, "\n\n")
	res := make([][][]int, len(groups)-1)
	for i, group := range groups[1:] {
		ls := strings.Split(group, "\n")
		for _, line := range ls[1:] {
			nums := strings.Split(line, " ")
			triplet := []int{}
			for _, n := range nums {
				num, err := strconv.Atoi(n)
				if err != nil {
					panic(err)
				}
				triplet = append(triplet, num)

			}
			a := res[i]
			a = append(a, triplet)
			res[i] = a
		}
	}

	return initNums, res
}

func partOne(input string) int {
	initNums, lookup := parse(input)
	res := make([]int, len(initNums))
	for i, num := range initNums {
		curNum := num
		for _, paths := range lookup {
			for _, path := range paths {
				dest := path[0]
				source := path[1]
				rang := path[2]
				maxSource := source + rang
				// in range
				if curNum <= maxSource && curNum >= source {
					delta := curNum - source
					dest += delta
					curNum = dest
					break
				}

			}

			res[i] = curNum
		}
	}

	return slices.Min(res)
}

func partTwo(input string) int {
	initNums, lookup := parse(input)
	res := []int{}

	for i := 0; i < len(initNums); i += 2 {
		var curStart, curEnd int

		curStart = initNums[i]
		curEnd = curStart + initNums[i+1]

		for j := curStart; j <= curEnd; j++ {
			curNum := j
			for k, paths := range lookup {
				for _, path := range paths {
					dest := path[0]
					source := path[1]
					rang := path[2]
					maxSource := source + rang
					// in range
					if curNum <= maxSource && curNum >= source {
						delta := curNum - source
						dest += delta
						curNum = dest
						break
					}

				}

				if k == len(lookup)-1 {
					res = append(res, curNum)
				}
			}
		}
	}
	return slices.Min(res)
}
