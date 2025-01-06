/*
Write a function that takes an array of numbers as input and returns the sum of all positive numbers in the array.

For example:
Input: [1, -4, 7, 12]
Output: 20

Напишите функцию, которая принимает массив чисел и возвращает сумму всех положительных чисел в массиве.

Например:
Ввод: [1, -4, 7, 12]
Вывод: 20

https://www.codewars.com/kata/5715eaedb436cf5606000381
*/

package main

import "fmt"

func SumOfPositives(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		if number > 0 {
			sum += number
		}
	}
	return sum
}

func main() {
	arr := []int{1, -4, 7, 12}
	fmt.Println(SumOfPositives(arr))

	arr = []int{-11, -12, -13}
	fmt.Println(SumOfPositives(arr))

	arr = []int{}
	fmt.Println(SumOfPositives(arr))
}
