package selectionSort

import "testing"

func TestSelectionSortWithReversedArray(t *testing.T) {
	array := []int{4, 3, 2, 1}
	IntSelectionSort(array)
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

func TestSelectionSortWithRandomArray(t *testing.T) {
	array := []int{2, 4, 1, 3}
	IntSelectionSort(array)
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
	IntSelectionSort(array)
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

func TestSelectionSortStillHasAllNumbers(t *testing.T) {
	array := []int{2, 4, 1, 3}
	IntSelectionSort(array)
	if !contains(array, 1) ||
		!contains(array, 2) ||
		!contains(array, 3) ||
		!contains(array, 4) {
		t.Errorf("Sort failed.\n"+
			"Contained these items: %v.\n"+
			"Should have had 1,2,3,4", array)
	}
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
