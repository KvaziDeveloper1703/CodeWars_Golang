/*
Write a function maxSequence that takes an array of integers and finds the maximum sum of any contiguous subarray.

Conditions:
+ If the array contains only positive numbers, return the sum of the entire array.
+ If the array contains only negative numbers or is empty, return 0.
+ For mixed arrays, find the subarray with the maximum sum.

Example:
Input: [-2, 1, -3, 4, -1, 2, 1, -5, 4]
Output: 6 (subarray: [4, -1, 2, 1])

Напишите функцию maxSequence, которая принимает массив целых чисел и находит максимальную сумму любой непрерывной подпоследовательности.

Условия:
+ Если массив содержит только положительные числа, вернуть сумму всего массива.
+ Если массив состоит только из отрицательных чисел или пустой, вернуть 0.
+ Для смешанных массивов найти подпоследовательность с максимальной суммой.

Пример:
Ввод: [-2, 1, -3, 4, -1, 2, 1, -5, 4]
Вывод: 6 (подпоследовательность: [4, -1, 2, 1])

https://www.codewars.com/kata/54521e9ec8e60bc4de000d6c
*/

package main

import (
	"fmt"
)

func maxSequence(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	maxSum := 0
	currentSum := 0

	for _, num := range arr {
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
