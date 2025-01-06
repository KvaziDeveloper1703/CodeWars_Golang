/*
Write a function that takes an array of numbers, squares each number, and then returns the sum of the squared values.

For example:
Input: [1, 2, 2] → Output: 9

Напишите функцию, которая принимает массив чисел, возводит каждое число в квадрат и возвращает сумму квадратов.

Например:
Ввод: [1, 2, 2] → Вывод: 9

https://www.codewars.com/kata/515e271a311df0350d00000f
*/

package main

import "fmt"

func SquareSum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number * number
	}
	return sum
}

func main() {
	fmt.Println(SquareSum([]int{1, 2, 2}))
	fmt.Println(SquareSum([]int{0, 3, 4, 5}))
	fmt.Println(SquareSum([]int{}))
}
