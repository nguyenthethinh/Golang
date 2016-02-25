package main

import (
	"fmt"
)

func findGreatestNumber(nums ...int) int {
	fmt.Println("Input number:", nums, " , ")
	max := 0
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	max := findGreatestNumber(1, 4, 2, 5, 3, 9)
	fmt.Println("Max: ", max)
	numbers := []int{4, 0, -1, 9, 8, 20}
	max = findGreatestNumber(numbers...)
	fmt.Println("Max: ", max)
}
