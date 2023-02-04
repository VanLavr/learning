/*

arr1 = 1 2 3 4 5  6  7  8  9 5
arr2 = 3 4 5 6 11 45 76 28 0 0

resarr = 3 4 5 6

*/

package UniqIntersection

import (
	"fmt"
)

func UniqIntersection(array1 []int, array2 []int) (resultArray []int) {

	for i := 0; i < len(array1); i++ {
		for j := 0; j < len(array2); j++ {
			if array1[i] == array2[j] {
				resultArray = append(resultArray, array1[i])
			}
		}
	}

	fmt.Printf("indexes                  [%d %d %d %d %d %d  %d  %d  %d  %d]", 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Print("\nnot unique intersection: ")
	fmt.Println(resultArray)
	resultArray = DeleteDuplicate(resultArray)

	return
}

func DeleteDuplicate(array []int) []int {
	var set []int

	for i := 0; i < len(array); i++ {
		var inputCounter int
		for j := 0; j < len(array); j++ {
			if array[i] != array[j] && i != j {
				inputCounter++
			}

			if inputCounter == len(array) - 1 {
				set = append(set, array[i])
			}
		}
	}



	return set
}