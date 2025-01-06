/*
Write a function that takes an array of integers as input. The array will always have at least three elements and will consist entirely of either odd or even integers, except for one "outlier" integer. The function should identify and return this outlier.

Examples:
Input: [2, 4, 0, 100, 4, 11, 2602, 36] → Output: 11
Input: [160, 3, 1719, 19, 11, 13, -21] → Output: 160

Напишите функцию, которая принимает массив целых чисел. Массив всегда будет содержать как минимум три элемента и будет состоять либо только из нечётных чисел, либо только из чётных чисел, за исключением одного "выбивающегося" числа. Функция должна определить и вернуть это выбивающееся число.

Примеры:
Ввод: [2, 4, 0, 100, 4, 11, 2602, 36] → Вывод: 11
Ввод: [160, 3, 1719, 19, 11, 13, -21] → Вывод: 160

https://www.codewars.com/kata/5526fc09a1bbd946250002dc
*/

package main

import (
	"fmt"
)

func FindOutlier(numbers []int) int {
	oddCount, evenCount := 0, 0

	for i := 0; i < 3; i++ {
		if numbers[i]%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}

	isMajorityEven := evenCount > oddCount

	for _, num := range numbers {
		if isMajorityEven && num%2 != 0 {
			return num
		} else if !isMajorityEven && num%2 == 0 {
			return num
		}
	}

	return 0
}

func main() {
	fmt.Println(FindOutlier([]int{2, 4, 0, 100, 4, 11, 2602, 36})) // Output: 11
	fmt.Println(FindOutlier([]int{160, 3, 1719, 19, 11, 13, -21})) // Output: 160
}
