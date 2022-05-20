package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	bs, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), "\r\n")

	digitsDisplay := [][]string{}

	nbs := map[int]int{}

	for _, line := range s {
		newline := strings.Split(line, " | ")
		digitsDisplay = append(digitsDisplay, newline)
		for _, digit := range strings.Split(newline[1], " ") {
			nbs[len(digit)] += 1
		}

	}

	fmt.Println(nbs[2] + nbs[3] + nbs[4] + nbs[7])

}
