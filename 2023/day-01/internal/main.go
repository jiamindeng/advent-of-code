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

var stringToInt = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func partTwo(input string) int {
	sum := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		a := 0
		b := 0
		// iterate forwards
	out1:
		for i, chr := range line {
			if _, err := strconv.Atoi(string(chr)); err == nil {
				parsed, _ := strconv.ParseInt(string(chr), 10, 64)
				a = int(parsed)
				break
			} else {
				// parse word here
				word := ""
				for j := i; j < len(line); j++ {
					if _, err := strconv.Atoi(string(line[j])); err != nil {
						word += string(line[j])
						parsed, ok := stringToInt[word]
						if ok {
							a = parsed
							break out1
						}
					}
				}
			}
		}

	out2:
		// iterate backwards
		for i := len(line) - 1; i >= 0; i-- {
			chr := line[i]
			if _, err := strconv.Atoi(string(chr)); err == nil {
				parsed, _ := strconv.ParseInt(string(chr), 10, 64)
				b = int(parsed)
				break
			} else {
				// parse word here
				word := ""
				for j := i; j >= 0; j-- {
					if _, err := strconv.Atoi(string(line[j])); err != nil {
						word += string(line[j])
						reversed := reverseString(word)
						parsed, ok := stringToInt[reversed]
						if ok {
							b = parsed
							break out2
						}
					}
				}
			}
		}
		num := a*10 + b
		sum += num
	}
	return sum
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func partOne(input string) int {
	sum := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		a := 0
		b := 0
		// iterate forwards
		for _, chr := range line {
			if _, err := strconv.Atoi(string(chr)); err != nil {
				continue
			} else {
				parsed, _ := strconv.ParseInt(string(chr), 10, 64)
				a = int(parsed)
				break
			}
		}
		// iterate backwards
		for i := len(line) - 1; i >= 0; i-- {
			chr := line[i]
			if _, err := strconv.Atoi(string(chr)); err != nil {
				continue
			} else {
				parsed, _ := strconv.ParseInt(string(chr), 10, 64)
				b = int(parsed)
				break
			}
		}
		num := a*10 + b
		sum += num
	}
	return sum
}
