package main

import (
	"fmt"

	"github.com/talkanbaev-artur/sorting-bench/algorithms"
)

func main() {
	arr := []int{5, 6, 2, 3, 1, 9}
	fmt.Printf("Mergesort: %+v\n", algorithms.Mergesort(arr))
	fmt.Printf("Quicksort: %+v\n", algorithms.QuickSort(arr))
}
