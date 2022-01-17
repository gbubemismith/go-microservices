package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	elements := []int{9, 7, 6, 5}

	elements = BubbleSort(elements)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 4, len(elements))
	assert.EqualValues(t, 5, elements[0])
	assert.EqualValues(t, 6, elements[1])
	assert.EqualValues(t, 7, elements[2])
	assert.EqualValues(t, 9, elements[3])

}
