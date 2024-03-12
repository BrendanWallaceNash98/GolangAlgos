package Packages

func Swap(numbers []int, i, j int) {
	// numbers is the list and i and j are the indexes for the windows

	temp := numbers[i]
	numbers[i] = numbers[j]
	numbers[j] = temp

}

func BubbleSort(numbers []int) []int {

	for j := 0; j < len(numbers); j++ {
		for i := 0; i < len(numbers)-1; i++ {
			if numbers[i] > numbers[i+1] {
				Swap(numbers, i, i+1)
			}
		}
	}
	return numbers
}
