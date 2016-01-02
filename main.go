package main

import (
	"fmt"
	"github.com/cyarie/algorithms/ch2"
	"github.com/cyarie/algorithms/ch6"
)

func main() {
	arr := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	var a []int = arr
	fmt.Println(a)

	ch2.InsertSort(a)

	fmt.Println(a)

	fmt.Println("Doin' heap thangs...")
	heapArr := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	th := ch6.Heap{
		A: heapArr,
		HeapSize: len(arr),
	}

	th.BuildMaxHeap()
}
