package main

import "fmt"

func main() {
	fmt.Println("Even number from 0 to 100:")
	var i int
	for i = 0; i <= 100;i++  {
		if (i % 2 == 0){
			fmt.Print(i, " , ")
		}
	}
}
