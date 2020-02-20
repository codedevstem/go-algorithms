package selectionSort

// IntSelectionSort has an O(n) complexity of ~(N^2)/2. Not brilliant, but acceptable
func IntSelectionSort(array []int) []int {
	length := len(array)
	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if array[j] < array[min] {
				min = j
			}
		}
		array[i], array[min] = array[min], array[i]
	}
	return array
}
