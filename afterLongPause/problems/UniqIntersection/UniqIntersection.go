/*

make an alphabet -> compare input intersection with alphabet
-> make uniqintersection

*/

package UniqIntersection

func MakeAnAlphabet(inputArray []int) (alphbet []int) {

	var min int = inputArray[0]
	var max int = inputArray[0]

	for _, v := range inputArray {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}

	}

	for i := min; i <= max; i++ {
		alphbet = append(alphbet, i)
	}

	return
}

func MakeAnIntersection(inputArray1 []int, inputArray2 []int) (intersection []int) {

	for i := 0; i < len(inputArray1); i++ {
		for j := 0; j < len(inputArray2); j++ {
			if inputArray1[i] == inputArray2[j] {
				intersection = append(intersection, inputArray1[i])
			}
		}
	}

	return
}

func MakeAnUniqIntersection(inputIntersection []int, inputAlphabet []int) (uniqIntersection []int) {
	/*
		take an element from inputIntersection and compare it to all elements in inputAlphabet, if
		this element is here in inputAlphabet then put it to uniqIntersection and delete this element from
		inputAlphabet (to escape the situation where we meet this element twice (or more) while passing throw
		inputIntersection)
	*/
	return
}
