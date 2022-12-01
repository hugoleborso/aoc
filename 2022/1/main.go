package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type scan []string
type elvesPackages []int

func main() {
	s := newScanFromFile("input.txt")
	// s.print()
	elvesPackages := elvesPackages{}
	currentElve := 0
	for _, textLine := range s {
		if textLine == "" {
			elvesPackages = append(elvesPackages, currentElve)
			currentElve = 0
		} else {
			intVar, _ := strconv.Atoi(string(textLine))
			currentElve += intVar
		}
	}
	elvesPackages = append(elvesPackages, currentElve)
	elvesPackages.print()
	fmt.Println(elvesPackages.max())
	fmt.Println(elvesPackages.sumTopThree())

}



func (l elvesPackages) max() int {
	max := 0
	for _,n := range l{
		if n>max {
			max=n
		}
	}
	return max
}

func (l elvesPackages) sumTopThree() int {
	max := [3]int{}
	for _,n := range l{
		if n>max[2] {
			if n<max[1]{
				max[2]=n
			} else if n<max[0]{
				max[2]=max[1]
				max[1]=n
			} else {
				max[2]=max[1]
				max[1]=max[0]
				max[0]=n
			}
		}
	}
	return max[0]+max[1]+max[2]
}

func (s scan) print() {
	for i, info := range s {
		fmt.Println(i, info)
	}
}

func (s elvesPackages) print() {
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

	s := strings.Split(string(bs), "\n")

	return s
}
