package ch6
import "fmt"

type Heap struct {
	A            []int
	HeapSize     int
}

// Doing heap things
func leftChild(i int) int {
	return (2 * i) + 1
}

func rightChild(i int) int {
	return (2 * i) + 2
}

func (h *Heap) MaxHeapify(a []int, index int) {
	var largest int
	heapSize := h.HeapSize
	left := leftChild(index)
	right := rightChild(index)

	if left <= heapSize && h.A[left] > h.A[index] {
		largest = left
	} else {
		largest = index
	}

	if right > len(a) - 1 {
		fmt.Println("Index error -- no node there")
	} else if right <= heapSize && h.A[right] > h.A[largest] {
		largest = right
	}

	if largest != index {
		temp_i := h.A[largest]
		temp_l := h.A[index]
		h.A[index] = temp_i
		h.A[largest] = temp_l
		h.MaxHeapify(h.A, largest)
	}
}

func (h *Heap) BuildMaxHeap() {
	fmt.Println("Starting tree:")
	fmt.Println(h.A)
	for i := (len(h.A) / 2) - 1; i >= 0; i-- {
		h.MaxHeapify(h.A, i)
	}
	fmt.Println("Ending tree:")
	fmt.Println(h.A)
}
