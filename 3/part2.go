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

	var scanOnes []string
	var scanZeros []string

	for j := 0; j < n; j++ {
		scanOnes = scanOnes[:0]
		scanZeros = scanZeros[:0]
		for _, bitsValue := range s {

			intBit, err := strconv.Atoi(string(bitsValue[j]))
			if err != nil {
				fmt.Println("Error: ", err)
				os.Exit(1)
			}
			if intBit == 0 {
				scanZeros = append(scanZeros, bitsValue)
			} else {
				scanOnes = append(scanOnes, bitsValue)
			}
		}
		if len(scanOnes) >= len(scanZeros) {
			s = scanOnes
		} else {
			s = scanZeros
		}
	}

	var oxygen float64

	for j, bit := range s[0] {
		intBit, err := strconv.Atoi(string(bit))
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
		if intBit != 0 {
			oxygen += math.Pow(2, float64(n-j-1))
		}
	}

	s = newScanFromFile("input.txt")

	for j := 0; j < n; j++ {
		scanOnes = scanOnes[:0]
		scanZeros = scanZeros[:0]
		for _, bitsValue := range s {
			intBit, err := strconv.Atoi(string(bitsValue[j]))
			if err != nil {
				fmt.Println("Error: ", err)
				os.Exit(1)
			}
			if intBit == 0 {
				scanZeros = append(scanZeros, bitsValue)
			} else {
				scanOnes = append(scanOnes, bitsValue)
			}
		}
		if len(scanZeros) == 0 {
			s = scanOnes
		} else if len(scanOnes) == 0 {
			s = scanZeros
		} else if len(scanZeros) <= len(scanOnes) {
			s = scanZeros
		} else {
			s = scanOnes
		}
	}

	var CO2 float64

	for j, bit := range s[0] {
		intBit, err := strconv.Atoi(string(bit))
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
		if intBit != 0 {
			CO2 += math.Pow(2, float64(n-j-1))
		}
	}

	fmt.Println(oxygen, CO2, int(oxygen*CO2))
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
