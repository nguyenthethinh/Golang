package main

import "fmt"

func main() {
	trueOrFalse := (true && false) || (false && true) || !(false && false)
	fmt.Println("The value of expression is: ", trueOrFalse)
}
