package algorithms_test

import (
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
