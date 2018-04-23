package selection

import (
	"os"

	"github.com/op/go-logging"
	"github.com/sbouchardet/algorithms/sort"
)

func getLogger(name string) *logging.Logger {
	var format = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{callpath} ▶ %{level} %{id:03x}%{color:reset}: %{message}`,
	)
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)
	backend2Formatter := logging.NewBackendFormatter(backend2, format)
	logging.SetBackend(backend2Formatter)
	return logging.MustGetLogger(name)
}

// MedianOfMedian Oraculo
func MedianOfMedian(a []int) int {
	logger := getLogger("MedianOfMedian")
	sizeGroup := 5
	medians := []int{}
	aux := []int{}
	for i := 0; i < len(a); i++ {
		if i%sizeGroup == 0 && len(aux) > 0 {
			aux = sort.MergeSort(aux)
			medians = append(medians, median(aux))
			aux = []int{}
		} else {
			aux = append(aux, a[i])
		}
	}
	if len(aux) > 0 {
		medians = append(medians, median(aux))
	}

	logger.Debugf("medians: %v\n", medians)

	if len(medians) <= sizeGroup {
		return median(medians)
	}

	return MedianOfMedian(medians)

}

func median(vec []int) int {
	n := len(vec)
	return vec[int(n/2)]
}

func partition(input []int, m int) ([]int, []int) {
	R := []int{}
	L := []int{}
	for i := 0; i < len(input); i++ {
		if input[i] < m {
			L = append(L, input[i])
		}
		if input[i] > m {
			R = append(R, input[i])
		}
	}
	return L, R
}

// ExotericSelect Select the k-esim element

func ExotericSelect(input []int, k int) int {

	logger := getLogger("ExotericSelect")

	mom := MedianOfMedian(input)
	L, R := partition(input, mom)

	if len(L) == k-1 {
		logger.Debugf("Find %d element: %d", k, mom)
		return mom
	} else if len(L) > k-1 {
		logger.Debugf("call ExotericSelect to L")
		return ExotericSelect(L, k)
	}

	logger.Debugf("call ExotericSelect to R")
	return ExotericSelect(R, k-len(L)-1)

}
