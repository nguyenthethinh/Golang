package main

import "fmt"

func half (num int)(int, bool){
	return num/2, num % 2 == 0
}

func main() {
	fmt.Println("Enter a number: ")
	var num int
	fmt.Scan(&num)
	d, e := half(num)
	fmt.Println(d, e)
}
