/*
Create a function that takes an integer as an argument and returns "Even" for even numbers or "Odd" for odd numbers.

https://www.codewars.com/kata/53da3dbb4a5168369a0000fe
*/

package main

import "fmt"

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	}
	return "Odd"
}

func main() {
	number := 4
	result := EvenOrOdd(number)
	fmt.Printf("The number %d is %s.\n", number, result)

	number = 9
	result = EvenOrOdd(number)
	fmt.Printf("The number %d is %s.\n", number, result)
}
