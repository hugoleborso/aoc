package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type scan []string

func main() {
	s := newScanFromFile("input.txt")
	//s.print()
	inputs := strings.Split(s[0], ",")

	result := findLastWinningBoard(s, inputs)

	fmt.Println("Result : ", result)

}

func findLastWinningBoard(s scan, inputs []string) int {
	res := 0
	winningBoardsList := []int{}
	bingoBoardsValues, bingoBoardsValuesFound := createBingoBoards(s)
	for _, n := range inputs {
		fmt.Println(n)
		m, _ := strconv.Atoi(n)
		for b, board := range bingoBoardsValues {
			if contains(winningBoardsList, b) == false {
				for i, line := range board {
					for j, val := range line {
						if m == val {
							bingoBoardsValuesFound[b][i][j] = true
						}
					}
				}
			}

		}
		winningBoards := checkBoards(bingoBoardsValuesFound)

		for _, winningBoard := range winningBoards {

			if contains(winningBoardsList, winningBoard) == false {
				res = m * unMarkedSum(winningBoard, bingoBoardsValues, bingoBoardsValuesFound)
				fmt.Println("Board : ", winningBoard, "Result : ", res)
				winningBoardsList = append(winningBoardsList, winningBoard)
			}
			bingoBoardsValuesFound[winningBoard] = [5][5]bool{}

		}

	}
	return res
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (s scan) print() {
	for i, info := range s {
		fmt.Println(i, info)
	}
}

func createBingoBoards(s scan) ([][5][5]int, [][5][5]bool) {
	bingoBoardsValues := [][5][5]int{}
	bingoBoardsValuesFound := [][5][5]bool{}

	for i, line := range s {
		if (i-2)%6 == 0 {
			bingoBoardsValues = append(bingoBoardsValues, [5][5]int{})
			bingoBoardsValuesFound = append(bingoBoardsValuesFound, [5][5]bool{})
		}

		if i > 1 && (i-1)%6 != 0 {
			cline := strings.Split(line, " ")
			for k, val := range cline {
				intBit, _ := strconv.Atoi(string(val))
				bingoBoardsValues[(i-2)/6][(i-2)%6][k] = intBit
			}
		}
	}
	return bingoBoardsValues, bingoBoardsValuesFound
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

func checkBoards(bingoBoardsValuesFound [][5][5]bool) []int {
	res := []int{}
	for n, board := range bingoBoardsValuesFound {
		for i, line := range board {
			if line == [5]bool{true, true, true, true, true} {
				res = append(res, n)
			} else if [5]bool{board[0][i], board[1][i], board[2][i], board[3][i], board[4][i]} == [5]bool{true, true, true, true, true} {
				res = append(res, n)
			}
		}
	}
	return res
}

func unMarkedSum(boardNb int, bingoBoardsValues [][5][5]int, bingoBoardsValuesFound [][5][5]bool) int {
	sum := 0
	for i, line := range bingoBoardsValuesFound[boardNb] {
		for j, val := range line {
			if val == false {
				sum = sum + bingoBoardsValues[boardNb][i][j]
			}
		}
	}
	return sum
}
