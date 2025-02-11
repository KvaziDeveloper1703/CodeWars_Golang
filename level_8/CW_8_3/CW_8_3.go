/*
Write a function that takes an array of numbers as input and returns the sum of all positive numbers in the array.

Напишите функцию, которая принимает массив чисел и возвращает сумму всех положительных чисел в массиве.

https://www.codewars.com/kata/5715eaedb436cf5606000381
*/

package main

import "fmt"

func SumOfPositives(givenNumbers []int) int {
	sum := 0
	for _, number := range givenNumbers {
		if number > 0 {
			sum += number
		}
	}
	return sum
}

func main() {
	array := []int{1, -4, 7, 12}
	result := SumOfPositives(array)
	fmt.Println(result)
}
