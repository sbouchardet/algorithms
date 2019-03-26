package sort

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/op/go-logging"
)

// TODO

type heap struct {
	size   int
	vector []int
	log    *logging.Logger
}

func (h *heap) indexFather(index int) int {
	return int(index / 2)
}

func (h *heap) indexSonLeft(index int) int {
	return (2 * (index + 1)) - 1
}

func (h *heap) indexSonRight(index int) int {
	return h.indexSonLeft(index) + 1
}

func (h *heap) removeMax() int {
	removed := h.vector[0]
	h.changeValues(0, h.size-1)
	h.size = h.size - 1
	h.down(0)
	return removed
}

func (h *heap) down(indexToDown int) {

	L := h.indexSonLeft(indexToDown)
	R := h.indexSonRight(indexToDown)

	maxIndex := indexToDown
	if L < h.size && h.vector[L] > h.vector[maxIndex] {
		maxIndex = L
	}

	if R < h.size && h.vector[R] > h.vector[maxIndex] {
		maxIndex = R
	}

	if maxIndex != indexToDown {
		h.log.Debugf("change value v[%d]=%d e v[%d]=%d\n",
			maxIndex,
			h.vector[maxIndex],
			indexToDown,
			h.vector[indexToDown])
		h.changeValues(maxIndex, indexToDown)
		h.log.Debugf("call descer(%d)\n\n", maxIndex)
		h.down(maxIndex)
	}

}

func (h *heap) changeValues(a int, b int) []int {
	aux := h.vector[a]
	h.vector[a] = h.vector[b]
	h.vector[b] = aux
	return h.vector
}

func (h *heap) heapfyMax() {
	for i := int(h.size / 2); i >= 0; i-- {
		h.down(i)
	}
}

func (h *heap) getLevel(i int) int {
	return int(math.Log2(float64(i + 1)))
}

func (h *heap) printHeap() {
	level := 0
	for i := 0; i < h.size; i++ {
		if level < h.getLevel(i) {
			level = h.getLevel(i)
			fmt.Print("\n")
		}
		countSpaces := float64(h.size) / math.Pow(2, float64(level+1))
		spaces := strings.Repeat(" ", int(countSpaces))

		fmt.Printf("%s%d", spaces, h.vector[i])
	}
}

func createHeap(input []int, loglevel int) *heap {
	var format = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{callpath} ▶ %{level} %{id:03x}%{color:reset}: %{message}`,
	)
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)
	backend2Formatter := logging.NewBackendFormatter(backend2, format)
	logging.SetBackend(backend2Formatter)
	return &heap{len(input), input, logging.MustGetLogger("HeapSort")}
}

// Heapsort sort
func Heapsort(input []int, loglevel int) []int {

	var heapS = createHeap(input, loglevel)
	heapS.heapfyMax()

	for i := 0; i < len(input); i++ {
		heapS.removeMax()
	}

	return heapS.vector
}
