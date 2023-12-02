package main

import (
	"fmt"
	"math"
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

type CubeSet map[string]int

// result is in order
func parse(input string) [][]CubeSet {
	games := [][]CubeSet{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		game := []CubeSet{}
		s := strings.Split(line, ": ")
		var gameNumber int
		_, err := fmt.Sscanf(s[0], "Game %d", &gameNumber)
		if err != nil {
			panic(err)
		}
		sets := strings.Split(s[1], "; ")
		for _, set := range sets {
			colors := strings.Split(set, ", ")
			cur := CubeSet{}
			for _, color := range colors {
				var colorNumber int
				var colorName string
				_, err := fmt.Sscanf(color, "%d %s", &colorNumber, &colorName)
				if err != nil {
					panic(err)
				}
				cur[colorName] = colorNumber
			}
			game = append(game, cur)
		}
		games = append(games, game)
	}

	return games
}

func gameIsPossible(game []CubeSet, cubes CubeSet) bool {
	for _, set := range game {
		for color, number := range cubes {
			if set[color] > number {
				return false
			}
		}
	}

	return true
}

func partOne(input string) int {
	sum := 0
	cubes := CubeSet{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	sets := parse(input)
	for i, set := range sets {
		if gameIsPossible(set, cubes) {
			sum += i + 1
		}
	}

	return sum
}

func minCubes(game []CubeSet) int {
	product := 1
	max := map[string]int{}
	for _, set := range game {
		for color, num := range set {
			max[color] = int(math.Max(float64(max[color]), float64(num)))
		}
	}

	for _, num := range max {
		product *= num
	}
	return product
}

func partTwo(input string) int {
	sum := 0

	games := parse(input)
	for _, game := range games {
		sum += minCubes(game)
	}

	return sum
}
