package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

type scan []string
type lines [][2][2]int
type grid [][]int

func main() {
	s := newScanFromFile("input.txt")
	//s.print()
	lines := getLinesFromScan(s)
	lines = filterHorVerDiagLines(lines)
	//fmt.Println(lines)
	fmt.Println("Result :", lines.getOverlaps(false))
}

func getLinesFromScan(s scan) lines {
	lines := lines{}
	for _, textLine := range s {
		s1 := strings.Split(textLine, " -> ")
		line := [2][2]int{}
		for i, s2 := range s1 {
			s3 := strings.Split(s2, ",")
			for j, strVar := range s3 {
				intVar, _ := strconv.Atoi(strVar)
				line[i][j] = intVar
			}
		}
		lines = append(lines, line)
	}
	return lines
}

func filterHorVerDiagLines(l lines) lines {
	outputLines := lines{}
	for _, line := range l {
		if line[0][0] == line[1][0] || line[0][1] == line[1][1] {
			if line[0][0] > line[1][0] || line[0][1] > line[1][1] {
				outputLines = append(outputLines, [2][2]int{line[1], line[0]})
			} else {
				outputLines = append(outputLines, line)
			}

		}
		if math.Abs(float64(line[0][0]-line[1][0])) == math.Abs(float64(line[0][1]-line[1][1])) {
			if line[0][0] > line[1][0] {
				outputLines = append(outputLines, [2][2]int{line[1], line[0]})
			} else {
				outputLines = append(outputLines, line)
			}
		}
	}
	return outputLines
}

func (l lines) getDims() (x, y int) {
	maxX := 0
	maxY := 0

	for _, line := range l {
		if line[0][0] > maxX {
			maxX = line[0][0]
		}
		if line[1][0] > maxX {
			maxX = line[1][0]
		}
		if line[0][1] > maxY {
			maxY = line[0][1]
		}
		if line[1][0] > maxY {
			maxY = line[1][1]
		}
	}
	return maxX, maxY
}

func (l lines) getOverlaps(print bool) int {
	x, y := l.getDims()

	grid := grid{}
	for i := 0; i < y+2; i++ {
		grid = append(grid, make([]int, x+2))
	}

	for _, line := range l {
		distX := line[1][0] - line[0][0]
		distY := line[1][1] - line[0][1]

		if math.Abs(float64(distX)) == math.Abs(float64(distY)) {
			if distY > 0 {
				for i := 0; i <= distX; i++ {
					grid[line[0][1]+i][line[0][0]+i]++
				}
			} else {
				for i := 0; i <= distX; i++ {
					grid[line[0][1]-i][line[0][0]+i]++
				}
			}

		} else {
			if distX > 0 {
				for i := 0; i <= distX; i++ {
					grid[line[0][1]][line[0][0]+i]++
				}
			}
			if distY > 0 {
				for j := 0; j <= distY; j++ {
					grid[line[0][1]+j][line[0][0]]++
				}
			}
		}

	}
	if print {
		grid.print()
	}

	count := 0
	for _, line := range grid {
		for _, val := range line {
			if val > 1 {
				count++
			}
		}
	}

	return count
}

func (g grid) print() {
	for _, l := range g {
		fmt.Println(l)
	}
}

func (s scan) print() {
	for i, info := range s {
		fmt.Println(i, info)
	}
}

func newScanFromFile(filename string) scan {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), "\r\n")

	return s
}
