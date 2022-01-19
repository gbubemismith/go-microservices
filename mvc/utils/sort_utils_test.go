package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortWorstCase(t *testing.T) {
	elements := []int{9, 7, 6, 5}

	elements = BubbleSort(elements)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 4, len(elements))
	assert.EqualValues(t, 5, elements[0])
	assert.EqualValues(t, 6, elements[1])
	assert.EqualValues(t, 7, elements[2])
	assert.EqualValues(t, 9, elements[3])

}

func TestBubbleSortBestCase(t *testing.T) {
	elements := []int{5, 6, 7, 9}

	elements = BubbleSort(elements)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 4, len(elements))
	assert.EqualValues(t, 5, elements[0])
	assert.EqualValues(t, 6, elements[1])
	assert.EqualValues(t, 7, elements[2])
	assert.EqualValues(t, 9, elements[3])

}

func TestBubbleSortNilSlice(t *testing.T) {
	BubbleSort(nil)
}

func BenchmarkBubbleSort10(b *testing.B) {
	elements := getElements(10)
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkBubbleSort1000(b *testing.B) {
	elements := getElements(1000)
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func getElements(n int) []int {
	result := make([]int, n)
	i := 0
	for j := n - 1; j >= 0; j-- {
		result[i] = j
		i++
	}
	return result
}
