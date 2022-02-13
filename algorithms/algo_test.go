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

func BenchmarkQuickSort100(b *testing.B) {
	arr := createArray(100000)

	for i := 0; i < b.N; i++ {
		b.Run("Name-100", func(b *testing.B) {
			algorithms.QuickSort(arr)
		})
	}
}

func BenchmarkMergeSort100(b *testing.B) {
	arr := createArray(100000)

	for i := 0; i < b.N; i++ {
		b.Run("Name-100", func(b *testing.B) {
			algorithms.Mergesort(arr)
		})
	}
}
