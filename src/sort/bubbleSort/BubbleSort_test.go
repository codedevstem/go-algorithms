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

func TestBubbleSortWithRandomArray(t *testing.T) {
	array := []int{2, 4, 1, 3}
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

func TestBubbleSortWithPreorderedArray(t *testing.T) {
	array := []int{2, 4, 1, 3}
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
