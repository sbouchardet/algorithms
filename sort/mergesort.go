package sort

//MergeSort sort
func MergeSort(input []int) []int {
	n := len(input)
	if n == 1 {
		return input
	}

	m := n / 2
	inputL := input[:m]
	inputR := input[m:]

	LSorted := MergeSort(inputL)
	RSorted := MergeSort(inputR)

	output := merge(LSorted, RSorted)

	return output

}

func merge(left []int, right []int) []int {
	nLeft, nRight := len(left), len(right)
	countL, countR := 0, 0
	output := []int{}

	for countL < nLeft && countR < nRight {
		if left[countL] < right[countR] {
			output = append(output, left[countL])
			countL++
		} else {
			output = append(output, right[countR])
			countR++
		}
	}

	if countL < nLeft {
		output = append(output, left[countL:]...)
	} else if countR < nRight {
		output = append(output, right[countR:]...)
	}

	return output

}
