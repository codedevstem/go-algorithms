package linearSearch

import "testing"

func TestIntLinearSearch(t *testing.T) {
	target := 45
	targetTrueIndex := 10
	a := []int{5, 2, 7, 5, 12, 8, 9, 2, 5, 7, 45, 2, 6, 5, 9, 6, 3, 1, 6, 54, 3, 45, 6, 54, 3, 45, 65, 43, 5, 68, 74, 56, 765, 47, 654, 8675}
	match, index, err := IntLinearSearch(a, target)
	if err != nil {
		t.Errorf("LinearSearch failed, Should have found the element")
	} else if match != target {
		t.Errorf("LinearSearch failed, found wrong element\n."+
			"Item was %d, but should have been %d", match, target)
	} else if index != targetTrueIndex {
		t.Errorf("LinearSearch failed, found right element, but wrong index.\n"+
			"Index was %d, but was %d", index, targetTrueIndex)
	}
}

func TestIntLinearSearchNotFoundError(t *testing.T) {
	target := 10
	a := []int{1, 2, 3, 4, 5}
	_, _, err := IntLinearSearch(a, target)
	if err == nil {
		t.Errorf("LinearSearch failed, Should not have found element: %d", target)
	}
}
