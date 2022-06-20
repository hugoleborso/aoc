package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	bs, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	scan := strings.Split(string(bs), "\r\n")

	mapp := make([][]int, len(scan))

	for i, line := range scan {
		for _, r := range line {
			n, err := strconv.Atoi(string(r))
			if err != nil {
				// handle error
				fmt.Println(err)
				os.Exit(2)
			}
			mapp[i] = append(mapp[i], n)
		}

	}
	l := len(mapp)
	L := len(mapp[0])
	pits := []int{}
	for i, line := range mapp {
		for j, n := range line {
			if isPit(n, i, j, mapp, l, L) {
				pits = append(pits, n)
			}
		}
	}
	s := 0
	for _, p := range pits {
		s += p + 1
	}
	fmt.Println(s)

}

func isPit(n, i, j int, mapp [][]int, l, L int) bool {
	for _, a := range []int{-1, 1} {
		if i+a < l && i+a >= 0 {
			if mapp[i+a][j] <= n {
				return false
			}

		}
		if j+a < L && j+a >= 0 {
			if mapp[i][j+a] <= n {
				return false
			}
		}
	}
	return true
}
