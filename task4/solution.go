package task4

/*
package main

import (
	"fmt"
	"sort"
)

var requiredAmount int
var numberOfNominals int
var nominals []int
var res []int

func main() {
	fmt.Scan(&requiredAmount, &numberOfNominals)

	for i := 0; i < numberOfNominals; i++ {
		var a int
		fmt.Scan(&a)
		nominals = append(nominals, a)
	}

	sort.Slice(nominals, func(i, j int) bool {
		return nominals[i] > nominals[j]
	})
	fmt.Println(nominals)

	result := make([]int, 0)
	used := make([]int, numberOfNominals)

	if Solve(result, used) {
		fmt.Println(len(res))
		fmt.Println(res)
	} else {
		fmt.Println(-1)
	}
}

func Solve(result []int, used []int) bool {
	if sum(result) == requiredAmount {
		res = result
		return true
	}

	if sum(used) == 2*numberOfNominals {
		return false
	}

	for i, nom := range nominals {
		if used[i] != 2 {
			nused := make([]int, numberOfNominals)
			copy(nused, used)
			nresult := make([]int, len(result))
			copy(nresult, result)
			nused[i]++
			nresult = append(nresult, nom)
			if Solve(nresult, nused) {
				return true
			}
		}
	}
	return false
}

func sum(arr []int) int {
	sum := 0
	for _, elem := range arr {
		sum += elem
	}
	return sum
}

*/
