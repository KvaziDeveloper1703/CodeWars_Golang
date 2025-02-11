/*
Write a function that takes an array of numbers, squares each number, and then returns the sum of the squared values.

Напишите функцию, которая принимает массив чисел, возводит каждое число в квадрат и возвращает сумму квадратов.

https://www.codewars.com/kata/515e271a311df0350d00000f
*/

package main

import "fmt"

func SquareSum(givenNumbers []int) int {
	sum := 0
	for _, number := range givenNumbers {
		sum += number * number
	}
	return sum
}

func main() {
	numbers := []int{1, 2, 2}
	result := SquareSum(numbers)
	fmt.Println(result)
}
