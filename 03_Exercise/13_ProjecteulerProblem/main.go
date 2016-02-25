/*
* https://projecteuler.net/problem=16
* Power digit sum Problem 16
*     2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.
*     What is the sum of the digits of the number 21000?
 */
package main

import (
	"fmt"
)

func sumOfDigit(n int) int {
	total := 0
	for n > 0 {
		total += n % 10
		n = n / 10
	}
	return total
}

func main() {
	fmt.Print("Enter number:")
	var number int
	fmt.Scan(&number)
	fmt.Println("Sum of digits is ", sumOfDigit(number))

}
