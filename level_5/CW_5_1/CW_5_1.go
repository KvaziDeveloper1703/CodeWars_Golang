/*
Write a function that takes an array of integers and moves all the zeros to the end of the array while preserving the order of the non-zero elements.

Example:
Input: [1, 2, 0, 1, 0, 1, 0, 3, 0, 1]
Output: [1, 2, 1, 1, 3, 1, 0, 0, 0, 0]

Напишите функцию, которая принимает массив целых чисел и перемещает все нули в конец массива, сохраняя порядок остальных элементов.

Пример:
Ввод: [1, 2, 0, 1, 0, 1, 0, 3, 0, 1]
Вывод: [1, 2, 1, 1, 3, 1, 0, 0, 0, 0]

https://www.codewars.com/kata/52597aa56021e91c93000cb0
*/

package main

import (
	"fmt"
)

func MoveZeros(arr []int) []int {
	result := make([]int, len(arr))
	idx := 0

	for _, num := range arr {
		if num != 0 {
			result[idx] = num
			idx++
		}
	}

	return result
}

func main() {
	input := []int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1}
	output := MoveZeros(input)
	fmt.Println(output)
}
