package main

import "fmt"

func main() {
	var smallNum int
	var largeNum int
	fmt.Println("Enter a number:")
	fmt.Scan(&smallNum);
	fmt.Println("Enter a larger number:")
	fmt.Scan(&largeNum);
	fmt.Println("The remainder is: ", largeNum % smallNum)
}
