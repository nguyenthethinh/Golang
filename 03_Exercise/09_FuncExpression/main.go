package main

import "fmt"

func main() {
	fmt.Println("Enter a number: ")
	var num int
	fmt.Scan(&num)
	twoValue := func(num int) (int, bool) {
		return num / 2, num%2 == 0
	}
	fmt.Println(twoValue(num))
}
