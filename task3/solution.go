package task3

import "fmt"

func SolutionTask3() {
	var numberOfCards int
	fmt.Scan(&numberOfCards)

	in := make([]int, numberOfCards)
	out := make([]int, numberOfCards)

	dict := make(map[int]int)

	for i := 0; i < numberOfCards; i++ {
		fmt.Scan(&in[i])
		dict[in[i]]++
	}
	for j := 0; j < numberOfCards; j++ {
		fmt.Scan(&out[j])
		if _, ok := dict[out[j]]; ok {
			dict[out[j]]--
		} else {
			fmt.Println("NO")
			return
		}
	}

	for i := range dict {
		if dict[i] != 0 {
			fmt.Println("NO")
			return
		}
	}

	begin := 0
	end := 0
	check := 0
	for i := 0; i < numberOfCards; i++ {
		if check == 0 {
			if in[i] != out[i] {
				begin = i
				check++
			}
		} else if check == 1 {
			if in[i] == out[i] {
				end = i
				check++
			}
		} else {
			if in[i] != out[i] {
				fmt.Println("NO")
				return
			}
		}
	}

	min := out[begin]
	for i := begin; i < end; i++ {
		if out[i] < min {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}
