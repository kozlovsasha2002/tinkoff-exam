package task2

import (
	"fmt"
	"math"
)

func SolutionOfTask2() {
	var str string
	fmt.Scan(&str)

	dict := map[uint8]int{
		's': 0,
		'h': 0,
		'e': 0,
		'r': 0,
		'i': 0,
		'f': 0,
	}

	for i := range str {
		char := str[i]
		if _, ok := dict[char]; ok {
			dict[char]++
		}
	}

	min := math.MaxInt
	for i := range dict {
		if dict[i] < min {
			min = dict[i]
		}
	}

	fmt.Println(min)
}
