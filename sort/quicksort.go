package sort

import "math/rand"

// Quicksort sort
func Quicksort(input []int) []int {
	n := len(input)

	if n <= 1 {
		return input
	}

	p := input[rand.Intn(n)]
	L, P, R := partition(input, p)

	LSorted := Quicksort(L)
	RSorted := Quicksort(R)

	concatAux := append(LSorted, P...)
	return append(concatAux, RSorted...)

}

func partition(vector []int, pivo int) ([]int, []int, []int) {
	L, P, R := []int{}, []int{}, []int{}
	for i := 0; i < len(vector); i++ {
		if vector[i] > pivo {
			R = append(R, vector[i])
		}
		if vector[i] < pivo {
			L = append(L, vector[i])
		}
		if vector[i] == pivo {
			P = append(P, vector[i])
		}
	}

	return L, P, R
}
