package main

import (
	"fmt"
	"math"
	"os"
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
	res := 0
	allH, allW := parse(input)
	for i := range allH {
		hand := allH[i]
		winning := allW[i]
		diff := setDiff(hand, winning)
		if diff == 1 || diff == 2 {
			res += diff
		} else if diff == 0 {
			continue
		} else {
			res += int(math.Pow(2, float64(diff-1)))
		}
	}
	return res
}

func partTwo(input string) int {
	allH, allW := parse(input)
	diffs := []int{}
	for i := range allH {
		hand := allH[i]
		winning := allW[i]
		diff := setDiff(hand, winning)
		diffs = append(diffs, diff)
	}

	counts := map[int]int{}
	sum := 0
	queue := []int{}
	for i := range diffs {
		queue = append(queue, i)
	}
	for len(queue) > 0 {
		sum += 1
		cur := queue[0]
		queue = queue[1:]
		counts[cur]++
		for i := 0; i < diffs[cur]; i++ {
			queue = append(queue, cur+i+1)
		}
	}
	return sum
}

func setDiff(hand map[string]bool, winning map[string]bool) int {
	count := 0
	for num := range hand {
		_, ok := winning[num]
		if ok {
			count++
		}
	}
	return count
}

func parse(input string) ([]map[string]bool, []map[string]bool) {
	allH := []map[string]bool{}
	allW := []map[string]bool{}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		h := map[string]bool{}
		w := map[string]bool{}
		parts := strings.Split(line, ": ")
		p := strings.Split(parts[1], " | ")
		winning := strings.Fields(p[0])
		hand := strings.Fields(p[1])

		for _, num := range hand {
			h[num] = true
		}

		for _, num := range winning {
			w[num] = true
		}

		allH = append(allH, h)
		allW = append(allW, w)
	}

	return allH, allW
}
