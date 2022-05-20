package main

import (
	"fmt"
)

func main() {
	days := 256
	lanternFishes := []int{3, 1, 5, 4, 4, 4, 5, 3, 4, 4, 1, 4, 2, 3, 1, 3, 3, 2, 3, 2, 5, 1, 1, 4, 4, 3, 2, 4, 2, 4, 1, 5, 3, 3, 2, 2, 2, 5, 5, 1, 3, 4, 5, 1, 5, 5, 1, 1, 1, 4, 3, 2, 3, 3, 3, 4, 4, 4, 5, 5, 1, 3, 3, 5, 4, 5, 5, 5, 1, 1, 2, 4, 3, 4, 5, 4, 5, 2, 2, 3, 5, 2, 1, 2, 4, 3, 5, 1, 3, 1, 4, 4, 1, 3, 2, 3, 2, 4, 5, 2, 4, 1, 4, 3, 1, 3, 1, 5, 1, 3, 5, 4, 3, 1, 5, 3, 3, 5, 4, 2, 3, 4, 1, 2, 1, 1, 4, 4, 4, 3, 1, 1, 1, 1, 1, 4, 2, 5, 1, 1, 2, 1, 5, 3, 4, 1, 5, 4, 1, 3, 3, 1, 4, 4, 5, 3, 1, 1, 3, 3, 3, 1, 1, 5, 4, 2, 5, 1, 1, 5, 5, 1, 4, 2, 2, 5, 3, 1, 1, 3, 3, 5, 3, 3, 2, 4, 3, 2, 5, 2, 5, 4, 5, 4, 3, 2, 4, 3, 5, 1, 2, 2, 4, 3, 1, 5, 5, 1, 3, 1, 3, 2, 2, 4, 5, 4, 2, 3, 2, 3, 4, 1, 3, 4, 2, 5, 4, 4, 2, 2, 1, 4, 1, 5, 1, 5, 4, 3, 3, 3, 3, 3, 5, 2, 1, 5, 5, 3, 5, 2, 1, 1, 4, 2, 2, 5, 1, 4, 3, 3, 4, 4, 2, 3, 2, 1, 3, 1, 5, 2, 1, 5, 1, 3, 1, 4, 2, 4, 5, 1, 4, 5, 5, 3, 5, 1, 5, 4, 1, 3, 4, 1, 1, 4, 5, 5, 2, 1, 3, 3}

	lanternFishesGroup := map[int]int{}
	for f := 0; f <= 8; f++ {
		lanternFishesGroup[f] = 0
	}

	for _, f := range lanternFishes {
		lanternFishesGroup[f] += 1
	}

	for d := 1; d <= days; d++ {
		newLanternFishesGroup := map[int]int{}
		for f := 1; f <= 8; f++ {
			newLanternFishesGroup[f-1] = lanternFishesGroup[f]
		}
		newLanternFishesGroup[8] = lanternFishesGroup[0]
		newLanternFishesGroup[6] += lanternFishesGroup[0]
		lanternFishesGroup = newLanternFishesGroup

	}
	sum := 0
	for f := 0; f <= 8; f++ {
		sum += lanternFishesGroup[f]
	}
	fmt.Println(sum)
}
