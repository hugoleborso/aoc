package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	direction string
	size      int
}

type scan []instruction

func main() {
	scan := newScanFromFile("input.txt")
	//scan.print()
	aim := 0
	x := 0
	y := 0

	for _, instr := range scan {
		if instr.direction == "up" {
			aim -= instr.size
		} else if instr.direction == "down" {
			aim += instr.size
		} else if instr.direction == "forward" {
			y += instr.size
			x += aim * instr.size
		}
	}
	fmt.Println("pos :", x, y)
	fmt.Println("prod:", x*y)
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

	s := scan{}
	stringScan := strings.Split(string(bs), "\r\n")

	for _, st := range stringScan {
		info := strings.Split(st, " ")
		sizeInt, err := strconv.Atoi(info[1])
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
		s = append(s, instruction{direction: info[0], size: sizeInt})
	}

	return s
}
