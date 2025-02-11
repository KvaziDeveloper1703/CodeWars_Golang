/*
Write a function that takes a number as input and ensures it is negative. If the number is already negative, return it as is.

Напишите функцию, которая принимает число и преобразует его в отрицательное. Если число уже отрицательное, просто верните его.

https://www.codewars.com/kata/55685cd7ad70877c23000102
*/

package main

import "fmt"

func MakeNegative(givenNumber int) int {
	if givenNumber > 0 {
		return -givenNumber
	}
	return givenNumber
}

func main() {
	number := 1703
	result := MakeNegative(number)
	fmt.Println(result)
}
