package linearSearch

import (
	"fmt"
)

// IntLinearSearch finds the first matching item (t) from array (a) in O(n/2) time-complexity.
// If found it returns the item and index
// If not found, it returns an error.
func IntLinearSearch(a []int, t int) (int, int, error) {
	for i := 0; i < len(a); i++ {
		if a[i] == t {
			return a[i], i, nil
		}
	}
	return 0, 0, fmt.Errorf("could not find %d in array", t)
}
