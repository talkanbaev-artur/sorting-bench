package algorithms_test

import (
	"math/rand"
	"testing"

	"github.com/talkanbaev-artur/sorting-bench/algorithms"

	"github.com/stretchr/testify/assert"
)

var (
	testArray        = []int{56, -100, 98, 1, 7, 42, -0, 0, 19, -400, 789}
	testCorrectArray = []int{-400, -100, 0, 0, 1, 7, 19, 42, 56, 98, 789}
)

func TestQuickSort(t *testing.T) {
	sorted := algorithms.QuickSort(testArray)
	assert.Equal(t, testCorrectArray, sorted)
}

func TestMergeSort(t *testing.T) {
	sorted := algorithms.Mergesort(testArray)
	assert.Equal(t, testCorrectArray, sorted)
}

func createArray(n int) []int {
	var res []int
	for i := 0; i < n; i++ {
		res = append(res, rand.Intn(2000000))
	}
	return res
}

var arr1m = createArray(1000000)

func BenchmarkQuickSort1M(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Run("Quick", func(b *testing.B) {
			algorithms.QuickSort(arr1m)
		})
	}
}

func BenchmarkMergeSort1M(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Run("Merge", func(b *testing.B) {
			algorithms.Mergesort(arr1m)
		})
	}
}

var arr1b = createArray(100000000)

func BenchmarkQuickSort100M(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Run("Quick", func(b *testing.B) {
			algorithms.QuickSort(arr1b)
		})
	}
}

func BenchmarkMergeSort100M(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Run("Merge", func(b *testing.B) {
			algorithms.Mergesort(arr1b)
		})
	}
}
