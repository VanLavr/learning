package main

import (
	"fmt"
	"myProblems/UniqIntersection"
)

func main() {

	var arr1 []int
	var arr2 []int

	for i := 0; i < 10; i++ {
		arr1 = append(arr1, i + 1)
		arr2 = append(arr2, i + 5)
	}

	arr1 = append(arr1, 13)
	arr1 = append(arr1, 14)
	arr1 = append(arr1, 13)
	arr1 = append(arr1, 14)

	for i := 0; i < len(arr1); i++ {
		fmt.Printf("array1[%d] = %d\n", i, arr1[i])
	}

	fmt.Println()
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("array2[%d] = %d\n", i, arr2[i])
	}

	arrRes := UniqIntersection.UniqIntersection(arr1, arr2)
	fmt.Println("\nIntersection: ")
	fmt.Println(arrRes)
}