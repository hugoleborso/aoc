package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strings"
)

func main() {

	bs, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	scan := strings.Split(string(bs), "\r\n")
	sum := 0

	for _, line := range scan {

		newline := strings.Split(line, " | ")
		corres := decode(strings.Split(newline[0], " "))
		digits := strings.Split(newline[1], " ")
		lineSum := 0
		for i, digit := range digits {
			lineSum += int(math.Pow(10, float64(3-i))) * corres[sortString(digit)]
		}
		//fmt.Println(digits, ":", lineSum, "     ", corres)
		sum += lineSum
	}

	fmt.Println(sum)
}

func decode(combinations []string) map[string]int {
	letters := []string{"a", "b", "c", "d", "e", "f", "g"}
	mix := map[string][]string{}
	result := map[string]int{}

	for _, c := range combinations {
		c = sortString(c)
		switch len(c) {
		case 2:
			result[c] = 1
			poss := []string{string(c[0]), string(c[1])}
			mix["c"] = poss
			mix["f"] = poss
		case 7:
			result[c] = 8
		}
	}
	for _, c := range combinations {
		c = sortString(c)
		switch len(c) {
		case 3:
			result[c] = 7
			mix["a"] = makeArrayWithWithout([]string{string(c[0]), string(c[1]), string(c[2])}, mix["c"])
		case 4:
			result[c] = 4
			poss := makeArrayWithWithout([]string{string(c[0]), string(c[1]), string(c[2]), string(c[3])}, mix["c"])
			mix["b"] = poss
			mix["d"] = poss
		}
	}

	for _, c := range combinations {
		c = sortString(c)
		if len(c) == 6 {
			for _, le := range letters {
				if !strings.Contains(c, le) && contains(mix["c"], le) {
					mix["c"] = []string{le}
					mix["f"] = makeArrayWithWithout(mix["f"], mix["c"])
					result[c] = 6
				}

				if !strings.Contains(c, le) && contains(mix["d"], le) {
					mix["d"] = []string{le}
					mix["b"] = makeArrayWithWithout(mix["b"], mix["d"])
					result[c] = 0
				}

				if !strings.Contains(c, le) && !contains(mix["d"], le) && !contains(mix["c"], le) {
					mix["e"] = []string{le}
					result[c] = 9
				}
			}
		}
	}

	for _, c := range combinations {
		c = sortString(c)
		if len(c) == 5 {
			if !strings.Contains(c, mix["b"][0]) && !strings.Contains(c, mix["e"][0]) {
				result[c] = 3
			}
			if !strings.Contains(c, mix["c"][0]) {
				result[c] = 5
			}
			if !strings.Contains(c, mix["f"][0]) {
				result[c] = 2
			}
		}
	}

	return result
}

func contains(list []string, s string) bool {
	for _, l := range list {
		if l == s {
			return true
		}
	}
	return false
}

func makeArrayWithWithout(with []string, without []string) []string {
	res := []string{}
	for _, el := range with {
		if !contains(without, el) {
			res = append(res, el)
		}
	}
	return res
}

func sortString(s string) string {
	comList := strings.Split(s, "")
	sort.Strings(comList)
	return strings.Join(comList, "")
}
