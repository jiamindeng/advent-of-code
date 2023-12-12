package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

var pipes = map[string]map[string]bool{
	"|": {
		"-1,0": true,
		"1,0":  true,
	},
	"-": {
		"0,-1": true,
		"0,1":  true,
	},
	"L": {
		"-1,0": true,
		"0,1":  true,
	},
	"J": {
		"-1,0": true,
		"0,-1": true,
	},
	"7": {
		"1,0":  true,
		"0,-1": true,
	},
	"F": {
		"1,0": true,
		"0,1": true,
	},
	"S": {
		"1,0":  true,
		"0,1":  true,
		"0,-1": true,
		"-1,0": true,
	},
}

var dirs = map[string]map[string]bool{
	"1,0": {
		"|": true,
		"L": true,
		"J": true,
	},
	"-1,0": {
		"|": true,
		"F": true,
		"7": true,
	},
	"0,1": {
		"-": true,
		"J": true,
		"7": true,
	},
	"0,-1": {
		"-": true,
		"L": true,
		"F": true,
	},
}

var delta = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func main() {
	f, _ := os.ReadFile("./input.txt")
	resOne, _ := partOne(string(f))
	fmt.Println(resOne)
	resTwo := partTwo(string(f))
	fmt.Println(resTwo)
}

func parse(input string) ([][]string, int, int) {
	res := [][]string{}
	c := 0
	r := 0
	lines := strings.Split(input, "\n")
	for cc, line := range lines {
		parsedLine := []string{}
		for rr, l := range line {
			chr := string(l)
			parsedLine = append(parsedLine, chr)
			if chr == "S" {
				c = cc
				r = rr
			}
		}
		res = append(res, parsedLine)
	}
	return res, r, c
}

func partOne(input string) (int, [][]int) {
	in, startC, startR := parse(input)
	steps := make([][]int, len(in))
	for i := range steps {
		row := make([]int, len(in[0]))
		for j := range row {
			row[j] = -1
		}
		steps[i] = row
	}

	steps[startR][startC] = 0
	queue := [][]int{{startR, startC}}
	for len(queue) != 0 {
		currR := queue[0][0]
		currC := queue[0][1]
		queue = queue[1:]
		curr := in[currR][currC]
		for _, d := range delta {
			newC := d[0] + currC
			newR := d[1] + currR
			inRange := newC >= 0 && newR >= 0 && newC < len(in[0]) && newR < len(in)
			if inRange && steps[newR][newC] == -1 {
				newPipe := in[newR][newC]
				dKey := fmt.Sprintf("%d,%d", d[1], d[0])
				canTravel := pipes[curr][dKey] && dirs[dKey][newPipe]
				if canTravel {
					next := []int{newR, newC}
					queue = append(queue, next)
					steps[newR][newC] = steps[currR][currC] + 1
				}
			}
		}
	}

	return getMax(steps), steps
}

func getMax(steps [][]int) int {
	max := -1
	for _, row := range steps {
		for _, step := range row {
			max = int(math.Max(float64(step), float64(max)))
		}
	}
	return max
}

func partTwo(input string) int {
	numIn := 0
	in, _, _ := parse(input)
	_, steps := partOne(input)

	edges := []string{"S", "F", "|", "J", ".", "-"}
	for r, row := range in {
		for c := range row {
			if steps[r][c] != -1 {
				continue
			}
			i := r
			j := c
			numCrosses := 0
			for i < len(in) && j < len(in[0]) {
				ray := in[i][j]
				if steps[i][j] != -1 && slices.Contains(edges, ray) {
					numCrosses++
				}
				i++
				j++
			}
			if numCrosses%2 == 1 {
				numIn++
			}
		}
	}

	return numIn
}
