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

func parse(input string) []string {
	return strings.Split(input, ",")
}

//
// Determine the ASCII code for the current character of the string.
// Increase the current value by the ASCII code you just determined.
// Set the current value to itself multiplied by 17.
// Set the current value to the remainder of dividing itself by 256.

func hash(word string) int {
	curr := 0
	for _, chr := range word {
		curr += int(chr)
		curr *= 17
		curr %= 256
	}
	return curr
}

func partOne(input string) int {
	sum := 0
	for _, word := range parse(input) {
		sum += hash(word)
	}
	return sum
}

func partTwo(input string) int {
	boxes := make([][]map[string]int, 256)
	for _, word := range parse(input) {
		wrd := ""
		box := 0
		val := 0
		if strings.Contains(word, "=") {
			w := strings.Split(word, "=")
			wrd = w[0]
			box = hash(w[0])
			val, _ = strconv.Atoi(w[1])
			if boxes[box] == nil {
				boxes[box] = []map[string]int{}
			}
			found := false
			for i, b := range boxes[box] {
				for k := range b {
					if k == wrd {
						boxes[box][i][k] = val
						found = true
					}
				}
			}
			if !found {
				boxes[box] = append(boxes[box], map[string]int{wrd: val})
			}
		} else if strings.Contains(word, "-") {
			w := strings.Split(word, "-")
			wrd = w[0]
			box = hash(w[0])
			if boxes[box] == nil {
				boxes[box] = []map[string]int{}
			}
			for i, b := range boxes[box] {
				for k := range b {
					if k == wrd {
						boxes[box] = append(boxes[box][:i], boxes[box][i+1:]...)
						break
					}
				}
			}
		}
		fmt.Println(boxes)
	}
	sum := 0
	for i, box := range boxes {
		for j, slot := range box {
			for _, v := range slot {
				sum += (i + 1) * (j + 1) * v
			}
		}
	}
	return sum
}
