package bubbleSort

import (
	"testing"
)

func TestBubbleSortWithReversedArray(t *testing.T) {
	array := []int{4, 3, 2, 1}
	IntBubbleSort(array)
	if array[0] > array[1] ||
		array[0] > array[2] ||
		array[0] > array[3] ||
		array[1] > array[2] ||
		array[1] > array[3] ||
		array[2] > array[3] {
		t.Errorf("Sort failed.\n"+
			"Order should have been acending but was %v", array)
	}
}
