package bubbleSort

// BubbleSort is a very basic sorting algorithm that has a Big(O) notation of n^2. horrible and should not be used.

// IntBubbleSort sorts an array of type int
func IntBubbleSort(array []int) []int {
	lengthOfArray := len(array)
	for i := 0; i < lengthOfArray/2; i++ {
		for j := lengthOfArray - 1; j > i; j-- {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	return array
}
