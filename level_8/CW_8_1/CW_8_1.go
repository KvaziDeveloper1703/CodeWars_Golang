/*
Write a program that takes an integer as input and returns "Even" if the number is even or "Odd" if the number is odd.

Напишите программу, которая принимает целое число и возвращает "Even", если число чётное, или "Odd", если число нечётное.

https://www.codewars.com/kata/53da3dbb4a5168369a0000fe
*/

package main

import "fmt"

func EvenOrOdd(givenNumber int) string {
	if givenNumber%2 == 0 {
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
