/*
Write a function that takes two lists, a and b, and returns a new list that contains all the elements of a that are not present in b, while maintaining their original order.

If a value is present in b, all occurrences of that value must be removed from a.

Examples:
Input: array_diff({1, 2}, 2, {1}, 1, *z) → Output: {2} (z == 1)
Input: array_diff({1, 2, 2, 2, 3}, 5, {2}, 1, *z) → Output: {1, 3} (z == 2)

Напишите функцию, которая принимает два списка, a и b, и возвращает новый список, содержащий все элементы из a, которые отсутствуют в b, сохраняя их исходный порядок.

Если значение присутствует в b, все его вхождения должны быть удалены из a.

Примеры:
Ввод: array_diff({1, 2}, 2, {1}, 1, *z) → Вывод: {2} (z == 1)
Ввод: array_diff({1, 2, 2, 2, 3}, 5, {2}, 1, *z) → Вывод: {1, 3} (z == 2)

https://www.codewars.com/kata/523f5d21c841566fde000009
*/

package main

import (
	"fmt"
)

func ArrayDiff(a, b []int) []int {
	result := []int{}
	bSet := make(map[int]bool)

	for _, num := range b {
		bSet[num] = true
	}

	for _, num := range a {
		if !bSet[num] {
			result = append(result, num)
		}
	}

	return result
}

func main() {
	a1 := []int{1, 2}
	b1 := []int{1}
	fmt.Println(ArrayDiff(a1, b1))

	a2 := []int{1, 2, 2, 2, 3}
	b2 := []int{2}
	fmt.Println(ArrayDiff(a2, b2))
}
