package task6

import "fmt"

func SolutionOfTask6() {
	var perfumeQuantity int
	var numberOfScreams int

	fmt.Scan(&perfumeQuantity, &numberOfScreams)
	perfumes := make([][]int, perfumeQuantity)
	for i := range perfumes {
		perfumes[i] = make([]int, 1)
		perfumes[i][0] = i + 1
	}

	amountOfGroup := make(map[int]int)
	for i := 0; i < perfumeQuantity; i++ {
		amountOfGroup[i+1] = 1
	}

	for i := 0; i < numberOfScreams; i++ {
		var char int
		fmt.Scan(&char)
		if char == 1 {
			var a, b int
			fmt.Scan(&a, &b)

			first := 0
			second := 0

			for j := 0; j < len(perfumes); j++ {
				for k := 0; k < len(perfumes[j]); k++ {
					if perfumes[j][k] == a {
						first = j
					}
					if perfumes[j][k] == b {
						second = j
					}
				}
				if first != 0 && second != 0 {
					break
				}
			}
			if first != second {
				for i := range perfumes[first] {
					amountOfGroup[perfumes[first][i]]++
				}

				for i := range perfumes[second] {
					perfumes[first] = append(perfumes[first], perfumes[second][i])
					amountOfGroup[perfumes[second][i]]++
				}
				perfumes = append(perfumes[:second], perfumes[(second+1):]...)
			}
		}
		if char == 2 {
			var a, b int
			fmt.Scan(&a, &b)
			result := 0
			check := true
			for j := 0; j < len(perfumes); j++ {
				for k := 0; k < len(perfumes[j]); k++ {
					if perfumes[j][k] == a {
						result++
					}
					if perfumes[j][k] == b {
						result++
					}
				}
				if result == 2 {
					check = false
					break
				}
				result = 0
			}
			if check {
				fmt.Println("NO")
			} else {
				fmt.Println("YES")
			}
		}
		if char == 3 {
			var a int
			fmt.Scan(&a)
			num, _ := amountOfGroup[a]
			fmt.Println(num)
		}
	}
}
