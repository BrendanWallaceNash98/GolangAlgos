package main

import (
	"GolangAlgos/Packages"
	"fmt"
)



func main() {
	numList := []int{7, 5, 6, 2, 4, 5}

	sortList := Packages.BubbleSort(numList)

	fmt.Println(sortList)
}
