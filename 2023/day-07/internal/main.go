package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

var order = map[string]int{
	"A": 0, "K": 1, "Q": 2, "J": 3, "T": 4, "9": 5, "8": 6, "7": 7, "6": 8, "5": 9, "4": 10, "3": 11, "2": 12,
}

func main() {
	f, _ := os.ReadFile("./input.txt")
	resOne := partOne(string(f))
	fmt.Println(resOne)
	resTwo := partTwo(string(f))
	fmt.Println(resTwo)
}

func sortAlpha(hands []string) []string {
	sort.Slice(hands, func(i, j int) bool {
		a := hands[i]
		b := hands[j]
		idx := 0
		for idx < len(a) && idx < len(b) {
			if order[string(a[idx])] < order[string(b[idx])] {
				return true
			}
			if order[string(a[idx])] > order[string(b[idx])] {
				return false
			}
			idx++
		}
		if len(a) < len(b) {
			return true
		}
		if len(a) > len(b) {
			return false
		}
		if len(a) == len(b) {
			return true
		}
		return true
	})
	return hands
}

func parse(input string) (map[string]int, error) {
	res := map[string]int{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		l := strings.Split(line, " ")
		num, err := strconv.Atoi(l[1])
		res[l[0]] = num
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

func partition(hands []string, isPartTwo bool) [][]string {
	if isPartTwo {
		order["J"] = 13
	}
	fiveOfAKind := []string{}
	fourOfAKind := []string{}
	fullHouse := []string{}
	threeOfAKind := []string{}
	twoPair := []string{}
	onePair := []string{}
	highCard := []string{}

	for _, hand := range hands {
		count := map[string]int{}
		for _, chr := range hand {
			count[string(chr)]++
		}
		sortedCount := []int{}
		for _, v := range count {
			sortedCount = append(sortedCount, v)
		}
		j := count["J"]
		slices.Sort(sortedCount)
		if slices.Equal(sortedCount, []int{5}) {
			fiveOfAKind = append(fiveOfAKind, hand)
		} else if slices.Equal(sortedCount, []int{1, 4}) {
			if isPartTwo && j == 1 || isPartTwo && j == 4 {
				fiveOfAKind = append(fiveOfAKind, hand)
			} else {
				fourOfAKind = append(fourOfAKind, hand)
			}

		} else if slices.Equal(sortedCount, []int{2, 3}) {
			if isPartTwo && j == 2 || isPartTwo && j == 3 {
				fiveOfAKind = append(fiveOfAKind, hand)
			} else {
				fullHouse = append(fullHouse, hand)
			}

		} else if slices.Equal(sortedCount, []int{1, 1, 3}) {
			if isPartTwo && j == 1 || isPartTwo && j == 3 {
				fourOfAKind = append(fourOfAKind, hand)
			} else {
				threeOfAKind = append(threeOfAKind, hand)
			}
		} else if slices.Equal(sortedCount, []int{1, 2, 2}) {
			if isPartTwo && j == 2 {
				fourOfAKind = append(fourOfAKind, hand)
			} else if isPartTwo && j == 1 {
				fullHouse = append(fullHouse, hand)
			} else {
				twoPair = append(twoPair, hand)
			}
		} else if slices.Equal(sortedCount, []int{1, 1, 1, 2}) {
			if isPartTwo && j == 1 || isPartTwo && j == 2 {
				threeOfAKind = append(threeOfAKind, hand)
			} else {
				onePair = append(onePair, hand)
			}
		} else {
			if isPartTwo && j == 1 {
				onePair = append(onePair, hand)
			} else {
				highCard = append(highCard, hand)
			}
		}
	}

	return [][]string{fiveOfAKind, fourOfAKind, fullHouse, threeOfAKind, twoPair, onePair, highCard}
}

func process(input string, isPartTwo bool) int {
	handMap, err := parse(input)
	if err != nil {
		panic(err)
	}
	hands := []string{}
	for hand := range handMap {
		hands = append(hands, hand)
	}
	partitionedHands := partition(hands, isPartTwo)
	finalHands := []string{}
	for _, hand := range partitionedHands {
		sortedHands := sortAlpha(hand)
		finalHands = append(finalHands, sortedHands...)
	}
	res := 0
	for i := len(finalHands) - 1; i >= 0; i-- {
		res += (len(finalHands) - i) * handMap[finalHands[i]]
	}

	return res
}

func partOne(input string) int {
	return process(input, false)
}

func partTwo(input string) int {
	return process(input, true)
}
