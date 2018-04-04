package sort

// BubbleSort sort
func BubbleSort(input []int) []int {

	n := len(input)
	for i := 0; i < n-1; i++ {
		for j := i; j < n-1; j++ {
			if input[j] > input[j+1] {
				aux := input[j]
				input[j] = input[j+1]
				input[j+1] = aux
			}
		}
	}

	return input
}
