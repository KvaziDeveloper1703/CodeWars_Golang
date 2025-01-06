/*
Write a function maxSequence that takes an array of integers and returns the maximum sum of any contiguous subarray.

Rules:
+ If the array contains only positive numbers, the maximum sum is the sum of the entire array.
+ If the array contains only negative numbers or is empty, return 0.
+ The function must work for arrays with both positive and negative numbers.

Example:
Input: [-2, 1, -3, 4, -1, 2, 1, -5, 4]
Output: 6 (subarray: [4, -1, 2, 1])

Напишите функцию maxSequence, которая принимает массив целых чисел и возвращает максимальную сумму любой непрерывной подпоследовательности.

Правила:
+ Если массив состоит только из положительных чисел, максимальная сумма равна сумме всего массива.
+ Если массив состоит только из отрицательных чисел или пустой, вернуть 0.
+ Функция должна работать для массивов, содержащих как положительные, так и отрицательные числа.

Пример:
Ввод: [-2, 1, -3, 4, -1, 2, 1, -5, 4]
Вывод: 6 (подмассив: [4, -1, 2, 1])

https://www.codewars.com/kata/54521e9ec8e60bc4de000d6c
*/

package main

import (
	"fmt"
)

func maxSequence(numbers []int) int {
	maxSum := 0
	currentSum := 0

	for _, num := range numbers {
		currentSum += num
		if currentSum < 0 {
			currentSum = 0
		}
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}

func main() {
	fmt.Println(maxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSequence([]int{-1, -2, -3}))
	fmt.Println(maxSequence([]int{1, 2, 3, 4}))
	fmt.Println(maxSequence([]int{}))
}
