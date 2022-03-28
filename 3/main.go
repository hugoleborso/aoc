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

func main() {
	s := newScanFromFile("input.txt")
	//s.print()
	n := len(s[0])
	bitSum := make([]int, n)

	for i, bitsValue := range s {
		for j, val := range bitsValue {
			intBit, err := strconv.Atoi(string(val))
			if err != nil {
				fmt.Println("Error: ", err)
				fmt.Println(i, j, val)
				os.Exit(1)
			}
			bitSum[j] += intBit
		}
	}
	var gamma float64
	var epsilon float64

	m := len(s)
	for j, bit := range bitSum {

		i := math.Round(float64(bit) / float64(m))
		if i != 0 {
			gamma += math.Pow(2, float64(n-j-1))
		} else {
			epsilon += math.Pow(2, float64(n-1-j))
		}

	}
	fmt.Println(gamma, epsilon, int(gamma*epsilon))

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
