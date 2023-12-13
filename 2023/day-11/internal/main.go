package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	f, _ := os.ReadFile("./input.txt")
	resOne := partOne(string(f), true)
	fmt.Println(resOne)
	resTwo := partTwo(string(f))
	fmt.Println(resTwo)
}

func parse(input string) ([][]string, map[int]bool, map[int]bool, [][]int) {
	lines := strings.Split(input, "\n")
	parsed := make([][]string, len(lines))
	emptyRows := map[int]bool{}
	emptyCols := map[int]bool{}

	galaxyCoords := [][]int{}

	for c := range lines[0] {
		emptyCols[c] = true
	}

	for r, line := range lines {
		if parsed[r] == nil {
			parsed[r] = []string{}
		}
		emptyRow := true
		for c, chr := range line {
			if string(chr) != "." {
				emptyRow = false
				emptyCols[c] = false
				galaxyCoords = append(galaxyCoords, []int{r, c})
			}
			parsed[r] = append(parsed[r], string(chr))
		}
		if emptyRow {
			emptyRows[r] = true
		}
	}

	return parsed, emptyRows, emptyCols, galaxyCoords
}

func partOne(input string, partOne bool) int {
	diff := 1000000 - 1
	_, emptyRows, emptyCols, galaxies := parse(input)
	sum := 0
	for _, galaxy1 := range galaxies {
		for _, galaxy2 := range galaxies {
			if galaxy2[0] == galaxy1[0] && galaxy2[1] == galaxy1[1] {
				continue
			}
			r1, c1 := galaxy1[0], galaxy1[1]
			r2, c2 := galaxy2[0], galaxy2[1]
			dR := int(math.Abs(float64(r2 - r1)))
			dC := int(math.Abs(float64(c2 - c1)))

			for emptyRow, e := range emptyRows {
				if e && (r1 <= emptyRow && emptyRow <= r2 || r2 <= emptyRow && emptyRow <= r1) {
					if partOne {
						dR++
					} else {
						dR += diff
					}
				}
			}

			for emptyCol, e := range emptyCols {
				if e && (c1 <= emptyCol && emptyCol <= c2 || c2 <= emptyCol && emptyCol <= c1) {
					if partOne {
						dC++
					} else {
						dC += diff
					}
				}
			}
			sum = sum + dR + dC
		}
	}
	return sum / 2
}

func partTwo(input string) int {
	return partOne(input, false)
}
