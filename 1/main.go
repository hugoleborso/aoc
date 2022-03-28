package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type scan []int

func main() {
	scan := newScanFromFile("input.txt")
	largerCount := 0
	for i := range scan {
		if i > 2 {
			if scan[i]+scan[i-1]+scan[i-2] > scan[i-1]+scan[i-2]+scan[i-3] {
				largerCount += 1
			}
		}
	}
	fmt.Println("larger :", largerCount)
}

func (s scan) print() {
	for i, depth := range s {
		fmt.Println(i, depth)
	}
}

func newScanFromFile(filename string) scan {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := scan{}
	stringScan := strings.Split(string(bs), "\r\n")

	for _, depth := range stringScan {
		intDepth, err := strconv.Atoi(depth)
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
		s = append(s, intDepth)
	}

	return s
}
