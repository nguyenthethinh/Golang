package main

import "fmt"

func main() {
	var x [5]int
	x[0] = 1
	x[1] = 2
	x[2] = 3
	x[3] = 4
	x[4] = 100
	fmt.Println(x[3])
	var total float64 =0;
	for i := 0; i < len(x) ; i++  {
		total += float64(x[i]);
	}
	fmt.Println(total)
}

