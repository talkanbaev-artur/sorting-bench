package algorithms

import "log"

func Mergesort(items []int) []int {
	if len(items) > 1 {
		log.Println("Starting mergesort partition")
		m := len(items) / 2
		L := items[:m]
		R := items[m:]
		items = merge(Mergesort(R), Mergesort(L))
	}
	return items
}

func merge(L, R []int) []int {
	log.Println("Allocating space for merge")
	A := make([]int, len(L)+len(R))
	i, j, k := 0, 0, 0
	for i < len(L) && j < len(R) {
		if L[i] < R[j] {
			A[k] = L[i]
			i++
		} else {
			A[k] = R[j]
			j++
		}
		k++
	}
	log.Println("merging left side")
	for i < len(L) {
		A[k] = L[i]
		i++
		k++
	}
	log.Println("merging right side")
	for j < len(R) {
		A[k] = R[j]
		j++
		k++
	}
	return A
}
