package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	f, _ := os.ReadFile("./input.txt")
	input := parse(string(f))
	resOne := partOne(input, Point{dir: 0, r: 0, c: -1})
	fmt.Println(resOne)
	resTwo := partTwo(input)
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

func mod(d, m int) int {
	var res int = d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}

var dirToType = map[int]string{
	0:   "horizontal",
	90:  "vertical",
	180: "horizontal",
	270: "vertical",
}

var deltaDirs = map[string]map[string][]int{
	"vertical": {
		"\\": []int{90},
		"/":  []int{-90},
		"|":  []int{0},
		"-":  []int{-90, 90}, // splits in 2 dirs
	},
	"horizontal": {
		"\\": []int{-90},
		"/":  []int{90},
		"|":  []int{-90, 90},
		"-":  []int{0}, // splits in 2 dirs
	},
}

var deltaCoords = map[int][]int{
	0:   {0, 1},  // >
	90:  {-1, 0}, // ^
	180: {0, -1}, // ^
	270: {1, 0},  // ^
}

type Point struct {
	dir int
	r   int
	c   int
}

func getKey(m [2]int) string {
	return fmt.Sprintf("%d,%d", m[0], m[1])
}

func getPointKey(p Point) string {
	return fmt.Sprintf("%d,%d,%d", p.r, p.c, p.dir)
}

func partOne(trap [][]string, start Point) int {
	maxR := len(trap)
	maxC := len(trap[0])
	currPoints := []Point{start}
	visited := map[string]map[int]int{}

	for len(currPoints) > 0 && len(currPoints) < maxR*maxC {
		currPoint := currPoints[0]
		currPoints = currPoints[1:]
		_, ok := visited[getKey([2]int{currPoint.r, currPoint.c})]
		if !ok {
			visited[getKey([2]int{currPoint.r, currPoint.c})] = map[int]int{currPoint.dir: 1}
		} else {
			_, k := visited[getKey([2]int{currPoint.r, currPoint.c})][currPoint.dir]
			if !k {
				visited[getKey([2]int{currPoint.r, currPoint.c})][currPoint.dir]++
			}
		}

		deltaCoord := deltaCoords[currPoint.dir]
		newCoord := []int{currPoint.r + deltaCoord[0], currPoint.c + deltaCoord[1]}
		if newCoord[0] < maxR && newCoord[1] < maxC && newCoord[0] >= 0 && newCoord[1] >= 0 {
			chr := trap[newCoord[0]][newCoord[1]]
			if chr != "." {
				deltaDir, ok := deltaDirs[dirToType[mod(currPoint.dir, 360)]][chr]
				if !ok {
					panic(fmt.Sprintf("currPoint.dir:%d, dirToType:%s, chr:%s", currPoint.dir, dirToType[currPoint.dir], chr))
				}

				for _, d := range deltaDir {
					if _, ok := visited[getKey([2]int{newCoord[0], newCoord[1]})][mod(currPoint.dir+d, 360)]; !ok {
						currPoints = append(currPoints, Point{dir: mod(currPoint.dir+d, 360), r: newCoord[0], c: newCoord[1]})
					}
				}
			} else {
				if _, ok := visited[getKey([2]int{newCoord[0], newCoord[1]})][mod(currPoint.dir, 360)]; !ok {
					currPoints = append(currPoints, Point{dir: mod(currPoint.dir, 360), r: newCoord[0], c: newCoord[1]})
				}
			}
		}
	}

	return len(visited) - 1
}

func partTwo(trap [][]string) int {
	scenarios := []Point{}
	for r := range trap {
		scenarios = append(scenarios, Point{0, r, -1})
		scenarios = append(scenarios, Point{180, r, len(trap[0])})
	}
	for c := range trap[0] {
		scenarios = append(scenarios, Point{270, -1, c})
		scenarios = append(scenarios, Point{90, len(trap), 0})
	}

	max := -1

	for _, scenario := range scenarios {
		max = int(math.Max(float64(partOne(trap, scenario)), float64(max)))
	}

	return max
}
