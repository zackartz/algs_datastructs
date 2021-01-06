package main

import (
	"algs_datastructs/datastructs"
	"fmt"
)

func main() {
	m := &datastructs.MaxHeap{}
	fmt.Printf("%d", m)

	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Printf("%d\n", m)
	}

	for _, _ = range buildHeap {
		m.Extract()
		fmt.Printf("%d\n", m)
	}
}
