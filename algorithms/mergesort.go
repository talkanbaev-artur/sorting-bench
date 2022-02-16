package algorithms

func Mergesort(items []int) []int {
	if len(items) > 1 {
		m := len(items) / 2
		L := items[:m]
		R := items[m:]
		items = merge(Mergesort(R), Mergesort(L))
	}
	return items
}

func merge(L, R []int) []int {
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
	for i < len(L) {
		A[k] = L[i]
		i++
		k++
	}
	for j < len(R) {
		A[k] = R[j]
		j++
		k++
	}
	return A
}
