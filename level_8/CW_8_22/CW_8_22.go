/*
Given a non-empty array of integers, return the result of multiplying all elements together in order.
Example: [1, 2, 3, 4] → 1 × 2 × 3 × 4 = 24

Дан непустой массив целых чисел. Необходимо вернуть результат перемножения всех элементов массива по порядку.
Пример: [1, 2, 3, 4] → 1 × 2 × 3 × 4 = 24
*/

package main

func Grow(arr []int) int {
	result := 1
	for _, v := range arr {
		result *= v
	}
	return result
}