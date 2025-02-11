/*
Given an array of integers, your task is to find and return the smallest integer in the array.

Дан массив целых чисел. Твоя задача — найти и вернуть наименьшее число в этом массиве.

https://www.codewars.com/kata/55a2d7ebe362935a210000b2
*/

package main

import "fmt"

func SmallestIntegerFinder(givenNumbers []int) int {
	if len(givenNumbers) == 0 {
		return 0
	}

	minimum := givenNumbers[0]

	for _, number := range givenNumbers {
		if number < minimum {
			minimum = number
		}
	}

	return minimum
}

func main() {
	numbers := []int{4, 15, 88, 2}
	result := SmallestIntegerFinder(numbers)
	fmt.Println(result)
}
