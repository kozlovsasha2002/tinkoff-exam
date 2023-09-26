package task1

import "fmt"

func SolutionOfTask1() {
	var n, s int
	fmt.Scan(&n, &s)

	maxCost := 0
	for i := 0; i < n; i++ {
		weaponCost := 0
		fmt.Scan(&weaponCost)
		if weaponCost > maxCost && weaponCost <= s {
			maxCost = weaponCost
		}
	}
	if maxCost != 0 {
		fmt.Println(maxCost)
	} else {
		fmt.Println(0)
	}
}
